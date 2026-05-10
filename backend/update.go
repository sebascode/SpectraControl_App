// Update checker / installer.
//
// Estrategia: la fuente de verdad son los GitHub Releases del repo. El backend
// consulta /releases/latest, compara con `Version` (inyectada en build) y, si
// hay actualización Y estamos corriendo desde un AppImage (env $APPIMAGE
// existe), descarga el nuevo .AppImage y lo reemplaza con `os.Rename` (atómico
// en el mismo filesystem). El proceso en ejecución sigue funcionando porque
// Linux mantiene el inode antiguo abierto; al reiniciar se ejecuta el nuevo.
//
// Si $APPIMAGE no está seteado (dev / deb / rpm), `install` no escribe nada y
// devuelve el URL del release para que el usuario actualice manualmente.

package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
	"sync"
	"time"
)

const (
	updateRepo    = "sebascode/SpectraControl_App"
	updateAPIBase = "https://api.github.com"
)

var updateClient = &http.Client{Timeout: 30 * time.Second}

type ghAsset struct {
	Name               string `json:"name"`
	Size               int64  `json:"size"`
	BrowserDownloadURL string `json:"browser_download_url"`
}

type ghRelease struct {
	TagName     string    `json:"tag_name"`
	Name        string    `json:"name"`
	HTMLURL     string    `json:"html_url"`
	Body        string    `json:"body"`
	PublishedAt string    `json:"published_at"`
	Prerelease  bool      `json:"prerelease"`
	Assets      []ghAsset `json:"assets"`
}

// Cache del último check para no martillar la API de GitHub.
var (
	updateMu        sync.Mutex
	updateCacheRel  *ghRelease
	updateCacheTime time.Time
)

const updateCacheTTL = 10 * time.Minute

func appimagePath() string {
	// AppImage runtime expone APPIMAGE con la ruta absoluta del .AppImage.
	return os.Getenv("APPIMAGE")
}

func runningChannel() string {
	if appimagePath() != "" {
		return "appimage"
	}
	if Version == "dev" {
		return "dev"
	}
	return "package"
}

// parseSemver convierte "v1.2.3" o "1.2.3-rc1" en (1,2,3,"rc1"). Devuelve
// ok=false si el string no parsea como semver mínimo.
func parseSemver(s string) (major, minor, patch int, pre string, ok bool) {
	s = strings.TrimSpace(strings.TrimPrefix(s, "v"))
	if i := strings.IndexAny(s, "-+"); i >= 0 {
		pre = s[i+1:]
		s = s[:i]
	}
	parts := strings.Split(s, ".")
	if len(parts) < 1 || len(parts) > 3 {
		return 0, 0, 0, "", false
	}
	get := func(i int) (int, bool) {
		if i >= len(parts) {
			return 0, true
		}
		n, err := strconv.Atoi(parts[i])
		if err != nil {
			return 0, false
		}
		return n, true
	}
	var ok1, ok2, ok3 bool
	major, ok1 = get(0)
	minor, ok2 = get(1)
	patch, ok3 = get(2)
	if !(ok1 && ok2 && ok3) {
		return 0, 0, 0, "", false
	}
	return major, minor, patch, pre, true
}

// isNewer reporta si `latest` es estrictamente mayor que `current`. Una
// pre-release (`-rc1`) se considera menor que el mismo número sin prerelease.
// Si current="dev" siempre devolvemos true cuando latest parsea como semver.
func isNewer(current, latest string) bool {
	if current == "dev" {
		_, _, _, _, ok := parseSemver(latest)
		return ok
	}
	cM, cm, cp, cpr, cok := parseSemver(current)
	lM, lm, lp, lpr, lok := parseSemver(latest)
	if !cok || !lok {
		return false
	}
	if lM != cM {
		return lM > cM
	}
	if lm != cm {
		return lm > cm
	}
	if lp != cp {
		return lp > cp
	}
	// Misma triple (X.Y.Z): "" > "rc1" (release final > prerelease).
	if cpr == lpr {
		return false
	}
	if cpr != "" && lpr == "" {
		return true
	}
	if cpr == "" && lpr != "" {
		return false
	}
	return lpr > cpr
}

// fetchLatestRelease consulta la API con cache de updateCacheTTL.
func fetchLatestRelease(force bool) (*ghRelease, error) {
	updateMu.Lock()
	defer updateMu.Unlock()
	if !force && updateCacheRel != nil && time.Since(updateCacheTime) < updateCacheTTL {
		return updateCacheRel, nil
	}
	url := fmt.Sprintf("%s/repos/%s/releases/latest", updateAPIBase, updateRepo)
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "application/vnd.github+json")
	req.Header.Set("X-GitHub-Api-Version", "2022-11-28")
	req.Header.Set("User-Agent", "SpectraControl/"+Version)
	resp, err := updateClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode == http.StatusNotFound {
		return nil, fmt.Errorf("no hay releases publicadas todavía")
	}
	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(io.LimitReader(resp.Body, 512))
		return nil, fmt.Errorf("github api %d: %s", resp.StatusCode, string(body))
	}
	var rel ghRelease
	if err := json.NewDecoder(resp.Body).Decode(&rel); err != nil {
		return nil, err
	}
	updateCacheRel = &rel
	updateCacheTime = time.Now()
	return &rel, nil
}

