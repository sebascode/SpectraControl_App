# Smoke test beta — runbook

Validación end-to-end del flujo completo de SpectraControl en una máquina
virgen, previo al tag de la beta pública. Bloqueante para `v0.2.8-beta1+`
(ver issue #28).

## Pre-requisitos en el host de prueba

- VM Linux limpia (Fedora / Ubuntu / Debian recientes) o segunda máquina
  física sin estado previo de SpectraControl.
- Acceso físico al **botón redondo** del Hue bridge.
- IP del bridge en la misma LAN.
- Sesión gráfica (X11 o Wayland) con el usuario que va a probar.
- `curl` y `jq` disponibles para los checks de verificación.

> ⚠️ Antes de arrancar, confirmar que no hay estado previo:
> ```bash
> rm -rf ~/.config/spectracontrol ~/.config/autostart/SpectraControl.desktop ~/.local/share/spectracontrol
> ```

---

## 1 · Descarga desde Releases

- [ ] Abrir https://github.com/sebascode/SpectraControl_App/releases en un
      navegador **sin estar logueado en GitHub**. Confirmar que el AppImage
      del último release está visible y descargable.
- [ ] Descargarlo:
      ```bash
      LATEST=$(curl -s https://api.github.com/repos/sebascode/SpectraControl_App/releases/latest | jq -r '.assets[] | select(.name | endswith(".AppImage")) | .browser_download_url')
      echo "URL: $LATEST"
      curl -L -o ~/SpectraControl.AppImage "$LATEST"
      chmod +x ~/SpectraControl.AppImage
      ```
- [ ] Verificar firma / checksum si el release los publica.

**Verificar:** el archivo descargado pesa lo esperado (~50–80 MB),
`file ~/SpectraControl.AppImage` reporta `ELF 64-bit ... executable`.

---

## 2 · Primer arranque · selector de idioma

- [ ] Arrancar: `~/SpectraControl.AppImage`
- [ ] Confirmar que aparece el **setup-screen** con selector de idioma
      visible (es / en).
- [ ] Cambiar el selector a `en` y verificar que los labels del setup
      cambian en vivo.
- [ ] Dejarlo en el idioma de preferencia para continuar.

**Verificar:** archivo de config aún **no existe**:
```bash
test ! -f ~/.config/spectracontrol/hue_config.json && echo "OK (sin estado previo)"
```

---

## 3 · Discovery del bridge

**Auto-discovery:**
- [ ] Click en "🔍 Descubrir automáticamente".
- [ ] Toast verde "Bridge encontrado en X.X.X.X" aparece dentro de ~3s.
- [ ] El campo IP se rellena solo.

**Manual fallback:**
- [ ] Borrar el campo IP, escribirla a mano, confirmar que el botón de pair
      se habilita.

---

## 4 · Pairing con botón físico

- [ ] **Presionar el botón redondo del bridge físicamente.**
- [ ] Dentro de los siguientes 30s: click en "🔗 Vincular".
- [ ] Toast de éxito; el setup-screen avanza al paso siguiente
      (selección de zona / habitaciones).

**Verificar:** ya quedó el config persistido con perms estrictos:
```bash
ls -la ~/.config/spectracontrol/hue_config.json
# Debe ser 0600 (-rw-------) y propiedad del usuario actual.
jq '. | keys' ~/.config/spectracontrol/hue_config.json
# ["api_key", "client_key", "ip"]
```

---

## 5 · Zona de Entertainment

- [ ] En la pantalla principal, navegar a la sección de Sync.
- [ ] Si no existe zona: usar la app oficial de Hue móvil para crear
      "Zona de entretenimiento" con 2–3 luces.
- [ ] Refrescar la lista en SpectraControl; la zona aparece.
- [ ] Seleccionar la zona.

**Verificar:**
```bash
curl -s http://localhost:8000/api/entertainment | jq '.[] | {id, name, lights: (.lights | length)}'
```

---

## 6 · Screen Sync arranca + luces responden

- [ ] Con la zona seleccionada, click "Iniciar Screen Sync".
- [ ] El selector de captura del SO debería aparecer (portal XDG en
      Wayland, o picker nativo en X11).
- [ ] Compartir una pantalla con contenido de colores fuertes
      (poner un video o un sitio con colores saturados).
- [ ] **Confirmar visualmente** que las luces cambian de color reflejando
      lo que hay en pantalla, con latencia ≲ 200 ms.
- [ ] La app Hue móvil muestra el bridge en modo "sincronizado".

---

## 7 · Stop sync + control individual y por habitación

- [ ] Click "Detener Screen Sync".
- [ ] Las luces se quedan en el último color o vuelven al estado previo
      (según preferencia).
- [ ] Ir a "Luces": toggle on/off de una luz responde dentro de ~500 ms.
- [ ] Mover el slider de brillo de una luz; cambia en vivo.
- [ ] Ir a "Habitaciones": toggle on/off por habitación responde.
- [ ] **Bonus** (cubierto por #15): cambiar el brillo desde la **app Hue móvil**
      → la UI de SpectraControl refleja el cambio en ≲ 1 s sin refrescar.

---

## 8 · Autostart de la app

- [ ] En Ajustes → toggle "Iniciar con el sistema" → ON.
- [ ] **Verificar archivo:**
      ```bash
      cat ~/.config/autostart/SpectraControl.desktop
      ```
      Debe contener `Exec=` apuntando al AppImage, `Hidden=false`,
      `X-GNOME-Autostart-enabled=true`.
- [ ] **Logout / login completo** de la sesión gráfica.
- [ ] Tras el login, SpectraControl arranca solo (window o tray).

---

## 9 · Autostart de Screen Sync

- [ ] Con SpectraControl abierto, Ajustes → toggle "Iniciar Screen Sync al
      arrancar" → ON. Confirmar que solo se habilita si #8 está ON.
- [ ] Salir de la app (tray → "Salir").
- [ ] **Logout / login**.
- [ ] Tras el login: SpectraControl arranca **y** el sync arranca solo
      en la zona previamente seleccionada, sin abrir la ventana.

---

## 10 · Botón "Chequear actualizaciones"

- [ ] Ajustes → "Chequear actualizaciones".
- [ ] Toast "Estás en la última versión" aparece (estamos corriendo el
      release más reciente).
- [ ] No deberían aparecer prompts de actualización.

> Si querés probar el path de "hay update disponible", correr una versión
> vieja (ej. `v0.2.7`) y repetir — debe ofrecer actualizar.

---

## 11 · Cambio de idioma en vivo

- [ ] Ajustes → cambiar idioma de `es` ↔ `en`.
- [ ] Confirmar que las labels de la UI cambian al instante.
- [ ] **Importante (memoria [[feedback-i18n-native-strings]]):** abrir el
      tray del sistema y verificar que los items del menú ("Mostrar",
      "Salir", etc.) **también** cambiaron de idioma. Si quedan en el
      idioma anterior, es bug — `syncNativeStrings()` no se está
      disparando.

---

## 12 · Export / import de perfil

**Export:**
- [ ] Ajustes → "Exportar perfil" → descarga `spectracontrol-profile.json`.
- [ ] Inspeccionar el JSON:
      ```bash
      jq '. | keys' ~/Downloads/spectracontrol-profile.json
      ```
      Debe incluir credenciales (`ip`, `api_key`, `client_key`) y
      preferencias (idioma, zona, brillo global, autostart flags).
- [ ] **Tratar el archivo como contraseña** — contiene credenciales del
      bridge.

**Import en otra instancia (o tras reset):**
- [ ] En otra máquina o tras `rm -rf ~/.config/spectracontrol`, arrancar
      SpectraControl, pasar el setup hasta el paso de bridge.
- [ ] Usar "Importar perfil" → seleccionar el JSON.
- [ ] La app salta al estado restaurado: bridge ya configurado, idioma
      correcto, zona seleccionada, sin pedir pairing de nuevo.

---

## 13 · Comportamiento al cerrar

**Minimize a tray:**
- [ ] Con `quit_on_close=false` (default), cerrar la ventana con la X.
- [ ] La ventana desaparece, el icono del tray sigue ahí.
- [ ] Click en el tray → la ventana vuelve.
- [ ] **Verificar que el backend sigue corriendo:**
      ```bash
      pgrep -af spectractl
      curl -s http://localhost:8000/api/config | jq .configured
      ```

**Quit total:**
- [ ] Tray → "Salir".
- [ ] La ventana se cierra, el tray desaparece.
- [ ] **Verificar que el backend murió:**
      ```bash
      pgrep -af spectractl     # no debería devolver nada
      sleep 2; ss -ltn | grep ':8000'   # puerto libre
      ```

**Quit-on-close mode:**
- [ ] Volver a abrir, Ajustes → toggle "Cerrar al pulsar X" → ON.
- [ ] Cerrar con la X.
- [ ] El backend también muere (mismo check que arriba).

---

## Reporte final

Al terminar, anotar en el issue #28:

```
✅ Smoke test completado en <fecha> en <distro / VM>.

Sin issues: <listar items que pasaron limpios>
Con issues: <listar items con bug, link a issue follow-up por cada uno>
```

Si todos los items pasan, el branch está listo para `git tag v0.2.8-beta1`
y publicar release.
