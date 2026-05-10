#!/usr/bin/env bash
# Build the SpectraControl AppImage (Linux x86_64).
#
# NO_STRIP=true skips linuxdeploy's strip step. Fedora 40+ ships .so files with
# packed relocations (.relr.dyn / DT_RELR) that the strip bundled in linuxdeploy
# does not recognise — without this flag the bundle fails on Bazzite/Fedora.

set -e
cd "$(dirname "${BASH_SOURCE[0]}")"

export PATH="/home/linuxbrew/.linuxbrew/bin:$PATH"

TARGETS="${1:-appimage}"
echo "→ Building targets: $TARGETS (NO_STRIP=true)"
echo ""

NO_STRIP=true cargo tauri build --bundles "$TARGETS"

echo ""
echo "→ Done. Artifacts under src-tauri/target/release/bundle/"
find src-tauri/target/release/bundle -maxdepth 3 -type f \
    \( -name "*.AppImage" -o -name "*.deb" -o -name "*.rpm" \) -exec ls -lh {} \;
