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
use log::{debug, error, info, warn};
use tauri::{
    menu::{Menu, MenuItem},
    tray::TrayIconBuilder,
    Emitter, Manager,
};

// Inicializa env_logger leyendo SPECTRA_LOG_LEVEL (misma env var que el backend
// Go), default "info". Filtros tipo "debug,hyper=warn" también funcionan,
// env_logger los parsea igual que RUST_LOG.
fn setup_logger() {
    let level = std::env::var("SPECTRA_LOG_LEVEL").unwrap_or_else(|_| "info".to_string());
    env_logger::Builder::new()
        .parse_filters(&level)
        .format_timestamp_millis()
        .try_init()
        .ok();
}

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
    info!("[capture] portal listo — node_id={node_id}");

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
                debug!("[capture/gst] {line}");
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
                        warn!("[capture] gst-launch cerró stdout");
                        return;
                    }
                    Ok(n) => filled += n,
                    Err(e) => {
                        error!("[capture] read error: {e}");
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

// ── AUDIO CAPTURE via pulsesrc + GStreamer ──────────────────────────────────
// El portal xdg-desktop-portal de pantalla en Linux todavía no expone toggle
// de audio (GNOME ≤46 al menos), así que getDisplayMedia({audio:true}) cae
// silenciosamente sin track. Para el audio sync corremos un pipeline
// pulsesrc apuntando al monitor del sink default — sin diálogo, sin permiso
// extra. La FFT se hace acá en Rust con rustfft, y devolvemos un Vec<u8>
// con la misma semántica que AnalyserNode.getByteFrequencyData (fftSize
// 1024 → 512 bins, dB clamp -100..-30 → 0..255), así audioFrameToRGB() en
// JS funciona sin cambios.
const AUDIO_RATE: u32 = 44100;
const AUDIO_FFT: usize = 1024;
const AUDIO_BINS: usize = AUDIO_FFT / 2;
const AUDIO_RING: usize = AUDIO_FFT * 2;

struct AudioRing {
    buf: Vec<f32>,
    head: usize,
}

impl AudioRing {
    fn new(cap: usize) -> Self {
        Self {
            buf: vec![0.0; cap],
            head: 0,
        }
    }
    fn push(&mut self, samples: &[f32]) {
        for &s in samples {
            self.buf[self.head] = s;
            self.head = (self.head + 1) % self.buf.len();
        }
    }
    fn copy_latest(&self, out: &mut [f32]) {
        let n = out.len().min(self.buf.len());
        let start = (self.head + self.buf.len() - n) % self.buf.len();
        for (i, slot) in out.iter_mut().enumerate().take(n) {
            *slot = self.buf[(start + i) % self.buf.len()];
        }
    }
}

struct AudioCaptureState {
    ring: Arc<Mutex<AudioRing>>,
    init_lock: tokio::sync::Mutex<bool>,
    gst_child: Mutex<Option<Child>>,
    // Smoothing IIR estado-por-bin (alpha = smoothingTimeConstant en
    // AnalyserNode). Inicialmente en cero → primer frame entra crudo.
    smoothed: Mutex<Vec<f32>>,
    // Plan FFT compartido — rustfft::Fft no es Send+Sync via Arc<dyn>, así
    // que lo guardamos detrás de un Mutex.
    fft: Mutex<Arc<dyn rustfft::Fft<f32>>>,
}

static AUDIO_CAPTURE: OnceLock<AudioCaptureState> = OnceLock::new();

fn audio_capture_state() -> &'static AudioCaptureState {
    AUDIO_CAPTURE.get_or_init(|| {
        let mut planner = rustfft::FftPlanner::<f32>::new();
        AudioCaptureState {
            ring: Arc::new(Mutex::new(AudioRing::new(AUDIO_RING))),
            init_lock: tokio::sync::Mutex::new(false),
            gst_child: Mutex::new(None),
            smoothed: Mutex::new(vec![0.0; AUDIO_BINS]),
            fft: Mutex::new(planner.plan_fft_forward(AUDIO_FFT)),
        }
    })
}

