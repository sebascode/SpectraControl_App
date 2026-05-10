// SpectraControl Backend — Go
// Controla Philips Hue Bridge via API local REST v1.
package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"math"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/gorilla/websocket"
)

// Version se inyecta en build con -ldflags "-X main.Version=<tag>".
// "dev" indica build local / sin tag.
var Version = "dev"

// ── CONFIG ───────────────────────────────────────────────────────────────────

var (
	bridgeIP  string
	apiKey    string
	clientKey string // PSK para DTLS Entertainment API (hex, 32 bytes)
	cfgMu     sync.RWMutex
)

func configFilePath() string {
	dir, err := os.UserConfigDir()
	if err != nil {
		dir = filepath.Join(os.Getenv("HOME"), ".config")
	}
	d := filepath.Join(dir, "spectracontrol")
	os.MkdirAll(d, 0o700)
	return filepath.Join(d, "hue_config.json")
}

type persistedCfg struct {
	IP        string `json:"ip"`
	APIKey    string `json:"api_key"`
	ClientKey string `json:"client_key,omitempty"`
}

func loadPersistedConfig() {
	data, err := os.ReadFile(configFilePath())
	if err != nil {
		return
	}
	var cfg persistedCfg
	if json.Unmarshal(data, &cfg) != nil {
		return
	}
	cfgMu.Lock()
	bridgeIP = cfg.IP
	apiKey = cfg.APIKey
	clientKey = cfg.ClientKey
	cfgMu.Unlock()

	// Migración desde el .hue_config viejo de Python (está en CWD cuando se corre desde backend/)
	if bridgeIP == "" {
		for _, old := range []string{".hue_config", filepath.Join(exeDir(), ".hue_config")} {
			if data2, err := os.ReadFile(old); err == nil {
				var old2 persistedCfg
				if json.Unmarshal(data2, &old2) == nil && old2.IP != "" {
					cfgMu.Lock()
					bridgeIP = old2.IP
					apiKey = old2.APIKey
					cfgMu.Unlock()
					saveConfig()
					os.Remove(old)
					log.Printf("Configuración migrada desde %s", old)
					break
				}
			}
		}
	}
}

func saveConfig() error {
	cfgMu.RLock()
	cfg := persistedCfg{IP: bridgeIP, APIKey: apiKey, ClientKey: clientKey}
	cfgMu.RUnlock()
	data, _ := json.Marshal(cfg)
	return os.WriteFile(configFilePath(), data, 0o600)
}

func isConfigured() bool {
	cfgMu.RLock()
	defer cfgMu.RUnlock()
	return bridgeIP != "" && apiKey != ""
}

// ── PATHS ────────────────────────────────────────────────────────────────────

// exeDir returns the directory of the running binary.
// Falls back to "backend" (relative to CWD) when running via `go run`.
// `go run` compila a $TMPDIR/go-build*. Match exacto sobre "go-build" para no
// confundir con AppImage, que se monta en /tmp/.mount_<random>/...
func exeDir() string {
	exe, err := os.Executable()
	if err != nil || strings.Contains(exe, "/go-build") {
		return "backend"
	}
	return filepath.Dir(exe)
}

func frontendPath(override string) string {
	if override != "" {
		return override
	}
	dir := exeDir()
	if dir == "backend" {
		// Dev: CWD is project root, frontend is at ./frontend
		return "frontend"
	}
	// Release: binary is at {resources}/backend/spectractl
	// Frontend is at {resources}/frontend/
	return filepath.Join(dir, "..", "frontend")
}

// ── HUE HTTP CLIENT ──────────────────────────────────────────────────────────

var hueClient = &http.Client{Timeout: 5 * time.Second}

func hueURL(path string) string {
	cfgMu.RLock()
	defer cfgMu.RUnlock()
	return fmt.Sprintf("http://%s/api/%s/%s", bridgeIP, apiKey, path)
}

func hueGet(path string) ([]byte, error) {
	resp, err := hueClient.Get(hueURL(path))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return io.ReadAll(resp.Body)
}

func huePut(path string, body any) error {
	data, _ := json.Marshal(body)
	req, err := http.NewRequest(http.MethodPut, hueURL(path), bytes.NewReader(data))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	resp, err := hueClient.Do(req)
	if err != nil {
		return err
	}
	resp.Body.Close()
	return nil
}

