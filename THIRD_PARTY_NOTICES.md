# Third-Party Notices

SpectraControl is distributed under the [MIT License](LICENSE) and incorporates
third-party open-source software. This file lists those components together with
their licenses, as required by their respective terms.

Every component listed here is distributed under an OSI-approved permissive
license (MIT / Apache-2.0 / BSD-2 / BSD-3 / ISC / Zlib / Unicode-3.0) or a
file-level weak-copyleft license (MPL-2.0). No GPL, LGPL, AGPL or SSPL
dependencies are linked into the shipped binaries.

## How this file is generated

- **Rust (Tauri shell + WebView host):** entries below come from `cargo metadata`
  run inside `src-tauri/` for `x86_64-unknown-linux-gnu`. To regenerate:

  ```bash
  cd src-tauri
  cargo metadata --format-version 1 --filter-platform x86_64-unknown-linux-gnu
  ```

- **Go (backend binary `spectractl`):** entries below come from `go list -deps`
  filtered to packages actually linked into the release binary. To regenerate:

  ```bash
  cd backend
  go list -deps ./...
  ```

---

## Backend (`spectractl`, Go)

Packages compiled into the release binary:

| Module | Version | License | Source |
|---|---|---|---|
| `github.com/go-chi/chi/v5` | v5.1.0 | MIT | https://github.com/go-chi/chi |
| `github.com/gorilla/websocket` | v1.5.3 | BSD-2-Clause | https://github.com/gorilla/websocket |
| `github.com/pion/dtls/v2` | v2.2.12 | MIT | https://github.com/pion/dtls |
| `github.com/pion/logging` | v0.2.2 | MIT | https://github.com/pion/logging |
| `github.com/pion/transport/v2` | v2.2.4 | MIT | https://github.com/pion/transport |
| `golang.org/x/crypto` | v0.18.0 | BSD-3-Clause | https://cs.opensource.google/go/x/crypto |
| `golang.org/x/net` | v0.20.0 | BSD-3-Clause | https://cs.opensource.google/go/x/net |
| `golang.org/x/sys` | v0.16.0 | BSD-3-Clause | https://cs.opensource.google/go/x/sys |

Test-only modules (`testify`, `objx`, `go-spew`, `go-difflib`, `goldmark`,
`golang.org/x/{mod,sync,term,text,tools,xerrors}`, `gopkg.in/check.v1`,
`gopkg.in/yaml.v3`) are not linked into the binary and are not redistributed.
They are all permissively licensed (MIT / BSD / ISC / Apache-2.0).

---

## Tauri shell & WebView host (Rust)

380 crates resolved for `x86_64-unknown-linux-gnu`. Each line lists the crate
name, version, SPDX license expression, and upstream source URL.

> ℹ️ License expressions of the form `MIT OR Apache-2.0` mean the user may pick
> either license. SpectraControl, being MIT-licensed itself, exercises the MIT
> option when redistributing those crates.

- **adler2** 2.0.1 — 0BSD OR MIT OR Apache-2.0
  https://github.com/oyvindln/adler2
- **aho-corasick** 1.1.4 — Unlicense OR MIT
  https://github.com/BurntSushi/aho-corasick
- **alloc-no-stdlib** 2.0.4 — BSD-3-Clause
  https://github.com/dropbox/rust-alloc-no-stdlib
- **alloc-stdlib** 0.2.2 — BSD-3-Clause
  https://github.com/dropbox/rust-alloc-no-stdlib
- **anyhow** 1.0.102 — MIT OR Apache-2.0
  https://github.com/dtolnay/anyhow
- **ashpd** 0.10.3 — MIT
  https://github.com/bilelmoussaoui/ashpd
- **async-broadcast** 0.7.2 — MIT OR Apache-2.0
  https://github.com/smol-rs/async-broadcast
- **async-channel** 2.5.0 — Apache-2.0 OR MIT
  https://github.com/smol-rs/async-channel
- **async-executor** 1.14.0 — Apache-2.0 OR MIT
  https://github.com/smol-rs/async-executor
- **async-io** 2.6.0 — Apache-2.0 OR MIT
  https://github.com/smol-rs/async-io
- **async-lock** 3.4.2 — Apache-2.0 OR MIT
  https://github.com/smol-rs/async-lock
- **async-process** 2.5.0 — Apache-2.0 OR MIT
  https://github.com/smol-rs/async-process
- **async-recursion** 1.1.1 — MIT OR Apache-2.0
  https://github.com/dcchut/async-recursion
- **async-signal** 0.2.14 — Apache-2.0 OR MIT
  https://github.com/smol-rs/async-signal
- **async-task** 4.7.1 — Apache-2.0 OR MIT
  https://github.com/smol-rs/async-task
- **async-trait** 0.1.89 — MIT OR Apache-2.0
  https://github.com/dtolnay/async-trait
- **atk** 0.18.2 — MIT
  https://github.com/gtk-rs/gtk3-rs
- **atk-sys** 0.18.2 — MIT
  https://github.com/gtk-rs/gtk3-rs
- **atomic-waker** 1.1.2 — Apache-2.0 OR MIT
  https://github.com/smol-rs/atomic-waker
- **auto-launch** 0.5.0 — MIT
  https://github.com/zzzgydi/auto-launch.git
- **autocfg** 1.5.0 — Apache-2.0 OR MIT
  https://github.com/cuviper/autocfg
- **base64** 0.22.1 — MIT OR Apache-2.0
  https://github.com/marshallpierce/rust-base64
- **bit-set** 0.8.0 — Apache-2.0 OR MIT
  https://github.com/contain-rs/bit-set
- **bit-vec** 0.8.0 — Apache-2.0 OR MIT
  https://github.com/contain-rs/bit-vec
- **bitflags** 1.3.2 — MIT/Apache-2.0
  https://github.com/bitflags/bitflags
- **bitflags** 2.11.1 — MIT OR Apache-2.0
  https://github.com/bitflags/bitflags
