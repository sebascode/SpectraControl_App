#![cfg_attr(not(debug_assertions), windows_subsystem = "windows")]

// In dev mode (`cargo tauri dev`), beforeDevCommand in tauri.conf.json starts the Go backend.
// In release builds, this binary spawns the bundled `spectractl` from resource_dir
// and kills it on window close.

use std::io::{BufRead, BufReader, Read};
use std::process::{Child, Command, Stdio};
use std::sync::atomic::{AtomicBool, Ordering};
use std::sync::{Arc, Mutex, OnceLock};
use std::thread;
use std::time::Duration;
use ashpd::desktop::screencast::{CursorMode, Screencast, SourceType};
use ashpd::desktop::PersistMode;
use tauri::{
    menu::{Menu, MenuItem},
    tray::TrayIconBuilder,
    Manager,
};

// linuxdeploy's AppRun prepends bundled lib/plugin paths to env vars before exec'ing
// the app. Subprocesses we spawn that come from the *host* (gst-launch-1.0) inherit
// these and end up dlopening AppImage-bundled glib/gstreamer libs that don't match
// their build → symbol lookup errors. Strip the polluting vars before exec.
fn sanitize_host_env(cmd: &mut Command) {
    const VARS: &[&str] = &[
        "LD_LIBRARY_PATH",
        "LD_PRELOAD",
        "GST_PLUGIN_PATH",
        "GST_PLUGIN_PATH_1_0",
        "GST_PLUGIN_SYSTEM_PATH",
        "GST_PLUGIN_SYSTEM_PATH_1_0",
        "GST_PLUGIN_SCANNER",
        "GST_REGISTRY",
        "GIO_MODULE_DIR",
        "GIO_EXTRA_MODULES",
        "GSETTINGS_SCHEMA_DIR",
        "GTK_PATH",
        "GTK_DATA_PREFIX",
        "GTK_EXE_PREFIX",
        "GTK_IM_MODULE_FILE",
        "GDK_PIXBUF_MODULE_FILE",
        "QT_PLUGIN_PATH",
        "PYTHONHOME",
        "PYTHONPATH",
        "PERLLIB",
        "XDG_DATA_DIRS",
    ];
    for var in VARS {
        cmd.env_remove(var);
    }
}

// ── SCREEN CAPTURE via ScreenCast portal + GStreamer ────────────────────────
// 1) ashpd negocia la sesión Screencast con xdg-desktop-portal (UNA sola vez).
//    El portal nos da un pipewire node id donde fluye el stream del compositor.
// 2) Spawn de gst-launch-1.0: pipewiresrc → videoscale → videoconvert RGB →
//    fdsink stdout. Frames raw RGB caen en un pipe a 320×180 @ 30 fps.
// 3) Un reader thread los empuja a un Mutex compartido (último gana).
// 4) sample_screen_regions lee del Mutex y samplea en cada posición pedida —
//    cero llamadas DBus.

const CAP_W: u32 = 320;
const CAP_H: u32 = 180;
const CAP_FPS: u32 = 30;
const FRAME_BYTES: usize = (CAP_W * CAP_H * 3) as usize;

struct CaptureState {
    latest: Arc<Mutex<Option<Vec<u8>>>>,
    init_lock: tokio::sync::Mutex<bool>,
    gst_child: Mutex<Option<Child>>,
}

static CAPTURE: OnceLock<CaptureState> = OnceLock::new();

fn capture_state() -> &'static CaptureState {
    CAPTURE.get_or_init(|| CaptureState {
        latest: Arc::new(Mutex::new(None)),
        init_lock: tokio::sync::Mutex::new(false),
        gst_child: Mutex::new(None),
    })
}

