// entertainment.go — Hue Entertainment API (DTLS streaming)
// Implementa el protocolo HueStream v2 sobre UDP/DTLS hacia el puerto 2100 del bridge.
// Cuando está activo, las luces se muestran como "sincronizadas" en la app Hue.
package main

import (
	"bytes"
	"crypto/tls"
	"encoding/binary"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"strings"
	"sync"
	"sync/atomic"
	"time"

	"github.com/pion/dtls/v2"
)

// ── CLIP v2 HTTP CLIENT ──────────────────────────────────────────────────────
// El bridge usa HTTPS con certificado autofirmado en la API v2.

var clipHTTP = &http.Client{
	Timeout: 10 * time.Second,
	Transport: &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true}, //nolint:gosec
	},
}

func clipURL(path string) string {
	cfgMu.RLock()
	defer cfgMu.RUnlock()
	return fmt.Sprintf("https://%s/clip/v2/resource/%s", bridgeIP, path)
}

// ── HUE-APPLICATION-ID ───────────────────────────────────────────────────────
// El streaming v2 usa el application_id (NO el api_key) como PSK identity
// del handshake DTLS. Se obtiene haciendo GET /auth/v1 con el header
// hue-application-key; el bridge devuelve el id en el header de respuesta
// hue-application-id. Es estable por app key, así que lo cacheamos.

var (
	cachedAppID string
	appIDMu     sync.Mutex
)

func fetchApplicationID() (string, error) {
	appIDMu.Lock()
	defer appIDMu.Unlock()
	if cachedAppID != "" {
		return cachedAppID, nil
	}

	cfgMu.RLock()
	url := fmt.Sprintf("https://%s/auth/v1", bridgeIP)
	key := apiKey
	cfgMu.RUnlock()

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return "", err
	}
	req.Header.Set("hue-application-key", key)

	resp, err := clipHTTP.Do(req)
	if err != nil {
		return "", fmt.Errorf("/auth/v1: %w", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("/auth/v1 → HTTP %d", resp.StatusCode)
	}
	id := resp.Header.Get("hue-application-id")
	if id == "" {
		return "", fmt.Errorf("/auth/v1 sin header hue-application-id")
	}
	cachedAppID = id
	log.Printf("[ent] application_id obtenido: %s", id)
	return id, nil
}

func clipDo(method, path string, body any) ([]byte, error) {
	var bodyReader io.Reader
	if body != nil {
		data, _ := json.Marshal(body)
		bodyReader = bytes.NewReader(data)
	}
	req, err := http.NewRequest(method, clipURL(path), bodyReader)
	if err != nil {
		return nil, err
	}
	cfgMu.RLock()
	req.Header.Set("hue-application-key", apiKey)
	cfgMu.RUnlock()
	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}
	resp, err := clipHTTP.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode >= 400 {
		return respBody, fmt.Errorf("CLIP v2 %s %s → HTTP %d: %s",
			method, path, resp.StatusCode, strings.TrimSpace(string(respBody)))
	}
	return respBody, nil
}

// ── CLIP v2 TYPES ────────────────────────────────────────────────────────────

type entConfig struct {
	ID       string `json:"id"`
	Metadata struct {
		Name string `json:"name"`
	} `json:"metadata"`
	Status   string `json:"status"` // "inactive" | "active"
	Channels []struct {
		ChannelID int `json:"channel_id"`
		Members   []struct {
			Service struct {
				RID   string `json:"rid"`
				RType string `json:"rtype"`
			} `json:"service"`
		} `json:"members"`
	} `json:"channels"`
}

// ── ENTERTAINMENT STATE ──────────────────────────────────────────────────────

type entState struct {
	mu         sync.Mutex
	conn       net.Conn          // DTLS connection (nil = inactivo)
	configID   string
	channelMap map[string]byte   // v1 light ID → channel ID
	colorCh    chan []chanColor
	seq        byte
}

var ent entState

// entPushSeq counts calls to pushToEntertainment; used to throttle diagnostic logs.
var entPushSeq int64

type chanColor struct {
	id      byte
	r, g, b uint16 // 0-65535
}

// ── HELPERS ──────────────────────────────────────────────────────────────────

