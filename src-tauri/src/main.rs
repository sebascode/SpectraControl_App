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
                let project_root = std::env::current_dir()
                    .expect("cannot determine working directory");
                let child = std::process::Command::new("uv")
                    .args([
                        "run", "uvicorn",
                        "--app-dir", "backend",
                        "main:app",
                        "--host", "127.0.0.1",
                        "--port", "8000",
                    ])
                    .current_dir(&project_root)
                    .spawn()
                    .expect("failed to start backend — is uv installed and in PATH?");
                *app.state::<BackendProcess>().0.lock().unwrap() = Some(child);
                // Give uvicorn time to bind before the webview loads
                std::thread::sleep(std::time::Duration::from_millis(1500));
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
