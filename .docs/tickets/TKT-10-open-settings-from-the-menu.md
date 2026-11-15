---
id: TKT-10
type: new-story
status: signed-off
story: US-10
---

# TKT-10: Open Settings from the main app menu

**Type:** new-story
**Story:** [US-10](../stories/US-10-open-settings-from-the-menu.md)
**Status:** signed-off

## Request

Move Environment and Project and Command to a settings, the Settings should be another page, so let's create the main app menu.

## Acceptance criteria

* AC-1: The window shows a main app menu the developer can open.
* AC-2: The main app menu includes a Settings item.
* AC-3: Choosing Settings opens a Settings page that is not the run page.
* AC-4: The Settings page shows Environment, Project path, and Suite command.
* AC-5: The run page does not show Environment, Project path, or Suite command controls.
* AC-6: From Settings, the developer can return to the run page.
* AC-7: On Settings, the developer can view and retry the loaded environment.
* AC-8: On Settings, the developer can enter and change the project path and the suite command.

## Action items

- [x] AI-1 (AC-1, AC-2): Add a main app menu that the developer can open and that includes a Settings item.
- [x] AI-2 (AC-3, AC-6): Add a Settings page separate from the run page, and a way to return to the run page from Settings.
- [x] AI-3 (AC-4, AC-5): Move Environment, Project path, and Suite command from the run page onto the Settings page so the run page no longer shows those controls.
- [x] AI-4 (AC-7, AC-8): Keep view/retry for the loaded environment and enter/change for project path and suite command working on Settings.
- [x] AI-5 (AC-1–AC-8): Align [docs/features/local-runner/run-suite.md](../../docs/features/local-runner/run-suite.md) with the menu and Settings page layout.

## Sign-off

- [x] Every action item above is checked
- [x] Every acceptance criterion above is met
- Date: 2026-11-15
- Result: Verified Menu → Settings / Back to run / Menu → Run; env/project/command only on Settings; run page starts suite with those values; `go test ./...` pass.
