# Troubleshooting

This guide covers the most common issues when using SpectraControl with a Hue
Bridge. If your problem isn't listed here, open a GitHub issue with the logs
attached (see the last section).

> Spanish version: [troubleshooting.es.md](troubleshooting.es.md)

## How to capture logs

Most diagnostics start by enabling `debug` mode and reproducing the problem.

### Development mode (`cargo tauri dev`)

```bash
SPECTRA_LOG_LEVEL=debug cargo tauri dev
```

Backend (Go) logs (`[ent]`, `[sse]` prefixes, etc.) and the Rust shell logs
appear interleaved in the same terminal.

### Installed app (native binary)

```bash
SPECTRA_LOG_LEVEL=debug spectra-control
```

Replace `spectra-control` with the actual executable path if it's not on your
`PATH`.

### Flatpak

```bash
flatpak run --env=SPECTRA_LOG_LEVEL=debug casa.scode.SpectraControl
```

---

## Common issues

### Lights stay on after closing the app while syncing

**Symptom:** you close the window while syncing and the lights remain lit in
the last frame's color.

**Cause:** the "Turn lights off when closing (while syncing)" toggle is
disabled. It's opt-in: the default behavior leaves the last state untouched.

**Fix:** Settings → Appearance → enable "Turn lights off when closing (while
syncing)".

---

### "DTLS handshake context deadline exceeded" error when starting sync

**Symptom:** pressing Sync doesn't start the stream and an error mentions
"DTLS handshake" or "deadline exceeded". The backend log shows
`[ent] DTLS dial fallido a <ip>:2100`.

**Cause:** the Hue Bridge only allows **one** active Entertainment stream at a
time, and another client owns it: typically a Hue Sync Box, the desktop Hue
Sync app, a game console with Sync, or another SpectraControl instance.

**Fix:**

1. Make sure the "Break other sync sessions on start" toggle is enabled
   (Settings → Screen Sync). It's on by default. This automatically stops any
   other session before starting ours.
2. If the toggle was already on and the problem persists, the Hue Sync Box
   hardware may be auto-reconnecting (firmware with aggressive auto-resume).
   Pause it from the Hue Sync app, switch it to HDMI passthrough, or unplug
   it temporarily.
3. You can also try the manual "Stop other apps' entertainment syncs now"
   button in Settings → Screen Sync.

To confirm the cause: run with `SPECTRA_LOG_LEVEL=debug` and look for
`[ent] config encontrada — id=... name=... status="active" ...` lines. If you
see one with `status="active"` that isn't yours, that's the conflict.

---

### Lights flicker or don't respond to manual commands

**Symptom:** moving brightness/color sliders in the UI doesn't update the
lights, or they update with a several-second delay.

**Cause:** an active Entertainment stream (yours, a Sync Box, etc.) is
overwriting individual lights via DTLS at ~30 fps. Manual HTTP commands apply
but get instantly overridden.

**Fix:** stop any sync (Stop button), and if the issue persists, use "Stop
other apps' entertainment syncs now" to close foreign streams.

---

### Bridge not detected / "no bridge found"

**Symptom:** on first launch (or after a router reboot) the app doesn't find
the bridge.

**Possible causes:**

- The bridge changed IPs (DHCP rotated the address). SpectraControl caches the
  IP from the initial pairing.
- The OS firewall is blocking mDNS (UDP port 5353).
- Bridge and PC are on different networks (common with guest Wi-Fi or VLANs).

**Fix:**

1. Verify the bridge has a steady blue network LED.
2. Ping the bridge IP from a terminal (`ping <ip>`).
3. If it changed, delete the config file and re-pair:
   - Linux: `~/.config/SpectraControl/config.json`
   - Flatpak: `~/.var/app/casa.scode.SpectraControl/config/SpectraControl/config.json`
4. Make sure your Wi-Fi/Ethernet is on the same subnet as the bridge.

---

### Audio sync captures nothing

**Symptom:** you enable Audio Sync but the lights don't react to system sound.

**Cause:** audio sync requires native system audio capture. This only works in
the **desktop app**, not the browser. Internally it spawns `gst-launch-1.0`
(GStreamer) to open a PulseAudio / PipeWire loopback.

**Fix:**

1. Confirm you're using the installed app, not `localhost:8000` in a browser.
2. On native Linux, make sure `gstreamer1.0-pulseaudio` is installed.
3. On Flatpak it's already bundled.
4. Verify your system is actually playing audio (active speakers/headphones).
   With everything muted there's nothing to react to.

---

### System notifications don't show up

**Symptom:** the app shows "sync started" in an in-app toast but no native
notification appears.

**Cause:** GNOME silences notifications from non-installed apps, or
`libnotify-bin` (`notify-send`) is missing on the host. SpectraControl shells
out to `notify-send` directly because the Tauri plugin panics under GNOME's
runtime and direct `notify-rust` is silenced.

**Fix:**

- Native Linux: `sudo dnf install libnotify` (Fedora) or
  `sudo apt install libnotify-bin` (Debian/Ubuntu).
- Flatpak: notifications go through the `xdg-desktop-portal` and should work
  as long as the portal is installed and running.

---

### "Start with system" toggle does nothing

**Symptom:** you enable autostart in Appearance but the app doesn't open at
login.

**Cause:** known issue — see
[#24](https://github.com/sebascode/SpectraControl_App/issues/24).

**Workaround:** manually create `~/.config/autostart/SpectraControl.desktop`
pointing at the binary.

---

## Reporting a bug

If your problem isn't listed here or the workaround doesn't help, open an
issue on
[GitHub](https://github.com/sebascode/SpectraControl_App/issues) and include:

1. **SpectraControl version** (visible in Settings → About, or
   `cat /var/lib/flatpak/app/casa.scode.SpectraControl/current/active/files/manifest.json`
   on Flatpak).
2. **Operating system** and desktop environment (e.g. Fedora 44, GNOME 47
   Wayland).
3. **Install mode** (Flatpak / binary / development).
4. **Hue Bridge model and firmware** (visible in the official Hue app,
   Settings → Hue Bridges → your bridge).
5. **Debug-mode logs** captured while reproducing the issue (see "How to
   capture logs" above). Paste only the relevant lines if the full log is
   long.
6. **Steps to reproduce** and what you expected to happen.

If you're going to post on a third-party forum (developers.meethue.com,
Reddit, etc.), strip bridge tokens before pasting logs: the `clientkey` from
`config.json` shouldn't appear in logs but it's worth double-checking.
