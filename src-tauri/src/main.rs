#![cfg_attr(not(debug_assertions), windows_subsystem = "windows")]

// In dev mode (`cargo tauri dev`), beforeDevCommand in tauri.conf.json starts uvicorn.
// In release builds, this binary spawns uvicorn itself and kills it on close.

use std::process::Child;
use std::sync::Mutex;
use tauri::Manager;

#[cfg(target_os = "linux")]
fn allow_display_capture(app: &tauri::App) {
    if let Some(window) = app.get_webview_window("main") {
        let _ = window.with_webview(|wv| {
            use webkit2gtk::{WebViewExt, PermissionRequestExt};
            wv.inner().connect_permission_request(|_, request| {
                request.allow();
                true
            });
        });
    }
}

struct BackendProcess(Mutex<Option<Child>>);

fn main() {
    tauri::Builder::default()
        .manage(BackendProcess(Mutex::new(None)))
        .setup(|app| {
            #[cfg(not(debug_assertions))]
            {
                let resource_dir = app.path().resource_dir()
                    .expect("cannot determine resource directory");
                let backend_bin = resource_dir.join("backend").join("spectractl");
                let child = std::process::Command::new(&backend_bin)
                    .args(["--addr", "127.0.0.1:8000"])
                    .spawn()
                    .expect("failed to start backend");
                *app.state::<BackendProcess>().0.lock().unwrap() = Some(child);
                // Go arranca mucho más rápido que uvicorn
                std::thread::sleep(std::time::Duration::from_millis(300));
            }
            #[cfg(target_os = "linux")]
            allow_display_capture(app);

            Ok(())
        })
        .on_window_event(|window, event| {
            if let tauri::WindowEvent::CloseRequested { .. } = event {
                let state = window.state::<BackendProcess>();
                let mut guard = state.0.lock().unwrap();
                if let Some(mut child) = guard.take() {
                    let _ = child.kill();
                }
            }
        })
        .run(tauri::generate_context!())
        .expect("error while running tauri application");
}
