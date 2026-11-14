---
name: ticket-start
description: Start development from a ticket in .docs/tickets/. Creates the ticket first when it is missing. A ticket creates a user story, updates one, or fixes a bug in one. Use when starting work, filing a story change, or fixing a story that does not behave as specified.
---

# Workflow: ticket-start

Development starts from a ticket. The ticket is the unit of work. The user story in `.docs/stories/` is the lasting requirement. Follow the phases in order. Copy [template.md](template.md). Do not invent a different ticket shape.

## What a ticket is

A ticket is one of these, and only one:

| Type | When | What changes in `.docs/stories/` |
| --- | --- | --- |
| `new-story` | No story covers the capability. | Add the next `US-<n>` file and a row in [.docs/stories/README.md](/.docs/stories/README.md). |
| `update-story` | A story exists, and its acceptance criteria need to change. | Edit that story. The ticket records the previous wording and the new wording. |
| `bug` | A story's acceptance criteria already say what should happen, and the product does something else. | Do not change those criteria. Quote them on the ticket. |

If a report is filed as a bug and no story states the expected behaviour, it is not a bug. Classify it as `update-story` or `new-story` and say why.

## Ticket file

- Path: `.docs/tickets/TKT-<n>-<slug>.md`
- `<n>` is the number in the `id` field. `<slug>` is a short kebab-case title.
- The body matches [template.md](template.md).
- Status is `draft`, `ready`, `in-progress`, or `signed-off`.
- Index every ticket in [.docs/tickets/README.md](/.docs/tickets/README.md).

## Acceptance criteria

- Each criterion is numbered `AC-1`, `AC-2`, and so on.
- Each criterion is one outcome that passes or fails. Name the actor, the action, and the observable result.
- A criterion that says "works" or "handles errors" is not done. Rewrite it until a reviewer can decide yes or no without asking what was meant.
- For `new-story` and `update-story`, the story file's acceptance criteria and the ticket's acceptance criteria say the same thing before any code is written. The story keeps them after the ticket is signed off.
- For `bug`, each ticket criterion points at the story criterion it restores, and states the observed result beside the required result.

## Action items

- Each item is numbered `AI-1`, `AI-2`, and so on, and names the criterion it serves: `AI-1 (AC-1)`.
- One item is one task a person can finish and check. Split an item that needs two reviews.
- The checkbox lives in the ticket file. Check it in the same change that finishes the task. An unchecked item is not done. A checked item with the task still open is a mistake; uncheck it.
- Verification is an action item when a criterion cannot be seen from the diff alone. Example: `AI-4 (AC-2): Run the suite command and record the pass/fail line in the ticket result.`
- Sign-off reads the checkboxes. It does not replace them.

## Phases

### 1. Resolve the ticket

Before editing application code:

1. Read the latest week of `docs/changelogs/` when that folder has entries.
2. Read the request and find the ticket it names: an id (`TKT-3`), a path under `.docs/tickets/`, or a slug.
3. When that file exists, use it. Do not create a second ticket for the same ask.
4. When no ticket is named, or the named ticket file is not in `.docs/tickets/`, create one before anything else. Allocate `<n>` as one greater than the highest `TKT-<n>` already in `.docs/tickets/`. When the user named an id, use that id.
5. Add the new row to `.docs/tickets/README.md`.

Creating the file is not approval to write code. Status stays `draft`.

### 2. Fill the ticket

1. Set `type` using the table above.
2. Link the story.
   - `new-story`: create `.docs/stories/US-<n>-<slug>.md` in the same shape as the existing stories (user value, source, acceptance criteria). Use the next story number. Add the row to `.docs/stories/README.md`. Point the feature area README at it when one exists. When the requirement stage already wrote that story, link the existing file and do not add a second one. Copy its acceptance criteria onto the ticket unchanged. The action items are the work that makes those criteria true in the product.
   - `update-story`: link the existing story. Under the ticket's acceptance criteria, write the previous criterion text, then the new text.
   - `bug`: link the existing story. Quote each acceptance criterion that fails. Write what happens instead.
3. Write the acceptance criteria and the action items.
4. On the story, add or extend a **Tickets** line with this ticket's id. For a bug, that line is the only edit to the story.

Stop. Show the user the ticket path, the type, the story link, the acceptance criteria, and the action items.

### 3. Approval

Wait for the user to accept the ticket or to change it.

- Accepted: set status to `ready`.
- Changes requested: edit the ticket and the story, then wait again.
- Do not write application code while status is `draft`.

A ticket that already has status `ready` or `in-progress`, and whose criteria the user is not revising, skips this wait.

### 4. Do the work

Set status to `in-progress` when the first action item starts.

- Work the action items in order.
- Check an item only after that task is finished.
- Stay inside the ticket. A new capability or a new failure gets its own ticket. Do not fold it into this one.
- Leave the sign-off boxes unchecked during this phase.

### 5. Check the criteria

For each `AC-*`, decide met or not met from the product behaviour, not from the plan.

- Met: leave the matching action items checked.
- Not met: uncheck the action item that was supposed to deliver it, fix the work, and check it again only when the criterion is met.
- A `bug` is fixed only when the quoted story criterion holds.

### 6. Sign off

Sign off only when every action item is checked and every acceptance criterion is met.

1. Check the two sign-off boxes.
2. Set `status` to `signed-off`.
3. Set **Date** to the calendar date and **Result** to one line naming what was verified.
4. Update the status column in `.docs/tickets/README.md`.

If any action item is unchecked, the ticket stays `in-progress`. Say which item is open. Do not write "signed off" in the result line.

## Rules

- One ticket, one type, one story.
- The ticket is the checklist. A separate progress note that the ticket does not link does not count.
- Do not sign off a ticket whose story link is missing.
- Do not delete a ticket. A withdrawn ticket stays on disk with status `draft` and a result line that says the user withdrew it, and its action items stay unchecked.
- Application code waits until status is `ready`.