async fn ensure_audio_capture_initialized() -> Result<(), String> {
    let state = audio_capture_state();
    let mut init = state.init_lock.lock().await;
    if *init {
        return Ok(());
    }

    // pulsesrc device=@DEFAULT_MONITOR@ captura el monitor del sink default,
    // lo que en PipeWire (con la capa de compatibilidad Pulse) es el mismo
    // audio que sale por los parlantes. Forzamos F32LE mono @ 44.1k para
    // tener un buffer denso y simple que alimentar a la FFT.
    let caps = format!(
        "audio/x-raw,format=F32LE,rate={AUDIO_RATE},channels=1,layout=interleaved"
    );
    let pipeline_args = [
        "-q",
        "pulsesrc",
        "device=@DEFAULT_MONITOR@",
        "!",
        "audioconvert",
        "!",
        "audioresample",
        "!",
        &caps,
        "!",
        "fdsink",
        "fd=1",
        "sync=false",
    ];
    let mut cmd = Command::new("gst-launch-1.0");
    cmd.args(
        pipeline_args
            .iter()
            .map(|s| s.to_string())
            .collect::<Vec<_>>(),
    )
    .stdout(Stdio::piped())
    .stderr(Stdio::piped());
    sanitize_host_env(&mut cmd);
    let mut child = cmd
        .spawn()
        .map_err(|e| format!("spawn gst-launch-1.0 (audio): {e}"))?;

    let mut stdout = child
        .stdout
        .take()
        .ok_or_else(|| "sin stdout (audio)".to_string())?;
    if let Some(stderr) = child.stderr.take() {
        thread::spawn(move || {
            for line in BufReader::new(stderr).lines().map_while(Result::ok) {
                debug!("[audio-capture/gst] {line}");
            }
        });
    }

    let ring = state.ring.clone();
    thread::spawn(move || {
        // Buffer de lectura: leemos hasta 4096 samples (16 KB) por iteración.
        let chunk_samples = 4096usize;
        let mut byte_buf = vec![0u8; chunk_samples * 4];
        loop {
            match stdout.read(&mut byte_buf) {
                Ok(0) => {
                    warn!("[audio-capture] gst-launch cerró stdout");
                    return;
                }
                Ok(n) => {
                    // F32LE → f32: ignoramos colas parciales (n % 4 != 0)
                    // hasta la próxima lectura, gst alinea normalmente.
                    let aligned = n - (n % 4);
                    if aligned == 0 {
                        continue;
                    }
                    let mut samples = Vec::with_capacity(aligned / 4);
                    for i in (0..aligned).step_by(4) {
                        let s = f32::from_le_bytes([
                            byte_buf[i],
                            byte_buf[i + 1],
                            byte_buf[i + 2],
                            byte_buf[i + 3],
                        ]);
                        samples.push(s);
                    }
                    ring.lock().unwrap().push(&samples);
                }
                Err(e) => {
                    error!("[audio-capture] read error: {e}");
                    return;
                }
            }
        }
    });

    *state.gst_child.lock().unwrap() = Some(child);
    *init = true;
    Ok(())
}