func huePostRaw(url string, body any) ([]byte, error) {
	data, _ := json.Marshal(body)
	resp, err := hueClient.Post(url, "application/json", bytes.NewReader(data))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return io.ReadAll(resp.Body)
}

// ── COLOR CONVERSION ─────────────────────────────────────────────────────────

func gamma(v float64) float64 {
	if v > 0.04045 {
		return math.Pow((v+0.055)/1.055, 2.4)
	}
	return v / 12.92
}

// rgbToXY converts sRGB (0–255) to CIE XY using the Philips Hue color matrix.
func rgbToXY(r, g, b float64) [2]float64 {
	r, g, b = gamma(r/255), gamma(g/255), gamma(b/255)
	X := r*0.664511 + g*0.154324 + b*0.162028
	Y := r*0.283881 + g*0.668433 + b*0.047685
	Z := r*0.000088 + g*0.072310 + b*0.986039
	t := X + Y + Z
	if t == 0 {
		return [2]float64{0, 0}
	}
	return [2]float64{
		math.Round(X/t*10000) / 10000,
		math.Round(Y/t*10000) / 10000,
	}
}

func hexToRGB(hex string) (float64, float64, float64) {
	hex = strings.TrimPrefix(hex, "#")
	var r, g, b int
	fmt.Sscanf(hex, "%02x%02x%02x", &r, &g, &b)
	return float64(r), float64(g), float64(b)
}

// ── MIDDLEWARE ───────────────────────────────────────────────────────────────

func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "*")
		w.Header().Set("Access-Control-Allow-Headers", "*")
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusNoContent)
			return
		}
		next.ServeHTTP(w, r)
	})
}

// ── RESPONSE HELPERS ─────────────────────────────────────────────────────────

func writeJSON(w http.ResponseWriter, v any) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(v)
}

func writeErr(w http.ResponseWriter, code int, msg string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(map[string]string{"detail": msg})
}

func requireConfig(w http.ResponseWriter) bool {
	if !isConfigured() {
		writeErr(w, http.StatusServiceUnavailable, "Bridge no configurado. Ve a /api/config primero.")
		return false
	}
	return true
}

// ── HANDLERS — Config ────────────────────────────────────────────────────────

func handleGetConfig(w http.ResponseWriter, r *http.Request) {
	cfgMu.RLock()
	ip, key := bridgeIP, apiKey
	cfgMu.RUnlock()
	writeJSON(w, map[string]any{
		"bridge_ip":  ip,
		"has_api_key": key != "",
		"configured": ip != "" && key != "",
	})
}

func handleSetConfig(w http.ResponseWriter, r *http.Request) {
	var body struct {
		IP     string `json:"ip"`
		APIKey string `json:"api_key"`
	}
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		writeErr(w, http.StatusBadRequest, err.Error())
		return
	}
	cfgMu.Lock()
	bridgeIP = body.IP
	apiKey = body.APIKey
	cfgMu.Unlock()
	if err := saveConfig(); err != nil {
		writeErr(w, http.StatusInternalServerError, err.Error())
		return
	}
	writeJSON(w, map[string]bool{"ok": true})
}

func handleDiscover(w http.ResponseWriter, r *http.Request) {
	resp, err := hueClient.Get("https://discovery.meethue.com/")
	if err != nil {
		writeErr(w, http.StatusServiceUnavailable, err.Error())
		return
	}
	defer resp.Body.Close()
	var bridges any
	json.NewDecoder(resp.Body).Decode(&bridges)
	writeJSON(w, map[string]any{"bridges": bridges})
}

