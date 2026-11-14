- Start development from a ticket under `.docs/tickets/`.

**always**
- follow [ticket-start](/.agents/workflows/ticket-start/SKILL.md) and copy [template.md](/.agents/workflows/ticket-start/template.md). Do not invent a different ticket shape
- the message sent with `/ticket-start` is the request. Do not ask the user to paste it again
- if that ticket file is not in `.docs/tickets/`, create it before any application code
- a ticket is exactly one of `new-story`, `update-story`, or `bug`
- acceptance criteria are pass/fail. Action items are checkboxes on the ticket. Sign off only when every action item is checked and every criterion is met

**flow**
1. Resolve the ticket (Phase 1). Create it when it is missing.
2. Fill type, story link, acceptance criteria, and action items (Phase 2). Stop and show them.
3. Wait for acceptance (Phase 3) before writing application code.
4. Do the action items, check each one in the ticket as it is finished, then sign off only when the criteria are met (Phases 4–6).
