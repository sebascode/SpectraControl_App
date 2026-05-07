#!/usr/bin/env bash
# Standalone run without Tauri (browser at http://localhost:8000)
# For Tauri dev mode: cargo tauri dev   (from project root)

set -e
cd "$(dirname "${BASH_SOURCE[0]}")"

echo "→ Starting SpectraControl..."
echo "→ Open http://localhost:8000 in your browser"
echo "→ Ctrl+C to stop"
echo ""
uv run python -m uvicorn --app-dir backend main:app --host 0.0.0.0 --port 8000 --reload
