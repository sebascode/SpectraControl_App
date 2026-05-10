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

## Installation

Pre-built bundles for each tagged release live in [GitHub Releases](https://github.com/sebascode/SpectraControl_App/releases) (AppImage, .deb, .rpm). Pick whichever matches your distro.

### AppImage (recommended — distro-agnostic, supports auto-update)

```bash
chmod +x SpectraControl_*_amd64.AppImage
./SpectraControl_*_amd64.AppImage
```

> Needs `libfuse2` to mount itself. On modern Fedora that is `fuse-libs`. As a fallback, run with `--appimage-extract-and-run`.

### .deb / .rpm

```bash
sudo dpkg -i SpectraControl_*_amd64.deb        # Debian/Ubuntu
sudo rpm -i SpectraControl-*.x86_64.rpm        # Fedora/RHEL
```

The bundles embed the Go backend (`spectractl`) and the frontend, so no separate install is needed. Config is written to `~/.config/spectracontrol/hue_config.json`. User scenes are picked up from `~/.config/spectracontrol/scenes/*.json` (override built-ins by `id`).

### Building locally

```bash
./build.sh appimage,deb,rpm
# → src-tauri/target/release/bundle/{appimage,deb,rpm}/
```

`build.sh` runs `cargo tauri build` with `NO_STRIP=true`, which is required on Fedora 40+ — Fedora's shared libraries use packed relocations (`.relr.dyn` / `DT_RELR`) that the `strip` bundled in `linuxdeploy` does not recognise.

## In-app updates

When you run the AppImage, on launch the app checks GitHub Releases for a newer tag. If one is available, a banner appears at the top with **Actualizar**. Clicking it:

1. Downloads the new `.AppImage` to a tmp file in the same directory as the running one.
2. Atomically replaces it (`rename(2)`) — your current process keeps running because Linux holds the old inode open.
3. Asks you to restart. Close the app and relaunch the same `.AppImage` path; the new version takes over.

`.deb` / `.rpm` users get a notification but no in-place install — use `dnf upgrade` / `apt upgrade` against the file you re-download.

Endpoints (handy for scripting):

| Endpoint | Description |
|---|---|
| `GET  /api/version`         | `{version, channel, os, arch}` for the running build |
| `GET  /api/update/check`    | queries GitHub `releases/latest`, returns `has_update`, `latest`, `release_url`, `asset_url`, `notes`, `can_install` |
| `POST /api/update/install`  | downloads the `.AppImage` asset and replaces `$APPIMAGE`. Returns `{ok, installed, restart_required}` or `{ok:false, reason:"not_appimage"}` for non-AppImage runs |

## Releasing

CI builds and publishes on tag push. To cut a release:

```bash
git tag v0.2.0
git push origin v0.2.0
```

`.github/workflows/release.yml` then:

1. Patches `tauri.conf.json` with the version.
2. Builds AppImage + deb + rpm on `ubuntu-22.04`, injecting the version into the Go binary via `-ldflags "-X main.Version=$SPECTRA_VERSION"`.
3. Uploads the bundles + a `SHA256SUMS` file to the GitHub Release for that tag.

You can dry-run the build (no publish) from the **Actions** tab → *Release* → *Run workflow*.

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
│   ├── scenes.go          ← dynamic scene runner (JSON-loaded presets)
│   ├── scenes/*.json      ← built-in scene presets (embedded via go:embed)
│   ├── brightness.go      ← global brightness state, applied to sync + scenes
│   └── go.mod / go.sum
├── frontend/
│   └── index.html         ← entire UI
├── src-tauri/
│   ├── src/main.rs        ← spawns Go backend, runs ashpd + gst capture
│   ├── Cargo.toml
│   └── tauri.conf.json
├── docs/                  ← Hue API v2 reference (offline copy of docs)
├── run.sh                 ← build + run the Go backend standalone
└── build.sh               ← build the AppImage (release)
```

## Troubleshooting

- **"DTLS conectado" but lights stay at neutral white** — the keepalive frame is rendering. The screen capture isn't producing frames. In Tauri mode, check the `cargo tauri dev` terminal for `[capture]` logs. In browser mode, only Firefox works on Wayland.
- **App says "sincronizado" but lights don't change color** — make sure all lights in the entertainment area are reachable in the Hue app. The streaming protocol does not turn lights on; SpectraControl turns them on before activating the stream, but powered-off bulbs (zigbee unreachable) won't respond.
- **`client_key` missing** — re-pair from the setup screen. Required for Entertainment v2 (DTLS PSK).
- **Build errors on Bazzite** referring to `libgbm` or `mesa-libgbm-devel` — these come from `xcap` / `libwayshot-xcap`. SpectraControl doesn't use them; if a `cargo build` complains, re-check `Cargo.toml` deps match the current tree.

## Known limitations

- Audio sync (#5) and per-light rectangular sections (#3) are not yet implemented.
- The AppImage build needs `NO_STRIP=true` on Fedora 40+ (handled automatically by `build.sh`); other distros may not need it.

## License

MIT — see [LICENSE](LICENSE).
