// Autostart vía xdg-desktop-portal cuando corremos en Flatpak. El plugin
// `tauri-plugin-autostart` escribe `~/.config/autostart/<app>.desktop` en el
// HOME del usuario; dentro del sandbox eso aterriza en
// `~/.var/app/<app-id>/config/autostart/`, invisible para el XDG autostart
// del host. El portal Background hace lo correcto: el frontend manda el
// toggle por acá cuando `is_flatpak()` es true, y `RequestBackground` con
// `autostart=true` registra el .desktop en el host.
//
// El portal no expone un "get current state": solo aplica. El frontend
// persiste el estado en localStorage (igual que `autostart.sync`) y nos
// trata como el lado mutador.

use ashpd::desktop::background::Background;
use log::warn;

#[tauri::command]
pub async fn set_autostart_portal(enabled: bool) -> Result<bool, String> {
    let response = Background::request()
        .reason("Sync your Hue lights with the screen in the background")
        .auto_start(enabled)
        .dbus_activatable(false)
        .send()
        .await
        .map_err(|e| {
            warn!("portal Background.RequestBackground send failed: {e}");
            format!("portal request failed: {e}")
        })?
        .response()
        .map_err(|e| {
            warn!("portal Background response error: {e}");
            format!("portal response error: {e}")
        })?;
    Ok(response.auto_start())
}