func handlePair(w http.ResponseWriter, r *http.Request) {
	var body map[string]string
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		writeErr(w, http.StatusBadRequest, err.Error())
		return
	}
	ip := body["ip"]
	if ip == "" {
		writeErr(w, http.StatusBadRequest, "Falta 'ip'")
		return
	}
	result, err := huePostRaw(
		fmt.Sprintf("http://%s/api", ip),
		map[string]any{"devicetype": "spectra_control#linux", "generateclientkey": true},
	)
	if err != nil {
		writeErr(w, http.StatusServiceUnavailable, err.Error())
		return
	}
	var arr []map[string]any
	if err := json.Unmarshal(result, &arr); err != nil || len(arr) == 0 {
		writeErr(w, http.StatusInternalServerError, string(result))
		return
	}
	if s, ok := arr[0]["success"].(map[string]any); ok {
		newKey := fmt.Sprint(s["username"])
		newClientKey := fmt.Sprint(s["clientkey"])
		cfgMu.Lock()
		apiKey = newKey
		if newClientKey != "" && newClientKey != "<nil>" {
			clientKey = newClientKey
		}
		cfgMu.Unlock()
		saveConfig()
		writeJSON(w, map[string]any{
			"api_key":    newKey,
			"has_clientkey": clientKey != "",
		})
		return
	}
	if e, ok := arr[0]["error"].(map[string]any); ok {
		writeErr(w, http.StatusForbidden, fmt.Sprint(e["description"]))
		return
	}
	writeErr(w, http.StatusInternalServerError, string(result))
}

// ── HANDLERS — Lights ────────────────────────────────────────────────────────

func handleGetLights(w http.ResponseWriter, r *http.Request) {
	if !requireConfig(w) {
		return
	}
	data, err := hueGet("lights")
	if err != nil {
		writeErr(w, http.StatusServiceUnavailable, err.Error())
		return
	}
	var raw map[string]map[string]any
	if err := json.Unmarshal(data, &raw); err != nil {
		writeErr(w, http.StatusInternalServerError, err.Error())
		return
	}
	lights := make([]map[string]any, 0, len(raw))
	for lid, info := range raw {
		state, _ := info["state"].(map[string]any)
		if state == nil {
			state = map[string]any{}
		}
		lights = append(lights, map[string]any{
			"id":         lid,
			"name":       info["name"],
			"type":       info["type"],
			"on":         state["on"],
			"bri":        state["bri"],
			"reachable":  state["reachable"],
			"color_mode": state["colormode"],
			"xy":         state["xy"],
			"ct":         state["ct"],
			"hue":        state["hue"],
			"sat":        state["sat"],
		})
	}
	writeJSON(w, map[string]any{"lights": lights})
}

func handleSetLightState(w http.ResponseWriter, r *http.Request) {
	if !requireConfig(w) {
		return
	}
	id := chi.URLParam(r, "id")
	var body map[string]any
	json.NewDecoder(r.Body).Decode(&body)
	if err := huePut("lights/"+id+"/state", body); err != nil {
		writeErr(w, http.StatusServiceUnavailable, err.Error())
		return
	}
	writeJSON(w, map[string]bool{"ok": true})
}

func handleSetLightColor(w http.ResponseWriter, r *http.Request) {
	if !requireConfig(w) {
		return
	}
	id := chi.URLParam(r, "id")
	var body struct {
		HexColor string `json:"hex_color"`
		Bri      int    `json:"bri"`
		TT       int    `json:"transitiontime"`
	}
	body.Bri, body.TT = 200, 4
	json.NewDecoder(r.Body).Decode(&body)
	rr, gg, bb := hexToRGB(body.HexColor)
	xy := rgbToXY(rr, gg, bb)
	if err := huePut("lights/"+id+"/state", map[string]any{
		"on": true, "xy": xy[:], "bri": body.Bri, "transitiontime": body.TT,
	}); err != nil {
		writeErr(w, http.StatusServiceUnavailable, err.Error())
		return
	}
	writeJSON(w, map[string]bool{"ok": true})
}

// ── HANDLERS — Groups ────────────────────────────────────────────────────────

func handleGetGroups(w http.ResponseWriter, r *http.Request) {
	if !requireConfig(w) {
		return
	}
	data, err := hueGet("groups")
	if err != nil {
		writeErr(w, http.StatusServiceUnavailable, err.Error())
		return
	}
	var raw map[string]map[string]any
	if err := json.Unmarshal(data, &raw); err != nil {
		writeErr(w, http.StatusInternalServerError, err.Error())
		return
	}
	grps := make([]map[string]any, 0, len(raw))
	for gid, info := range raw {
		action, _ := info["action"].(map[string]any)
		if action == nil {
			action = map[string]any{}
		}
		rawLights, _ := info["lights"].([]any)
		lids := make([]string, 0, len(rawLights))
		for _, l := range rawLights {
			if s, ok := l.(string); ok {
				lids = append(lids, s)
			}
		}
		grps = append(grps, map[string]any{
			"id":     gid,
			"name":   info["name"],
			"type":   info["type"],
			"lights": lids,
			"on":     action["on"],
			"bri":    action["bri"],
			"ct":     action["ct"],
			"xy":     action["xy"],
		})
	}
	writeJSON(w, map[string]any{"groups": grps})
}

