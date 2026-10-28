// Guard contra fuentes COLRv1 (emoji de color vectoriales).
//
// El WebKitGTK/Skia que va dentro del AppImage es de mayo de 2026 y se cae con
// un assertion en `colrv1_configure_skpaint` al rasterizar versiones nuevas de
// Noto-COLRv1.ttf (Fedora 44 / Bazzite desde 2026-09). El WebKitWebProcess
// muere al instante: la ventana queda en negro/blanco y GNOME solo muestra
// "SpectraControl está lista". Excluimos esas fuentes para ESTE proceso (y sus
// hijos) con un fontconfig propio; los emojis caen a Twemoji u otra fuente de
// bitmap. Si el usuario ya definió FONTCONFIG_FILE, no tocamos nada.

use std::fs;
use std::path::PathBuf;

use log::{info, warn};

const CONF: &str = r#"<?xml version="1.0"?>
<!DOCTYPE fontconfig SYSTEM "urn:fontconfig:fonts.dtd">
<fontconfig>
  <include ignore_missing="yes">/etc/fonts/fonts.conf</include>
  <selectfont>
    <rejectfont>
      <glob>*COLRv1*</glob>
      <glob>*colrv1*</glob>
    </rejectfont>
  </selectfont>
</fontconfig>
"#;

fn conf_path() -> Option<PathBuf> {
    let base = std::env::var_os("XDG_CONFIG_HOME")
        .map(PathBuf::from)
        .or_else(|| std::env::var_os("HOME").map(|h| PathBuf::from(h).join(".config")))?;
    Some(base.join("spectracontrol").join("fonts-nocolrv1.conf"))
}

/// Debe llamarse al inicio de `main`, antes de crear el webview.
pub fn guard_colrv1() {
    if std::env::var_os("FONTCONFIG_FILE").is_some()
        || std::env::var_os("SPECTRA_KEEP_COLRV1").is_some()
    {
        return;
    }
    let Some(path) = conf_path() else { return };
    let up_to_date = fs::read_to_string(&path).map(|c| c == CONF).unwrap_or(false);
    if !up_to_date {
        if let Some(dir) = path.parent() {
            let _ = fs::create_dir_all(dir);
        }
        if let Err(e) = fs::write(&path, CONF) {
            warn!("fonts: no pude escribir {}: {e}", path.display());
            return;
        }
    }
    // SAFETY-equivalente: aún estamos en un solo hilo (inicio de main).
    std::env::set_var("FONTCONFIG_FILE", &path);
    info!("fonts: COLRv1 excluido vía {}", path.display());
}