- **block-buffer** 0.10.4 — MIT OR Apache-2.0
  https://github.com/RustCrypto/utils
- **blocking** 1.6.2 — Apache-2.0 OR MIT
  https://github.com/smol-rs/blocking
- **brotli** 8.0.2 — BSD-3-Clause AND MIT
  https://github.com/dropbox/rust-brotli
- **brotli-decompressor** 5.0.0 — BSD-3-Clause/MIT
  https://github.com/dropbox/rust-brotli-decompressor
- **byteorder** 1.5.0 — Unlicense OR MIT
  https://github.com/BurntSushi/byteorder
- **bytes** 1.11.1 — MIT
  https://github.com/tokio-rs/bytes
- **cairo-rs** 0.18.5 — MIT
  https://github.com/gtk-rs/gtk-rs-core
- **cairo-sys-rs** 0.18.2 — MIT
  https://github.com/gtk-rs/gtk-rs-core
- **camino** 1.2.2 — MIT OR Apache-2.0
  https://github.com/camino-rs/camino
- **cargo-platform** 0.1.9 — MIT OR Apache-2.0
  https://github.com/rust-lang/cargo
- **cargo_metadata** 0.19.2 — MIT
  https://github.com/oli-obk/cargo_metadata
- **cargo_toml** 0.22.3 — Apache-2.0 OR MIT
  https://gitlab.com/lib.rs/cargo_toml
- **cc** 1.2.61 — MIT OR Apache-2.0
  https://github.com/rust-lang/cc-rs
- **cfb** 0.7.3 — MIT
  https://github.com/mdsteele/rust-cfb
- **cfg-expr** 0.15.8 — MIT OR Apache-2.0
  https://github.com/EmbarkStudios/cfg-expr
- **cfg-if** 1.0.4 — MIT OR Apache-2.0
  https://github.com/rust-lang/cfg-if
- **chrono** 0.4.44 — MIT OR Apache-2.0
  https://github.com/chronotope/chrono
- **concurrent-queue** 2.5.0 — Apache-2.0 OR MIT
  https://github.com/smol-rs/concurrent-queue
- **cookie** 0.18.1 — MIT OR Apache-2.0
  https://github.com/SergioBenitez/cookie-rs
- **cpufeatures** 0.2.17 — MIT OR Apache-2.0
  https://github.com/RustCrypto/utils
- **crc32fast** 1.5.0 — MIT OR Apache-2.0
  https://github.com/srijs/rust-crc32fast
- **crossbeam-channel** 0.5.15 — MIT OR Apache-2.0
  https://github.com/crossbeam-rs/crossbeam
- **crossbeam-utils** 0.8.21 — MIT OR Apache-2.0
  https://github.com/crossbeam-rs/crossbeam
- **crypto-common** 0.1.7 — MIT OR Apache-2.0
  https://github.com/RustCrypto/traits
- **cssparser** 0.36.0 — MPL-2.0
  https://github.com/servo/rust-cssparser
- **cssparser-macros** 0.6.1 — MPL-2.0
  https://github.com/servo/rust-cssparser
- **ctor** 0.8.0 — Apache-2.0 OR MIT
  https://github.com/mmastrac/rust-ctor
- **ctor-proc-macro** 0.0.7 — Apache-2.0 OR MIT
  https://github.com/mmastrac/rust-ctor
- **darling** 0.23.0 — MIT
  https://github.com/TedDriggs/darling
- **darling_core** 0.23.0 — MIT
  https://github.com/TedDriggs/darling
- **darling_macro** 0.23.0 — MIT
  https://github.com/TedDriggs/darling
- **dbus** 0.9.11 — Apache-2.0/MIT
  https://github.com/diwic/dbus-rs
- **deranged** 0.5.8 — MIT OR Apache-2.0
  https://github.com/jhpratt/deranged
- **derive_more** 2.1.1 — MIT
  https://github.com/JelteF/derive_more
- **derive_more-impl** 2.1.1 — MIT
  https://github.com/JelteF/derive_more
- **digest** 0.10.7 — MIT OR Apache-2.0
  https://github.com/RustCrypto/traits
- **dirs** 4.0.0 — MIT OR Apache-2.0
  https://github.com/soc/dirs-rs
- **dirs** 6.0.0 — MIT OR Apache-2.0
  https://github.com/soc/dirs-rs
- **dirs-sys** 0.3.7 — MIT OR Apache-2.0
  https://github.com/dirs-dev/dirs-sys-rs
- **dirs-sys** 0.5.0 — MIT OR Apache-2.0
  https://github.com/dirs-dev/dirs-sys-rs
- **displaydoc** 0.2.5 — MIT OR Apache-2.0
  https://github.com/yaahc/displaydoc
- **dlopen2** 0.8.2 — MIT
  https://github.com/OpenByteDev/dlopen2
- **dlopen2_derive** 0.4.3 — MIT
  https://github.com/OpenByteDev/dlopen2
- **dom_query** 0.27.0 — MIT
  https://github.com/niklak/dom_query
- **dpi** 0.1.2 — Apache-2.0 AND MIT
  https://github.com/rust-windowing/winit
- **dtoa** 1.0.11 — MIT OR Apache-2.0
  https://github.com/dtolnay/dtoa
- **dtoa-short** 0.3.5 — MPL-2.0
  https://github.com/upsuper/dtoa-short
- **dtor** 0.3.0 — Apache-2.0 OR MIT
  https://github.com/mmastrac/rust-ctor
- **dtor-proc-macro** 0.0.6 — Apache-2.0 OR MIT
  https://github.com/mmastrac/rust-ctor
- **dunce** 1.0.5 — CC0-1.0 OR MIT-0 OR Apache-2.0
  https://gitlab.com/kornelski/dunce
- **dyn-clone** 1.0.20 — MIT OR Apache-2.0
  https://github.com/dtolnay/dyn-clone
- **embed-resource** 3.0.9 — MIT
  https://github.com/nabijaczleweli/rust-embed-resource
- **endi** 1.1.1 — MIT
  https://github.com/zeenix/endi
- **enumflags2** 0.7.12 — MIT OR Apache-2.0
  https://github.com/meithecatte/enumflags2
- **enumflags2_derive** 0.7.12 — MIT OR Apache-2.0
  https://github.com/meithecatte/enumflags2
- **env_filter** 1.0.1 — MIT OR Apache-2.0
  https://github.com/rust-cli/env_logger
- **env_logger** 0.11.10 — MIT OR Apache-2.0
  https://github.com/rust-cli/env_logger
- **equivalent** 1.0.2 — Apache-2.0 OR MIT
  https://github.com/indexmap-rs/equivalent
- **erased-serde** 0.4.10 — MIT OR Apache-2.0
  https://github.com/dtolnay/erased-serde
- **errno** 0.3.14 — MIT OR Apache-2.0
  https://github.com/lambda-fairy/rust-errno
- **event-listener** 5.4.1 — Apache-2.0 OR MIT
  https://github.com/smol-rs/event-listener
- **event-listener-strategy** 0.5.4 — Apache-2.0 OR MIT
  https://github.com/smol-rs/event-listener-strategy
- **fastrand** 2.4.1 — Apache-2.0 OR MIT
  https://github.com/smol-rs/fastrand
- **fdeflate** 0.3.7 — MIT OR Apache-2.0
  https://github.com/image-rs/fdeflate
- **field-offset** 0.3.6 — MIT OR Apache-2.0
  https://github.com/Diggsey/rust-field-offset
- **filetime** 0.2.28 — MIT/Apache-2.0
  https://github.com/alexcrichton/filetime
- **find-msvc-tools** 0.1.9 — MIT OR Apache-2.0
  https://github.com/rust-lang/cc-rs
- **flate2** 1.1.9 — MIT OR Apache-2.0
  https://github.com/rust-lang/flate2-rs
- **fnv** 1.0.7 — Apache-2.0 / MIT
  https://github.com/servo/rust-fnv
- **foldhash** 0.2.0 — Zlib
  https://github.com/orlp/foldhash
- **form_urlencoded** 1.2.2 — MIT OR Apache-2.0
  https://github.com/servo/rust-url
- **futures-channel** 0.3.32 — MIT OR Apache-2.0
  https://github.com/rust-lang/futures-rs
- **futures-core** 0.3.32 — MIT OR Apache-2.0
  https://github.com/rust-lang/futures-rs
- **futures-executor** 0.3.32 — MIT OR Apache-2.0
  https://github.com/rust-lang/futures-rs
- **futures-io** 0.3.32 — MIT OR Apache-2.0
  https://github.com/rust-lang/futures-rs
- **futures-lite** 2.6.1 — Apache-2.0 OR MIT
  https://github.com/smol-rs/futures-lite
- **futures-macro** 0.3.32 — MIT OR Apache-2.0
  https://github.com/rust-lang/futures-rs
- **futures-sink** 0.3.32 — MIT OR Apache-2.0
  https://github.com/rust-lang/futures-rs
- **futures-task** 0.3.32 — MIT OR Apache-2.0
  https://github.com/rust-lang/futures-rs
- **futures-util** 0.3.32 — MIT OR Apache-2.0
  https://github.com/rust-lang/futures-rs
- **gdk** 0.18.2 — MIT
  https://github.com/gtk-rs/gtk3-rs
- **gdk-pixbuf** 0.18.5 — MIT
  https://github.com/gtk-rs/gtk-rs-core
- **gdk-pixbuf-sys** 0.18.0 — MIT
  https://github.com/gtk-rs/gtk-rs-core
- **gdk-sys** 0.18.2 — MIT
  https://github.com/gtk-rs/gtk3-rs
- **gdkwayland-sys** 0.18.2 — MIT
  https://github.com/gtk-rs/gtk3-rs
- **gdkx11** 0.18.2 — MIT
  https://github.com/gtk-rs/gtk3-rs
- **gdkx11-sys** 0.18.2 — MIT
  https://github.com/gtk-rs/gtk3-rs
- **generic-array** 0.14.7 — MIT
  https://github.com/fizyk20/generic-array.git
- **getrandom** 0.2.17 — MIT OR Apache-2.0
  https://github.com/rust-random/getrandom
- **getrandom** 0.3.4 — MIT OR Apache-2.0
  https://github.com/rust-random/getrandom
- **getrandom** 0.4.2 — MIT OR Apache-2.0
  https://github.com/rust-random/getrandom
- **gio** 0.18.4 — MIT
  https://github.com/gtk-rs/gtk-rs-core
- **gio-sys** 0.18.1 — MIT
  https://github.com/gtk-rs/gtk-rs-core
- **glib** 0.18.5 — MIT
  https://github.com/gtk-rs/gtk-rs-core
- **glib-macros** 0.18.5 — MIT
  https://github.com/gtk-rs/gtk-rs-core
- **glib-sys** 0.18.1 — MIT
  https://github.com/gtk-rs/gtk-rs-core
- **glob** 0.3.3 — MIT OR Apache-2.0
  https://github.com/rust-lang/glob
- **gobject-sys** 0.18.0 — MIT
  https://github.com/gtk-rs/gtk-rs-core
- **gtk** 0.18.2 — MIT
  https://github.com/gtk-rs/gtk3-rs
- **gtk-sys** 0.18.2 — MIT
  https://github.com/gtk-rs/gtk3-rs
- **gtk3-macros** 0.18.2 — MIT
  https://github.com/gtk-rs/gtk3-rs
- **hashbrown** 0.12.3 — MIT OR Apache-2.0
  https://github.com/rust-lang/hashbrown
- **hashbrown** 0.17.0 — MIT OR Apache-2.0
  https://github.com/rust-lang/hashbrown
- **heck** 0.4.1 — MIT OR Apache-2.0
  https://github.com/withoutboats/heck
