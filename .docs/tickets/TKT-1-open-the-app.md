---
id: TKT-1
type: new-story
status: draft
story: US-1
---

# TKT-1: Open the installed app

**Type:** new-story
**Story:** [US-1](../stories/US-1-open-the-app.md)
**Status:** draft

## Request

with the user story we created and our target for a shippable product, can you create all the tickets we needed?

This ticket makes US-1 true in the shipped app. The story file already exists from the requirement stage. Do not add a second story. Build this ticket before TKT-2.

## Acceptance criteria

* AC-1: On a Mac, the developer can install the app and open it from the Dock, Launchpad, or Finder.
* AC-2: On Windows, the developer can install the app and open it from the Start menu.
* AC-3: The run is shown in the app's own window.
* AC-4: Opening the app does not open a browser to show the run.

## Action items

- [ ] AI-1 (AC-3): Create the Go application with a system-webview window, and show the run in that window.
- [ ] AI-2 (AC-4): Opening the app does not open a browser to show the run.
- [ ] AI-3 (AC-1): Produce an installable Mac app and open it from the Dock, Launchpad, or Finder.
- [ ] AI-4 (AC-2): Produce an installable Windows app and open it from the Start menu.

## Sign-off

- [ ] Every action item above is checked
- [ ] Every acceptance criterion above is met
- Date:
- Result:
