---
id: TKT-2
type: new-story
status: draft
story: US-2
---

# TKT-2: Load the shell environment

**Type:** new-story
**Story:** [US-2](../stories/US-2-load-the-shell-environment.md)
**Status:** draft

## Request

with the user story we created and our target for a shippable product, can you create all the tickets we needed?

This ticket makes US-2 true in the app from TKT-1. The story file already exists. Do not add a second story. Suite and agent commands in later tickets must use the environment this ticket loads.

## Acceptance criteria

* AC-1: On open, the runner loads the developer's login-shell environment, including the PATH a new terminal would have.
* AC-2: A suite command started later in the session runs with that loaded environment.
* AC-3: An agent CLI started later in the session runs with that loaded environment.
* AC-4: The developer can open a view of the environment the runner loaded.
* AC-5: The developer can retry the load from that view.
* AC-6: When the environment cannot be loaded, the window says so.
* AC-7: When the environment cannot be loaded, the runner does not start a suite and does not start an agent request.
* AC-8: After a failed load, a successful retry makes the suite and agent actions available again.

## Action items

- [ ] AI-1 (AC-1): On open, load the developer's login-shell environment, including the PATH a new terminal would have.
- [ ] AI-2 (AC-2): Start a suite command with that loaded environment.
- [ ] AI-3 (AC-3): Start an agent CLI with that loaded environment.
- [ ] AI-4 (AC-4): Add a view that shows the environment the runner loaded.
- [ ] AI-5 (AC-5): From that view, the developer can retry the load.
- [ ] AI-6 (AC-6): When the load fails, the window says the environment could not be loaded.
- [ ] AI-7 (AC-7): When the load has failed, starting a suite and starting an agent request do nothing.
- [ ] AI-8 (AC-8): After a failed load, a successful retry makes the suite and agent actions available again.

## Sign-off

- [ ] Every action item above is checked
- [ ] Every acceptance criterion above is met
- Date:
- Result:
