// Detección del entorno de ejecución. Hoy solo distingue Flatpak vs todo lo
// demás (AppImage, .deb/.rpm, `cargo run`). Lo usamos para:
//   - Apagar tauri-plugin-updater: Flathub prohíbe auto-update, las apps las
//     actualiza el repo.
//   - Enrutar autostart y notificaciones por los portales XDG cuando el
//     sandbox no permite escribir `~/.config/autostart` ni shellear a
//     notify-send.

use serde::Serialize;

pub fn is_flatpak() -> bool {
    std::env::var_os("FLATPAK_ID").is_some()
}

#[derive(Serialize)]
pub struct RuntimeInfo {
    pub is_flatpak: bool,
}

#[tauri::command]
pub fn runtime_environment() -> RuntimeInfo {
    RuntimeInfo {
        is_flatpak: is_flatpak(),
    }
}