func handleSetGroupAction(w http.ResponseWriter, r *http.Request) {
	if !requireConfig(w) {
		return
	}
	id := chi.URLParam(r, "id")
	var body map[string]any
	json.NewDecoder(r.Body).Decode(&body)
	if err := huePut("groups/"+id+"/action", body); err != nil {
		writeErr(w, http.StatusServiceUnavailable, err.Error())
		return
	}
	writeJSON(w, map[string]bool{"ok": true})
}

func handleSetGroupColor(w http.ResponseWriter, r *http.Request) {
	if !requireConfig(w) {
		return
	}
	id := chi.URLParam(r, "id")
	var body struct {
		HexColor string `json:"hex_color"`
		Bri      int    `json:"bri"`
		TT       int    `json:"transitiontime"`
	}
	body.Bri, body.TT = 200, 4
	json.NewDecoder(r.Body).Decode(&body)
	rr, gg, bb := hexToRGB(body.HexColor)
	xy := rgbToXY(rr, gg, bb)
	if err := huePut("groups/"+id+"/action", map[string]any{
		"on": true, "xy": xy[:], "bri": body.Bri, "transitiontime": body.TT,
	}); err != nil {
		writeErr(w, http.StatusServiceUnavailable, err.Error())
		return
	}
	writeJSON(w, map[string]bool{"ok": true})
}

// ── HANDLER — Sync status (compat) ──────────────────────────────────────────

func handleGetSync(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, map[string]any{"running": false, "interval": 1.0, "group_ids": []string{}})
}

func handleSetSync(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, map[string]bool{"ok": true})
}

// ── WEBSOCKET — Screen Sync ──────────────────────────────────────────────────

var wsUpgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true },
}

type lightCmd struct {
	ID string  `json:"id"`
	R  float64 `json:"r"`
	G  float64 `json:"g"`
	B  float64 `json:"b"`
}

type wsMsg struct {
	GroupIDs []string   `json:"group_ids"`
	Lights   []lightCmd `json:"lights"`
	R        *float64   `json:"r"`
	G        *float64   `json:"g"`
	B        *float64   `json:"b"`
	Bri      *int       `json:"bri"`
	TT       *int       `json:"transitiontime"`
}

type colorUpdate struct {
	lights   []lightCmd
	groupIDs []string
	bri      int
	tt       int
	r, g, b  *float64
}

func sendColorUpdate(u colorUpdate) {
	if len(u.lights) > 0 {
		// Si el stream de entertainment está activo, usamos DTLS (más rápido, muestra
		// "sincronizado" en la app Hue). Si no, usamos HTTP PUT normal.
		if pushToEntertainment(u.lights) {
			return
		}
		var wg sync.WaitGroup
		for _, l := range u.lights {
			wg.Add(1)
			go func(l lightCmd) {
				defer wg.Done()
				xy := rgbToXY(l.R, l.G, l.B)
				huePut("lights/"+l.ID+"/state", map[string]any{
					"on": true, "xy": xy[:], "bri": u.bri, "transitiontime": u.tt,
				})
			}(l)
		}
		wg.Wait()
	} else if u.r != nil && len(u.groupIDs) > 0 {
		xy := rgbToXY(*u.r, *u.g, *u.b)
		var wg sync.WaitGroup
		for _, gid := range u.groupIDs {
			wg.Add(1)
			go func(gid string) {
				defer wg.Done()
				huePut("groups/"+gid+"/action", map[string]any{
					"on": true, "xy": xy[:], "bri": u.bri, "transitiontime": u.tt,
				})
			}(gid)
		}
		wg.Wait()
	}
}

