---
date: 2026-11-14
request: Are snapshot/restore/apply and full-suite rerun with previous+new for the named test acceptable for first shippable — or REQUIRE_CHANGE?
attendees: [david-kelley, jeff-gothelf, erika-hall]
status: decided
---
# Discovery: TKT-8 accept/discard clarifications

**Request (verbatim):** Are (1) snapshot/restore/apply and (2) full-suite rerun with previous+new for the named test acceptable for first shippable — or REQUIRE_CHANGE?

Clarifications from TKT-8 session: (1) Agent CLIs write in-place → runner snapshots before AskAgent, restores after CLI exits so nothing remains until Accept, then applies proposal on Accept / drops on Discard. (2) No single-test filter in product → “Run this test again” re-runs the full suite command; UI shows previous + new result for the named test.

**Brief:** Small clarification of an existing flow (one member per pillar: Kelley brainstorm, Gothelf explore, Hall research). Actor is a developer on their own Mac/Windows using the local runner. Existing flow: run suite → ask installed agent → accept or discard → after accept, run that test again with previous beside new. Paths: [/.docs/stories/US-8-accept-or-discard.md](/.docs/stories/US-8-accept-or-discard.md), [docs/features/local-runner/run-suite.md](docs/features/local-runner/run-suite.md), [/.docs/tickets/TKT-8-SESSION_REPORT.md](/.docs/tickets/TKT-8-SESSION_REPORT.md), original [/.agents/teams/discovery/sessions/2026-11-14-local-e2e-platform.md](/.agents/teams/discovery/sessions/2026-11-14-local-e2e-platform.md). Constraints: sign-off gate only; optional feature one-liners; no US-8 AC rewrite unless REQUIRE_CHANGE; no application code. Evidence: US-8 ACs and run-suite.md (user-visible outcomes); TKT-8 report (documented assumptions + AC verification claims). No user research on single-test filter expectations or mid-ask transient writes. Known unknowns: whether “same test again” implies a filtered run in the user’s mind; whether transient in-place writes during the CLI (restored before decision) violate trust if read as “nothing remains until accept.”

## Framing

Point of view (Kelley): A developer who asked an agent from a result needs a clear keep-or-throw moment before anything stays in the project, then a way to see whether *that* test’s outcome moved.

How Might We (kept for this gate): make “nothing stays until I accept” feel true at decision time; keep discard as “never asked”; make “Run this test again” mean re-check *this* test’s outcome even when the suite command is the whole suite; put new beside previous so “did accept help?” is obvious.

Business / user outcomes (Gothelf): (a) After Ask, the project only keeps an agent edit when the developer Accepts; Discard leaves the pre-request tree. (b) After Accept, the developer can see whether that test’s result changed, previous beside new. Not “ship snapshot” or “ship single-test filter.”

## Idea board

**Trust the tree until I choose:** K1–K8, K14–K16 (decision after restore; deliberate Accept; discard returns to original row). G1 — decision-time restored tree + held proposal. Hall — ACs bind end state at decision UI, not “never writes during CLI.”

**See whether this test moved:** K9–K13 (named-test previous|new; optional honesty that full suite ran). G2 — same suite command + previous+new merge. Hall — ACs require same test again + previous beside new; no product single-test filter.

**Parked reopen:** REQUIRE_CHANGE only if stakeholder forbids any transient in-place write as a product rule, or demands a real single-test filter as first-shippable scope (K19, G3, Hall shortlist 2).

Counts: Kelley K1–K20; Gothelf G1–G3; Hall H1–H5.

**Shortlist:**

1. **ACCEPT_SNAPSHOT_AND_RERUN** — first-shippable means for US-8. Why: user moment preserved; ACs met as end-state; snapshot/restore/apply and full-suite rerun are means, not new user jobs.
2. **REQUIRE_CHANGE** — park. Why: only if stakeholder bans transient writes or demands product single-test filter before ship.

## Hypotheses and riskiest assumptions

| ID | Hypothesis | XYZ | Risk | Value | Evidence |
|---|---|---|---|---|---|
| G1 | Decision-time restored tree + held proposal satisfies “nothing remains until accept” even with in-place CLI writes during the request. | — | Low | High | US-8 ACs 1–4; TKT-8 restore-before-proposal claim |
| G2 | Same suite command + previous+new for the named test satisfies “run that same test again / see new beside previous.” | — | Low–med (latency) | High | US-8 ACs 5–6; original G2 |
| G3 | Requiring a single-test filter or forbidding transient in-place writes before first shippable is an output bet without evidence. | — | High if forced now | Low until stakeholder rule | No user research (brief) |
| H2 | “Nothing remains until accept” means end state when the developer decides, not zero transient writes during the CLI. | — | Low for AC reading | High for this gate | US-8 / run-suite.md plain language |

## Cheapest tests

| Assumption | Pretotype / experiment | Pass if (set in advance) | Cost |
|---|---|---|---|
| G1 restore-before-decision | One real Ask → inspect tree before Accept; Accept once; Discard once | Before Accept: matches pre-request; Accept leaves change; Discard leaves pre-request | ~30–60 min smoke |
| G2 previous+new readable | After Accept, “Run this test again”; note wait + whether previous+new for that name is readable | Developer can state old vs new for that test without asking for a filter | One suite duration |
| G3 / stakeholder flip | Stakeholder yes/no: forbid transient writes? demand single-test filter for first ship? | Explicit yes → REQUIRE_CHANGE; else stay parked | <15 min |

Nothing above blocks this sign-off; G1/G2 are post-ship measure.

## Research plan

Research questions H1–H4 closed from US-8, run-suite.md, and TKT-8 report. H5 (observed keep/discard episode) parked — same as platform session; another conversation will not invent one. Method for this gate: heuristic AC reading only. Stop rule (Hall): just enough when each US-8 AC is met as a user-visible end state under the two assumptions, neither invents a new user-facing job, and no AC rewrite is needed — default ACCEPT_SNAPSHOT_AND_RERUN. Bias noted: sponsor bias on TKT-8 “ACs met”; confirmation bias on end-state reading of “nothing remains”; sampling gap on filter expectations — mitigated by named flip criteria, not by inventing preference.

Interview guide: none for this sign-off.

## Dissent

None. Kelley PRESERVE; Gothelf ACCEPT_SNAPSHOT_AND_RERUN; Hall stop rule met → ACCEPT. Mid-ask disk peek (HMW6) and full-suite wait remain known unknowns for later measure, not reopeners without a stakeholder product rule.

## Questions for the user

None required for this sign-off. Escalate only if the stakeholder forbids any transient in-place write as a product rule, or demands a product single-test filter before first shippable.

## Outcome

**DECISION: ACCEPT_SNAPSHOT_AND_RERUN.**

- Snapshot before Ask → CLI may write in place → restore after CLI exits → held proposal → Accept applies / Discard drops is acceptable for first shippable. “Nothing remains until accept” is the end state at the decision UI.
- “Run this test again” may re-run the full suite command; the window shows previous + new for the named test. A product single-test filter is not required for this outcome.
- No US-8 AC rewrite. Optional one-liners only in [docs/features/local-runner/run-suite.md](docs/features/local-runner/run-suite.md).
- Parked: REQUIRE_CHANGE if a stakeholder later bans transient writes as a product rule or demands a single-test filter as first-shippable scope.
- Post-ship measure: one Accept + one Discard tree check; one post-Accept rerun for previous+new readability under full-suite wait.