- **heck** 0.5.0 — MIT OR Apache-2.0
  https://github.com/withoutboats/heck
- **hex** 0.4.3 — MIT OR Apache-2.0
  https://github.com/KokaKiwi/rust-hex
- **html5ever** 0.38.0 — MIT OR Apache-2.0
  https://github.com/servo/html5ever
- **http** 1.4.0 — MIT OR Apache-2.0
  https://github.com/hyperium/http
- **http-body** 1.0.1 — MIT
  https://github.com/hyperium/http-body
- **http-body-util** 0.1.3 — MIT
  https://github.com/hyperium/http-body
- **httparse** 1.10.1 — MIT OR Apache-2.0
  https://github.com/seanmonstar/httparse
- **hyper** 1.9.0 — MIT
  https://github.com/hyperium/hyper
- **hyper-rustls** 0.27.9 — Apache-2.0 OR ISC OR MIT
  https://github.com/rustls/hyper-rustls
- **hyper-util** 0.1.20 — MIT
  https://github.com/hyperium/hyper-util
- **iana-time-zone** 0.1.65 — MIT OR Apache-2.0
  https://github.com/strawlab/iana-time-zone
- **ico** 0.5.0 — MIT
  https://github.com/mdsteele/rust-ico
- **icu_collections** 2.2.0 — Unicode-3.0
  https://github.com/unicode-org/icu4x
- **icu_locale_core** 2.2.0 — Unicode-3.0
  https://github.com/unicode-org/icu4x
- **icu_normalizer** 2.2.0 — Unicode-3.0
  https://github.com/unicode-org/icu4x
- **icu_normalizer_data** 2.2.0 — Unicode-3.0
  https://github.com/unicode-org/icu4x
- **icu_properties** 2.2.0 — Unicode-3.0
  https://github.com/unicode-org/icu4x
- **icu_properties_data** 2.2.0 — Unicode-3.0
  https://github.com/unicode-org/icu4x
- **icu_provider** 2.2.0 — Unicode-3.0
  https://github.com/unicode-org/icu4x
- **ident_case** 1.0.1 — MIT/Apache-2.0
  https://github.com/TedDriggs/ident_case
- **idna** 1.1.0 — MIT OR Apache-2.0
  https://github.com/servo/rust-url/
- **idna_adapter** 1.2.2 — Apache-2.0 OR MIT
  https://github.com/hsivonen/idna_adapter
- **indexmap** 1.9.3 — Apache-2.0 OR MIT
  https://github.com/bluss/indexmap
- **indexmap** 2.14.0 — Apache-2.0 OR MIT
  https://github.com/indexmap-rs/indexmap
- **infer** 0.19.0 — MIT
  https://github.com/bojand/infer
- **ipnet** 2.12.0 — MIT OR Apache-2.0
  https://github.com/krisprice/ipnet
- **itoa** 1.0.18 — MIT OR Apache-2.0
  https://github.com/dtolnay/itoa
- **javascriptcore-rs** 1.1.2 — MIT
  https://github.com/tauri-apps/javascriptcore-rs
- **javascriptcore-rs-sys** 1.1.1 — MIT
  https://github.com/tauri-apps/javascriptcore-rs
- **jiff** 0.2.24 — Unlicense OR MIT
  https://github.com/BurntSushi/jiff
- **json-patch** 3.0.1 — MIT/Apache-2.0
  https://github.com/idubrov/json-patch
- **jsonptr** 0.6.3 — MIT OR Apache-2.0
  https://github.com/chanced/jsonptr
- **keyboard-types** 0.7.0 — MIT OR Apache-2.0
  https://github.com/pyfisch/keyboard-types
- **libappindicator** 0.9.0 — Apache-2.0 OR MIT
- **libappindicator-sys** 0.9.0 — Apache-2.0 OR MIT
- **libc** 0.2.186 — MIT OR Apache-2.0
  https://github.com/rust-lang/libc
- **libdbus-sys** 0.2.7 — Apache-2.0/MIT
  https://github.com/diwic/dbus-rs
- **libloading** 0.7.4 — ISC
  https://github.com/nagisa/rust_libloading/
- **linux-raw-sys** 0.12.1 — Apache-2.0 WITH LLVM-exception OR Apache-2.0 OR MIT
  https://github.com/sunfishcode/linux-raw-sys
- **litemap** 0.8.2 — Unicode-3.0
  https://github.com/unicode-org/icu4x
- **lock_api** 0.4.14 — MIT OR Apache-2.0
  https://github.com/Amanieu/parking_lot
- **log** 0.4.29 — MIT OR Apache-2.0
  https://github.com/rust-lang/log
- **markup5ever** 0.38.0 — MIT OR Apache-2.0
  https://github.com/servo/html5ever
- **memchr** 2.8.0 — Unlicense OR MIT
  https://github.com/BurntSushi/memchr
- **memoffset** 0.9.1 — MIT
  https://github.com/Gilnaa/memoffset
- **mime** 0.3.17 — MIT OR Apache-2.0
  https://github.com/hyperium/mime
- **minisign-verify** 0.2.5 — MIT
  https://github.com/jedisct1/rust-minisign-verify
- **miniz_oxide** 0.8.9 — MIT OR Zlib OR Apache-2.0
  https://github.com/Frommi/miniz_oxide/tree/master/miniz_oxide
- **mio** 1.2.0 — MIT
  https://github.com/tokio-rs/mio
- **muda** 0.19.1 — Apache-2.0 OR MIT
  https://github.com/tauri-apps/muda
- **new_debug_unreachable** 1.0.6 — MIT
  https://github.com/mbrubeck/rust-debug-unreachable
- **notify-rust** 4.17.0 — MIT/Apache-2.0
  https://github.com/hoodie/notify-rust
- **num-conv** 0.2.1 — MIT OR Apache-2.0
  https://github.com/jhpratt/num-conv
- **num-traits** 0.2.19 — MIT OR Apache-2.0
  https://github.com/rust-num/num-traits
