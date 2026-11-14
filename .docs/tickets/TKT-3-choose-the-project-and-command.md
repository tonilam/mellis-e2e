---
id: TKT-3
type: new-story
status: signed-off
story: US-3
---

# TKT-3: Choose the project and the command

**Type:** new-story
**Story:** [US-3](../stories/US-3-choose-the-project-and-command.md)
**Status:** signed-off

## Request

with the user story we created and our target for a shippable product, can you create all the tickets we needed?

This ticket makes US-3 true. The story file already exists. Do not add a second story. The path and command chosen here are the ones TKT-4 runs.

## Acceptance criteria

* AC-1: The developer can enter a project path.
* AC-2: The window shows that path for as long as the session lasts.
* AC-3: The developer can enter the command that starts the E2E suite.
* AC-4: The window shows that command before a run starts.
* AC-5: The developer can change the command before the run starts.
* AC-6: When this machine already has a path and command from an earlier session, the next session offers that same path and command for the developer to confirm or change.
* AC-7: When the project path does not exist, the window says so and the runner does not start a command.
* AC-8: When the command cannot be started, the window says so and the runner shows no invented test results.
* AC-9: When the command is missing from the loaded environment, the window names that command and the runner shows no invented test results.

## Action items

- [x] AI-1 (AC-1): The developer can enter a project path.
- [x] AI-2 (AC-2): The window keeps that path visible until the session ends.
- [x] AI-3 (AC-3): The developer can enter the command that starts the E2E suite.
- [x] AI-4 (AC-4): The window shows that command before a run starts.
- [x] AI-5 (AC-5): The developer can change the command before the run starts.
- [x] AI-6 (AC-6): Save the path and command on this machine, and offer that same pair in the next session for the developer to confirm or change.
- [x] AI-7 (AC-7): When the path does not exist, the window says so and the runner does not start a command.
- [x] AI-8 (AC-8): When the command cannot be started, the window says so and the runner shows no invented test results.
- [x] AI-9 (AC-9): When the command is missing from the loaded environment, the window names that command and the runner shows no invented test results.

## Sign-off

- [x] Every action item above is checked
- [x] Every acceptance criterion above is met
- Date: 2026-11-14
- Result: Verified path/command UI, persistence across load, and StartSuite gates for missing path, missing PATH command, and start failure (`go test ./...`).