// buildEntIDMap devuelve un mapa: v2 entertainment resource UUID → v1 light ID ("3")
func buildEntIDMap() (map[string]string, error) {
	data, err := clipDo(http.MethodGet, "entertainment", nil)
	if err != nil {
		return nil, err
	}
	var resp struct {
		Data []struct {
			ID   string `json:"id"`
			IDv1 string `json:"id_v1"` // "/lights/3"
		} `json:"data"`
	}
	if err := json.Unmarshal(data, &resp); err != nil {
		return nil, err
	}
	m := make(map[string]string, len(resp.Data))
	for _, item := range resp.Data {
		// "/lights/3" → "3"
		if after, ok := strings.CutPrefix(item.IDv1, "/lights/"); ok {
			m[item.ID] = after
		}
	}
	return m, nil
}

// getEntertainmentConfigs lista todas las áreas de entretenimiento con sus canales.
func getEntertainmentConfigs() ([]map[string]any, error) {
	data, err := clipDo(http.MethodGet, "entertainment_configuration", nil)
	if err != nil {
		return nil, err
	}
	var resp struct {
		Data []entConfig `json:"data"`
	}
	if err := json.Unmarshal(data, &resp); err != nil {
		return nil, err
	}
	entIDMap, _ := buildEntIDMap()

	result := make([]map[string]any, 0, len(resp.Data))
	for _, cfg := range resp.Data {
		channels := make([]map[string]any, 0, len(cfg.Channels))
		for _, ch := range cfg.Channels {
			lights := make([]string, 0)
			for _, m := range ch.Members {
				if v1, ok := entIDMap[m.Service.RID]; ok {
					lights = append(lights, v1)
				}
			}
			channels = append(channels, map[string]any{
				"channel_id": ch.ChannelID,
				"lights":     lights,
			})
		}
		result = append(result, map[string]any{
			"id":       cfg.ID,
			"name":     cfg.Metadata.Name,
			"status":   cfg.Status,
			"channels": channels,
		})
	}
	return result, nil
}

// buildChannelMap construye el mapa v1 light ID → channel_id para una config dada.
func buildChannelMap(configID string) (map[string]byte, error) {
	data, err := clipDo(http.MethodGet, "entertainment_configuration/"+configID, nil)
	if err != nil {
		return nil, err
	}
	var resp struct {
		Data []entConfig `json:"data"`
	}
	if err := json.Unmarshal(data, &resp); err != nil || len(resp.Data) == 0 {
		return nil, fmt.Errorf("config %s no encontrada", configID)
	}
	cfg := resp.Data[0]

	entIDMap, err := buildEntIDMap()
	if err != nil {
		return nil, err
	}

	result := make(map[string]byte)
	for _, ch := range cfg.Channels {
		for _, m := range ch.Members {
			if v1, ok := entIDMap[m.Service.RID]; ok {
				result[v1] = byte(ch.ChannelID)
			}
		}
	}
	return result, nil
}

// ── START / STOP ─────────────────────────────────────────────────────────────