- **once_cell** 1.21.4 — MIT OR Apache-2.0
  https://github.com/matklad/once_cell
- **openssl-probe** 0.2.1 — MIT OR Apache-2.0
  https://github.com/rustls/openssl-probe
- **option-ext** 0.2.0 — MPL-2.0
  https://github.com/soc/option-ext.git
- **ordered-stream** 0.2.0 — MIT OR Apache-2.0
  https://github.com/danieldg/ordered-stream
- **pango** 0.18.3 — MIT
  https://github.com/gtk-rs/gtk-rs-core
- **pango-sys** 0.18.0 — MIT
  https://github.com/gtk-rs/gtk-rs-core
- **parking** 2.2.1 — Apache-2.0 OR MIT
  https://github.com/smol-rs/parking
- **parking_lot** 0.12.5 — MIT OR Apache-2.0
  https://github.com/Amanieu/parking_lot
- **parking_lot_core** 0.9.12 — MIT OR Apache-2.0
  https://github.com/Amanieu/parking_lot
- **percent-encoding** 2.3.2 — MIT OR Apache-2.0
  https://github.com/servo/rust-url/
- **phf** 0.13.1 — MIT
  https://github.com/rust-phf/rust-phf
- **phf_codegen** 0.13.1 — MIT
  https://github.com/rust-phf/rust-phf
- **phf_generator** 0.13.1 — MIT
  https://github.com/rust-phf/rust-phf
- **phf_macros** 0.13.1 — MIT
  https://github.com/rust-phf/rust-phf
- **phf_shared** 0.13.1 — MIT
  https://github.com/rust-phf/rust-phf
- **pin-project-lite** 0.2.17 — Apache-2.0 OR MIT
  https://github.com/taiki-e/pin-project-lite
- **piper** 0.2.5 — MIT OR Apache-2.0
  https://github.com/smol-rs/piper
- **pkg-config** 0.3.33 — MIT OR Apache-2.0
  https://github.com/rust-lang/pkg-config-rs
- **plist** 1.9.0 — MIT
  https://github.com/ebarnard/rust-plist/
- **png** 0.17.16 — MIT OR Apache-2.0
  https://github.com/image-rs/image-png
- **png** 0.18.1 — MIT OR Apache-2.0
  https://github.com/image-rs/image-png
- **polling** 3.11.0 — Apache-2.0 OR MIT
  https://github.com/smol-rs/polling
- **potential_utf** 0.1.5 — Unicode-3.0
  https://github.com/unicode-org/icu4x
- **powerfmt** 0.2.0 — MIT OR Apache-2.0
  https://github.com/jhpratt/powerfmt
- **ppv-lite86** 0.2.21 — MIT OR Apache-2.0
  https://github.com/cryptocorrosion/cryptocorrosion
- **precomputed-hash** 0.1.1 — MIT
  https://github.com/emilio/precomputed-hash
- **proc-macro-crate** 1.3.1 — MIT OR Apache-2.0
  https://github.com/bkchr/proc-macro-crate
- **proc-macro-crate** 2.0.2 — MIT OR Apache-2.0
  https://github.com/bkchr/proc-macro-crate
- **proc-macro-crate** 3.5.0 — MIT OR Apache-2.0
  https://github.com/bkchr/proc-macro-crate
- **proc-macro-error** 1.0.4 — MIT OR Apache-2.0
  https://gitlab.com/CreepySkeleton/proc-macro-error
- **proc-macro-error-attr** 1.0.4 — MIT OR Apache-2.0
  https://gitlab.com/CreepySkeleton/proc-macro-error
- **proc-macro2** 1.0.106 — MIT OR Apache-2.0
  https://github.com/dtolnay/proc-macro2
- **quick-xml** 0.39.3 — MIT
  https://github.com/tafia/quick-xml
- **quote** 1.0.45 — MIT OR Apache-2.0
  https://github.com/dtolnay/quote
- **rand** 0.8.6 — MIT OR Apache-2.0
  https://github.com/rust-random/rand
- **rand** 0.9.4 — MIT OR Apache-2.0
  https://github.com/rust-random/rand
- **rand_chacha** 0.3.1 — MIT OR Apache-2.0
  https://github.com/rust-random/rand
- **rand_chacha** 0.9.0 — MIT OR Apache-2.0
  https://github.com/rust-random/rand
- **rand_core** 0.6.4 — MIT OR Apache-2.0
  https://github.com/rust-random/rand
- **rand_core** 0.9.5 — MIT OR Apache-2.0
  https://github.com/rust-random/rand
- **raw-window-handle** 0.6.2 — MIT OR Apache-2.0 OR Zlib
  https://github.com/rust-windowing/raw-window-handle
- **ref-cast** 1.0.25 — MIT OR Apache-2.0
  https://github.com/dtolnay/ref-cast
- **ref-cast-impl** 1.0.25 — MIT OR Apache-2.0
  https://github.com/dtolnay/ref-cast
- **regex** 1.12.3 — MIT OR Apache-2.0
  https://github.com/rust-lang/regex
- **regex-automata** 0.4.14 — MIT OR Apache-2.0
  https://github.com/rust-lang/regex
- **regex-syntax** 0.8.10 — MIT OR Apache-2.0
  https://github.com/rust-lang/regex
- **reqwest** 0.13.3 — MIT OR Apache-2.0
  https://github.com/seanmonstar/reqwest
- **rfd** 0.16.0 — MIT
  https://github.com/PolyMeilex/rfd
- **ring** 0.17.14 — Apache-2.0 AND ISC
  https://github.com/briansmith/ring
- **rustc-hash** 2.1.2 — Apache-2.0 OR MIT
  https://github.com/rust-lang/rustc-hash
- **rustc_version** 0.4.1 — MIT OR Apache-2.0
  https://github.com/djc/rustc-version-rs
- **rustix** 1.1.4 — Apache-2.0 WITH LLVM-exception OR Apache-2.0 OR MIT
  https://github.com/bytecodealliance/rustix
