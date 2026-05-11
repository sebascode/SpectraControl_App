// SpectraControl — Hue v2 SSE event stream → WebSocket fan-out.
//
// El backend mantiene una sola conexión SSE (`/eventstream/clip/v2`) al bridge
// y reenvía cada delta a todos los WS suscritos en `/ws/state`. El frontend
// usa esos eventos para que sliders, toggles y previews reflejen lo que pasa
// "afuera" (app Hue, asistente de voz, otro cliente) sin tener que poller.
package main

import (
	"bufio"
	"context"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"math"
	"net/http"
	"strings"
	"sync"
	"time"

	"github.com/gorilla/websocket"
)

// ── Hue v2 SSE payload ──────────────────────────────────────────────────────

type hueV2Envelope struct {
	CreationTime string        `json:"creationtime"`
	Data         []hueV2Update `json:"data"`
	ID           string        `json:"id"`
	Type         string        `json:"type"` // "update" | "add" | "delete"
}

type hueV2Update struct {
	ID    string `json:"id"`
	IDV1  string `json:"id_v1"` // e.g. "/lights/3" o "/groups/1"
	Type  string `json:"type"`  // "light" | "grouped_light" | ...
	Owner struct {
		RID   string `json:"rid"`
		RType string `json:"rtype"`
	} `json:"owner"`
	On *struct {
		On bool `json:"on"`
	} `json:"on,omitempty"`
	Dimming *struct {
		Brightness float64 `json:"brightness"` // 0–100 en v2
	} `json:"dimming,omitempty"`
	Color *struct {
		XY struct {
			X float64 `json:"x"`
			Y float64 `json:"y"`
		} `json:"xy"`
	} `json:"color,omitempty"`
	ColorTemperature *struct {
		Mirek *int `json:"mirek"`
	} `json:"color_temperature,omitempty"`
}

// stateEvent es lo que se envía al frontend. Usa IDs v1 para no romper la UI
// existente (que ya está montada sobre v1).
type stateEvent struct {
	Type string      `json:"type"` // "light" | "group"
	ID   string      `json:"id"`   // v1 id (string)
	On   *bool       `json:"on,omitempty"`
	Bri  *int        `json:"bri,omitempty"` // 1–254 (escala v1)
	XY   *[2]float64 `json:"xy,omitempty"`
	CT   *int        `json:"ct,omitempty"` // mirek
}

// ── Hub ─────────────────────────────────────────────────────────────────────

type stateHub struct {
	mu      sync.RWMutex
	clients map[chan stateEvent]struct{}
}

func newStateHub() *stateHub {
	return &stateHub{clients: map[chan stateEvent]struct{}{}}
}

func (h *stateHub) subscribe() chan stateEvent {
	ch := make(chan stateEvent, 32)
	h.mu.Lock()
	h.clients[ch] = struct{}{}
	h.mu.Unlock()
	return ch
}

func (h *stateHub) unsubscribe(ch chan stateEvent) {
	h.mu.Lock()
	if _, ok := h.clients[ch]; ok {
		delete(h.clients, ch)
		close(ch)
	}
	h.mu.Unlock()
}

func (h *stateHub) broadcast(ev stateEvent) {
	h.mu.RLock()
	defer h.mu.RUnlock()
	for ch := range h.clients {
		// Buffer de 32; si el cliente no consume a tiempo, droppeamos el evento
		// en vez de bloquear todo el stream.
		select {
		case ch <- ev:
		default:
		}
	}
}

var bridgeStateHub = newStateHub()

// ── SSE client ──────────────────────────────────────────────────────────────

// El bridge expone v2 solo por HTTPS con un certificado auto-firmado por
// Signify. Es una conexión a una IP de LAN ya emparejada vía pairing button;
// pinear el cert agrega fricción sin ganancia real, así que confiamos en que
// `bridgeIP` apunta al equipo correcto.
var hueV2Client = &http.Client{
	Transport: &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		// SSE es long-lived; no queremos que el Transport corte el stream
		// por idle timeout.
		IdleConnTimeout: 0,
	},
	Timeout: 0,
}

// startBridgeEventStream mantiene la conexión SSE viva en background. Cuando
// el bridge no está configurado, espera. Cuando el stream cae, reconecta con
// backoff exponencial.
func startBridgeEventStream(ctx context.Context) {
	go func() {
		backoff := time.Second
		for {
			select {
			case <-ctx.Done():
				return
			default:
			}
			if !isConfigured() {
				select {
				case <-ctx.Done():
					return
				case <-time.After(5 * time.Second):
				}
				continue
			}
			err := streamBridgeEvents(ctx)
			if err != nil {
				logWarnf("[sse] stream dropped: %v (reconnect in %s)", err, backoff)
			}
			select {
			case <-ctx.Done():
				return
			case <-time.After(backoff):
			}
			if backoff < 30*time.Second {
				backoff *= 2
			} else {
				backoff = 30 * time.Second
			}
		}
	}()
}

