# Flatpak packaging

Manifest, AppStream metainfo and desktop entry for shipping SpectraControl on
[Flathub](https://flathub.org).

## Files

| File | Role |
| --- | --- |
| `casa.scode.SpectraControl.yml` | flatpak-builder manifest |
| `casa.scode.SpectraControl.metainfo.xml` | AppStream component (store listing + releases) |
| `casa.scode.SpectraControl.desktop` | XDG desktop entry installed into `/app/share/applications/` |
| `generate-sources.sh` | Produces the offline build inputs (`cargo-sources.json`, `backend-vendor.tar.gz`) |

## Local build & validation

The Flathub build runs offline, so every dependency must be materialised
beforehand by the generator script.

```bash
# 1. From a network-enabled checkout, generate offline inputs.
packaging/flatpak/generate-sources.sh

# 2. Install build tooling (one-time).
flatpak install -y flathub \
    org.gnome.Platform//47 \
    org.gnome.Sdk//47 \
    org.freedesktop.Sdk.Extension.rust-stable//24.08 \
    org.freedesktop.Sdk.Extension.golang//24.08
sudo dnf install -y flatpak-builder appstream desktop-file-utils  # Fedora
# (Debian/Ubuntu: apt install flatpak-builder appstream desktop-file-utils)

# 3. Build into a local repo and install the result for testing.
flatpak-builder --force-clean --install --user build-dir \
    packaging/flatpak/casa.scode.SpectraControl.yml

# 4. Run the sandboxed build.
flatpak run casa.scode.SpectraControl
```

### Linters

```bash
# AppStream component
appstreamcli validate packaging/flatpak/casa.scode.SpectraControl.metainfo.xml

# Desktop entry
desktop-file-validate packaging/flatpak/casa.scode.SpectraControl.desktop

# Manifest + finish-args (Flathub policy)
flatpak run --command=flatpak-builder-lint org.flatpak.Builder \
    manifest packaging/flatpak/casa.scode.SpectraControl.yml
flatpak run --command=flatpak-builder-lint org.flatpak.Builder \
    appstream packaging/flatpak/casa.scode.SpectraControl.metainfo.xml
```

## Flathub submission

The Flathub submission lives in a separate repo
(`flathub/casa.scode.SpectraControl`) created from
[flathub/flathub#new-pr-template](https://github.com/flathub/flathub).

The release workflow (`.github/workflows/release.yml`) publishes the offline
build inputs as release assets, so the Flathub submission manifest can
reference them by URL+sha256 without needing to regenerate them.

The manifest in the Flathub repo is the same as the one here with the
sources block swapped:

```yaml
sources:
  - type: git
    url: https://github.com/sebascode/SpectraControl_App.git
    tag: v0.2.15
    commit: <full SHA of the tag>

  - type: file
    url: https://github.com/sebascode/SpectraControl_App/releases/download/v0.2.15/cargo-sources.json
    sha256: <from SHA256SUMS in the release>
    dest-filename: cargo-sources.json

  - type: archive
    url: https://github.com/sebascode/SpectraControl_App/releases/download/v0.2.15/backend-vendor.tar.gz
    sha256: <from SHA256SUMS in the release>
    dest: backend/vendor
```

The `SHA256SUMS` file is also in the release; copy the two relevant lines
into the manifest.

## Sandbox notes

The Rust side already detects the sandbox at runtime (`FLATPAK_ID` env var,
`src-tauri/src/runtime.rs`) and adjusts behaviour:

- **Updater**: disabled — Flathub owns the update cycle.
- **Autostart**: routed through `org.freedesktop.portal.Background`
  (`src-tauri/src/autostart.rs`).
- **Notifications**: routed through `org.freedesktop.portal.Notification`
  (`src-tauri/src/notifications.rs`).

That means the manifest does not need `--filesystem=xdg-config/autostart`
nor `--talk-name=org.freedesktop.Notifications` — the portals are
auto-talkable.
