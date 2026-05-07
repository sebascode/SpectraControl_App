# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## What this is

**SpectraControl** — Linux desktop app for controlling Philips Hue lights. Replaces the official Hue Sync app which doesn't exist for Linux. Targets Bazzite (Fedora immutable) + KDE Plasma on Wayland.

## Running the app

```bash
# Standalone (browser at http://localhost:8000)
./run.sh
# or explicitly:
uv run uvicorn --app-dir backend main:app --host 0.0.0.0 --port 8000 --reload

# Tauri dev mode (desktop window, auto-starts backend)
cargo tauri dev   # from project root, requires Rust + tauri-cli
```

## Stack

- **Backend:** Python 3.14 · FastAPI · uvicorn · httpx — `backend/main.py`
- **Frontend:** Single `frontend/index.html` — vanilla HTML/CSS/JS, no framework, no bundler
- **Desktop shell:** Tauri 2 (Rust) — `src-tauri/`
- **Package manager:** `uv` (not pip/poetry)
- **Hue API:** Philips Hue Bridge v2, local REST API v1 (`http://<ip>/api/<key>/...`)

## Project structure

```
SpectraControl/
├── backend/
│   ├── main.py        ← entire FastAPI backend
│   └── .hue_config    ← bridge IP + API key (gitignored)
├── frontend/
│   └── index.html     ← entire UI (vanilla JS, no bundler)
├── src-tauri/
│   ├── src/main.rs    ← spawns uvicorn in release, handles shutdown
│   ├── tauri.conf.json
│   └── Cargo.toml
├── pyproject.toml     ← uv dependencies (Python)
└── run.sh             ← standalone run script
```

## Architecture

### Backend (`backend/main.py`)

All logic in one file. Key sections:
- `_HERE` — absolute path to `backend/`, used for `.hue_config` and locating `frontend/`
- `hue_request()` — thin async httpx wrapper that proxies to the local bridge
- `sync_state` dict — mutable shared state for the screen sync background thread
- Routes grouped by: `/api/config`, `/api/discover`, `/api/pair`, `/api/lights`, `/api/groups`, `/api/sync`
- FastAPI serves `frontend/` as StaticFiles at `/`

Bridge config is persisted as JSON in `backend/.hue_config` and loaded at startup.

### Frontend (`frontend/index.html`)

All UI in one file. The `const API` detection at the top handles dev (`:5173`) vs prod (same origin as FastAPI). State is kept in `lights[]` and `groups[]` module-level arrays.

### Tauri shell (`src-tauri/`)

- **Dev** (`cargo tauri dev`): `beforeDevCommand` in `tauri.conf.json` starts uvicorn; webview loads from `http://localhost:8000`
- **Release**: `main.rs` spawns `uv run uvicorn --app-dir backend main:app` using `current_dir()`, waits 1.5s, then the webview opens; kills the process on window close

## Color pipeline

```
RGB (0-255)
  → sRGB gamma correction
  → Philips Hue color matrix
  → CIE XY  [rgb_to_xy() in backend/main.py]
  → PUT /api/{key}/lights/{id}/state {"xy": [x, y]}
```

## Screen sync — why `getDisplayMedia()`

All native capture fails on KDE Wayland: `mss`, `scrot`, `ImageMagick import`, `grim` (wlroots only), `pywayland` (won't build on Python 3.14). The browser Web API `getDisplayMedia()` is the only capture path KDE Wayland permits. The planned implementation:
- Frontend: `getDisplayMedia()` → `<canvas>` sampling → WebSocket `/ws/color`
- Backend: receives RGB per light, converts to XY, pushes to bridge

The current `sync_loop()` in `backend/main.py` is legacy (uses xrandr/mss, broken on Wayland). The WebSocket endpoint is not yet implemented.

## Key constraints

- **Python 3.14** — check compatibility before adding deps (`pywayland` won't build)
- **KDE Wayland** — X11-only tools don't work for screen capture
- **Rust/Tauri** — install via `rustup` on Bazzite (not rpm-ostree); `tauri-cli` via `cargo install tauri-cli`
- Icons not yet generated — run `cargo tauri icon path/to/source.png` to create `src-tauri/icons/`
