---
id: TKT-5
type: new-story
status: draft
story: US-5
---

# TKT-5: Watch the run

**Type:** new-story
**Story:** [US-5](../stories/US-5-watch-the-run.md)
**Status:** draft

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

- [ ] AI-1 (AC-1): While the command is running, move each reported test from waiting, to running, to passed or failed.
- [ ] AI-2 (AC-2): While the command is running, show how many tests are finished, running, and still waiting.
- [ ] AI-3 (AC-3): The developer can stop the run.
- [ ] AI-4 (AC-4): After a stop, keep results already reported in the window.
- [ ] AI-5 (AC-5): After a stop, the command is no longer running.
- [ ] AI-6 (AC-6): When the command exits on its own, mark the run finished.
- [ ] AI-7 (AC-7): When the developer closes the window during a run, stop the command.

## Sign-off

- [ ] Every action item above is checked
- [ ] Every acceptance criterion above is met
- Date:
- Result:
