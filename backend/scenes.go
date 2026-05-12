// Dynamic scenes: animated color sequences played on a set of lights.
// Scenes are loaded from JSON files (embedded built-ins + user dir), so
// adding a new preset does NOT require recompiling. The runner reuses
// sendColorUpdate, so DTLS is preferred when active and HTTP PUT is the
// fallback (same path as screen sync).
package main

import (
	"context"
	"embed"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strings"
	"sync"
	"time"

	"github.com/go-chi/chi/v5"
)

// IDs de escenas son nombre de archivo en el user dir, así que se filtran a
// caracteres seguros para path: minúsculas, dígitos y guion. Sin puntos
// (evita ".." y oculta), sin slashes, sin espacios.
var sceneIDRegex = regexp.MustCompile(`^[a-z0-9-]+$`)

//go:embed scenes/*.json
var embeddedScenes embed.FS

type sceneKeyframe struct {
	T     float64  `json:"t"`     // 0..1 normalized within the cycle
	Color [3]uint8 `json:"color"` // R, G, B
}

type scenePreset struct {
	ID        string          `json:"id"`
	Name      string          `json:"name"`
	Duration  float64         `json:"duration"` // seconds for one full cycle
	Loop      bool            `json:"loop"`
	Stagger   float64         `json:"stagger"` // phase offset per light index, 0..1
	Keyframes []sceneKeyframe `json:"keyframes"`
	Source    string          `json:"source,omitempty"` // "builtin" | "user"
}

func userScenesDir() string {
	dir, err := os.UserConfigDir()
	if err != nil {
		dir = filepath.Join(os.Getenv("HOME"), ".config")
	}
	d := filepath.Join(dir, "spectracontrol", "scenes")
	os.MkdirAll(d, 0o755)
	return d
}

// loadScenePresets reads all JSON scenes — first the embedded built-ins,
// then the user dir. User scenes with the same ID override built-ins.
// Invalid files are logged and skipped; one bad file never breaks the list.
func loadScenePresets() []scenePreset {
	var all []scenePreset

	if entries, err := embeddedScenes.ReadDir("scenes"); err == nil {
		for _, e := range entries {
			if e.IsDir() || !strings.HasSuffix(e.Name(), ".json") {
				continue
			}
			data, err := embeddedScenes.ReadFile("scenes/" + e.Name())
			if err != nil {
				continue
			}
			var s scenePreset
			if err := json.Unmarshal(data, &s); err != nil {
				logWarnf("[scenes] embedded %s: %v", e.Name(), err)
				continue
			}
			s.Source = "builtin"
			all = append(all, s)
		}
	}

	userDir := userScenesDir()
	if entries, err := os.ReadDir(userDir); err == nil {
		for _, e := range entries {
			if e.IsDir() || !strings.HasSuffix(e.Name(), ".json") {
				continue
			}
			data, err := os.ReadFile(filepath.Join(userDir, e.Name()))
			if err != nil {
				continue
			}
			var s scenePreset
			if err := json.Unmarshal(data, &s); err != nil {
				logWarnf("[scenes] user %s: %v", e.Name(), err)
				continue
			}
			s.Source = "user"
			replaced := false
			for i := range all {
				if all[i].ID == s.ID {
					all[i] = s
					replaced = true
					break
				}
			}
			if !replaced {
				all = append(all, s)
			}
		}
	}

	for i := range all {
		sort.Slice(all[i].Keyframes, func(a, b int) bool {
			return all[i].Keyframes[a].T < all[i].Keyframes[b].T
		})
	}
	sort.Slice(all, func(i, j int) bool { return all[i].Name < all[j].Name })
	return all
}

var (
	sceneMu       sync.Mutex
	sceneCancel   context.CancelFunc
	sceneActiveID string
	sceneLightIDs []string
)

func interpolateScene(s *scenePreset, phase float64) (uint8, uint8, uint8) {
	phase -= float64(int(phase))
	if phase < 0 {
		phase += 1
	}
	kfs := s.Keyframes
	if len(kfs) == 0 {
		return 0, 0, 0
	}
	prev := kfs[0]
	next := kfs[len(kfs)-1]
	for i := 0; i < len(kfs)-1; i++ {
		if phase >= kfs[i].T && phase <= kfs[i+1].T {
			prev, next = kfs[i], kfs[i+1]
			break
		}
	}
	span := next.T - prev.T
	t := 0.0
	if span > 0 {
		t = (phase - prev.T) / span
	}
	lerp := func(a, b uint8) uint8 {
		return uint8(float64(a) + (float64(b)-float64(a))*t)
	}
	return lerp(prev.Color[0], next.Color[0]),
		lerp(prev.Color[1], next.Color[1]),
		lerp(prev.Color[2], next.Color[2])
}

func sceneRunner(ctx context.Context, preset scenePreset, lightIDs []string) {
	ticker := time.NewTicker(50 * time.Millisecond)
	defer ticker.Stop()
	start := time.Now()

	for {
		select {
		case <-ctx.Done():
			return
		case now := <-ticker.C:
			elapsed := now.Sub(start).Seconds()
			if !preset.Loop && elapsed >= preset.Duration {
				stopScene()
				return
			}
			cycleT := elapsed / preset.Duration
			cycleT -= float64(int(cycleT))

			cmds := make([]lightCmd, 0, len(lightIDs))
			for i, lid := range lightIDs {
				phase := cycleT + float64(i)*preset.Stagger
				r, g, b := interpolateScene(&preset, phase)
				cmds = append(cmds, lightCmd{
					ID: lid, R: float64(r), G: float64(g), B: float64(b),
				})
			}
			if isConfigured() {
				sendColorUpdate(colorUpdate{
					lights: cmds,
					bri:    int(getBri()),
					tt:     2,
				})
			}
		}
	}
}

