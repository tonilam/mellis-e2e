---
id: TKT-5
type: new-story
status: signed-off
story: US-5
---

# TKT-5: Watch the run

**Type:** new-story
**Story:** [US-5](../stories/US-5-watch-the-run.md)
**Status:** signed-off

## Request

with the user story we created and our target for a shippable product, can you create all the tickets we needed?

This ticket makes US-5 true for a suite started by TKT-4. The story file already exists. Do not add a second story.

## Acceptance criteria

* AC-1: While the command is running, each test the output reports moves from waiting, to running, to passed or failed.
* AC-2: While the command is running, the window shows how many tests are finished, how many are running, and how many are still waiting.
* AC-3: The developer can stop the run.
* AC-4: After a stop, results already reported stay in the window.
* AC-5: After a stop, the command is no longer running.
* AC-6: When the command exits on its own, the run is finished.
* AC-7: When the developer closes the window while a run is going, the runner stops the command.

## Action items

- [x] AI-1 (AC-1): While the command is running, move each reported test from waiting, to running, to passed or failed.
- [x] AI-2 (AC-2): While the command is running, show how many tests are finished, running, and still waiting.
- [x] AI-3 (AC-3): The developer can stop the run.
- [x] AI-4 (AC-4): After a stop, keep results already reported in the window.
- [x] AI-5 (AC-5): After a stop, the command is no longer running.
- [x] AI-6 (AC-6): When the command exits on its own, mark the run finished.
- [x] AI-7 (AC-7): When the developer closes the window during a run, stop the command.

## Sign-off

- [x] Every action item above is checked
- [x] Every acceptance criterion above is met
- Date: 2026-11-14
- Result: Verified async progress, counts, stop (keep results), finish on exit, stop on close (`go test ./...`). Discovery [ACCEPT_PROTOCOL](../../.agents/teams/discovery/sessions/2026-11-14-tkt5-reporter.md): mellis:test interim until a project reporter is named.
