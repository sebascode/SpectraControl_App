# Solución de problemas

Esta guía cubre los problemas más comunes al usar SpectraControl con un Hue
Bridge. Si tu problema no aparece acá, abrí un issue en GitHub adjuntando los
logs (ver la última sección).

> Versión en inglés: [troubleshooting.en.md](troubleshooting.en.md)

## Cómo capturar logs

Casi todo el diagnóstico empieza por activar el modo `debug` y reproducir el
problema.

### Modo desarrollo (`cargo tauri dev`)

```bash
SPECTRA_LOG_LEVEL=debug cargo tauri dev
```

Los logs del backend Go (prefijo `[ent]`, `[sse]`, etc.) y del shell Rust salen
mezclados en la misma terminal.

### App instalada (binario nativo)

```bash
SPECTRA_LOG_LEVEL=debug spectra-control
```

Sustituí `spectra-control` por la ruta real del ejecutable si no está en el
`PATH`.

### Flatpak

```bash
flatpak run --env=SPECTRA_LOG_LEVEL=debug casa.scode.SpectraControl
```

---

## Problemas comunes

### Las luces no se apagan al cerrar la app durante un sync

**Síntoma:** cerrás la ventana mientras estás sincronizando y las luces quedan
prendidas en el último color del frame.

**Causa:** el toggle "Apagar las luces al cerrar (si está sincronizando)" está
desactivado. Es opt-in: la preferencia por defecto es respetar el último
estado.

**Solución:** Settings → Apariencia → activá "Apagar las luces al cerrar (si
está sincronizando)".

---

### Error "DTLS handshake context deadline exceeded" al iniciar sync

**Síntoma:** al apretar Sync no arranca el stream y aparece un error mencionando
"DTLS handshake" o "deadline exceeded". En el log del backend se ve
`[ent] DTLS dial fallido a <ip>:2100`.

**Causa:** el Hue Bridge solo permite **un** stream de Entertainment activo a la
vez, y otro cliente lo tiene tomado: típicamente una Hue Sync Box, la app Hue
Sync de escritorio, una consola de juegos con Sync, u otra instancia de
SpectraControl.

**Solución:**

1. Asegurate de tener activo el toggle "Romper otras sincronizaciones al
   iniciar" (Settings → Screen Sync). Está activado por defecto. Esto stoppea
   automáticamente cualquier otra sesión antes de iniciar la nuestra.
2. Si el toggle ya estaba activo y el problema persiste, posiblemente la Hue
   Sync Box hardware está reconectándose automáticamente (firmware con
   auto-resume agresivo). Pausala desde la app Hue Sync, ponela en modo
   passthrough HDMI, o desconectala temporalmente.
3. También podés probar el botón manual "Detener sincronizaciones de otras apps
   ahora" en Settings → Screen Sync.

Para confirmar que es esto: corré con `SPECTRA_LOG_LEVEL=debug` y mirá las
líneas `[ent] config encontrada — id=... name=... status="active" ...`. Si ves
una con `status="active"` que no es la tuya, ese es el conflicto.

---

### Las luces parpadean o no responden a comandos manuales

**Síntoma:** al mover sliders de brillo/color desde la UI, las luces no
cambian o cambian con retraso de varios segundos.

**Causa:** hay un stream de Entertainment activo (puede ser nuestro, una Sync
Box, etc.) que está pisando los frames de color por DTLS a ~30 fps. Los
comandos HTTP individuales se aplican pero quedan sobrescritos al instante.

**Solución:** detené cualquier sync (botón Stop) y, si el problema persiste,
usá "Detener sincronizaciones de otras apps ahora" para cerrar streams ajenos.

---

### El bridge no se detecta / "no se encuentra puente"

**Síntoma:** al abrir la app por primera vez (o tras un reinicio del router) no
encuentra el bridge.

**Causas posibles:**

