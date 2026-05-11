# Architecture overview

SpectraControl is three processes that cooperate:

```mermaid
flowchart TD
    T["<b>Tauri shell (Rust)</b><br/>spawns spectractl<br/>tray · autostart · updater · notifications<br/>screen capture: ashpd → gst-launch"]
    W["<b>Webview</b><br/>index.html"]
    B["<b>spectractl (Go)</b><br/>127.0.0.1:8000"]
    H["<b>Hue bridge v2</b>"]

    T -- "embeds / IPC" --> W
    T -. "spawns" .-> B
    W <-- "HTTP / WebSocket" --> B
    B <== "DTLS UDP :2100<br/>(HueStream v2)" ==> H
    B -- "HTTPS (CLIP v2)<br/>entertainment configs, app id" --> H
```

## Process roles

### `spectractl` — Go backend (`backend/`)
- HTTP server on `127.0.0.1:8000` (chi router).
- REST endpoints for lights, groups, scenes, profile, entertainment configs.
- WebSocket endpoint that the frontend uses to push per-frame color updates during Screen Sync.
- DTLS streamer (`pion/dtls`) that talks the HueStream v2 protocol on UDP 2100. Interpolates incoming color updates and sends ~30 fps to the bridge.
- Loads built-in scene presets via `go:embed`, and merges user scenes from `~/.config/spectracontrol/scenes/`.
- Persists bridge credentials to `~/.config/spectracontrol/hue_config.json`.

### Webview — frontend (`frontend/index.html` + `i18n/`)
- Single HTML file, no bundler, no framework. Strings flow through a runtime i18n dict loaded from `frontend/i18n/<locale>.json`.
- Talks to `spectractl` over HTTP for reads/writes, WebSocket for the sync loop.
- In Tauri mode, calls `core.invoke('sample_screen_regions', ...)` once per sync tick to get sampled pixels from the screen capture pipeline. In browser mode, captures via `getDisplayMedia()` and samples in-page on a hidden `<canvas>`.
- Preferences (theme, locale, sync target zone, autostart-sync flag) live in `localStorage`.

### Tauri shell (`src-tauri/src/main.rs`)
- In release builds, spawns `spectractl` as a child process from the bundled `backend/` resource and kills it on `RunEvent::Exit`.
- Owns the screen capture pipeline (Linux):
  1. Negotiates a Screencast session with `xdg-desktop-portal` through `ashpd`.
  2. Spawns `gst-launch-1.0` with `pipewiresrc → videoscale → fdsink fd=1`, reads raw RGB frames from stdout into a `Mutex`'d buffer.
  3. Exposes `sample_screen_regions(positions)` so the frontend can pull averaged colors for each light without doing IPC per region.
- Hosts the system tray (`tray-icon` feature), autostart plugin, notification plugin, updater plugin, and process plugin.
- Window close behavior is governed by two `AtomicBool` states: `IsQuitting` and `QuitOnClose`. Default (`quit_on_close=false`) hides to tray on X.

## Boundaries

| Concern                      | Lives in           |
|------------------------------|--------------------|
| Hue protocol (REST + DTLS)   | `backend/`         |
| Frontend UI + state          | `frontend/`        |
| Screen capture (Wayland)     | `src-tauri/`       |
| Tray, autostart, updates     | `src-tauri/`       |
| Translations                 | `frontend/i18n/`   |
| Native widget strings        | `src-tauri/` + `frontend/index.html` (via `syncNativeStrings()`) |
| Bridge credentials persistence | `backend/` writes `~/.config/spectracontrol/hue_config.json` |
| User preferences persistence | `frontend/` via `localStorage` |

The webview never speaks DTLS directly — it always goes through the Go backend over WebSocket. The Tauri shell never touches the bridge — it only hosts the webview and provides Wayland-only capabilities (capture, tray, autostart) that the webview can't do on its own.

## Data flow during Screen Sync (Tauri mode)

```mermaid
sequenceDiagram
    autonumber
    actor User
    participant FE as Webview (index.html)
    participant R as Tauri (Rust)
    participant GO as spectractl (Go)
    participant HB as Hue bridge

    User->>FE: Tap entertainment zone card
    FE->>GO: POST /api/entertainment/activate?id=...
    GO->>HB: Open DTLS session (PSK = client_key)
    GO-->>GO: Start HueStream interpolator goroutine
    FE->>GO: WebSocket /ws/sync (open)

    loop every ~33 ms
        FE->>R: invoke("sample_screen_regions", positions)
        R-->>FE: [r,g,b] per light (from gst capture buffer)
        FE->>GO: { lights, bri, transitiontime }
    end

    loop every ~20 ms
        GO->>HB: HueStream packet (DTLS)
    end

    User->>FE: Tap to stop
    FE->>GO: WebSocket close
    GO->>HB: Tear down DTLS session
```

## Update flow

```mermaid
sequenceDiagram
    autonumber
    actor User
    participant FE as Webview
    participant R as Tauri (Rust)
    participant GH as GitHub Releases
    participant FS as Local filesystem
    participant GO as spectractl (Go)

    FE->>R: updater.check()
    R->>GH: GET releases/latest/download/latest.json
    GH-->>R: { version, signature, url }
    R-->>FE: pendingUpdate (if newer)
    FE->>User: Show update banner

    User->>FE: Click "Actualizar"
    FE->>R: pendingUpdate.downloadAndInstall()
    R->>GH: GET .AppImage + .sig
    R-->>R: Verify minisign signature
    R->>FS: Atomically replace $APPIMAGE

    FE->>FE: Show "Update ready — reopen the app"
    FE->>R: notification.send(updateInstalled)
    FE->>R: process.exit(0)
    R->>GO: kill (RunEvent::Exit → cleanup_children)
    Note over R,GO: Port :8000 is freed before the user reopens
```

The old `process.relaunch()` path is intentionally avoided: it raced the new instance against the old one for `:8000`. See [#25](https://github.com/sebascode/SpectraControl_App/issues/25).
