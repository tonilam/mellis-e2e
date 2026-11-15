---
id: TKT-12
type: new-story
status: signed-off
story: US-11
---

# TKT-12: Build on Wails V3

**Type:** new-story
**Story:** [US-11](../stories/US-11-build-on-wails-v3.md)
**Status:** signed-off

## Request

We are still in early stage, let's switch to V3. create a infrastructure ticket.

## Acceptance criteria

* AC-1: The Mellis Go module depends on Wails V3 and does not depend on Wails V2.
* AC-2: A developer can install, develop, and package Mellis using the Wails V3 toolchain, and the project README documents those commands.
* AC-3: On a Mac, the developer can still install the app and open it from the Dock, Launchpad, or Finder, with the run shown in the app's own window and without a browser tab for the run.
* AC-4: On Windows, the developer can still produce an installable app (including a Start-menu installer when NSIS is available).
* AC-5: After the migration, the main app menu, Settings, login-shell environment load, suite progress events, and agent progress events still work in the app window.

## Action items

- [x] AI-1 (AC-1): Replace the Wails V2 module with Wails V3 in `go.mod` and remove V2 imports from the Go code.
- [x] AI-2 (AC-1, AC-5): Port `main.go` and the App backend to the V3 application/service model (lifecycle, services, events, before-close).
- [x] AI-3 (AC-2, AC-5): Regenerate frontend bindings for V3, update frontend imports/events, and remove the V2 `wailsjs` tree.
- [x] AI-4 (AC-2): Update Wails project config and READMEs so install/dev/build/NSIS steps use the V3 toolchain (`wails3`).
- [x] AI-5 (AC-3): Build the Mac app, install it, and verify it opens in a system-webview window with no browser tab for the run.
- [x] AI-6 (AC-4): Produce the Windows installable build (and NSIS installer when `makensis` is available) and record the output paths.
- [x] AI-7 (AC-5): Smoke-check main menu → Settings, environment load/retry, and that suite/agent progress events still update the UI.
- [x] AI-8 (AC-1–AC-5): Run `go test ./...` and record pass/fail in the ticket result.

## Sign-off

- [x] Every action item above is checked
- [x] Every acceptance criterion above is met
- Date: 2026-11-15
- Result: Migrated to Wails V3.0.0 (`go.mod` has no V2). `go test ./...` pass. Mac `wails3 package` → `bin/mellis.app` opens a system-webview window. Windows `bin/mellis.exe`, `bin/mellis-amd64-installer.exe`, `bin/mellis-arm64-installer.exe`. READMEs use `wails3`.
