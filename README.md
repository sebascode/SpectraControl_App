# SpectraControl

Linux desktop app for Philips Hue lights, with screen-to-light sync over the official Hue Entertainment API v2. Fills the gap left by the Hue Sync app, which has no Linux version.

Tested on Bazzite (Fedora 44 immutable) + GNOME Wayland with NVIDIA. Should also work on KDE Plasma Wayland and any compositor that exposes `xdg-desktop-portal` Screencast.

## Features

- Per-light and per-room control (on/off, brightness, color)
- **Screen Sync over Entertainment v2** — DTLS streaming directly to the bridge on UDP 2100 (HueStream v2 protocol). Shows as "synchronized" in the official Hue app.
- **Native Wayland screen capture** — uses `xdg-desktop-portal` Screencast + GStreamer (`pipewiresrc`) under Tauri. Bypasses the WebKitGTK `getDisplayMedia` path which is broken on KDE/GNOME Wayland.
- **Browser fallback** — works in Firefox at `http://localhost:8000` using `getDisplayMedia()` for environments without Tauri.

## Requirements

- Linux with Wayland and `xdg-desktop-portal` (KDE or GNOME)
- Philips Hue Bridge v2 paired with `generateclientkey: true`
- Go 1.21+
- Rust + `cargo install tauri-cli` (for the desktop app)
- GStreamer with `pipewiresrc` (`gstreamer1-plugin-pipewire`)
- Build-only (Bazzite/Fedora 44):
  ```bash
  sudo rpm-ostree install pipewire-devel clang-devel glibc-devel \
    wayland-devel libxcb-devel webkit2gtk4.1-devel
  sudo rpm-ostree apply-live --allow-replacement   # or reboot
  ```

## Running

```bash
# Tauri desktop window (recommended)
cargo tauri dev

# Browser mode — Firefox at http://localhost:8000
./run.sh
```

`run.sh` builds and starts the Go backend on `:8000`. `cargo tauri dev` does the same automatically via `beforeDevCommand`.

> Chromium-based browsers and the WebKitGTK webview don't deliver frames from `getDisplayMedia()` on Wayland. Use Tauri (native portal capture) or Firefox (which uses pipewire correctly).

## First-time setup

1. Launch the app.
2. Settings → enter your bridge IP, or click **Descubrir** to auto-detect.
3. Press the physical button on the bridge, then click **Vincular**. The bridge returns an API key + a 32-byte `client_key` (PSK for DTLS).
4. Lights and rooms load automatically. Config is persisted to `~/.config/spectracontrol/hue_config.json`.

If the `client_key` is missing (older pairings without `generateclientkey`), re-pair from the setup screen — Entertainment streaming requires it.

## Screen Sync

1. Open the **Inicio** tab. Each entertainment area defined in the Hue app appears as a card.
2. Click a card. Approve the screen-capture portal prompt the first time (Tauri mode) — after that, no prompt.
3. The card turns green and shows "Sincronizando". Lights mirror the colors of your screen.
4. Click again to stop.

The capture pipeline downscales the screen to 320×180 RGB in GStreamer, so sampling stays cheap regardless of display resolution. The bridge receives ~30 fps of color updates per channel via DTLS.

## Stack

| Layer | Tech |
|---|---|
| Backend | Go · chi · gorilla/websocket · pion/dtls |
| Frontend | Vanilla HTML/CSS/JS (single file, no bundler) |
| Desktop shell | Tauri 2 (Rust) |
| Screen capture (Tauri) | `ashpd` (xdg-desktop-portal Screencast) + `gst-launch-1.0` (`pipewiresrc → videoscale → fdsink`) |
| Screen capture (browser) | `getDisplayMedia()` |
| Hue API v1 | Local REST (lights, groups, config) |
| Hue API v2 (CLIP) | HTTPS (entertainment configs, application id) |
| Entertainment streaming | UDP/DTLS · HueStream v2 · bridge port 2100 |

## Project structure

```
SpectraControl/
├── backend/
│   ├── main.go            ← HTTP routes, WebSocket, color conversion
│   ├── entertainment.go   ← DTLS streaming, CLIP v2, HueStream packet builder
│   └── go.mod / go.sum
├── frontend/
│   └── index.html         ← entire UI
├── src-tauri/
│   ├── src/main.rs        ← spawns Go backend, runs ashpd + gst capture
│   ├── Cargo.toml
│   └── tauri.conf.json
├── docs/                  ← Hue API v2 reference (offline copy of docs)
├── CLAUDE.md              ← architecture + dev internals
└── run.sh                 ← build + run the Go backend standalone
```

## Troubleshooting

- **"DTLS conectado" but lights stay at neutral white** — the keepalive frame is rendering. The screen capture isn't producing frames. In Tauri mode, check the `cargo tauri dev` terminal for `[capture]` logs. In browser mode, only Firefox works on Wayland.
- **App says "sincronizado" but lights don't change color** — make sure all lights in the entertainment area are reachable in the Hue app. The streaming protocol does not turn lights on; SpectraControl turns them on before activating the stream, but powered-off bulbs (zigbee unreachable) won't respond.
- **`client_key` missing** — re-pair from the setup screen. Required for Entertainment v2 (DTLS PSK).
- **Build errors on Bazzite** referring to `libgbm` or `mesa-libgbm-devel` — these come from `xcap` / `libwayshot-xcap`. SpectraControl doesn't use them; if a `cargo build` complains, re-check `Cargo.toml` deps match the current tree.

## Known limitations

- `src-tauri/main.rs` still references launching a legacy Python backend in release builds — pending update to spawn `spectractl` directly.
- App icons not generated. Run `cargo tauri icon <source.png>` to create them.
- Light positions inside an entertainment area are not yet honored — sampling currently uses a generic grid based on light count, not the `{x, y, z}` positions defined in the Hue app.
