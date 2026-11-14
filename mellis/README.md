# Mellis

Installable local E2E runner (Go + Wails). The UI runs in a system webview window; opening the app does not open a browser for the run.

## License

This project is licensed under the [GNU General Public License v3.0](../LICENSE).

## Prerequisites

- Go 1.23+
- [Wails CLI](https://wails.io/docs/gettingstarted/installation) v2
- Node.js / npm (frontend build)
- macOS: Xcode command line tools
- Windows packaging (optional on this machine): NSIS for a Start-menu installer

## Development

```bash
cd mellis
wails dev
```

## Build Mac app

```bash
cd mellis
wails build
```

Output: `build/bin/mellis.app`

Install and open:

```bash
cp -R build/bin/mellis.app /Applications/
open /Applications/mellis.app
```

You can also open the same `.app` from Finder, Dock, or Launchpad after it is in `/Applications` (or after dragging it there from Finder).

## Build Windows app

Cross-compile from macOS:

```bash
cd mellis
wails build -platform windows/amd64
```

Output: `build/bin/mellis.exe`

To produce a Start-menu installer (requires [NSIS](https://wails.io/docs/guides/windows-installer/) / `makensis` on the build machine):

```bash
wails build -platform windows/amd64 -nsis
```

Outputs: `build/bin/mellis.exe` and `build/bin/mellis-amd64-installer.exe`.

### Verify on Windows (Start menu)

This host is macOS, so Start-menu launch was not run here. On a Windows machine:

1. Copy `build/bin/mellis-amd64-installer.exe` to the Windows machine.
2. Run the installer and complete the wizard (installs under Program Files and registers a Start-menu shortcut).
3. Open **Start**, search for **mellis**, and launch it.
4. Confirm the run UI appears in the app window and that no browser tab opens for the run.
