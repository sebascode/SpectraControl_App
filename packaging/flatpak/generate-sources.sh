#!/usr/bin/env bash
# Materialise offline build inputs for the Flatpak manifest:
#   - cargo-sources.json  (every crate listed in src-tauri/Cargo.lock)
#   - backend-vendor.tar.gz  (Go module cache as a vendor tree)
#
# Run this from the repo root *with network access* before invoking
# flatpak-builder. Both outputs land in packaging/flatpak/ alongside the
# manifest. They are not committed to this repo; the release workflow
# regenerates them per tag and the Flathub submission repo ships them.
#
# Requires: python3, git, go, curl.

set -euo pipefail

REPO_ROOT="$(git rev-parse --show-toplevel)"
OUT_DIR="$REPO_ROOT/packaging/flatpak"
GENERATOR_URL="https://raw.githubusercontent.com/flatpak/flatpak-builder-tools/master/cargo/flatpak-cargo-generator.py"
GENERATOR="$OUT_DIR/.flatpak-cargo-generator.py"

echo "→ Fetching flatpak-cargo-generator.py"
curl -fsSL "$GENERATOR_URL" -o "$GENERATOR"

echo "→ Generating cargo-sources.json from src-tauri/Cargo.lock"
python3 "$GENERATOR" "$REPO_ROOT/src-tauri/Cargo.lock" -o "$OUT_DIR/cargo-sources.json"

echo "→ Vendoring Go modules from backend/go.mod"
(
  cd "$REPO_ROOT/backend"
  rm -rf vendor
  go mod vendor
  tar -czf "$OUT_DIR/backend-vendor.tar.gz" vendor
  rm -rf vendor
)

rm -f "$GENERATOR"

echo
echo "Generated:"
echo "  $OUT_DIR/cargo-sources.json"
echo "  $OUT_DIR/backend-vendor.tar.gz"