func handleWsColor(w http.ResponseWriter, r *http.Request) {
	conn, err := wsUpgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("[ws/color] upgrade:", err)
		return
	}
	defer conn.Close()
	log.Println("[ws/color] conectado")

	// Canal de tamaño 1: el sender toma frames, el reader nunca bloquea.
	// Si llega un frame nuevo mientras el sender está ocupado, reemplaza al pendiente
	// (semántica "último frame gana") — frames intermedios se descartan, no se acumulan.
	ch := make(chan colorUpdate, 1)
	defer close(ch)

	go func() {
		for u := range ch {
			if isConfigured() {
				sendColorUpdate(u)
			}
		}
	}()

	var groupIDs []string

	for {
		_, raw, err := conn.ReadMessage()
		if err != nil {
			if !websocket.IsCloseError(err, websocket.CloseNormalClosure, websocket.CloseGoingAway) {
				log.Println("[ws/color] error:", err)
			}
			break
		}
		var msg wsMsg
		if json.Unmarshal(raw, &msg) != nil {
			continue
		}
		if msg.GroupIDs != nil {
			groupIDs = msg.GroupIDs
		}
		if msg.Bri != nil {
			setBri(*msg.Bri)
		}

		if len(msg.Lights) == 0 && msg.R == nil {
			continue // mensaje de configuración, sin color
		}

		// Sync arranca → detener cualquier escena dinámica activa.
		if id, _ := sceneStatus(); id != "" {
			stopScene()
		}

		tt := 2
		if msg.TT != nil {
			tt = *msg.TT
		}

		u := colorUpdate{
			lights: msg.Lights, groupIDs: groupIDs,
			bri: int(getBri()), tt: tt,
			r: msg.R, g: msg.G, b: msg.B,
		}
		// Descarta el frame pendiente si el sender aún no lo procesó,
		// y envía el frame más reciente.
		select {
		case <-ch:
		default:
		}
		ch <- u
	}
	log.Println("[ws/color] desconectado")
}

// ── MAIN ─────────────────────────────────────────────────────────────────────

func main() {
	var (
		addr        = flag.String("addr", ":8000", "dirección de escucha")
		frontendDir = flag.String("frontend", "", "ruta al directorio frontend (auto-detectada si vacía)")
	)
	flag.Parse()

	loadPersistedConfig()

	r := chi.NewRouter()
	r.Use(corsMiddleware)

	// Config
	r.Get("/api/config", handleGetConfig)
	r.Post("/api/config", handleSetConfig)
	r.Post("/api/discover", handleDiscover)
	r.Post("/api/pair", handlePair)

	// Lights
	r.Get("/api/lights", handleGetLights)
	r.Put("/api/lights/{id}/state", handleSetLightState)
	r.Put("/api/lights/{id}/color", handleSetLightColor)

	// Groups
	r.Get("/api/groups", handleGetGroups)
	r.Put("/api/groups/{id}/action", handleSetGroupAction)
	r.Put("/api/groups/{id}/color", handleSetGroupColor)

	// Sync (compat con frontend existente)
	r.Get("/api/sync", handleGetSync)
	r.Post("/api/sync", handleSetSync)

	// Entertainment API (DTLS streaming)
	r.Get("/api/entertainment", handleListEntertainment)
	r.Post("/api/entertainment/start", handleStartEntertainment)
	r.Post("/api/entertainment/stop", handleStopEntertainment)

	// Dynamic scenes
	r.Get("/api/scenes", handleListScenes)
	r.Post("/api/scenes/start", handleStartScene)
	r.Post("/api/scenes/stop", handleStopScene)

	// Global brightness (compartido por sync y escenas)
	r.Get("/api/brightness", handleGetBrightness)
	r.Post("/api/brightness", handleSetBrightness)

	// Updater (consulta GitHub Releases y reemplaza el AppImage in-place)
	r.Get("/api/version", handleGetVersion)
	r.Get("/api/update/check", handleCheckUpdate)
	r.Post("/api/update/install", handleInstallUpdate)

	// WebSocket
	r.Get("/ws/color", handleWsColor)

	// Frontend estático — sin cache para que los cambios en index.html se reflejen al recargar.
	fe := frontendPath(*frontendDir)
	log.Printf("Sirviendo frontend desde: %s", fe)
	fileServer := http.FileServer(http.Dir(fe))
	r.Handle("/*", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Cache-Control", "no-store, must-revalidate")
		fileServer.ServeHTTP(w, r)
	}))

	log.Printf("SpectraControl %s (%s) escuchando en http://localhost%s", Version, runningChannel(), *addr)
	log.Fatal(http.ListenAndServe(*addr, r))
}