async fn ensure_capture_initialized() -> Result<(), String> {
    let state = capture_state();
    let mut init = state.init_lock.lock().await;
    if *init {
        return Ok(());
    }

    // 1) Negociar Screencast con el portal
    let proxy = Screencast::new()
        .await
        .map_err(|e| format!("portal proxy: {e}"))?;
    let session = proxy
        .create_session()
        .await
        .map_err(|e| format!("create_session: {e}"))?;
    proxy
        .select_sources(
            &session,
            CursorMode::Hidden,
            SourceType::Monitor.into(),
            false,
            None,
            PersistMode::DoNot,
        )
        .await
        .map_err(|e| format!("select_sources: {e}"))?;
    let response = proxy
        .start(&session, None)
        .await
        .map_err(|e| format!("start: {e}"))?
        .response()
        .map_err(|e| format!("start response: {e}"))?;
    let streams: Vec<_> = response.streams().to_vec();
    let stream_info = streams.first().ok_or_else(|| "sin streams".to_string())?;
    let node_id = stream_info.pipe_wire_node_id();
    eprintln!("[capture] portal listo — node_id={node_id}");

    // 2) Lanzar gst-launch leyendo de pipewire y escupiendo RGB raw a stdout
    let pipeline_args = [
        "-q",
        "pipewiresrc",
        &format!("path={node_id}"),
        "!",
        "videorate",
        "!",
        &format!("video/x-raw,framerate={CAP_FPS}/1"),
        "!",
        "videoscale",
        "!",
        &format!("video/x-raw,width={CAP_W},height={CAP_H}"),
        "!",
        "videoconvert",
        "!",
        "video/x-raw,format=RGB",
        "!",
        "fdsink",
        "fd=1",
        "sync=false",
    ];
    let mut cmd = Command::new("gst-launch-1.0");
    cmd.args(pipeline_args.iter().map(|s| s.to_string()).collect::<Vec<_>>())
        .stdout(Stdio::piped())
        .stderr(Stdio::piped());
    sanitize_host_env(&mut cmd);
    let mut child = cmd
        .spawn()
        .map_err(|e| format!("spawn gst-launch-1.0: {e}"))?;

    let mut stdout = child.stdout.take().ok_or_else(|| "sin stdout".to_string())?;
    if let Some(stderr) = child.stderr.take() {
        thread::spawn(move || {
            for line in BufReader::new(stderr).lines().map_while(Result::ok) {
                eprintln!("[capture/gst] {line}");
            }
        });
    }

    // 3) Reader thread: lee frames completos y los publica
    let latest = state.latest.clone();
    thread::spawn(move || {
        let mut buf = vec![0u8; FRAME_BYTES];
        loop {
            let mut filled = 0;
            while filled < FRAME_BYTES {
                match stdout.read(&mut buf[filled..]) {
                    Ok(0) => {
                        eprintln!("[capture] gst-launch cerró stdout");
                        return;
                    }
                    Ok(n) => filled += n,
                    Err(e) => {
                        eprintln!("[capture] read error: {e}");
                        return;
                    }
                }
            }
            *latest.lock().unwrap() = Some(buf.clone());
        }
    });

    *state.gst_child.lock().unwrap() = Some(child);
    *init = true;
    Ok(())
}

#[tauri::command]
async fn sample_screen_regions(positions: Vec<[f32; 2]>) -> Result<Vec<[u8; 3]>, String> {
    ensure_capture_initialized().await?;

    // Esperar el primer frame (hasta 1.5s la primera vez tras negociar el portal)
    let state = capture_state();
    let mut tries = 0;
    let frame = loop {
        if let Some(f) = state.latest.lock().unwrap().clone() {
            break f;
        }
        if tries > 75 {
            return Err("frame no disponible (gst-launch no produjo)".to_string());
        }
        tries += 1;
        tokio::time::sleep(Duration::from_millis(20)).await;
    };

    let sample_size = (CAP_W / 8).max(8);
    let half = sample_size / 2;
    let mut colors: Vec<[u8; 3]> = Vec::with_capacity(positions.len());

    for [px, py] in positions {
        let cx = (px.clamp(0.0, 1.0) * CAP_W as f32) as u32;
        let cy = (py.clamp(0.0, 1.0) * CAP_H as f32) as u32;
        let x0 = cx.saturating_sub(half);
        let y0 = cy.saturating_sub(half);
        let x1 = (cx + half).min(CAP_W);
        let y1 = (cy + half).min(CAP_H);

        let mut r: u64 = 0;
        let mut g: u64 = 0;
        let mut b: u64 = 0;
        let mut count: u64 = 0;
        for y in y0..y1 {
            for x in x0..x1 {
                let off = ((y * CAP_W + x) * 3) as usize;
                r += frame[off] as u64;
                g += frame[off + 1] as u64;
                b += frame[off + 2] as u64;
                count += 1;
            }
        }
        let count = count.max(1);
        colors.push([(r / count) as u8, (g / count) as u8, (b / count) as u8]);
    }
    Ok(colors)
}

#[cfg(target_os = "linux")]
fn allow_display_capture(app: &tauri::App) {
    if let Some(window) = app.get_webview_window("main") {
        let _ = window.with_webview(|wv| {
            use webkit2gtk::{WebViewExt, SettingsExt, PermissionRequestExt};
            let wv = wv.inner();
            // Enable getUserMedia / getDisplayMedia in the embedded WebView.
            if let Some(settings) = wv.settings() {
                settings.set_enable_media(true);
                settings.set_enable_media_stream(true);
            }
            // Auto-allow any browser permission request (camera, mic, screen capture).
            wv.connect_permission_request(|_, request| {
                request.allow();
                true
            });
        });
    }
}

struct BackendProcess(Mutex<Option<Child>>);

