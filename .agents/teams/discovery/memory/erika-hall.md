---
owner: erika-hall
team: discovery
updated: 2026-11-14
---

# Memory: Erika Hall

> [!warning] Owner-only
> Only the `erika-hall` agent writes this file. Other members don't read it. The facilitator may read it but never writes it. Rules: [Memory rules](/.agents/teams/discovery/TEAM.md#memory-rules).

## Domain

- 2026-11-14 — Local runner for a developer on their own machine: the project's application is already running, its E2E suite starts with a CLI command, and this runner starts the suite by invoking that command. The runner does not start, stop, or restart the application. ([/.agents/teams/discovery/sessions/2026-11-14-local-e2e-platform.md](/.agents/teams/discovery/sessions/2026-11-14-local-e2e-platform.md))

## User decisions and preferences

- 2026-11-14 — The runner is an installable Mac app and an installable Windows app with its own window. The developer opens it from the Dock, Launchpad, Finder, or the Start menu. On open it loads the developer's login-shell environment, and the suite command and agent CLIs run with that environment. Docker stays dropped because a container cannot execute the project's host toolchain or the agent CLI installed on the machine. ([/.agents/teams/discovery/sessions/2026-11-14-local-e2e-platform.md](/.agents/teams/discovery/sessions/2026-11-14-local-e2e-platform.md))
- 2026-11-14 — Dropped: Docker as where the UI, the suite, or the agent CLI runs; a terminal-only launch; a separate browser tab; the runner booting the application under test. ([/.agents/teams/discovery/sessions/2026-11-14-local-e2e-platform.md](/.agents/teams/discovery/sessions/2026-11-14-local-e2e-platform.md))
- 2026-11-14 — A user story is shown on a test when that test's metadata already contains one. Performance for this version is how long the test took in this run. These follow Hall's defaults because questions 1–7 are still unanswered. The decision is a requirement the user set, not a past episode of someone running a suite. ([/.agents/teams/discovery/sessions/2026-11-14-local-e2e-platform.md](/.agents/teams/discovery/sessions/2026-11-14-local-e2e-platform.md))
- 2026-11-14 — Accept or discard before an agent change stays in the draft feature doc as a default the user can still strike. The build is a Go binary; the feature doc states behaviour, not the language. No application code until the user approves `docs/features/local-runner/run-suite.md`. ([/.agents/teams/discovery/sessions/2026-11-14-local-e2e-platform.md](/.agents/teams/discovery/sessions/2026-11-14-local-e2e-platform.md))
- 2026-11-14 — ACCEPT_PROTOCOL for TKT-5 sign-off: `mellis:test` + extension points; unmatched ignored; empty progress OK. Do not pick Playwright/Jest/TAP (or any real reporter) until a project is named. Optional interim-protocol doc note only; no US-5 rewrite. ([/.agents/teams/discovery/sessions/2026-11-14-tkt5-reporter.md](/.agents/teams/discovery/sessions/2026-11-14-tkt5-reporter.md))
- 2026-11-14 — ACCEPT_CLI_DEFAULTS for TKT-7 sign-off: interim `claude` / `cursor-agent` + `-p <prompt>`, LookPath on loaded PATH, product labels when present. No US-7 AC rewrite. Optional interim binary/flag note in `docs/features/local-runner/run-suite.md` only. ([/.agents/teams/discovery/sessions/2026-11-14-tkt7-agent-cli.md](/.agents/teams/discovery/sessions/2026-11-14-tkt7-agent-cli.md))
- 2026-11-14 — ACCEPT_SNAPSHOT_AND_RERUN for TKT-8 sign-off: snapshot/restore/apply for in-place agent CLI writes; full-suite rerun with previous+new for the named test (no product single-test filter). US-8 ACs read as end state at decision UI. No US-8 AC rewrite. Optional one-liners in `docs/features/local-runner/run-suite.md` only. ([/.agents/teams/discovery/sessions/2026-11-14-tkt8-accept-discard.md](/.agents/teams/discovery/sessions/2026-11-14-tkt8-accept-discard.md))

## Worked / didn't

- 2026-11-14 — Defaults closed the open display questions when no episode existed, and the user kept them. The day-1 inquiry and the day-2 container check did not run; the user resolved packaging directly. ([/.agents/teams/discovery/sessions/2026-11-14-local-e2e-platform.md](/.agents/teams/discovery/sessions/2026-11-14-local-e2e-platform.md))
- 2026-11-14 — Same stop-rule pattern (docs/ACs answer user-visible outcomes; no named project / no episode → documented interim means) closed TKT-5 (ACCEPT_PROTOCOL), TKT-7 (ACCEPT_CLI_DEFAULTS), and TKT-8 (ACCEPT_SNAPSHOT_AND_RERUN) without fieldwork. Flip only if stakeholder forbids interim or demands new first-shippable scope. ([/.agents/teams/discovery/sessions/2026-11-14-tkt5-reporter.md](/.agents/teams/discovery/sessions/2026-11-14-tkt5-reporter.md), [/.agents/teams/discovery/sessions/2026-11-14-tkt7-agent-cli.md](/.agents/teams/discovery/sessions/2026-11-14-tkt7-agent-cli.md), [/.agents/teams/discovery/sessions/2026-11-14-tkt8-accept-discard.md](/.agents/teams/discovery/sessions/2026-11-14-tkt8-accept-discard.md))

## Parked

- 2026-11-14 — Still open, and not evidence of behaviour: report questions 1–7, which project, which command, what a user-story annotation looks like in their files, and which real reporter/adapter maps that project’s output (prior H1; confirmed parked again). ([/.agents/teams/discovery/sessions/2026-11-14-local-e2e-platform.md](/.agents/teams/discovery/sessions/2026-11-14-local-e2e-platform.md), [/.agents/teams/discovery/sessions/2026-11-14-tkt5-reporter.md](/.agents/teams/discovery/sessions/2026-11-14-tkt5-reporter.md))
- 2026-11-14 — Parked: story as the front door, run history, flakiness, traces, and a 50% XYZ over an uncounted group. ([/.agents/teams/discovery/sessions/2026-11-14-local-e2e-platform.md](/.agents/teams/discovery/sessions/2026-11-14-local-e2e-platform.md))
- 2026-11-14 — Parked: agent CLI override picker / teach-once; hard lock of binary names/flags after observed installs; user binary census until a named project fails detection. ([/.agents/teams/discovery/sessions/2026-11-14-tkt7-agent-cli.md](/.agents/teams/discovery/sessions/2026-11-14-tkt7-agent-cli.md))
- 2026-11-14 — Parked REQUIRE_CHANGE for TKT-8 unless stakeholder bans transient in-place writes as a product rule or demands a product single-test filter for first ship. ([/.agents/teams/discovery/sessions/2026-11-14-tkt8-accept-discard.md](/.agents/teams/discovery/sessions/2026-11-14-tkt8-accept-discard.md))