#[tauri::command]
async fn sample_audio_bins(smoothing: Option<f32>) -> Result<Vec<u8>, String> {
    ensure_audio_capture_initialized().await?;
    let state = audio_capture_state();
    // 0.6 imita AnalyserNode default (bueno para color stability del audio
    // sync). Para detección de beats — scene-audio — el caller pasa 0 para
    // recibir magnitudes sin suavizar.
    let smoothing = smoothing.unwrap_or(0.6).clamp(0.0, 0.99);

    // Esperar a tener al menos un buffer completo (la primera vez tras
    // arrancar gst). Bound bajo: si en ~1 s no hay datos, error.
    let mut samples = vec![0.0f32; AUDIO_FFT];
    let mut tries = 0;
    loop {
        state.ring.lock().unwrap().copy_latest(&mut samples);
        // Heurística rápida: si todos los samples siguen siendo 0.0 y no
        // pasaron muchos intentos, esperar — todavía no llegó audio.
        let any_nonzero = samples.iter().any(|&s| s != 0.0);
        if any_nonzero || tries > 50 {
            break;
        }
        tries += 1;
        tokio::time::sleep(Duration::from_millis(20)).await;
    }

    // Hann window — mismo windowing implícito que AnalyserNode.
    for (i, s) in samples.iter_mut().enumerate() {
        let w = 0.5
            * (1.0
                - (2.0 * std::f32::consts::PI * i as f32 / (AUDIO_FFT - 1) as f32)
                    .cos());
        *s *= w;
    }

    let mut fft_buf: Vec<rustfft::num_complex::Complex<f32>> = samples
        .iter()
        .map(|&s| rustfft::num_complex::Complex { re: s, im: 0.0 })
        .collect();
    {
        let fft = state.fft.lock().unwrap();
        fft.process(&mut fft_buf);
    }

    // dB clamp -100..-30 idéntico al default de AnalyserNode; el smoothing
    // ahora viene del caller (parámetro arriba).
    const MIN_DB: f32 = -100.0;
    const MAX_DB: f32 = -30.0;
    let mut smoothed = state.smoothed.lock().unwrap();
    let mut out = vec![0u8; AUDIO_BINS];
    for i in 0..AUDIO_BINS {
        let re = fft_buf[i].re;
        let im = fft_buf[i].im;
        let mag = (re * re + im * im).sqrt() / AUDIO_FFT as f32;
        let sm = smoothing * smoothed[i] + (1.0 - smoothing) * mag;
        smoothed[i] = sm;
        let db = 20.0 * sm.max(1e-12).log10();
        let v = ((db - MIN_DB) / (MAX_DB - MIN_DB)).clamp(0.0, 1.0);
        out[i] = (v * 255.0).round() as u8;
    }
    Ok(out)
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

// graceful_quit: cuando el usuario pide cerrar (tray "Salir" o X con
// quit_on_close=true), no salimos inmediato. Emitimos `spectra:before-quit`
// y dejamos al frontend apagar las luces del ambiente sincronizado, que
// después invoca `confirm_quit`. Un watchdog garantiza salida aunque el JS
// no responda. QuitRequested evita re-entrada si el usuario insiste con
// otro click.
struct QuitRequested(AtomicBool);

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

// withGlobalTauri=true no inyecta los wrappers JS de los plugins (no hay
// bundler), así que window.__TAURI__.notification es undefined desde el
// frontend. Exponemos el envío en Rust — el frontend pasa los strings ya
// traducidos y nosotros disparamos la notificación nativa.
//
// Usamos notify-rust directamente en vez del plugin: tauri-plugin-notification
// 2.x corre el .show() interno con async_runtime::spawn → cae en un worker
// de tokio, donde el block_on de zbus dentro de notify-rust panic'ea con
// "Cannot start a runtime from within a runtime". std::thread::spawn aísla
// la llamada en un OS thread sin contexto tokio.
// Shell-out a `notify-send` (libnotify-bin):
// - tauri-plugin-notification 2.x panic'ea en Linux: spawnea notify-rust en
//   un worker de tokio donde el block_on interno de zbus dispara "Cannot
//   start a runtime from within a runtime".
// - Llamar notify-rust directamente desde un std::thread evita el panic y
//   dbus responde OK con un id, pero GNOME Shell descarta la notificación
//   silenciosamente incluso con DesktopEntry / urgency / icon hints
//   (parece ser credenciales del sender que libnotify maneja y notify-rust
//   no).
// - notify-send es parte de libnotify-bin, ubicuo en escritorios Linux, y
//   verificado que renderiza en GNOME Shell.
#[tauri::command]
fn notify_native(title: String, body: String) -> Result<(), String> {
    let output = std::process::Command::new("notify-send")
        .arg("--app-name=SpectraControl")
        .arg("--icon=casa.scode.SpectraControl")
        .arg("--expire-time=5000")
        .arg(&title)
        .arg(&body)
        .output()
        .map_err(|e| format!("spawn notify-send: {}", e))?;
    if !output.status.success() {
        let stderr = String::from_utf8_lossy(&output.stderr);
        return Err(format!("notify-send exit {}: {}", output.status, stderr));
    }
    Ok(())
}

fn force_exit(app: &tauri::AppHandle) {
    app.state::<IsQuitting>().0.store(true, Ordering::SeqCst);
    cleanup_children(app);
    app.exit(0);
}

fn request_graceful_quit(app: &tauri::AppHandle) {
    let already = app
        .state::<QuitRequested>()
        .0
        .swap(true, Ordering::SeqCst);
    if already {
        return;
    }
    // Ocultar la ventana enseguida — el usuario ya pidió cerrar y no debería
    // ver la UI mientras corre el shutdown del bridge.
    if let Some(window) = app.get_webview_window("main") {
        let _ = window.hide();
    }
    let _ = app.emit("spectra:before-quit", ());
    let handle = app.clone();
    thread::spawn(move || {
        // Watchdog: si el frontend no llamó a confirm_quit en este margen,
        // forzamos salida. Suficiente para mandar PUTs de "on:false" al
        // bridge local (latencia típica ~50ms) sin hacer esperar al usuario.
        thread::sleep(Duration::from_millis(1500));
        force_exit(&handle);
    });
}

#[tauri::command]
fn confirm_quit(app: tauri::AppHandle) {
    force_exit(&app);
}

// Redimensiona la ventana principal a un tamaño lógico (DPI-independent).
// Lo expone Rust porque `window.__TAURI__.window.LogicalSize` no siempre
// llega cuando `withGlobalTauri:true` está activo.
#[tauri::command]
fn resize_window(
    window: tauri::WebviewWindow,
    width: f64,
    height: f64,
) -> Result<(), String> {
    window
        .set_size(tauri::Size::Logical(tauri::LogicalSize {
            width,
            height,
        }))
        .map_err(|e| e.to_string())
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
    if let Some(audio) = AUDIO_CAPTURE.get() {
        if let Some(mut gst) = audio.gst_child.lock().unwrap().take() {
            let _ = gst.kill();
        }
    }
}

fn main() {
    setup_logger();
    tauri::Builder::default()
        .plugin(tauri_plugin_updater::Builder::new().build())
        .plugin(tauri_plugin_process::init())
        .plugin(tauri_plugin_dialog::init())
        .plugin(tauri_plugin_autostart::init(
            tauri_plugin_autostart::MacosLauncher::LaunchAgent,
            None,
        ))
        .manage(BackendProcess(Mutex::new(None)))
        .manage(IsQuitting(AtomicBool::new(false)))
        .manage(QuitOnClose(AtomicBool::new(false)))
        .manage(QuitRequested(AtomicBool::new(false)))
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
                        request_graceful_quit(app);
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
                    // El usuario eligió que la X cierre la app. Cancelamos el
                    // cierre de la ventana mientras el frontend apaga las luces;
                    // confirm_quit (o el watchdog) hará el exit real luego.
                    api.prevent_close();
                    request_graceful_quit(&app);
                    return;
                }
                // Default: ocultar al tray.
                api.prevent_close();
                let _ = window.hide();
            }
        })
        .invoke_handler(tauri::generate_handler![
            sample_screen_regions,
            sample_audio_bins,
            set_tray_menu_labels,
            set_quit_on_close,
            write_text_file,
            notify_native,
            confirm_quit,
            resize_window
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