- **rustls** 0.23.40 — Apache-2.0 OR ISC OR MIT
  https://github.com/rustls/rustls
- **rustls-native-certs** 0.8.3 — Apache-2.0 OR ISC OR MIT
  https://github.com/rustls/rustls-native-certs
- **rustls-pki-types** 1.14.1 — MIT OR Apache-2.0
  https://github.com/rustls/pki-types
- **rustls-platform-verifier** 0.7.0 — MIT OR Apache-2.0
  https://github.com/rustls/rustls-platform-verifier
- **rustls-webpki** 0.103.13 — ISC
  https://github.com/rustls/webpki
- **same-file** 1.0.6 — Unlicense/MIT
  https://github.com/BurntSushi/same-file
- **schemars** 0.8.22 — MIT
  https://github.com/GREsau/schemars
- **schemars** 0.9.0 — MIT
  https://github.com/GREsau/schemars
- **schemars** 1.2.1 — MIT
  https://github.com/GREsau/schemars
- **schemars_derive** 0.8.22 — MIT
  https://github.com/GREsau/schemars
- **scopeguard** 1.2.0 — MIT OR Apache-2.0
  https://github.com/bluss/scopeguard
- **selectors** 0.36.1 — MPL-2.0
  https://github.com/servo/stylo
- **semver** 1.0.28 — MIT OR Apache-2.0
  https://github.com/dtolnay/semver
- **serde** 1.0.228 — MIT OR Apache-2.0
  https://github.com/serde-rs/serde
- **serde-untagged** 0.1.9 — MIT OR Apache-2.0
  https://github.com/dtolnay/serde-untagged
- **serde_core** 1.0.228 — MIT OR Apache-2.0
  https://github.com/serde-rs/serde
- **serde_derive** 1.0.228 — MIT OR Apache-2.0
  https://github.com/serde-rs/serde
- **serde_derive_internals** 0.29.1 — MIT OR Apache-2.0
  https://github.com/serde-rs/serde
- **serde_json** 1.0.149 — MIT OR Apache-2.0
  https://github.com/serde-rs/json
- **serde_repr** 0.1.20 — MIT OR Apache-2.0
  https://github.com/dtolnay/serde-repr
- **serde_spanned** 0.6.9 — MIT OR Apache-2.0
  https://github.com/toml-rs/toml
- **serde_spanned** 1.1.1 — MIT OR Apache-2.0
  https://github.com/toml-rs/toml
- **serde_with** 3.19.0 — MIT OR Apache-2.0
  https://github.com/jonasbb/serde_with/
- **serde_with_macros** 3.19.0 — MIT OR Apache-2.0
  https://github.com/jonasbb/serde_with/
- **serialize-to-javascript** 0.1.2 — MIT OR Apache-2.0
  https://github.com/chippers/serialize-to-javascript
- **serialize-to-javascript-impl** 0.1.2 — MIT OR Apache-2.0
  https://github.com/chippers/serialize-to-javascript
- **servo_arc** 0.4.3 — MIT OR Apache-2.0
  https://github.com/servo/stylo
- **sha2** 0.10.9 — MIT OR Apache-2.0
  https://github.com/RustCrypto/hashes
- **shlex** 1.3.0 — MIT OR Apache-2.0
  https://github.com/comex/rust-shlex
- **signal-hook-registry** 1.4.8 — MIT OR Apache-2.0
  https://github.com/vorner/signal-hook
- **simd-adler32** 0.3.9 — MIT
  https://github.com/mcountryman/simd-adler32
- **siphasher** 1.0.3 — MIT/Apache-2.0
  https://github.com/jedisct1/rust-siphash
- **slab** 0.4.12 — MIT
  https://github.com/tokio-rs/slab
- **smallvec** 1.15.1 — MIT OR Apache-2.0
  https://github.com/servo/rust-smallvec
- **socket2** 0.6.3 — MIT OR Apache-2.0
  https://github.com/rust-lang/socket2
- **soup3** 0.5.0 — MIT
  https://gitlab.gnome.org/World/Rust/soup3-rs
- **soup3-sys** 0.5.0 — MIT
  https://gitlab.gnome.org/World/Rust/soup3-rs
- **stable_deref_trait** 1.2.1 — MIT OR Apache-2.0
  https://github.com/storyyeller/stable_deref_trait
- **string_cache** 0.9.0 — MIT OR Apache-2.0
  https://github.com/servo/string-cache
- **string_cache_codegen** 0.6.1 — MIT OR Apache-2.0
  https://github.com/servo/string-cache
- **strsim** 0.11.1 — MIT
  https://github.com/rapidfuzz/strsim-rs
- **subtle** 2.6.1 — BSD-3-Clause
  https://github.com/dalek-cryptography/subtle
- **syn** 1.0.109 — MIT OR Apache-2.0
  https://github.com/dtolnay/syn
- **syn** 2.0.117 — MIT OR Apache-2.0
  https://github.com/dtolnay/syn
- **sync_wrapper** 1.0.2 — Apache-2.0
  https://github.com/Actyx/sync_wrapper
- **synstructure** 0.13.2 — MIT
  https://github.com/mystor/synstructure
- **system-deps** 6.2.2 — MIT OR Apache-2.0
  https://github.com/gdesmott/system-deps
- **tao** 0.35.2 — Apache-2.0
  https://github.com/tauri-apps/tao
- **tar** 0.4.45 — MIT OR Apache-2.0
  https://github.com/alexcrichton/tar-rs
- **target-lexicon** 0.12.16 — Apache-2.0 WITH LLVM-exception
  https://github.com/bytecodealliance/target-lexicon
- **tauri** 2.11.1 — Apache-2.0 OR MIT
  https://github.com/tauri-apps/tauri
- **tauri-build** 2.6.1 — Apache-2.0 OR MIT
  https://github.com/tauri-apps/tauri
