# Contributing to SpectraControl

Thanks for your interest. This is a small project — the easiest way to help is to file detailed bug reports or send focused PRs. The notes below should cover everything you need to get a dev build running and to know what the maintainers expect from a patch.

## Quick start

```bash
git clone https://github.com/sebascode/SpectraControl_App.git
cd SpectraControl_App

# 1) Go backend (you can run it standalone for browser mode)
./run.sh
# Browse http://localhost:8000 in Firefox.

# 2) Full Tauri shell (recommended for any change that touches the desktop side)
cargo tauri dev
# This runs the Go backend automatically (beforeDevCommand in tauri.conf.json)
# and opens the desktop window.
```

System requirements are documented in [`README.md`](README.md#requirements). Bazzite / Fedora 40+ users need `NO_STRIP=true` when building release bundles; `build.sh` already sets it.

## Repo layout

```
backend/        Go 1.21+ — HTTP, WebSocket, DTLS streaming to the Hue bridge
frontend/       Vanilla HTML/CSS/JS in a single index.html (no bundler, no framework)
  i18n/         One JSON file per locale
src-tauri/      Tauri 2 shell — spawns the Go backend, owns the tray, autostart,
                screen capture pipeline (ashpd + gst-launch)
docs/           Hue API v2 reference (offline copy), architecture notes
.github/        CI workflows, issue/PR templates
```

Detailed component responsibilities live in [`docs/architecture.md`](docs/architecture.md).

## Working on different layers

### Logging & verbosity

Both layers honor a single env var: **`SPECTRA_LOG_LEVEL`** (`debug | info | warn | error`, default `info`). Set it before launching the AppImage and it propagates to the Tauri shell *and* the spawned Go backend (`spectractl` inherits the parent env).

```bash
SPECTRA_LOG_LEVEL=debug ./SpectraControl_0.x.y.AppImage
```

- Go: leveled `slog` via `backend/logger.go`. Use the helpers (`logDebugf`, `logInfof`, `logWarnf`, `logErrorf`, `logFatalf`). Do **not** add raw `log.Printf` / `fmt.Println` for diagnostics — they bypass the level filter.
- Rust: `log` + `env_logger`. Use `log::debug!`, `info!`, `warn!`, `error!`. Same rule: no `eprintln!` for routine logs. Panic paths can still use `eprintln!`/`expect()` since the logger is not guaranteed during a panic.

The Go side also accepts a `-log-level` flag for `spectractl` standalone runs (`./run.sh` doesn't set one — env var is preferred for parity with the AppImage).

### Backend (Go)

- Stdlib + chi router + gorilla/websocket + pion/dtls. No ORM, no codegen.
- Add a route in `main.go`, keep the handler small, push protocol details (DTLS, color math) into the matching file.
- Run it standalone with `./run.sh` and curl the endpoints from a terminal while you iterate — much faster than restarting the full Tauri build.
- Tests live alongside their target: `go test ./...` from `backend/`. The suite is small on purpose — pure functions only (no network, no globals). New protocol code should land with a unit test for its byte layout or math, the same way `entertainment_test.go` covers `buildHueStreamPacket`.

### Frontend (HTML/JS)

- One file: `frontend/index.html`. Yes, on purpose: keeps the bundle small, lets the Go backend serve it as a static resource, and avoids a JS toolchain.
- New UI strings **must** go through the i18n helpers (`t("key")` / `tf("key", vars)`). Add the key to both `frontend/i18n/es.json` and `frontend/i18n/en.json` in the same PR.
- Native widgets (tray menu, future native dialogs) need a Rust-side label setter wired through `syncNativeStrings()` — these strings don't pass through `applyTranslations()` because they live outside the webview.

### Rust shell (`src-tauri/`)

- Plugins (updater, autostart, notification, dialog, tray-icon, process) are registered in `main.rs` and granted in `capabilities/default.json`. If you add a plugin or a new permission, update both.
- The frontend reaches Tauri APIs via `withGlobalTauri: true` (no JS bundler). For external plugins that don't expose a JS namespace, use `core.invoke('plugin:<name>|<cmd>', args)` — see `setAutostart` in `index.html` for the pattern.
- Exit cleanup goes through `RunEvent::Exit` (`main.rs`). Any new child process you spawn should be tracked in a `manage`d state and killed in `cleanup_children`.
- Tests: `cargo test` from `src-tauri/`. Today this covers `sanitize_host_env` (env-var stripping for spawned host processes). Add tests for any new pure helper you introduce.

## Commit conventions

We follow [Conventional Commits](https://www.conventionalcommits.org/) loosely:

```
feat(scope): short imperative summary

Optional body that explains *why* — the diff already shows *what*. Reference
issues and PRs by number. End with a Closes #N or Refs #N when relevant.
```

Common scopes: `backend`, `frontend`, `ui`, `tauri`, `bundle`, `ci`, `docs`, `updater`, `i18n`, `settings`. Keep the subject under ~72 chars.

## Pull request flow

1. Fork, create a topic branch from `master`.
2. Keep the PR focused — one logical change. If you're refactoring as you fix a bug, split them.
3. Make sure tests pass: `go test ./...` in `backend/` and `cargo test` in `src-tauri/`. Manual smoke testing is still expected for UI / capture changes — the suite covers protocol-level math (HueStream packet layout, scene interpolation, color lerp, brightness clamp, env sanitization) but not the runtime pipeline.
4. Open the PR against `master` with:
   - a short description of the user-facing change,
   - what you tested manually (which distro, which compositor, which Hue setup),
   - screenshots / a short clip for UI changes.

A maintainer will review and squash-merge once it's green.

## Release flow

Maintainers only. See the "Releasing" section of `README.md` for the full procedure. For signing-key handling specifically (including rotation), see [`docs/release-keys.md`](docs/release-keys.md). Short version:

1. Bump version in `src-tauri/Cargo.toml` and `src-tauri/tauri.conf.json` (Cargo.lock regenerates).
2. Commit `chore: bump version to X.Y.Z`.
3. Tag `vX.Y.Z` (or `vX.Y.Z-betaN` for prereleases).
4. Push tag → `.github/workflows/release.yml` builds and publishes the bundles + signed `latest.json`.

## Security policy

- **Vulnerabilities**: do **not** open a public issue. Use GitHub's [private security advisories](https://github.com/sebascode/SpectraControl_App/security/advisories/new) so we can coordinate a fix before disclosure.
- **External PRs are not processed with autonomous AI agents.** Maintainers may use AI as a research assistant while reviewing, but every line that lands in `master` from an outside contribution is read and approved by a human. If you suspect a PR is adversarial (prompt injection in description, hidden control characters, unusual binary blobs), flag it in the PR thread.
- **Supply chain**: third-party GitHub Actions are pinned to commit SHAs. If you bump one in a PR, include the SHA → tag mapping in the commit body so the reviewer can verify against the action's release notes.
- **Dependencies**: `cargo audit` and `govulncheck` run on every PR that touches lockfiles. They block merge on findings; fix the dep or document why the advisory does not apply.

## Reporting bugs

Use the bug issue template. The most useful reports include:

- The exact version (Settings → Acerca de).
- Distro + compositor + GPU.
- Steps to reproduce, and what you expected vs. what happened.
- Output from running the AppImage from a terminal so the stderr/Go logs are captured.

## Code of conduct

Be kind. We're here to make Hue work on Linux, not to win arguments.
