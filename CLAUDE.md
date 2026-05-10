# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## What this is

**SpectraControl** — Linux desktop app for controlling Philips Hue lights. Replaces the official Hue Sync app which doesn't exist for Linux. Targets Bazzite (Fedora immutable) + KDE Plasma on Wayland.

## Running the app

```bash
# Standalone (browser at http://localhost:8000)
./run.sh
# run.sh compiles the Go binary if needed, then runs it directly:
#   cd backend && go build -o spectractl .
#   ./backend/spectractl -addr :8000

# Tauri dev mode (desktop window, auto-starts backend)
cargo tauri dev   # from project root, requires Rust + tauri-cli
```

## Stack

- **Backend:** Go · net/http · chi · gorilla/websocket · pion/dtls — `backend/main.go` + `backend/entertainment.go`
- **Frontend:** Single `frontend/index.html` — vanilla HTML/CSS/JS, no framework, no bundler
- **Desktop shell:** Tauri 2 (Rust) — `src-tauri/`
- **Hue API v1:** local REST `http://<ip>/api/<key>/...` (lights, groups, config)
- **Hue API v2 (CLIP):** HTTPS `https://<ip>/clip/v2/resource/...` (entertainment configs)
- **Entertainment streaming:** UDP/DTLS to bridge port 2100, HueStream v2 protocol

> `backend/main.py` is the **legacy Python/FastAPI backend**. It is kept for reference but is NOT used. Do not modify or run it.

## Project structure

```
SpectraControl/
├── backend/
│   ├── main.go            ← entire Go backend (HTTP routes, WS, color conversion)
│   ├── entertainment.go   ← Hue Entertainment API (DTLS streaming, CLIP v2)
│   ├── go.mod / go.sum
│   ├── spectractl         ← compiled binary (gitignored)
│   └── main.py            ← legacy Python backend (unused)
├── frontend/
│   └── index.html         ← entire UI (vanilla JS, no bundler)
├── src-tauri/
│   ├── src/main.rs        ← spawns backend process in release, handles shutdown
│   ├── tauri.conf.json
│   └── Cargo.toml
└── run.sh                 ← standalone run script (compiles + starts Go backend)
```

## Config file

Stored at `~/.config/spectracontrol/hue_config.json`:

```json
{
  "ip": "192.168.x.x",
  "api_key": "<hue v1 username>",
  "client_key": "<PSK hex for DTLS, 32 bytes>"
}
```

`client_key` is only populated if the bridge was paired with `generateclientkey: true`. Without it, Entertainment API (DTLS) is unavailable. If it's missing, re-pair by pressing the bridge button and calling `POST /api/pair`.

## Architecture

### Backend (`backend/main.go`)

Key sections:
- `loadPersistedConfig()` / `saveConfig()` — reads/writes `~/.config/spectracontrol/hue_config.json`; migrates from legacy `backend/.hue_config` on first run
- `hueGet()` / `huePut()` — thin HTTP wrappers to Hue v1 API
- `rgbToXY()` — sRGB → CIE XY via Philips color matrix
- `handleWsColor()` — WebSocket `/ws/color`; receives `{lights:[{id,r,g,b}]}` frames from frontend, routes to DTLS or HTTP PUT
- `sendColorUpdate()` — calls `pushToEntertainment()` first; falls back to HTTP PUT per light if entertainment is inactive
- Routes: `/api/config`, `/api/discover`, `/api/pair`, `/api/lights`, `/api/groups`, `/api/sync` (compat stub), `/api/entertainment`, `/ws/color`

### Entertainment API (`backend/entertainment.go`)

Implements HueStream v2 over UDP/DTLS to bridge port 2100. Key parts:
- `startEntertainmentStreaming(configID)` — activates streaming mode on bridge via CLIP v2 PUT, opens DTLS PSK connection, starts `entertainmentSender` goroutine
- `stopEntertainmentStreaming()` — closes DTLS conn, resets all state, deactivates streaming on bridge
- `entertainmentSender(colorCh)` — goroutine running at 20 fps; on write error **calls `stopEntertainmentStreaming()` before returning** to prevent stale `ent.conn` state
- `pushToEntertainment(lights)` — maps v1 light IDs → DTLS channel IDs, sends frame; returns `false` if inactive so caller falls back to HTTP PUT
- `buildHueStreamPacket()` — builds 16-byte header + 7 bytes/channel HueStream v2 packet

**Critical:** if `ent.conn` is non-nil but the DTLS connection is dead, `pushToEntertainment` returns `true` and the colors are lost (no HTTP fallback). The `entertainmentSender` goroutine now auto-cleans on write error to prevent this. If the backend is restarted while entertainment was active, the bridge may still be in streaming mode — call `POST /api/entertainment/stop` to reset.

### Frontend (`frontend/index.html`)

All UI in one file. Key sections:
- `const API` — detects dev (`:5173`) vs prod (same origin)
- `lights[]` / `groups[]` — module-level state; groups filtered to exclude `type === "Entertainment"`
- **Sync flow:** `toggleSync()` → `getDisplayMedia()` → `_startSyncWithStream()` → WebSocket `/ws/color` → `scheduleSample()` → canvas sampling → `sampleRegions()` → send `{lights:[{id,r,g,b}]}`
- `sampleRegions()` — divides canvas into a grid matching `syncLightIds`; samples ~8×8 px per region; applies `boostSaturation()` (×4 HSV saturation push)
- `syncLightIds` — ordered list of light IDs for sync regions; user-configurable via the order panel; persisted to `localStorage`
- Entertainment section: `loadEntertainmentConfigs()` / `startEntertainment()` / `stopEntertainment()` — must activate BEFORE starting sync for DTLS path

### Tauri shell (`src-tauri/`)

- **Dev** (`cargo tauri dev`): `beforeDevCommand` in `tauri.conf.json` starts the Go backend; webview loads from `http://localhost:8000`
- **Release**: `main.rs` needs updating — currently spawns `uv run uvicorn` (Python, legacy). Should spawn `./backend/spectractl` instead.

## Screen sync flow

```
getDisplayMedia() → MediaStream
  → <video> element (hidden, off-screen)
  → <canvas> 64×36 px drawImage every N ms
  → sampleRegions() per light → {id, r, g, b}
  → WebSocket /ws/color  {lights:[...], bri:200, transitiontime:N}
  → Go backend
      ├── if Entertainment active → pushToEntertainment() → DTLS HueStream packet → bridge:2100
      └── else → huePut /lights/{id}/state {xy:[x,y], bri, transitiontime} → bridge:80
```

## Color pipeline

```
RGB (0–255)
  → boostSaturation() in frontend (×4 HSV saturation)
  → WebSocket to Go backend
  → rgbToXY(): sRGB gamma → Philips Hue matrix → CIE XY
  → PUT /api/{key}/lights/{id}/state {"xy": [x, y]}
  OR
  → HueStream packet (channel_id + R/G/B as uint16 0–65535)
```

## Key constraints

- **KDE Wayland** — X11 screen capture tools don't work. `getDisplayMedia()` is the only viable capture path.
- **Go backend** — use `go build` from `backend/`; do NOT use `uv`/Python for the backend.
- **`client_key` required for Entertainment** — must pair with `generateclientkey: true`. Check `~/.config/spectracontrol/hue_config.json`.
- **Rust/Tauri** — install via `rustup` on Bazzite (not rpm-ostree); `tauri-cli` via `cargo install tauri-cli`
- Icons not yet generated — run `cargo tauri icon path/to/source.png` to create `src-tauri/icons/`
- `src-tauri/main.rs` still launches the Python backend in release mode — needs updating to launch `spectractl`