// is_quitting=false → cerrar la ventana solo la oculta (sigue corriendo en el
// tray; la sync continúa porque el backend Go sigue vivo). El menú "Salir" del
// tray pone el flag en true antes de pedir exit, y entonces sí limpiamos.
struct IsQuitting(AtomicBool);

// quit_on_close=true → la X de la ventana cierra la app entera (igual que el
// "Salir" del tray). false → la X oculta al tray. El frontend lo empuja desde
// el toggle de Apariencia con set_quit_on_close.
struct QuitOnClose(AtomicBool);

#[tauri::command]
fn set_quit_on_close(on: bool, state: tauri::State<'_, QuitOnClose>) {
    state.0.store(on, Ordering::SeqCst);
}

// Referencias a los items del menú del tray para poder reescribir sus labels
// cuando el frontend cambia de idioma (los strings de Rust no pasan por el
// helper i18n del frontend).
struct TrayMenuItems {
    show: MenuItem<tauri::Wry>,
    quit: MenuItem<tauri::Wry>,
}

#[tauri::command]
fn set_tray_menu_labels(
    show: String,
    quit: String,
    items: tauri::State<'_, TrayMenuItems>,
) -> Result<(), String> {
    items.show.set_text(show).map_err(|e| e.to_string())?;
    items.quit.set_text(quit).map_err(|e| e.to_string())?;
    Ok(())
}

// write_text_file complementa al save-dialog de tauri-plugin-dialog: el JS
// pide al usuario una ruta, y luego nos llama acá para volcar el contenido.
// Sin scope adicional porque la ruta ya pasó por el diálogo nativo del SO.
#[tauri::command]
fn write_text_file(path: String, contents: String) -> Result<(), String> {
    std::fs::write(&path, contents).map_err(|e| e.to_string())
}

fn cleanup_children(app: &tauri::AppHandle) {
    if let Some(state) = app.try_state::<BackendProcess>() {
        if let Some(mut child) = state.0.lock().unwrap().take() {
            let _ = child.kill();
        }
    }
    if let Some(cap) = CAPTURE.get() {
        if let Some(mut gst) = cap.gst_child.lock().unwrap().take() {
            let _ = gst.kill();
        }
    }
}

fn main() {
    tauri::Builder::default()
        .plugin(tauri_plugin_updater::Builder::new().build())
        .plugin(tauri_plugin_process::init())
        .plugin(tauri_plugin_notification::init())
        .plugin(tauri_plugin_dialog::init())
        .plugin(tauri_plugin_autostart::init(
            tauri_plugin_autostart::MacosLauncher::LaunchAgent,
            None,
        ))
        .manage(BackendProcess(Mutex::new(None)))
        .manage(IsQuitting(AtomicBool::new(false)))
        .manage(QuitOnClose(AtomicBool::new(false)))
        .setup(|app| {
            #[cfg(not(debug_assertions))]
            {
                let resource_dir = app.path().resource_dir()
                    .expect("cannot determine resource directory");
                let backend_bin = resource_dir.join("backend").join("spectractl");
                let child = std::process::Command::new(&backend_bin)
                    .args(["-addr", "127.0.0.1:8000"])
                    .spawn()
                    .expect("failed to start backend");
                *app.state::<BackendProcess>().0.lock().unwrap() = Some(child);
                std::thread::sleep(std::time::Duration::from_millis(300));
            }
            #[cfg(target_os = "linux")]
            allow_display_capture(app);

            // System tray. Mantiene la app viva al ocultar la ventana — la
            // sync sigue corriendo en background. Los labels se reescriben
            // desde JS en cuanto initI18n() resuelve (set_tray_menu_labels),
            // los placeholders quedan en español por si JS no llega a llamar.
            let show_item = MenuItem::with_id(app, "show", "Mostrar", true, None::<&str>)?;
            let quit_item = MenuItem::with_id(app, "quit", "Salir", true, None::<&str>)?;
            let menu = Menu::with_items(app, &[&show_item, &quit_item])?;
            app.manage(TrayMenuItems {
                show: show_item.clone(),
                quit: quit_item.clone(),
            });

            let icon = app
                .default_window_icon()
                .cloned()
                .expect("default window icon missing");

            TrayIconBuilder::with_id("main")
                .icon(icon)
                .tooltip("SpectraControl")
                .menu(&menu)
                .show_menu_on_left_click(false)
                .on_menu_event(|app, event| match event.id.as_ref() {
                    "show" => {
                        if let Some(window) = app.get_webview_window("main") {
                            let _ = window.unminimize();
                            let _ = window.show();
                            let _ = window.set_focus();
                        }
                    }
                    "quit" => {
                        app.state::<IsQuitting>().0.store(true, Ordering::SeqCst);
                        cleanup_children(app);
                        app.exit(0);
                    }
                    _ => {}
                })
                .on_tray_icon_event(|tray, event| {
                    // Click izquierdo en el icono = mostrar la ventana.
                    if let tauri::tray::TrayIconEvent::Click {
                        button: tauri::tray::MouseButton::Left,
                        button_state: tauri::tray::MouseButtonState::Up,
                        ..
                    } = event
                    {
                        let app = tray.app_handle();
                        if let Some(window) = app.get_webview_window("main") {
                            let _ = window.unminimize();
                            let _ = window.show();
                            let _ = window.set_focus();
                        }
                    }
                })
                .build(app)?;

            Ok(())
        })
        .on_window_event(|window, event| {
            if let tauri::WindowEvent::CloseRequested { api, .. } = event {
                let app = window.app_handle();
                if app.state::<IsQuitting>().0.load(Ordering::SeqCst) {
                    // Salida real: la limpieza ya la hizo el handler del menú,
                    // dejamos pasar el close.
                    return;
                }
                if app.state::<QuitOnClose>().0.load(Ordering::SeqCst) {
                    // El usuario eligió que la X cierre la app: limpiamos y
                    // marcamos is_quitting para que el siguiente CloseRequested
                    // (si llega) pase derecho.
                    app.state::<IsQuitting>().0.store(true, Ordering::SeqCst);
                    cleanup_children(&app);
                    app.exit(0);
                    return;
                }
                // Default: ocultar al tray.
                api.prevent_close();
                let _ = window.hide();
            }
        })
        .invoke_handler(tauri::generate_handler![
            sample_screen_regions,
            set_tray_menu_labels,
            set_quit_on_close,
            write_text_file
        ])
        .build(tauri::generate_context!())
        .expect("error while building tauri application")
        .run(|app, event| {
            // Cualquier path de salida — tray "Salir", X con quit_on_close,
            // process.exit() del frontend tras instalar update — pasa por
            // RunEvent::Exit. Garantizamos que el backend Go muera acá para
            // que no quede huérfano ocupando el puerto.
            if let tauri::RunEvent::Exit = event {
                cleanup_children(app);
            }
        });
}