func streamBridgeEvents(ctx context.Context) error {
	cfgMu.RLock()
	ip, key := bridgeIP, apiKey
	cfgMu.RUnlock()
	if ip == "" || key == "" {
		return fmt.Errorf("bridge not configured")
	}

	url := fmt.Sprintf("https://%s/eventstream/clip/v2", ip)
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return err
	}
	req.Header.Set("hue-application-key", key)
	req.Header.Set("Accept", "text/event-stream")

	resp, err := hueV2Client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("HTTP %d from bridge", resp.StatusCode)
	}
	logInfof("[sse] conectado al eventstream del bridge")

	scanner := bufio.NewScanner(resp.Body)
	scanner.Buffer(make([]byte, 64*1024), 1024*1024)
	for scanner.Scan() {
		line := scanner.Text()
		if !strings.HasPrefix(line, "data:") {
			continue
		}
		payload := strings.TrimSpace(strings.TrimPrefix(line, "data:"))
		if payload == "" || payload == "[]" {
			continue
		}
		dispatchSSEPayload(payload)
	}
	if err := scanner.Err(); err != nil {
		return err
	}
	return fmt.Errorf("EOF")
}

// dispatchSSEPayload parsea un `data:` line del bridge y emite stateEvents
// hacia el hub. Separado de streamBridgeEvents para poder testearlo sin red.
func dispatchSSEPayload(payload string) {
	var envs []hueV2Envelope
	if err := json.Unmarshal([]byte(payload), &envs); err != nil {
		logWarnf("[sse] parse error: %v", err)
		return
	}
	for _, env := range envs {
		if env.Type != "" && env.Type != "update" {
			continue
		}
		for _, u := range env.Data {
			if ev, ok := mapV2Update(u); ok {
				bridgeStateHub.broadcast(ev)
			}
		}
	}
}

// mapV2Update convierte un update v2 a stateEvent v1-style. Devuelve ok=false
// cuando el update no aporta nada útil (sin id_v1 o sin campos relevantes).
func mapV2Update(u hueV2Update) (stateEvent, bool) {
	if u.IDV1 == "" {
		return stateEvent{}, false
	}
	parts := strings.Split(strings.Trim(u.IDV1, "/"), "/")
	if len(parts) != 2 {
		return stateEvent{}, false
	}
	var entType string
	switch parts[0] {
	case "lights":
		entType = "light"
	case "groups":
		entType = "group"
	default:
		return stateEvent{}, false
	}
	ev := stateEvent{Type: entType, ID: parts[1]}
	if u.On != nil {
		v := u.On.On
		ev.On = &v
	}
	if u.Dimming != nil {
		// v2: 0–100 → v1: 1–254
		bri := int(math.Round(u.Dimming.Brightness/100.0*253)) + 1
		if bri < 1 {
			bri = 1
		} else if bri > 254 {
			bri = 254
		}
		ev.Bri = &bri
	}
	if u.Color != nil {
		xy := [2]float64{
			math.Round(u.Color.XY.X*10000) / 10000,
			math.Round(u.Color.XY.Y*10000) / 10000,
		}
		ev.XY = &xy
	}
	if u.ColorTemperature != nil && u.ColorTemperature.Mirek != nil {
		ct := *u.ColorTemperature.Mirek
		ev.CT = &ct
	}
	if ev.On == nil && ev.Bri == nil && ev.XY == nil && ev.CT == nil {
		return stateEvent{}, false
	}
	return ev, true
}

// ── WebSocket fan-out ───────────────────────────────────────────────────────

func handleWsState(w http.ResponseWriter, r *http.Request) {
	conn, err := wsUpgrader.Upgrade(w, r, nil)
	if err != nil {
		logWarnf("[ws/state] upgrade: %v", err)
		return
	}
	defer conn.Close()
	logDebugf("[ws/state] conectado")

	ch := bridgeStateHub.subscribe()
	defer bridgeStateHub.unsubscribe(ch)

	conn.SetReadDeadline(time.Now().Add(60 * time.Second))
	conn.SetPongHandler(func(string) error {
		conn.SetReadDeadline(time.Now().Add(60 * time.Second))
		return nil
	})

	// Reader: solo nos importa detectar close/error.
	done := make(chan struct{})
	go func() {
		defer close(done)
		for {
			if _, _, err := conn.ReadMessage(); err != nil {
				return
			}
		}
	}()

	ticker := time.NewTicker(30 * time.Second)
	defer ticker.Stop()
	for {
		select {
		case <-done:
			logDebugf("[ws/state] desconectado")
			return
		case ev, ok := <-ch:
			if !ok {
				return
			}
			conn.SetWriteDeadline(time.Now().Add(5 * time.Second))
			if err := conn.WriteJSON(ev); err != nil {
				return
			}
		case <-ticker.C:
			conn.SetWriteDeadline(time.Now().Add(5 * time.Second))
			if err := conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				return
			}
		}
	}
}