func pickAppImageAsset(rel *ghRelease) *ghAsset {
	for i := range rel.Assets {
		a := &rel.Assets[i]
		n := strings.ToLower(a.Name)
		if strings.HasSuffix(n, ".appimage") {
			return a
		}
	}
	return nil
}

// ── HANDLERS ────────────────────────────────────────────────────────────────

func handleGetVersion(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, map[string]any{
		"version": Version,
		"channel": runningChannel(),
		"os":      runtime.GOOS,
		"arch":    runtime.GOARCH,
	})
}

func handleCheckUpdate(w http.ResponseWriter, r *http.Request) {
	force := r.URL.Query().Get("force") == "1"
	rel, err := fetchLatestRelease(force)
	if err != nil {
		writeErr(w, http.StatusBadGateway, err.Error())
		return
	}
	asset := pickAppImageAsset(rel)
	resp := map[string]any{
		"current":      Version,
		"latest":       strings.TrimPrefix(rel.TagName, "v"),
		"tag":          rel.TagName,
		"has_update":   isNewer(Version, rel.TagName),
		"release_url":  rel.HTMLURL,
		"published_at": rel.PublishedAt,
		"notes":        rel.Body,
		"channel":      runningChannel(),
		"can_install":  asset != nil && appimagePath() != "",
	}
	if asset != nil {
		resp["asset_name"] = asset.Name
		resp["asset_url"] = asset.BrowserDownloadURL
		resp["asset_size"] = asset.Size
	}
	writeJSON(w, resp)
}

func handleInstallUpdate(w http.ResponseWriter, r *http.Request) {
	current := appimagePath()
	if current == "" {
		writeJSON(w, map[string]any{
			"ok":         false,
			"reason":     "not_appimage",
			"message":    "El auto-update solo funciona corriendo desde el AppImage. Descarga manualmente.",
			"channel":    runningChannel(),
		})
		return
	}

	rel, err := fetchLatestRelease(true)
	if err != nil {
		writeErr(w, http.StatusBadGateway, err.Error())
		return
	}
	if !isNewer(Version, rel.TagName) {
		writeJSON(w, map[string]any{"ok": true, "already_latest": true, "version": Version})
		return
	}
	asset := pickAppImageAsset(rel)
	if asset == nil {
		writeErr(w, http.StatusNotFound, "el release no incluye un .AppImage")
		return
	}

	if err := downloadAndReplaceAppImage(current, asset); err != nil {
		writeErr(w, http.StatusInternalServerError, err.Error())
		return
	}
	writeJSON(w, map[string]any{
		"ok":               true,
		"installed":        strings.TrimPrefix(rel.TagName, "v"),
		"appimage_path":    current,
		"restart_required": true,
	})
}

// downloadAndReplaceAppImage descarga `asset` a un archivo temporal junto al
// AppImage actual y luego hace un rename atómico. Mantenerse en el mismo
// directorio garantiza que ambos paths estén en el mismo filesystem (rename(2)
// sobre filesystems distintos falla con EXDEV).
func downloadAndReplaceAppImage(currentPath string, asset *ghAsset) error {
	dir := filepath.Dir(currentPath)
	tmp, err := os.CreateTemp(dir, ".spectracontrol-update-*.AppImage")
	if err != nil {
		return fmt.Errorf("no se pudo crear archivo temporal en %s: %w", dir, err)
	}
	tmpPath := tmp.Name()
	cleanup := func() { _ = os.Remove(tmpPath) }

	req, err := http.NewRequest(http.MethodGet, asset.BrowserDownloadURL, nil)
	if err != nil {
		tmp.Close()
		cleanup()
		return err
	}
	req.Header.Set("User-Agent", "SpectraControl/"+Version)

	dlClient := &http.Client{Timeout: 10 * time.Minute}
	resp, err := dlClient.Do(req)
	if err != nil {
		tmp.Close()
		cleanup()
		return fmt.Errorf("descarga fallida: %w", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode/100 != 2 {
		tmp.Close()
		cleanup()
		return fmt.Errorf("descarga: status %d", resp.StatusCode)
	}

	written, err := io.Copy(tmp, resp.Body)
	if err != nil {
		tmp.Close()
		cleanup()
		return fmt.Errorf("escribiendo descarga: %w", err)
	}
	if asset.Size > 0 && written != asset.Size {
		tmp.Close()
		cleanup()
		return fmt.Errorf("tamaño descargado (%d) ≠ esperado (%d)", written, asset.Size)
	}
	if err := tmp.Sync(); err != nil {
		tmp.Close()
		cleanup()
		return err
	}
	if err := tmp.Close(); err != nil {
		cleanup()
		return err
	}
	if err := os.Chmod(tmpPath, 0o755); err != nil {
		cleanup()
		return err
	}
	if err := os.Rename(tmpPath, currentPath); err != nil {
		cleanup()
		return fmt.Errorf("reemplazando AppImage: %w (¿permisos sobre %s?)", err, currentPath)
	}
	return nil
}
