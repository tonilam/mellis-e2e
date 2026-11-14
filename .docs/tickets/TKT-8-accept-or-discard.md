---
id: TKT-8
type: new-story
status: signed-off
story: US-8
---

# TKT-8: Accept or discard the change

**Type:** new-story
**Story:** [US-8](../stories/US-8-accept-or-discard.md)
**Status:** signed-off

## Request

with the user story we created and our target for a shippable product, can you create all the tickets we needed?

This ticket makes US-8 true for a change proposed by the agent request in TKT-7. The story file already exists. Do not add a second story. The rerun uses the same command and environment as TKT-4.

## Acceptance criteria

* AC-1: The agent shows a proposed change before that change remains in the project.
* AC-2: The developer can accept the proposed change. After accept, the change is in the project.
* AC-3: The developer can discard the proposed change. After discard, the project files match the tree from before the request.
* AC-4: Nothing from the request remains in the project until the developer accepts.
* AC-5: After an accepted change, the developer can run that same test again from the window.
* AC-6: After that rerun, the window shows the new result beside the previous result for that test.

## Action items

- [x] AI-1 (AC-1): The agent shows a proposed change before that change remains in the project.
- [x] AI-2 (AC-2): The developer can accept the proposed change, and after accept the change is in the project.
- [x] AI-3 (AC-3): The developer can discard the proposed change, and after discard the project files match the tree from before the request.
- [x] AI-4 (AC-4): Nothing from the request remains in the project until the developer accepts.
- [x] AI-5 (AC-5): After an accepted change, the developer can run that same test again from the window.
- [x] AI-6 (AC-6): After that rerun, the window shows the new result beside the previous result for that test.

## Sign-off

- [x] Every action item above is checked
- [x] Every acceptance criterion above is met
- Date: 2026-11-14
- Result: Verified snapshot/restore/accept/discard and rerun previous+new (`go test ./...`). Discovery [ACCEPT_SNAPSHOT_AND_RERUN](../../.agents/teams/discovery/sessions/2026-11-14-tkt8-accept-discard.md).