func startEntertainmentStreaming(configID string) error {
	ent.mu.Lock()
	defer ent.mu.Unlock()

	if ent.conn != nil {
		return fmt.Errorf("streaming ya activo en config %s", ent.configID)
	}

	channelMap, err := buildChannelMap(configID)
	if err != nil {
		return fmt.Errorf("buildChannelMap: %w", err)
	}
	if len(channelMap) == 0 {
		return fmt.Errorf("la config %s no tiene luces mapeadas — ¿necesita re-parear con generateclientkey?", configID)
	}

	// Encender todas las luces de la zona antes del stream. El protocolo
	// Entertainment no controla el on/off — solo el color. Si las luces
	// están off cuando se activa el stream, el bridge les pasa los frames
	// pero quedan invisibles, y los toggles manuales se encolan detrás
	// del stream (delay de varios segundos).
	var wgOn sync.WaitGroup
	for v1ID := range channelMap {
		wgOn.Add(1)
		go func(id string) {
			defer wgOn.Done()
			huePut("lights/"+id+"/state", map[string]any{"on": true, "bri": 254}) //nolint:errcheck
		}(v1ID)
	}
	wgOn.Wait()

	// Forzar cierre de cualquier sesión existente (ej. Hue Sync Box en modo fijo).
	// Si hay otro cliente con DTLS abierto sobreescribirá nuestros colores.
	// Ignoramos el error — el área puede no estar activa aún.
	clipDo(http.MethodPut, "entertainment_configuration/"+configID, map[string]string{"action": "stop"}) //nolint:errcheck
	time.Sleep(150 * time.Millisecond)

	// Activar modo streaming en el bridge (CLIP v2)
	if _, err := clipDo(http.MethodPut,
		"entertainment_configuration/"+configID,
		map[string]string{"action": "start"},
	); err != nil {
		return fmt.Errorf("activar streaming: %w", err)
	}
	log.Printf("[ent] modo streaming activo — config %s, canales: %v", configID, channelMap)

	// Abrir conexión DTLS PSK al bridge:2100
	// CRÍTICO: en v2 la PSK identity es el hue-application-id (NO el api_key).
	// Antes el bridge aceptaba el handshake con api_key y mostraba "sincronizado",
	// pero descartaba todos los frames v2 silenciosamente.
	appID, err := fetchApplicationID()
	if err != nil {
		return fmt.Errorf("application_id: %w", err)
	}

	cfgMu.RLock()
	ip := bridgeIP
	pskHex := clientKey
	cfgMu.RUnlock()

	pskBytes, err := hex.DecodeString(pskHex)
	if err != nil {
		return fmt.Errorf("clientkey hex inválido: %w", err)
	}

	addr := &net.UDPAddr{IP: net.ParseIP(ip), Port: 2100}
	dtlsCfg := &dtls.Config{
		PSK: func(_ []byte) ([]byte, error) {
			return pskBytes, nil
		},
		PSKIdentityHint: []byte(appID),
		CipherSuites:    []dtls.CipherSuiteID{dtls.TLS_PSK_WITH_AES_128_GCM_SHA256},
	}

	conn, err := dtls.Dial("udp4", addr, dtlsCfg)
	if err != nil {
		// Revertir streaming mode en el bridge
		clipDo(http.MethodPut, "entertainment_configuration/"+configID, map[string]string{"action": "stop"}) //nolint:errcheck
		return fmt.Errorf("DTLS dial: %w", err)
	}

	ent.conn = conn
	ent.configID = configID
	ent.channelMap = channelMap
	ent.colorCh = make(chan []chanColor, 1)
	ent.seq = 0

	// Frame inicial en blanco al ~50% para todos los canales: mantiene la sesión
	// DTLS viva mientras el usuario aprueba getDisplayMedia, y a la vez es visible
	// (negro = luz aparentemente apagada, hace imposible distinguir "stream activo
	// pero sin frames" de "stream caído").
	initial := make([]chanColor, 0, len(channelMap))
	for _, chID := range channelMap {
		initial = append(initial, chanColor{id: chID, r: 0x8000, g: 0x8000, b: 0x8000})
	}
	ent.colorCh <- initial

	go entertainmentSender(ent.colorCh)
	log.Printf("[ent] DTLS conectado a %s:2100", ip)
	return nil
}

func stopEntertainmentStreaming() {
	ent.mu.Lock()
	configID := ent.configID
	conn := ent.conn
	ch := ent.colorCh
	ent.conn = nil
	ent.configID = ""
	ent.channelMap = nil
	ent.colorCh = nil
	ent.mu.Unlock()

	if ch != nil {
		close(ch)
	}
	if conn != nil {
		conn.Close()
	}
	if configID != "" {
		clipDo(http.MethodPut, "entertainment_configuration/"+configID, map[string]string{"action": "stop"}) //nolint:errcheck
		log.Println("[ent] streaming detenido")
	}
}

// ── DTLS SENDER GOROUTINE ────────────────────────────────────────────────────

