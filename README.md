# Mellis

Installable local E2E runner for macOS and Windows. Built with [Go](https://go.dev/) and [Wails](https://v3.wails.io/) V3; the UI runs in a system webview window (no browser tab for the run).

Application source lives in [`mellis/`](mellis/).

## Prerequisites

- Go 1.26+
- [Wails CLI](https://v3.wails.io/getting-started/installation/) v3 (`wails3`)
- Node.js / npm (frontend build)
- macOS: Xcode command line tools
- Windows installer (optional): [NSIS](https://v3.wails.io/guides/windows-installer/) / `makensis`

## Development

```bash
cd mellis
wails3 dev
```

## Build

### macOS

```bash
cd mellis
wails3 build
wails3 package
```

Output: `mellis/bin/mellis.app`

Install and open:

```bash
cp -R bin/mellis.app /Applications/
open /Applications/mellis.app
```

### Windows

Cross-compile from macOS (or build natively on Windows):

```bash
cd mellis
wails3 build GOOS=windows
```

Output: `mellis/bin/mellis.exe`

Start-menu installer (requires NSIS on the build machine):

```bash
cd mellis
wails3 build GOOS=windows GOARCH=amd64
wails3 package GOOS=windows GOARCH=amd64
```

Outputs: `mellis/bin/mellis.exe` and `mellis/bin/mellis-amd64-installer.exe`.

## License

[GNU General Public License v3.0](LICENSE)
