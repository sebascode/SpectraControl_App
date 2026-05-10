#!/usr/bin/env bash
# Standalone run without Tauri (browser at http://localhost:8000)
# For Tauri dev mode: cargo tauri dev   (from project root)

set -e
cd "$(dirname "${BASH_SOURCE[0]}")"

export PATH="/home/linuxbrew/.linuxbrew/bin:$PATH"

# Recompila si el binario no existe o main.go es más nuevo
if [ ! -f backend/spectractl ] || [ backend/main.go -nt backend/spectractl ]; then
  echo "→ Compilando backend Go..."
  (cd backend && go build -o spectractl .)
fi

echo "→ Starting SpectraControl..."
echo "→ Open http://localhost:8000 in your browser"
echo "→ Ctrl+C to stop"
echo ""
exec ./backend/spectractl -addr :8000
