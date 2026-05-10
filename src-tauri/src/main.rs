#![cfg_attr(not(debug_assertions), windows_subsystem = "windows")]

// In dev mode (`cargo tauri dev`), beforeDevCommand in tauri.conf.json starts uvicorn.
// In release builds, this binary spawns uvicorn itself and kills it on close.

use std::io::Read;
use std::process::{Child, Command, Stdio};
use std::sync::{Arc, Mutex, OnceLock};
use std::thread;
use std::time::Duration;
use ashpd::desktop::screencast::{CursorMode, Screencast, SourceType};
use ashpd::desktop::PersistMode;
use tauri::Manager;

// ── SCREEN CAPTURE via ScreenCast portal + GStreamer ────────────────────────
// 1) ashpd negocia la sesión Screencast con xdg-desktop-portal (UNA sola vez).
//    El portal nos da un pipewire node id donde fluye el stream del compositor.
// 2) Spawn de gst-launch-1.0: pipewiresrc → videoscale → videoconvert RGB →
//    fdsink stdout. Frames raw RGB caen en un pipe a 320×180 @ 30 fps.
// 3) Un reader thread los empuja a un Mutex compartido (último gana).
// 4) sample_screen_regions lee del Mutex y samplea — cero llamadas DBus.

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
    let mut child = Command::new("gst-launch-1.0")
        .args(pipeline_args.iter().map(|s| s.to_string()).collect::<Vec<_>>())
        .stdout(Stdio::piped())
        .stderr(Stdio::null())
        .spawn()
        .map_err(|e| format!("spawn gst-launch-1.0: {e}"))?;

    let mut stdout = child.stdout.take().ok_or_else(|| "sin stdout".to_string())?;

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
async fn sample_screen_regions(cols: u32, rows: u32) -> Result<Vec<[u8; 3]>, String> {
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

    let cols = cols.max(1);
    let rows = rows.max(1);
    let sample_size = (CAP_W / 8).max(8);
    let half = sample_size / 2;
    let mut colors: Vec<[u8; 3]> = Vec::with_capacity((cols * rows) as usize);

    for idx in 0..(cols * rows) {
        let col = idx % cols;
        let row = idx / cols;
        let cx = ((col as f32 + 0.5) / cols as f32 * CAP_W as f32) as u32;
        let cy = ((row as f32 + 0.5) / rows as f32 * CAP_H as f32) as u32;
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
                // Cerrar también el gst-launch del screen capture si está vivo
                if let Some(cap) = CAPTURE.get() {
                    if let Some(mut gst) = cap.gst_child.lock().unwrap().take() {
                        let _ = gst.kill();
                    }
                }
            }
        })
        .invoke_handler(tauri::generate_handler![sample_screen_regions])
        .run(tauri::generate_context!())
        .expect("error while running tauri application");
}
