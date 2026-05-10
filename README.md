# SpectraControl

Linux desktop app for controlling Philips Hue lights. Fills the gap left by the official Hue Sync app, which has no Linux version. Built for Bazzite (Fedora immutable) + KDE Plasma on Wayland.

## Features

- Turn lights on/off, set brightness and color per light or room
- **Screen Sync** — captures screen via `getDisplayMedia()`, samples regions per light, streams colors at up to ~20 fps
- **Entertainment Mode** — uses the Hue Entertainment API (DTLS/HueStream v2) for low-latency sync that shows "synchronized" in the Hue app; falls back to HTTP PUT when inactive

## Requirements

- Go 1.21+
- A Philips Hue Bridge on the local network
- KDE Plasma on Wayland (or any browser that supports `getDisplayMedia`)
- Rust + tauri-cli (only for the Tauri desktop shell)

## Running

```bash
# Browser mode (http://localhost:8000)
./run.sh

# Tauri desktop window
cargo tauri dev
```

`run.sh` compiles the Go backend if the binary is missing or outdated, then starts it on port 8000.

## First-time setup

1. Open http://localhost:8000
2. Click the settings icon — enter your bridge IP (or use auto-discover)
3. Press the physical button on the bridge, then click **Vincular**
4. Save — lights and rooms will load automatically

The config is saved to `~/.config/spectracontrol/hue_config.json`. Pairing with the bridge button generates a `client_key` (PSK) required for Entertainment/DTLS mode. If the key is missing, re-pair.

## Screen Sync

1. Select rooms in the **Screen Sync** panel
2. Arrange light positions in the order grid (saved per light set in localStorage)
3. *(Optional)* In the **Modo Entertainment** section, select an area and click **Activar** — enables DTLS streaming for lower latency
4. Click **▶ Iniciar Sync** — a screen-capture dialog appears; select your display

The color dot and hex value in the sync panel show the live average color being sent.

## Stack

| Layer | Tech |
|---|---|
| Backend | Go · chi · gorilla/websocket · pion/dtls |
| Frontend | Vanilla HTML/CSS/JS (single file, no bundler) |
| Desktop shell | Tauri 2 (Rust) |
| Hue API v1 | Local REST (lights, groups) |
| Hue API v2 (CLIP) | HTTPS (entertainment configs) |
| Entertainment streaming | UDP/DTLS · HueStream v2 · bridge port 2100 |

## Project structure

```
SpectraControl/
├── backend/
│   ├── main.go            ← HTTP routes, WebSocket, color conversion
│   ├── entertainment.go   ← DTLS streaming, CLIP v2, HueStream packet builder
│   ├── go.mod / go.sum
│   └── spectractl         ← compiled binary (gitignored)
├── frontend/
│   └── index.html         ← entire UI
├── src-tauri/             ← Tauri desktop shell
└── run.sh                 ← build + run script
```

## Known limitations

- `src-tauri/main.rs` still launches the legacy Python backend in release builds — needs updating to spawn `spectractl`
- Icons not generated yet — run `cargo tauri icon <source.png>` to create them
