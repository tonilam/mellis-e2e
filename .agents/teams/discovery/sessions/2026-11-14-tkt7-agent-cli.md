---
date: 2026-11-14
request: For the first shippable runner, are claude / cursor-agent with -p acceptable interim defaults, or must we lock different binary names/flags now?
attendees: [david-kelley, jeff-gothelf, erika-hall]
status: decided
---
# Discovery: TKT-7 agent CLI defaults

**Request (verbatim):** For the first shippable runner, are `claude` / `cursor-agent` with `-p <prompt>` acceptable interim defaults (detect via LookPath on loaded PATH; show those labels when present), or must we lock different binary names/flags now?

**Brief:** Small clarification of an existing flow (US-7 / Ask an installed agent). Actors: developer on own Mac/Windows. Flow: from a suite result, choose an installed agent CLI (Claude Code CLI / Cursor CLI examples), send create/fix; accept/discard is TKT-8. Paths: [2026-11-14-local-e2e-platform.md](./2026-11-14-local-e2e-platform.md), [docs/features/local-runner/run-suite.md](../../../../docs/features/local-runner/run-suite.md), [.docs/stories/US-7-ask-an-agent.md](../../../../.docs/stories/US-7-ask-an-agent.md), [.docs/tickets/TKT-7-SESSION_REPORT.md](../../../../.docs/tickets/TKT-7-SESSION_REPORT.md). Constraints: login-shell env (TKT-2); product ACs name labels + availability, not binaries; TKT-7 assumed `claude` / `cursor-agent` + `-p`. Evidence: docs/tickets/code; no user research on real binaries. Changelog through TKT-6 only.

**Attendees (one per pillar) and why:** Kelley (brainstorm) — frame the real need vs argv lock; Gothelf (explore) — smallest bet / outcome vs output; Hall (research) — just-enough / stop rule for sign-off. Osborn, Savoia, Fitzpatrick not convened (quantity, pretotype, interview design not needed for this gate).

## Framing

POV (Kelley): developer needs to ask an installed agent after a suite result; ship needs honest availability and recognizable labels, not a forever binary lock.

How Might We (kept): continue their workflow; show only honestly available agents; ship without freezing an unresearched CLI contract; treat interim detection as revisable.

## Idea board

**recognize my agent:** K3, K16, K20, G3, H1  
**only offer what’s really there:** K1, K2, K8, K23, G1, H1  
**ship the ask without freezing argv:** K4, K5, K7, K15, K18, K21, G2, H2, H5  
**make the interim visible:** K6, K22, H3  
**later override / teach (parked):** K9, K12, K17, K24  
**wild remove-CLI (not shortlisted):** K10, K11

**Shortlist:**
1. **ACCEPT_CLI_DEFAULTS** — sign off TKT-7 with interim `claude` / `cursor-agent` + `-p <prompt>`, LookPath, product labels; document interim. Why: ACs lock labels + availability; same ACCEPT_PROTOCOL pattern as TKT-5.
2. **REQUIRE_CHANGE** — lock different binary/flags before sign-off. Why considered: Cursor binary never product-specified. Parked: no user evidence those names are wrong; delays learning the ask outcome.

## Hypotheses and riskiest assumptions

| ID | Hypothesis | XYZ | Risk | Value | Evidence |
|---|---|---|---|---|---|
| G1 | Interim defaults deliver ask-when-present / clear miss | — | Med | High | US-7; TKT-7 assumptions |
| G2 | Locking names/flags now does not raise outcome over interim | — | Low | Low | No user research on binaries |
| G3 | Product labels matter more than binary strings | — | Low | High | US-7 AC wording |
| H stop | More binary research won’t change this gate | — | — | — | Hall stop rule |

## Cheapest tests

| Assumption | Pretotype / experiment | Pass if (set in advance) | Cost |
|---|---|---|---|
| G1 after ACCEPT | One real machine with known Claude/Cursor install; one Ask from a result | Choice appears; invoke starts; miss names binary when absent | Next working session |
| Systematic miss | Only if labeled CLIs never found under interim map | Then revise that CLI’s binary/flag only | Hours if triggered |

## Research plan

Already answered: ACs and feature doc do not lock binaries/flags; interim map is documented in TKT-7 / code. Stop rule: enough for ACCEPT_CLI_DEFAULTS. Do not census installs for this gate. Park user binary census until a named project fails detection or product later locks.

## Dissent

None on the decision. All three voted ACCEPT_CLI_DEFAULTS. Board keeps a flip condition (Hall): REQUIRE_CHANGE only if stakeholder forbids interim, or light check shows named CLIs systematically never found.

## Questions for the user

None for this gate (unanimous; ACs already omit binaries).

## Outcome

**DECISION: ACCEPT_CLI_DEFAULTS** — sign off TKT-7 as-is on interim defaults (`claude` / `cursor-agent` with `-p <prompt>`; LookPath on loaded PATH; show “Claude Code CLI” / “Cursor CLI” when present). No US-7 AC rewrite. Optional product note in [run-suite.md](../../../../docs/features/local-runner/run-suite.md): binary names and `-p` are interim and revisable. Parked: machine override, teach-once picker, hard lock after observed installs.