func startScene(id string, lightIDs []string) error {
	presets := loadScenePresets()
	var preset *scenePreset
	for i := range presets {
		if presets[i].ID == id {
			preset = &presets[i]
			break
		}
	}
	if preset == nil {
		return fmt.Errorf("escena no encontrada: %s", id)
	}

	sceneMu.Lock()
	if sceneCancel != nil {
		sceneCancel()
	}
	ctx, cancel := context.WithCancel(context.Background())
	sceneCancel = cancel
	sceneActiveID = id
	sceneLightIDs = append([]string(nil), lightIDs...)
	sceneMu.Unlock()

	go sceneRunner(ctx, *preset, lightIDs)
	return nil
}

func stopScene() {
	sceneMu.Lock()
	defer sceneMu.Unlock()
	if sceneCancel != nil {
		sceneCancel()
		sceneCancel = nil
	}
	sceneActiveID = ""
	sceneLightIDs = nil
}

func sceneStatus() (string, []string) {
	sceneMu.Lock()
	defer sceneMu.Unlock()
	return sceneActiveID, append([]string(nil), sceneLightIDs...)
}

// ── HANDLERS ─────────────────────────────────────────────────────────────────

func handleListScenes(w http.ResponseWriter, r *http.Request) {
	id, lids := sceneStatus()
	writeJSON(w, map[string]any{
		"scenes":        loadScenePresets(),
		"active_id":     id,
		"active_lights": lids,
		"user_dir":      userScenesDir(),
	})
}

func handleStartScene(w http.ResponseWriter, r *http.Request) {
	if !requireConfig(w) {
		return
	}
	var body struct {
		SceneID  string   `json:"scene_id"`
		LightIDs []string `json:"light_ids"`
	}
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		writeErr(w, http.StatusBadRequest, err.Error())
		return
	}
	if body.SceneID == "" || len(body.LightIDs) == 0 {
		writeErr(w, http.StatusBadRequest, "scene_id y light_ids son requeridos")
		return
	}
	if err := startScene(body.SceneID, body.LightIDs); err != nil {
		writeErr(w, http.StatusBadRequest, err.Error())
		return
	}
	logInfof("[scenes] iniciada %s sobre %d luces", body.SceneID, len(body.LightIDs))
	writeJSON(w, map[string]bool{"ok": true})
}

func handleStopScene(w http.ResponseWriter, r *http.Request) {
	stopScene()
	logInfof("[scenes] detenida")
	writeJSON(w, map[string]bool{"ok": true})
}

// handleCreateScene crea o actualiza una escena de usuario. Las built-in
// embebidas son read-only; si llega un id que coincide con una built-in se
// escribe igualmente en el user dir y `loadScenePresets()` ya las sobreescribe
// al leer disco (el override convive sin tocar el embed).
func handleCreateScene(w http.ResponseWriter, r *http.Request) {
	var s scenePreset
	if err := json.NewDecoder(r.Body).Decode(&s); err != nil {
		writeErr(w, http.StatusBadRequest, "json inválido: "+err.Error())
		return
	}
	if !sceneIDRegex.MatchString(s.ID) {
		writeErr(w, http.StatusBadRequest, "id debe ser [a-z0-9-]+")
		return
	}
	if len(s.ID) > 64 {
		writeErr(w, http.StatusBadRequest, "id demasiado largo (máx 64)")
		return
	}
	if strings.TrimSpace(s.Name) == "" {
		writeErr(w, http.StatusBadRequest, "name requerido")
		return
	}
	if s.Duration <= 0 {
		writeErr(w, http.StatusBadRequest, "duration debe ser > 0")
		return
	}
	if len(s.Keyframes) < 2 {
		writeErr(w, http.StatusBadRequest, "se requieren al menos 2 keyframes")
		return
	}
	for _, k := range s.Keyframes {
		if k.T < 0 || k.T > 1 {
			writeErr(w, http.StatusBadRequest, "keyframe.t fuera de [0,1]")
			return
		}
	}
	// Persistir source="user" siempre — el archivo vive en user dir y el
	// loader lo sobreescribe sobre cualquier built-in homónima.
	s.Source = "user"

	path := filepath.Join(userScenesDir(), s.ID+".json")
	data, err := json.MarshalIndent(s, "", "    ")
	if err != nil {
		writeErr(w, http.StatusInternalServerError, "marshal: "+err.Error())
		return
	}
	if err := os.WriteFile(path, data, 0o644); err != nil {
		writeErr(w, http.StatusInternalServerError, "write: "+err.Error())
		return
	}
	logInfof("[scenes] guardada %s (%d keyframes)", s.ID, len(s.Keyframes))
	writeJSON(w, s)
}

// handleDeleteScene borra el archivo de una escena de usuario. Las built-in
// no se pueden borrar (viven embebidas en el binario); si alguien pide
// borrar un id que solo existe como built-in, devolvemos 404.
func handleDeleteScene(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if !sceneIDRegex.MatchString(id) {
		writeErr(w, http.StatusBadRequest, "id inválido")
		return
	}
	path := filepath.Join(userScenesDir(), id+".json")
	if err := os.Remove(path); err != nil {
		if os.IsNotExist(err) {
			writeErr(w, http.StatusNotFound, "escena de usuario no encontrada")
			return
		}
		writeErr(w, http.StatusInternalServerError, err.Error())
		return
	}
	// Si la escena borrada estaba activa, frenar el runner.
	activeID, _ := sceneStatus()
	if activeID == id {
		stopScene()
	}
	logInfof("[scenes] borrada %s", id)
	writeJSON(w, map[string]bool{"ok": true})
}
