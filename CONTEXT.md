# SpectraControl — Contexto para Claude Code

## Qué es
App de escritorio para controlar luces **Philips Hue** en Linux (Bazzite/KDE Wayland).
Reemplaza la app oficial Hue Sync que no existe para Linux.

## Stack actual
- **Backend:** Python 3.14 + FastAPI + uvicorn
- **Frontend:** HTML/CSS/JS puro (sin framework), servido por FastAPI como StaticFiles
- **Gestor de paquetes:** `uv`
- **Control de versiones:** git (rama `master`)

## Estructura del proyecto
```
Hue-Sync-App/          ← renombrar a SpectraControl
├── main.py            ← toda la lógica backend
├── index.html         ← frontend completo
├── pyproject.toml
├── uv.lock
└── .hue_config        ← config persistida (bridge IP + API key), gitignoreado
```

## Features implementadas
1. **Configuración de bridge** — auto-descubrimiento via `discovery.meethue.com` + pairing (botón físico)
2. **Control de habitaciones (groups)** — on/off, brillo global, color (presets + picker)
3. **Control de luces individuales** — on/off, brillo, color
4. **Screen Sync** — sincroniza colores de pantalla con las luces en tiempo real

## Decisiones técnicas importantes

### Screen Sync — arquitectura WebSocket + browser capture
El sync de pantalla tuvo varios intentos fallidos en KDE Wayland:
- ❌ `ImageMagick import` — solo X11
- ❌ `scrot` — solo X11  
- ❌ `grim` — solo wlroots (Sway/Hyprland), no KDE
- ❌ `mss` — captura negra en KDE Wayland (XWayland rootless no tiene framebuffer)
- ❌ `xdg-desktop-portal-kde` — no instalado, usa portal GNOME
- ❌ `pywayland` — falla build en Python 3.14

**Solución final:** El frontend usa `navigator.mediaDevices.getDisplayMedia()` (Web API nativa que KDE Wayland sí permite con diálogo de permiso), captura frames en un `<canvas>`, samplea colores por posición horizontal, y los envía al backend via **WebSocket** (`/ws/color`). El backend convierte RGB→CIE XY y manda al bridge Hue.

### Conversión de color
RGB → CIE XY implementado manualmente en `main.py` (función `rgb_to_xy`) con gamma correction y matriz de color Philips Hue.

### Resolución de grupos
El sync opera por **habitación** (group), no por luz individual. El backend resuelve las luces de cada grupo y distribuye puntos de sampleo horizontalmente en la pantalla.

## Entorno
- **OS:** Bazzite (Fedora immutable) con KDE Plasma en Wayland
- **Hardware:** monitor ultrawide, varias habitaciones Hue configuradas
- **Hue Bridge:** v2, API v1 local (REST)
- **Python:** 3.14 (importante — algunas libs no compilan, ej: pywayland)

## Dependencias actuales (pyproject.toml)
```
fastapi
uvicorn[standard]
httpx
pydantic
python-multipart
mss              ← instalado pero NO funciona en KDE Wayland, puede removerse
```

## Cómo correr actualmente
```bash
cd Hue-Sync-App
uv run uvicorn main:app --host 0.0.0.0 --port 8000 --reload
# Abrir http://localhost:8000
```

## Próximos pasos — Migración a Tauri

### Objetivo
Convertir el proyecto en una **app de escritorio nativa con Tauri** para Linux.
- Nombre: **SpectraControl**
- No se va a distribuir en stores, solo instalación manual para el dev y amigos
- El backend Python corre como **sidecar** de Tauri
- El frontend actual (HTML/JS) se usa como webview de Tauri sin cambios mayores
- Agregar ícono, `.desktop` file, instalador simple

### Arquitectura Tauri propuesta
```
spectra-control/
├── src-tauri/          ← Tauri (Rust)
│   ├── src/
│   │   └── main.rs     ← inicia el sidecar Python + abre webview
│   ├── tauri.conf.json
│   └── Cargo.toml
├── frontend/
│   └── index.html      ← frontend actual, sin cambios
├── backend/
│   └── main.py         ← backend actual, sin cambios
└── pyproject.toml
```

### Sidecar Python en Tauri
Tauri puede lanzar el proceso uvicorn como sidecar y matarlo al cerrar la app.
El webview apunta a `http://localhost:8000` donde corre FastAPI.

### Consideraciones
- Tauri requiere Rust + `tauri-cli`
- En Bazzite (immutable) instalar Rust via `rustup` (no rpm-ostree)
- El `getDisplayMedia()` del sync funciona desde el webview de Tauri (Webkit2GTK)
- Necesita ícono SVG/PNG para el `.desktop` y la barra de título

## Configuración del bridge
Se persiste en `.hue_config` (JSON con `ip` y `api_key`) en el directorio de trabajo.
En Tauri esto debería moverse a `$XDG_CONFIG_HOME/spectra-control/config.json`.
