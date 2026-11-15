# US-11: Build on Wails V3

* **User value:** As a developer maintaining Mellis while it is still early, I want the installable app built on Wails V3, so that we ship on the current stable desktop stack before more product work lands.
* **Source:** Request to switch from Wails V2 to V3 (infrastructure / early-stage migration)
* **Tickets:** [TKT-12](../tickets/TKT-12-build-on-wails-v3.md)

## Acceptance criteria

* The Mellis Go module depends on Wails V3 and does not depend on Wails V2.
* A developer can install, develop, and package Mellis using the Wails V3 toolchain, and the project README documents those commands.
* On a Mac, the developer can still install the app and open it from the Dock, Launchpad, or Finder, with the run shown in the app's own window and without a browser tab for the run.
* On Windows, the developer can still produce an installable app (including a Start-menu installer when NSIS is available).
* After the migration, the main app menu, Settings, login-shell environment load, suite progress events, and agent progress events still work in the app window.