- **tauri-codegen** 2.6.1 — Apache-2.0 OR MIT
  https://github.com/tauri-apps/tauri
- **tauri-macros** 2.6.1 — Apache-2.0 OR MIT
  https://github.com/tauri-apps/tauri
- **tauri-plugin** 2.6.1 — Apache-2.0 OR MIT
  https://github.com/tauri-apps/tauri
- **tauri-plugin-autostart** 2.5.1 — Apache-2.0 OR MIT
  https://github.com/tauri-apps/plugins-workspace
- **tauri-plugin-dialog** 2.7.1 — Apache-2.0 OR MIT
  https://github.com/tauri-apps/plugins-workspace
- **tauri-plugin-fs** 2.5.1 — Apache-2.0 OR MIT
  https://github.com/tauri-apps/plugins-workspace
- **tauri-plugin-notification** 2.3.3 — Apache-2.0 OR MIT
  https://github.com/tauri-apps/plugins-workspace
- **tauri-plugin-process** 2.3.1 — Apache-2.0 OR MIT
  https://github.com/tauri-apps/plugins-workspace
- **tauri-plugin-updater** 2.10.1 — Apache-2.0 OR MIT
  https://github.com/tauri-apps/plugins-workspace
- **tauri-runtime** 2.11.1 — Apache-2.0 OR MIT
  https://github.com/tauri-apps/tauri
- **tauri-runtime-wry** 2.11.1 — Apache-2.0 OR MIT
  https://github.com/tauri-apps/tauri
- **tauri-utils** 2.9.1 — Apache-2.0 OR MIT
  https://github.com/tauri-apps/tauri
- **tauri-winres** 0.3.6 — MIT
  https://github.com/tauri-apps/winres
- **tempfile** 3.27.0 — MIT OR Apache-2.0
  https://github.com/Stebalien/tempfile
- **tendril** 0.5.0 — MIT OR Apache-2.0
  https://github.com/servo/html5ever
- **thiserror** 1.0.69 — MIT OR Apache-2.0
  https://github.com/dtolnay/thiserror
- **thiserror** 2.0.18 — MIT OR Apache-2.0
  https://github.com/dtolnay/thiserror
- **thiserror-impl** 1.0.69 — MIT OR Apache-2.0
  https://github.com/dtolnay/thiserror
- **thiserror-impl** 2.0.18 — MIT OR Apache-2.0
  https://github.com/dtolnay/thiserror
- **time** 0.3.47 — MIT OR Apache-2.0
  https://github.com/time-rs/time
- **time-core** 0.1.8 — MIT OR Apache-2.0
  https://github.com/time-rs/time
- **time-macros** 0.2.27 — MIT OR Apache-2.0
  https://github.com/time-rs/time
- **tinystr** 0.8.3 — Unicode-3.0
  https://github.com/unicode-org/icu4x
- **tokio** 1.52.2 — MIT
  https://github.com/tokio-rs/tokio
- **tokio-rustls** 0.26.4 — MIT OR Apache-2.0
  https://github.com/rustls/tokio-rustls
- **tokio-util** 0.7.18 — MIT
  https://github.com/tokio-rs/tokio
- **toml** 0.8.2 — MIT OR Apache-2.0
  https://github.com/toml-rs/toml
- **toml** 0.9.12+spec-1.1.0 — MIT OR Apache-2.0
  https://github.com/toml-rs/toml
- **toml** 1.1.2+spec-1.1.0 — MIT OR Apache-2.0
  https://github.com/toml-rs/toml
- **toml_datetime** 0.6.3 — MIT OR Apache-2.0
  https://github.com/toml-rs/toml
- **toml_datetime** 0.7.5+spec-1.1.0 — MIT OR Apache-2.0
  https://github.com/toml-rs/toml
- **toml_datetime** 1.1.1+spec-1.1.0 — MIT OR Apache-2.0
  https://github.com/toml-rs/toml
- **toml_edit** 0.19.15 — MIT OR Apache-2.0
  https://github.com/toml-rs/toml
- **toml_edit** 0.20.2 — MIT OR Apache-2.0
  https://github.com/toml-rs/toml
- **toml_edit** 0.25.11+spec-1.1.0 — MIT OR Apache-2.0
  https://github.com/toml-rs/toml
- **toml_parser** 1.1.2+spec-1.1.0 — MIT OR Apache-2.0
  https://github.com/toml-rs/toml
- **toml_writer** 1.1.1+spec-1.1.0 — MIT OR Apache-2.0
  https://github.com/toml-rs/toml
- **tower** 0.5.3 — MIT
  https://github.com/tower-rs/tower
- **tower-http** 0.6.10 — MIT
  https://github.com/tower-rs/tower-http
- **tower-layer** 0.3.3 — MIT
  https://github.com/tower-rs/tower
- **tower-service** 0.3.3 — MIT
  https://github.com/tower-rs/tower
- **tracing** 0.1.44 — MIT
  https://github.com/tokio-rs/tracing
- **tracing-attributes** 0.1.31 — MIT
  https://github.com/tokio-rs/tracing
- **tracing-core** 0.1.36 — MIT
  https://github.com/tokio-rs/tracing
- **tray-icon** 0.23.1 — MIT OR Apache-2.0
  https://github.com/tauri-apps/tray-icon
- **try-lock** 0.2.5 — MIT
  https://github.com/seanmonstar/try-lock
- **typeid** 1.0.3 — MIT OR Apache-2.0
  https://github.com/dtolnay/typeid
- **typenum** 1.20.0 — MIT OR Apache-2.0
  https://github.com/paholg/typenum
- **unic-char-property** 0.9.0 — MIT/Apache-2.0
  https://github.com/open-i18n/rust-unic/
- **unic-char-range** 0.9.0 — MIT/Apache-2.0
  https://github.com/open-i18n/rust-unic/
- **unic-common** 0.9.0 — MIT/Apache-2.0
  https://github.com/open-i18n/rust-unic/
