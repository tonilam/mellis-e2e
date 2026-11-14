---
id: TKT-4
type: new-story
status: signed-off
story: US-4
---

# TKT-4: Run the suite

**Type:** new-story
**Story:** [US-4](../stories/US-4-run-the-suite.md)
**Status:** signed-off

## Request

with the user story we created and our target for a shippable product, can you create all the tickets we needed?

This ticket makes US-4 true. The story file already exists. Do not add a second story. The runner starts the command from TKT-3 in the environment from TKT-2. It does not start the project's application.

## Acceptance criteria

* AC-1: The developer can start the suite from the window.
* AC-2: The runner executes the command shown for the session, in the project path shown for the session, with the environment loaded on open.
* AC-3: Starting the suite does not start the project's application.
* AC-4: Starting the suite does not stop or restart the project's application.
* AC-5: When the project's application is not running, the runner still executes the command.
* AC-6: When the project's application is not running, a failure on screen is the failure the command reported.
* AC-7: When the project's application is not running, the runner does not start that application.

## Action items

- [x] AI-1 (AC-1): The developer can start the suite from the window.
- [x] AI-2 (AC-2): The runner executes the shown command in the shown project path with the environment loaded on open.
- [x] AI-3 (AC-3): Starting the suite does not start the project's application.
- [x] AI-4 (AC-4): Starting the suite does not stop or restart the project's application.
- [x] AI-5 (AC-5): When the project's application is not running, the runner still executes the command.
- [x] AI-6 (AC-6): When the project's application is not running, a failure on screen is the failure the command reported.
- [x] AI-7 (AC-7): When the project's application is not running, the runner does not start that application.

## Sign-off

- [x] Every action item above is checked
- [x] Every acceptance criterion above is met
- Date: 2026-11-14
- Result: Verified Start suite from window runs shown command/path/env only; command failure text surfaced; no project-app lifecycle (`go test ./...`).