// lerpColors interpolates linearly between two color slices by factor t (0→1).
// Channels are matched by ID; unmatched target channels are returned as-is.
func lerpColors(from, to []chanColor, t float64) []chanColor {
	fromMap := make(map[byte]chanColor, len(from))
	for _, c := range from {
		fromMap[c.id] = c
	}
	out := make([]chanColor, len(to))
	for i, c := range to {
		if f, ok := fromMap[c.id]; ok {
			out[i] = chanColor{
				id: c.id,
				r:  uint16(float64(f.r) + (float64(c.r)-float64(f.r))*t),
				g:  uint16(float64(f.g) + (float64(c.g)-float64(f.g))*t),
				b:  uint16(float64(f.b) + (float64(c.b)-float64(f.b))*t),
			}
		} else {
			out[i] = c
		}
	}
	return out
}

// smoothstep returns a smooth interpolation factor using the 3t²-2t³ curve.
// Starts and ends slowly, fast in the middle — more natural than linear.
func smoothstep(t float64) float64 {
	return t * t * (3 - 2*t)
}

// entertainmentSender envía paquetes HueStream al bridge a 40 fps.
// Interpola con smoothstep entre el color anterior y el nuevo target para producir
// transiciones suaves aunque el WebSocket mande frames a menor frecuencia.
func entertainmentSender(colorCh <-chan []chanColor) {
	const fadeFrames = 12 // 12 × 25ms = 300ms de fade entre colores

	ticker := time.NewTicker(25 * time.Millisecond) // 40 fps
	defer ticker.Stop()

	var from, to []chanColor
	var fadeFrame int // frame actual dentro del fade (0..fadeFrames)

	for {
		select {
		case colors, ok := <-colorCh:
			if !ok {
				return // canal cerrado → goroutine termina
			}
			if to == nil {
				// primer frame: sin fade
				from, to = colors, colors
				fadeFrame = fadeFrames
			} else {
				// nueva target: arrancar fade desde el color interpolado actual
				if fadeFrame < fadeFrames {
					t := float64(fadeFrame) / float64(fadeFrames)
					from = lerpColors(from, to, t)
				} else {
					from = to
				}
				to = colors
				fadeFrame = 0
			}
		case <-ticker.C:
		}

		if to == nil {
			continue
		}

		ent.mu.Lock()
		conn := ent.conn
		configID := ent.configID
		seq := ent.seq
		ent.seq++
		ent.mu.Unlock()

		if conn == nil {
			continue
		}

		var frame []chanColor
		if fadeFrame >= fadeFrames {
			frame = to
		} else {
			t := smoothstep(float64(fadeFrame) / float64(fadeFrames))
			frame = lerpColors(from, to, t)
			fadeFrame++
		}

		pkt := buildHueStreamPacket(seq, configID, frame)
		if _, err := conn.Write(pkt); err != nil {
			log.Println("[ent] write error — limpiando estado:", err)
			stopEntertainmentStreaming()
			return
		}
		// Heartbeat cada ~200 paquetes (≈5s a 40fps) para confirmar que el sender sigue vivo
		if seq%200 == 0 {
			log.Printf("[ent] sender vivo — seq=%d canales=%d", seq, len(frame))
		}
	}
}

// ── PACKET BUILDER ───────────────────────────────────────────────────────────

// buildHueStreamPacket construye un paquete HueStream v2 (color space RGB).
//
// Header (16 bytes):
//   0-8  "HueStream" (9 bytes)
//   9    version major = 0x02
//   10   version minor = 0x00
//   11   sequence number
//   12-13 reserved
//   14   color space: 0x00 = RGB, 0x01 = XY+brillo
//   15   reserved
//
// Entertainment configuration UUID (36 bytes ASCII, posiciones 16-51):
//   El UUID con guiones, ej: "6eaf3b98-418d-48f3-89e4-a374cf9ef290"
//   Sin esto el bridge acepta el stream pero descarta los frames.
//
// Por canal (7 bytes, empezando en byte 52):
//   0    channel ID
//   1-2  R (big-endian uint16, 0-65535)
//   3-4  G
//   5-6  B
func buildHueStreamPacket(seq byte, configID string, channels []chanColor) []byte {
	pkt := make([]byte, 52+len(channels)*7)
	copy(pkt[:9], "HueStream")
	pkt[9] = 0x02
	pkt[10] = 0x00
	pkt[11] = seq
	// pkt[12], [13], [14], [15] ya son 0x00 → byte 14 = 0x00 = RGB
	copy(pkt[16:52], configID) // UUID ASCII (36 bytes)
	for i, ch := range channels {
		base := 52 + i*7
		pkt[base] = ch.id
		binary.BigEndian.PutUint16(pkt[base+1:], ch.r)
		binary.BigEndian.PutUint16(pkt[base+3:], ch.g)
		binary.BigEndian.PutUint16(pkt[base+5:], ch.b)
	}
	return pkt
}