- **unic-ucd-ident** 0.9.0 — MIT/Apache-2.0
  https://github.com/open-i18n/rust-unic/
- **unic-ucd-version** 0.9.0 — MIT/Apache-2.0
  https://github.com/open-i18n/rust-unic/
- **unicode-ident** 1.0.24 — (MIT OR Apache-2.0) AND Unicode-3.0
  https://github.com/dtolnay/unicode-ident
- **unicode-segmentation** 1.13.2 — MIT OR Apache-2.0
  https://github.com/unicode-rs/unicode-segmentation
- **untrusted** 0.9.0 — ISC
  https://github.com/briansmith/untrusted
- **url** 2.5.8 — MIT OR Apache-2.0
  https://github.com/servo/rust-url
- **urlpattern** 0.3.0 — MIT
  https://github.com/denoland/rust-urlpattern
- **utf-8** 0.7.6 — MIT OR Apache-2.0
  https://github.com/SimonSapin/rust-utf8
- **utf8_iter** 1.0.4 — Apache-2.0 OR MIT
  https://github.com/hsivonen/utf8_iter
- **uuid** 1.23.1 — Apache-2.0 OR MIT
  https://github.com/uuid-rs/uuid
- **version-compare** 0.2.1 — MIT
  https://gitlab.com/timvisee/version-compare
- **version_check** 0.9.5 — MIT/Apache-2.0
  https://github.com/SergioBenitez/version_check
- **walkdir** 2.5.0 — Unlicense/MIT
  https://github.com/BurntSushi/walkdir
- **want** 0.3.1 — MIT
  https://github.com/seanmonstar/want
- **web_atoms** 0.2.4 — MIT OR Apache-2.0
  https://github.com/servo/html5ever
- **webkit2gtk** 2.0.2 — MIT
  https://github.com/tauri-apps/webkit2gtk-rs
- **webkit2gtk-sys** 2.0.2 — MIT
  https://github.com/tauri-apps/webkit2gtk-rs
- **winnow** 0.5.40 — MIT
  https://github.com/winnow-rs/winnow
- **winnow** 0.7.15 — MIT
  https://github.com/winnow-rs/winnow
- **winnow** 1.0.2 — MIT
  https://github.com/winnow-rs/winnow
- **writeable** 0.6.3 — Unicode-3.0
  https://github.com/unicode-org/icu4x
- **wry** 0.55.1 — Apache-2.0 OR MIT
  https://github.com/tauri-apps/wry
- **x11** 2.21.0 — MIT
  https://github.com/AltF02/x11-rs.git
- **x11-dl** 2.21.0 — MIT
  https://github.com/AltF02/x11-rs.git
- **xattr** 1.6.1 — MIT OR Apache-2.0
  https://github.com/Stebalien/xattr
- **yoke** 0.8.2 — Unicode-3.0
  https://github.com/unicode-org/icu4x
- **yoke-derive** 0.8.2 — Unicode-3.0
  https://github.com/unicode-org/icu4x
- **zbus** 5.15.0 — MIT
  https://github.com/z-galaxy/zbus/
- **zbus_macros** 5.15.0 — MIT
  https://github.com/z-galaxy/zbus/
- **zbus_names** 4.3.2 — MIT
  https://github.com/z-galaxy/zbus/
- **zerocopy** 0.8.48 — BSD-2-Clause OR Apache-2.0 OR MIT
  https://github.com/google/zerocopy
- **zerofrom** 0.1.7 — Unicode-3.0
  https://github.com/unicode-org/icu4x
- **zerofrom-derive** 0.1.7 — Unicode-3.0
  https://github.com/unicode-org/icu4x
- **zeroize** 1.8.2 — Apache-2.0 OR MIT
  https://github.com/RustCrypto/utils
- **zerotrie** 0.2.4 — Unicode-3.0
  https://github.com/unicode-org/icu4x
- **zerovec** 0.11.6 — Unicode-3.0
  https://github.com/unicode-org/icu4x
- **zerovec-derive** 0.11.3 — Unicode-3.0
  https://github.com/unicode-org/icu4x
- **zmij** 1.0.21 — MIT
  https://github.com/dtolnay/zmij
- **zvariant** 5.11.0 — MIT
  https://github.com/z-galaxy/zbus/
- **zvariant_derive** 5.11.0 — MIT
  https://github.com/z-galaxy/zbus/
- **zvariant_utils** 3.3.1 — MIT
  https://github.com/z-galaxy/zbus/

---

## License texts

Every license referenced above is reproduced here in its canonical form, or
is available at <https://spdx.org/licenses/>. Components covered by an
`A OR B` expression may be redistributed under either; SpectraControl
exercises the MIT option where available.

### MIT License

```
Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
```

### Apache License 2.0

Full text: <https://www.apache.org/licenses/LICENSE-2.0>

### BSD-2-Clause ("Simplified BSD")

Full text: <https://opensource.org/license/bsd-2-clause>

### BSD-3-Clause ("New BSD")

Full text: <https://opensource.org/license/bsd-3-clause>

### ISC License

Full text: <https://opensource.org/license/isc-license-txt>

### Mozilla Public License 2.0 (MPL-2.0)

Full text: <https://www.mozilla.org/en-US/MPL/2.0/>

> MPL-2.0 is a **file-level** weak-copyleft license. SpectraControl uses the
> affected crates (`cssparser`, `cssparser-macros`, `dtoa-short`, `option-ext`,
> `selectors`) as unmodified upstream dependencies. Source for each crate is
> available at the URL listed in the inventory above; users who wish to
> exercise MPL-2.0 rights can retrieve the upstream code there or via
> `cargo vendor`.

### Other permissive licenses

Crates released under `Unicode-3.0`, `Zlib`, `0BSD`, `Unlicense`, `MIT-0`,
`CC0-1.0`, and `Apache-2.0 WITH LLVM-exception` are similarly permissive and
compatible with SpectraControl's MIT distribution. Their canonical texts are
available at <https://spdx.org/licenses/>.
