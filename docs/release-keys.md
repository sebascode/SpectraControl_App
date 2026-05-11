# Release signing keys

SpectraControl signs AppImage releases with [minisign](https://jedisct1.github.io/minisign/) so the in-app updater can verify the artifact before replacing the binary on disk. The verification is done client-side by `tauri-plugin-updater` against a public key compiled into every build.

## Where the keys live

| Key half | Location | Notes |
|---|---|---|
| Public | `src-tauri/tauri.conf.json` → `plugins.updater.pubkey` (base64) | Compiled into every release. Public; safe to commit. |
| Private | GitHub repo secret `TAURI_SIGNING_PRIVATE_KEY` (and password in `TAURI_SIGNING_PRIVATE_KEY_PASSWORD`) | Only the release workflow has access. **Never** check it into the repo. |

The CI workflow (`.github/workflows/release.yml`) reads the secret env vars and `cargo tauri build` consumes them through standard Tauri channels. No special invocation is required at the command line.

## Threat model in one sentence

If the private key leaks, an attacker can sign a malicious `.AppImage`, host it where the updater can reach it, and your installed users will install it without warning. The pubkey is in every binary, so existing users will only trust signatures from whoever holds the private key — including a leaked copy.

## Rotation procedure

Triggered when:

- You suspect the private key (or its password) was exposed.
- You're migrating to a new maintainer who needs to control releases.
- Routine rotation hygiene (no firm cadence — sign-once-and-forget keys are common in practice for small projects).

The procedure has a hard tradeoff: **existing installs cannot auto-update past the rotation point**. They have the old pubkey baked in and will reject signatures from the new key. Plan accordingly.

### Step 1 — Generate the new keypair

```bash
# In a clean directory you do not commit:
mkdir -p ~/.spectracontrol-release-key
cd ~/.spectracontrol-release-key
cargo tauri signer generate -w minisign.key
# Pick a strong password; you will need it for every release.
# minisign.key      → private key (KEEP SECRET, encrypted with the password you set)
# minisign.key.pub  → public key
```

### Step 2 — Update the public key in the repo

```bash
# Copy the contents of minisign.key.pub (single line) and replace the
# value of plugins.updater.pubkey in src-tauri/tauri.conf.json. The
# pubkey field is the base64 form of the file.
```

Commit and push the change as part of the rotation release. The binary built from this commit will only trust the new key.

### Step 3 — Update the GitHub secrets

In **Settings → Secrets and variables → Actions**:

- `TAURI_SIGNING_PRIVATE_KEY` → paste the contents of `minisign.key` (the full encrypted blob, including the `untrusted comment:` header).
- `TAURI_SIGNING_PRIVATE_KEY_PASSWORD` → the password you set when generating.

### Step 4 — Cut a release with the new key

Bump version, tag, push. CI builds and signs with the new private key, and clients built from this tag onwards have the matching public key.

### Step 5 — Communicate the break to users

Existing installs **cannot** receive this update through the auto-updater — their embedded pubkey is the old one, so the signature on the new `latest.json` will be rejected and the banner will silently never appear (or appear and fail on install).

Mitigations:

- **Release notes**: clearly state "one-time manual install required" and link to the AppImage in the release.
- **README** banner during the transition period.
- **Past release**: optionally publish a final point release on the **old** key that contains a notification or in-app message pointing users to the new download URL. This is best-effort — once they install the new key version, future updates work as normal.

### Step 6 — Destroy the old private key

After confirming the new key works end-to-end (build → sign → verify on a clean install), securely wipe the old `minisign.key` from every machine and offline backup. Update the GitHub Secret if the old value is still present (the rotation in Step 3 already overwrites it, but double-check).

## Reading minisign output (for the curious)

The signature embedded in `latest.json` is a base64-wrapped minisign signature block:

```
untrusted comment: signature from tauri secret key
RWQoZ...   ← Ed25519 signature over the AppImage bytes
trusted comment: timestamp:1778480053  file:SpectraControl_0.2.6_amd64.AppImage
6I6P...    ← Ed25519 signature over the trusted comment line
```

The updater verifies the first signature against the AppImage download, and the second against the trusted-comment line, using the pubkey baked into the binary. Both must validate for the install to proceed.
