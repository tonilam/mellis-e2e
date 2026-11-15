# Mellis

Installable local E2E runner (Go + Wails V3). The UI runs in a system webview window; opening the app does not open a browser for the run.

## License

This project is licensed under the [GNU General Public License v3.0](../LICENSE).

## Prerequisites

- Go 1.26+
- [Wails CLI](https://v3.wails.io/getting-started/installation/) v3 (`wails3`)
- Node.js / npm (frontend build)
- macOS: Xcode command line tools
- Windows packaging (optional on this machine): NSIS for a Start-menu installer

## Development

```bash
cd mellis
wails3 dev
```

## Build Mac app

```bash
cd mellis
wails3 build
wails3 package
```

Output: `bin/mellis.app`

Install and open:

```bash
cp -R bin/mellis.app /Applications/
open /Applications/mellis.app
```

You can also open the same `.app` from Finder, Dock, or Launchpad after it is in `/Applications` (or after dragging it there from Finder).

## Build Windows app

Cross-compile from macOS:

```bash
cd mellis
wails3 build GOOS=windows
```

Output: `bin/mellis.exe`

To produce a Start-menu installer (requires [NSIS](https://v3.wails.io/guides/windows-installer/) / `makensis` on the build machine):

```bash
cd mellis
wails3 build GOOS=windows
wails3 package GOOS=windows
```

Outputs (examples): `bin/mellis.exe`, `bin/mellis-amd64-installer.exe`, and/or `bin/mellis-arm64-installer.exe` depending on `GOARCH`.

### Verify on Windows (Start menu)

This host is macOS, so Start-menu launch was not run here. On a Windows machine:

1. Copy the Windows installer to the Windows machine.
2. Run the installer and complete the wizard (installs under Program Files and registers a Start-menu shortcut).
3. Open **Start**, search for **mellis**, and launch it.
4. Confirm the run UI appears in the app window and that no browser tab opens for the run.
