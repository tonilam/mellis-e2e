- Orchestrate pending tickets under `.docs/tickets/` as sequential ticket-start sessions.

**always**
- follow [ticket-batch](/.agents/workflows/ticket-batch/SKILL.md)
- you are the orchestrator: manage progress, `docs/changelogs/`, discovery clarifications, verify, and sign-off
- each ticket runs in a **new** session that follows [ticket-start](/.agents/workflows/ticket-start/SKILL.md); do not implement ticket application code in this session
- the message sent with `/ticket-batch` is the request (which tickets, or all pending)

**flow**
1. Read the latest week of `docs/changelogs/` and list non-`signed-off` tickets in build order.
2. For each next ticket: set `ready` when the batch is pre-approved, spawn a ticket-start session, wait, verify ACs.
3. On `CLARIFICATION_NEEDED`, convene the analyst (small: one member per pillar); then continue.
4. Sign off only when every action item is checked and every criterion is met; append `docs/changelogs/<date>.md`.
5. When the batch ends, summarize signed-off tickets and any parked follow-ups.