- El bridge cambió de IP (DHCP rotó la dirección). SpectraControl guarda la IP
  desde el primer pareo.
- El firewall del SO está bloqueando mDNS (puerto 5353/UDP).
- El bridge y la PC están en redes distintas (típico con Wi-Fi de invitados o
  VLANs).

**Solución:**

1. Verificá que el bridge tenga su LED de red encendido (azul fijo).
2. Pingueá la IP del bridge desde la terminal (`ping <ip>`).
3. Si cambió, borrá el archivo de config y re-pareá:
   - Linux: `~/.config/SpectraControl/config.json`
   - Flatpak: `~/.var/app/casa.scode.SpectraControl/config/SpectraControl/config.json`
4. Asegurate de que tu Wi-Fi/Ethernet esté en la misma subred que el bridge.

---

### El sync de audio no captura nada

**Síntoma:** activás Audio Sync pero las luces no reaccionan al sonido del
sistema.

**Causa:** el sync de audio requiere captura nativa del audio del sistema. Esto
solo funciona en la **app de escritorio**, no en el navegador. Internamente
spawnea `gst-launch-1.0` (GStreamer) para abrir un loopback de PulseAudio /
PipeWire.

**Solución:**

1. Confirmá que estás usando la app instalada, no `localhost:8000` en el
   navegador.
2. En Linux nativo, asegurate de tener `gstreamer1.0-pulseaudio` instalado.
3. En Flatpak ya viene incluido.
4. Verificá que tu sistema esté reproduciendo audio (un parlante / auriculares
   activos). Si todo está silenciado, no hay nada que reaccionar.

---

### Las notificaciones del sistema no aparecen

**Síntoma:** la app dice "sync iniciado" en un toast pero no aparece la
notificación nativa.

**Causa:** GNOME silencia las notificaciones de apps no instaladas, o falta
`libnotify-bin` (`notify-send`) en el host. SpectraControl shellea a
`notify-send` directamente porque el plugin de Tauri panic'ea con el runtime de
GNOME y `notify-rust` directo es silenciado.

**Solución:**

- Linux nativo: `sudo dnf install libnotify` (Fedora) o
  `sudo apt install libnotify-bin` (Debian/Ubuntu).
- Flatpak: las notificaciones van por el portal `xdg-desktop-portal` y deberían
  funcionar siempre que el portal esté instalado y corriendo.

---

### "Iniciar con el sistema" no hace nada

**Síntoma:** activás el toggle de autostart en Apariencia pero la app no se
abre al iniciar sesión.

**Causa:** issue conocido — ver
[#24](https://github.com/sebascode/SpectraControl_App/issues/24).

**Workaround:** crear manualmente `~/.config/autostart/SpectraControl.desktop`
apuntando al binario.

---

## Reportar un bug

Si el problema no aparece acá o el workaround no funciona, abrí un issue en
[GitHub](https://github.com/sebascode/SpectraControl_App/issues) e incluí:

1. **Versión de SpectraControl** (visible en Settings → Acerca de, o
   `cat /var/lib/flatpak/app/casa.scode.SpectraControl/current/active/files/manifest.json`
   en Flatpak).
2. **Sistema operativo** y entorno de escritorio (ej: Fedora 44, GNOME 47
   Wayland).
3. **Modo de instalación** (Flatpak / binario / desarrollo).
4. **Modelo y firmware del Hue Bridge** (visible en la app oficial de Hue,
   Settings → Hue Bridges → tu bridge).
5. **Logs en modo debug** capturados al reproducir el problema (ver "Cómo
   capturar logs" arriba). Pegá las líneas relevantes — no hace falta el log
   completo si es muy largo.
6. **Pasos para reproducir** y qué esperabas que pasara.

Si vas a compartir en un foro de terceros (developers.meethue.com, Reddit,
etc.), remové tokens del bridge antes de pegar logs: el `clientkey` del
`config.json` no debería aparecer en logs pero conviene revisar.
