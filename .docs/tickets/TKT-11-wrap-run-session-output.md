---
id: TKT-11
type: update-story
status: signed-off
story: US-5
---

# TKT-11: Wrap run session output

**Type:** update-story
**Story:** [US-5](../stories/US-5-watch-the-run.md)
**Status:** signed-off

## Request

create a bug ticket, the output in the run session has overflow, please fix it by warping long string to fit max width.

## Acceptance criteria

* AC-1: When the run session shows a line longer than the window width, that line wraps so the full line stays within the window's max width without horizontal overflow.

Previous: (none — US-5 did not require long output lines to wrap inside the window.)
New: When the run session shows a line longer than the window width, that line wraps so the full line stays within the window's max width without horizontal overflow.

## Action items

- [x] AI-1 (AC-1): Wrap long unbroken strings in the run session output so they stay within the window's max width.
- [x] AI-2 (AC-1): Verify with a long PATH-style string that the run session output stays inside the panel with no horizontal overflow.

## Sign-off

- [x] Every action item above is checked
- [x] Every acceptance criterion above is met
- Date: 2026-11-15
- Result: Verified long PATH in `.action-result` wraps (`overflow-wrap: anywhere`); scrollWidth equals clientWidth (576px); no shell overflow. `go test ./...` pass.
