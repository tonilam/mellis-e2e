# Workflows

Phased work an agent runs in order. The skill file is the source of truth. Do not improvise a second version of the steps.

| Workflow | Use it when |
| --- | --- |
| [ticket-start](ticket-start/SKILL.md) | Development starts. The work is a new user story, a change to an existing story, or a bug in a story. The ticket lives in `.docs/tickets/`. |

## ticket-start

Start from a ticket. When the ticket file is not in `.docs/tickets/`, create it from [ticket-start/template.md](ticket-start/template.md) before writing code.

The ticket is one of:

1. `new-story` — add a user story under `.docs/stories/`
2. `update-story` — change an existing story's acceptance criteria
3. `bug` — the story already says what should happen, and the product does something else

Every ticket has pass/fail acceptance criteria and checkbox action items. Sign off only when every action item is checked and every criterion is met.