// ── PUSH FROM WEBSOCKET ──────────────────────────────────────────────────────

// pushToEntertainment mapea los comandos de luz del WebSocket a canales DTLS.
// Retorna true si el stream de entertainment está activo (para evitar HTTP PUT).
func pushToEntertainment(lights []lightCmd) bool {
	ent.mu.Lock()
	active := ent.conn != nil
	chMap := ent.channelMap
	colorCh := ent.colorCh
	ent.mu.Unlock()

	if !active {
		return false
	}

	colors := make([]chanColor, 0, len(lights))
	briScale := float64(getBri()) / 255.0
	for _, l := range lights {
		chID, ok := chMap[l.ID]
		if !ok {
			continue
		}
		// HueStream v2 has no brightness channel — scale RGB pre-encoding.
		// 0-255 → 0-65535 (×257 = 65535/255 exacto)
		colors = append(colors, chanColor{
			id: chID,
			r:  uint16(l.R*briScale) * 257,
			g:  uint16(l.G*briScale) * 257,
			b:  uint16(l.B*briScale) * 257,
		})
	}

	// Diagnostic: log first 3 pushes + every ~30 s so the terminal shows what's happening.
	n := atomic.AddInt64(&entPushSeq, 1)
	if n <= 3 || n%600 == 0 {
		ids := make([]string, 0, len(lights))
		for _, l := range lights {
			ids = append(ids, l.ID)
		}
		mapKeys := make([]string, 0, len(chMap))
		for k := range chMap {
			mapKeys = append(mapKeys, k)
		}
		log.Printf("[ent] push #%d: IDs=%v  chMap=%v  → %d/%d mapeadas",
			n, ids, mapKeys, len(colors), len(lights))
	}

	if len(colors) == 0 {
		// None of the requested lights are in this zone's channel map.
		// Return false so the caller can fall back to HTTP PUT instead of silently dropping colors.
		return false
	}

	// Latest-frame-wins: descarta el frame pendiente si el sender está ocupado
	select {
	case <-colorCh:
	default:
	}
	colorCh <- colors
	return true
}

// ── HTTP HANDLERS ────────────────────────────────────────────────────────────

func handleListEntertainment(w http.ResponseWriter, r *http.Request) {
	if !requireConfig(w) {
		return
	}
	cfgMu.RLock()
	hasKey := clientKey != ""
	cfgMu.RUnlock()
	if !hasKey {
		writeErr(w, http.StatusPreconditionFailed,
			"Se requiere clientkey. Vuelve a vincular el bridge presionando el botón físico.")
		return
	}
	configs, err := getEntertainmentConfigs()
	if err != nil {
		writeErr(w, http.StatusServiceUnavailable, err.Error())
		return
	}
	ent.mu.Lock()
	activeID := ent.configID
	ent.mu.Unlock()
	writeJSON(w, map[string]any{
		"configs":   configs,
		"active_id": activeID,
	})
}

func handleStartEntertainment(w http.ResponseWriter, r *http.Request) {
	if !requireConfig(w) {
		return
	}
	var body struct {
		ConfigID string `json:"config_id"`
	}
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil || body.ConfigID == "" {
		writeErr(w, http.StatusBadRequest, "Falta config_id")
		return
	}
	if err := startEntertainmentStreaming(body.ConfigID); err != nil {
		writeErr(w, http.StatusInternalServerError, err.Error())
		return
	}
	writeJSON(w, map[string]bool{"ok": true})
}

func handleStopEntertainment(w http.ResponseWriter, r *http.Request) {
	stopEntertainmentStreaming()
	writeJSON(w, map[string]bool{"ok": true})
}
