# Mellis

Installable local E2E runner for macOS and Windows. Built with [Go](https://go.dev/) and [Wails](https://wails.io/); the UI runs in a system webview window (no browser tab for the run).

Application source lives in [`mellis/`](mellis/).

## Prerequisites

- Go 1.23+
- [Wails CLI](https://wails.io/docs/gettingstarted/installation) v2
- Node.js / npm (frontend build)
- macOS: Xcode command line tools
- Windows installer (optional): [NSIS](https://wails.io/docs/guides/windows-installer/) / `makensis`

## Development

```bash
cd mellis
wails dev
```

## Build

### macOS

```bash
cd mellis
wails build
```

Output: `mellis/build/bin/mellis.app`

Install and open:

```bash
cp -R build/bin/mellis.app /Applications/
open /Applications/mellis.app
```

### Windows

Cross-compile from macOS (or build natively on Windows):

```bash
cd mellis
wails build -platform windows/amd64
```

Output: `mellis/build/bin/mellis.exe`

Start-menu installer (requires NSIS on the build machine):

```bash
wails build -platform windows/amd64 -nsis
```

Outputs: `mellis/build/bin/mellis.exe` and `mellis/build/bin/mellis-amd64-installer.exe`.

## License

[GNU General Public License v3.0](LICENSE)
