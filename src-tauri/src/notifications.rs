// Notificaciones nativas. Dos paths según el runtime:
//
// - Host (AppImage / .deb / .rpm / dev): shell-out a `notify-send`
//   (libnotify-bin). Es lo único que GNOME Shell renderiza de forma
//   confiable; ver historia en los comentarios abajo.
//
// - Flatpak: el sandbox no expone el bus de notify-send hacia el host de
//   forma directa, y meter libnotify-bin en el manifest sería un
//   workaround que no escala. Usamos org.freedesktop.portal.Notification.
//   Si el portal falla por algo (no esperable, pero protocolos cambian),
//   intentamos notify-send como fallback antes de devolver error.
//
// Historial de por qué NO usamos tauri-plugin-notification ni notify-rust
// directo:
// - tauri-plugin-notification 2.x: spawnea notify-rust en un worker de
//   tokio donde el block_on interno de zbus dispara "Cannot start a
//   runtime from within a runtime" → panic.
// - notify-rust desde un std::thread aislado: dbus responde OK con un id,
//   pero GNOME Shell descarta la notificación silenciosamente incluso con
//   DesktopEntry / urgency / icon hints (parece ser credenciales del
//   sender que libnotify maneja y notify-rust no).
// - notify-send (libnotify-bin) es ubicuo en escritorios Linux y
//   verificado que renderiza.

use ashpd::desktop::notification::{Notification, NotificationProxy, Priority};
use log::warn;

use crate::runtime;

#[tauri::command]
pub async fn notify_native(title: String, body: String) -> Result<(), String> {
    if runtime::is_flatpak() {
        match notify_via_portal(&title, &body).await {
            Ok(()) => return Ok(()),
            Err(e) => warn!("portal notification failed, falling back to notify-send: {e}"),
        }
    }
    notify_via_shell(&title, &body)
}

async fn notify_via_portal(title: &str, body: &str) -> Result<(), String> {
    let proxy = NotificationProxy::new()
        .await
        .map_err(|e| format!("notification portal connect: {e}"))?;
    // El portal requiere un id por notificación. No las actualizamos ni
    // retiramos desde acá, así que un id único por llamada es suficiente.
    let id = format!(
        "spectra-{}",
        std::time::SystemTime::now()
            .duration_since(std::time::UNIX_EPOCH)
            .map(|d| d.as_millis())
            .unwrap_or(0)
    );
    let notification = Notification::new(title)
        .body(Some(body))
        .priority(Priority::Normal);
    proxy
        .add_notification(&id, notification)
        .await
        .map_err(|e| format!("add_notification: {e}"))?;
    Ok(())
}

fn notify_via_shell(title: &str, body: &str) -> Result<(), String> {
    let output = std::process::Command::new("notify-send")
        .arg("--app-name=SpectraControl")
        .arg("--icon=casa.scode.SpectraControl")
        .arg("--expire-time=5000")
        .arg(title)
        .arg(body)
        .output()
        .map_err(|e| format!("spawn notify-send: {}", e))?;
    if !output.status.success() {
        let stderr = String::from_utf8_lossy(&output.stderr);
        return Err(format!("notify-send exit {}: {}", output.status, stderr));
    }
    Ok(())
}