#[cfg(test)]
mod tests {
    use super::*;
    use std::ffi::OsString;

    #[test]
    fn sanitize_host_env_removes_all_polluting_vars() {
        // Build a Command pre-populated with every var sanitize_host_env claims to strip.
        let polluting = [
            "LD_LIBRARY_PATH",
            "LD_PRELOAD",
            "GST_PLUGIN_PATH",
            "GST_PLUGIN_PATH_1_0",
            "GST_PLUGIN_SYSTEM_PATH",
            "GST_PLUGIN_SYSTEM_PATH_1_0",
            "GST_PLUGIN_SCANNER",
            "GST_REGISTRY",
            "GIO_MODULE_DIR",
            "GIO_EXTRA_MODULES",
            "GSETTINGS_SCHEMA_DIR",
            "GTK_PATH",
            "GTK_DATA_PREFIX",
            "GTK_EXE_PREFIX",
            "GTK_IM_MODULE_FILE",
            "GDK_PIXBUF_MODULE_FILE",
            "QT_PLUGIN_PATH",
            "PYTHONHOME",
            "PYTHONPATH",
            "PERLLIB",
            "XDG_DATA_DIRS",
        ];

        let mut cmd = Command::new("/bin/true");
        for v in &polluting {
            cmd.env(v, "appimage-poison");
        }
        // A non-polluting var should survive sanitize_host_env untouched.
        cmd.env("KEEP_ME", "ok");

        sanitize_host_env(&mut cmd);

        // Command::get_envs yields (key, None) for env_remove calls and
        // (key, Some(value)) for env(...) calls. After sanitize_host_env, each
        // polluting var must appear as a removal.
        let envs: Vec<(OsString, Option<OsString>)> = cmd
            .get_envs()
            .map(|(k, v)| (k.to_owned(), v.map(|v| v.to_owned())))
            .collect();

        for v in &polluting {
            let entry = envs.iter().find(|(k, _)| k == v);
            match entry {
                Some((_, None)) => {} // removed — good
                Some((_, Some(val))) => panic!(
                    "{v} still set to {val:?} after sanitize_host_env"
                ),
                None => panic!("{v} missing from Command envs (expected a removal entry)"),
            }
        }

        // KEEP_ME should still be present with its value.
        let keep = envs.iter().find(|(k, _)| k == "KEEP_ME");
        assert!(
            matches!(keep, Some((_, Some(v))) if v == "ok"),
            "KEEP_ME was disturbed by sanitize_host_env: {keep:?}"
        );
    }
}
