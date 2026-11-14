---
date: 2026-11-14
request: For the first shippable runner, is mellis:test + extension points acceptable until a project/reporter is named, or must we pick Playwright/Jest/TAP now?
attendees: [david-kelley, jeff-gothelf, erika-hall]
status: decided
---
# Discovery: TKT-5 reporter clarification

**Request (verbatim):** For the first shippable runner, is the mellis:test protocol + extension points acceptable until a real project/reporter is named — OR must we pick Playwright/Jest/TAP (etc.) now?

**Brief:** Small clarification of an existing flow (Watch the run / US-5 / TKT-5), not a new capability. One member per pillar. Actors: developer on their machine. Entities: suite command, live progress rows, line protocol. Existing flow: runner executes the project’s CLI command; tests the output reports move waiting→running→passed/failed ([docs/features/local-runner/run-suite.md](/docs/features/local-runner/run-suite.md), [.docs/stories/US-5-watch-the-run.md](/.docs/stories/US-5-watch-the-run.md)). TKT-5 shipped `mellis:test <status> <name>` plus adapter extension points; unmatched lines ignored ([.docs/tickets/TKT-5-SESSION_REPORT.md](/.docs/tickets/TKT-5-SESSION_REPORT.md)). Prior discovery left the reporter unset until a project is named ([sessions/2026-11-14-local-e2e-platform.md](/.agents/teams/discovery/sessions/2026-11-14-local-e2e-platform.md)). No named project; Questions 1–7 still open. Evidence: those docs only — no user research.

## Attendees (why)

| Pillar | Member | Why |
|---|---|---|
| Brainstorm | Kelley | POV / HMW on trusting progress without locking a stack; Osborn’s quantity would expand the gate |
| Explore | Gothelf | Smallest bet and outcomes for sign-off; Savoia’s “build or not” is already past TKT-5 |
| Research | Hall | Stop rule / just enough to decide; Fitzpatrick’s interview guide waits on a named project |

## Framing

POV (Kelley): a developer shipping the first local runner needs to trust that live progress means what the suite reported — not invented tests — because without a named project they cannot know which reporter their suite will speak.

HMW kept: see live progress without naming a project; keep “what the suite reported” as the only source of truth; treat a temporary language as a bridge; finish honestly when the suite is silent; do not define E2E as one framework’s reporter.

## Idea board / synthesis

Accept interim protocol: K1, K4, K5, K11, K17 · Empty progress as honest finish: K2, K7, K8, K15, K16 · Extension points for later project: K9, K10, K18 · Gothelf G1 ACCEPT_PROTOCOL as the sign-off bet; G2 REQUIRE_REPORTER falsifies if first project differs; G3 outcome is the seam, not owning a native reporter · Hall H1–H5: US-5 does not name a reporter; no project named; empty progress documented; prior standing decision still holds.

**Shortlist (gate only):** ACCEPT_PROTOCOL — why: matches prior discovery, US-5 wording, and Hall’s stop rule. REQUIRE_REPORTER — parked until a project is named.

## Hypotheses and riskiest assumptions

| ID | Hypothesis | Risk | Value | Evidence |
|---|---|---|---|---|
| G1 | ACCEPT_PROTOCOL achieves first-shippable sign-off without inventing a framework | Low for this gate | High | Prior discovery; TKT-5; US-5 |
| G2 | REQUIRE_REPORTER is needed for watched runs on first ship | High (wrong lock) | Speculative | No named project |
| H-stop | Enough to decide when docs answer H1/H4 and no project is named → default ACCEPT | — | — | Hall Round 1 |

## Cheapest tests

| Assumption | Experiment | Pass if | Cost |
|---|---|---|---|
| Real suite can report progress after a project is named | Name one project + command; one watched run (Gothelf §8) | Progress via protocol or thin adapter; fail only if reportable lines are impossible without a product-level lock | One session — **after** this gate |

## Research plan

Already answered from docs: US-5 / run-suite do not require a named reporter; TKT-5 empty progress + no invented tests; prior stack note left reporter unset. Do not research Playwright vs Jest vs TAP until a project is named. Stop rule (Hall): if no project named and ACs do not require a reporter → ACCEPT_PROTOCOL.

## Dissent

None on the gate. All three support ACCEPT_PROTOCOL. Residual risk (Gothelf): protocol lines may never appear in a real suite until a project is named — that is the next learning bet, not a reason to lock a reporter now.

## Questions for the user

None shape-changing for this gate. Still open from the parent session (not reopened here): name the first project and its suite command.

## Outcome

**Decision: ACCEPT_PROTOCOL.**

- Sign off TKT-5 as-is: `mellis:test` + extension points; unmatched ignored; empty progress OK.
- Do **not** pick Playwright/Jest/TAP now.
- Real-reporter / first adapter is a follow-up clarification once a project is named (same park as prior H1).
- Optional doc tweak (not required for sign-off): one line in run-suite or the ticket that mellis:test is the interim way output can report progress until a project’s reporter is mapped — US-5 ACs unchanged.
- No story or feature-doc rewrite required by this team.

### 2026-11-14 addendum: TKT-6 optional fields

**Decision: ACCEPT_PROTOCOL_FIELDS.** TKT-6’s optional `duration=` / `story=` / `failure=` on the same `mellis:test` line are interim wire shape only — they carry fields run-suite already shows when the output includes them; unmatched lines stay ignored; Playwright/Jest metadata encoding remains unspecified until a project is named. No new discovery round. Sign off TKT-6 as-is. Optional doc tweak: note those optional keys on the interim protocol line (same place as the TKT-5 interim note); no AC rewrite.
