---
name: ticket-batch
description: >-
  Orchestrate pending tickets in .docs/tickets/ as sequential ticket-start
  sessions. Manages progress, docs/changelogs, discovery clarifications,
  verification, and sign-off. Use when the user asks to work all pending
  tickets, run a ticket batch, ship the next tickets, or repeat the
  ticket-batch procedure.
---

# Workflow: ticket-batch

The orchestrator does **not** implement tickets. It drives one [ticket-start](../ticket-start/SKILL.md) session per ticket, verifies results, updates `docs/changelogs/`, routes clarifications to discovery, and signs off.

## When to use

- Pending tickets in `.docs/tickets/` need delivery in build order
- User asks to batch-work tickets, manage sessions, and sign off
- After a prior batch, repeat for remaining `draft` / `ready` tickets

## Roles

| Role | Who | Does |
| --- | --- | --- |
| Orchestrator | This session | Queue, spawn sessions, verify, changelog, discovery, sign-off, skill/process notes |
| Ticket session | Fresh subagent per ticket | [ticket-start](../ticket-start/SKILL.md) Phase 4–5 only (when pre-approved) |
| Discovery | [analyst](../../roles/product.md) + members | Clarifications only; one member per pillar for small extensions |

Do not play discovery members in the orchestrator. Do not implement application code in the orchestrator except tiny doc fixes discovery requested (feature-doc one-liners).

## Prep

1. Read the latest week of `docs/changelogs/` (create the folder if missing).
2. Read `.docs/tickets/README.md` and every non-`signed-off` ticket.
3. Honour declared build order (e.g. “Build before TKT-N”). Default: ascending `TKT-<n>`.
4. Skip tickets already `signed-off`. Prefer `ready` / `in-progress` next; for a user-requested batch of filled `draft` tickets, set the next one to `ready` when starting it (Phase 3 pre-approved — see below).

## Phase 3 in a batch

When the user asks this workflow to **work, verify, and sign off** pending tickets whose AC and action items are already filled:

- Treat that ask as Phase 3 acceptance for each ticket in the batch.
- Set status `ready` before spawning the session.
- Tell the session: do **not** wait for another user confirmation; do **not** check Sign-off boxes; leave status `in-progress` when AIs are done.

If a ticket is still empty (no AC/AI), run full ticket-start Phase 2 in a session and stop for user acceptance before coding.

## Per-ticket loop (sequential)

For each next ticket:

### 1. Spawn a ticket-start session

New subagent (`generalPurpose`), background OK. Prompt must include:

- Repo root, ticket path, story path, feature doc, latest changelog paths
- Prior signed-off tickets and where the app lives (paths only)
- Environment tips learned here (e.g. `export PATH="$HOME/go/bin:$PATH"` for `wails`; `CGO_LDFLAGS="-framework UniformTypeIdentifiers"` on this Mac SDK if build fails on UTType)
- Stay inside this ticket only
- Check each `AI-*` when finished; leave Sign-off unchecked
- On product ambiguity: stop with `CLARIFICATION_NEEDED` (question, known facts, options) — do not invent product decisions
- Final `SESSION_REPORT`: status, AIs, how each AC was verified, files, clarifications, one-line changelog suggestion
- Do not write `docs/changelogs/`; do not commit unless the ticket requires it

### 2. Wait for the session; then verify

Re-run the session’s verification commands yourself (at least `go test ./...` when Go code changed). Spot-check AC evidence against the product and tests. If an AI was checked but the criterion is not met, uncheck it and resume the same session with the gap — do not sign off.

Cross-platform ACs on a single host: packaging + installer wiring can count when live OS open is impossible; record that limit in the Result line.

### 3. Clarifications → discovery

If the session (or your verify) raises product ambiguity:

1. Convene the [analyst](../../roles/product.md).
2. Small extension of an existing flow: **one member per pillar** (three subagents), not the full six.
3. Same clarification family as an open discovery session: resume the analyst instead of a duplicate full round.
4. Analyst writes `/.agents/teams/discovery/sessions/<YYYY-MM-DD>-<slug>.md` and returns a `DECISION` block.
5. `ACCEPT_*`: optional feature-doc one-liner only; then sign off. `REQUIRE_CHANGE`: resume the ticket session with the decision; re-verify.

Bug restores need no discovery. New capability (not a ticket clarification) uses the full discovery team before requirements.

### 4. Sign off and changelog

Only when every action item is checked and every AC is met (after discovery if needed):

1. Check both Sign-off boxes; set `status: signed-off`; Date; one-line Result (what was verified + discovery decision link if any).
2. Update `.docs/tickets/README.md` status column.
3. Append under `docs/changelogs/<YYYY-MM-DD>.md` a `## TKT-n — <title>` heading and the session’s changelog bullet (or your edit of it).
4. Mark the next ticket `ready` and spawn its session.

### 5. Close the batch

When no pending tickets remain (or the user scoped a subset):

- Summarize signed-off ids and changelog path.
- Note open follow-ups discovery parked (e.g. real reporter adapters, CLI binary map).
- If the user asked for a repeatable skill/command, update this workflow from new lessons (keep it short).

## Session prompt skeleton

```
You are a ticket-start development session.
Follow /.agents/workflows/ticket-start/SKILL.md for TKT-<n> only.
Repo: <root>
Ticket: <.docs/tickets/...>
Story / feature / changelog: <paths>
Prior work: <paths only>
Status is ready (Phase 3 pre-approved for this batch). Do not wait for the user.
Set in-progress when AI-1 starts; update tickets README.
Check AI-* when done. Do NOT sign off / do NOT set signed-off.
Stay in this ticket. CLARIFICATION_NEEDED on product ambiguity.
PATH tip: export PATH="$HOME/go/bin:$PATH" before wails build.
Deliver SESSION_REPORT. No docs/changelogs. No commits.
Start Phase 4 now.
```

## Lessons locked in (2026-11-14 batch)

- One ticket → one new session; orchestrator never implements the ticket body.
- Sequential when tickets declare order; do not parallelize dependent tickets.
- `CLARIFICATION_NEEDED` is a gate: discovery before sign-off, not after.
- Interim protocols (e.g. `mellis:test`, CLI binary defaults) are product decisions via discovery `ACCEPT_*`, then one-liners in `docs/features/`, not silent invention.
- Sessions write optional `.docs/tickets/TKT-n-SESSION_REPORT.md`; orchestrator owns changelog and Sign-off.
- Avoid repo-wide `find` for tooling; prefer known paths (`$(go env GOPATH)/bin`).

## Related

- Single ticket: [ticket-start](../ticket-start/SKILL.md), command [/.cursor/commands/ticket-start.md](../../../.cursor/commands/ticket-start.md)
- Discovery: [/.agents/teams/discovery/TEAM.md](../../teams/discovery/TEAM.md)
- Command entry: [/.cursor/commands/ticket-batch.md](../../../.cursor/commands/ticket-batch.md)
