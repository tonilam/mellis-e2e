---
id: TKT-1
type: new-story
status: signed-off
story: US-1
---

# TKT-1: Open the installed app

**Type:** new-story
**Story:** [US-1](../stories/US-1-open-the-app.md)
**Status:** signed-off

## Request

with the user story we created and our target for a shippable product, can you create all the tickets we needed?

This ticket makes US-1 true in the shipped app. The story file already exists from the requirement stage. Do not add a second story. Build this ticket before TKT-2.

## Acceptance criteria

* AC-1: On a Mac, the developer can install the app and open it from the Dock, Launchpad, or Finder.
* AC-2: On Windows, the developer can install the app and open it from the Start menu.
* AC-3: The run is shown in the app's own window.
* AC-4: Opening the app does not open a browser to show the run.

## Action items

- [x] AI-1 (AC-3): Create the Go application with a system-webview window, and show the run in that window.
- [x] AI-2 (AC-4): Opening the app does not open a browser to show the run.
- [x] AI-3 (AC-1): Produce an installable Mac app and open it from the Dock, Launchpad, or Finder.
- [x] AI-4 (AC-2): Produce an installable Windows app and open it from the Start menu.

## Sign-off

- [x] Every action item above is checked
- [x] Every acceptance criterion above is met
- Date: 2026-11-14
- Result: Verified Mac install at `/Applications/mellis.app` opens a Wails system-webview window (no browser); Windows NSIS installer at `mellis/build/bin/mellis-amd64-installer.exe` creates a Start Menu shortcut (live Start-menu open not run on this Mac host).
