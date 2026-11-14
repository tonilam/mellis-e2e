---
owner: jeff-gothelf
team: discovery
updated: 2026-11-14
---

# Memory: Jeff Gothelf

> [!warning] Owner-only
> Only the `jeff-gothelf` agent writes this file. Other members don't read it. The facilitator may read it but never writes it. Rules: [Memory rules](/.agents/teams/discovery/TEAM.md#memory-rules).

## Domain

2026-11-14 — Local runner for a developer on their own machine: point at a project that already has E2E tests, start the suite by invoking that project's CLI command, and show pass, fail, how long each test took in this run, and a user story when that test's metadata already contains one. ([/.agents/teams/discovery/sessions/2026-11-14-local-e2e-platform.md](/.agents/teams/discovery/sessions/2026-11-14-local-e2e-platform.md))

2026-11-14 — The project's application must already be running locally. This runner does not start, stop, or restart it. ([/.agents/teams/discovery/sessions/2026-11-14-local-e2e-platform.md](/.agents/teams/discovery/sessions/2026-11-14-local-e2e-platform.md))

2026-11-14 — Packaging is an installable Mac app and an installable Windows app with its own window. The developer opens it from the Dock, Launchpad, Finder, or the Start menu. On open it loads the developer's login-shell environment, and the suite command and agent CLIs run with that environment. ([/.agents/teams/discovery/sessions/2026-11-14-local-e2e-platform.md](/.agents/teams/discovery/sessions/2026-11-14-local-e2e-platform.md))

2026-11-14 — An agent CLI already installed on the machine may be asked to create a test, fix the code of a failed test, or fix a test. Claude Code CLI and Cursor CLI were examples. ([/.agents/teams/discovery/sessions/2026-11-14-local-e2e-platform.md](/.agents/teams/discovery/sessions/2026-11-14-local-e2e-platform.md))

2026-11-14 — First-shippable Ask detects Claude Code CLI / Cursor CLI via LookPath as interim `claude` / `cursor-agent` with `-p <prompt>` on the loaded PATH; product labels are the choice names; binaries/flags are revisable. ([/.agents/teams/discovery/sessions/2026-11-14-tkt7-agent-cli.md](/.agents/teams/discovery/sessions/2026-11-14-tkt7-agent-cli.md))

2026-11-14 — Decisions in this session are requirements the requester set, not an observed run. Questions 1–7 are unanswered, and no project, suite command, or story-annotation format is known yet. ([/.agents/teams/discovery/sessions/2026-11-14-local-e2e-platform.md](/.agents/teams/discovery/sessions/2026-11-14-local-e2e-platform.md))

2026-11-14 — First-shippable live progress uses an interim `mellis:test <waiting|running|passed|failed> <name>` line protocol plus adapter extension points; unmatched lines are ignored and empty progress is OK when a suite emits none. ([/.agents/teams/discovery/sessions/2026-11-14-tkt5-reporter.md](/.agents/teams/discovery/sessions/2026-11-14-tkt5-reporter.md))

2026-11-14 — First-shippable accept/discard means: agent CLIs may write in-place; runner snapshots before Ask, restores after CLI exit so the tree matches pre-request until Accept; Accept applies the held proposal; Discard drops it. “Run this test again” re-runs the full suite command and shows previous+new for the named test (no product single-test filter). ([/.agents/teams/discovery/sessions/2026-11-14-tkt8-accept-discard.md](/.agents/teams/discovery/sessions/2026-11-14-tkt8-accept-discard.md))

## User decisions and preferences

2026-11-14 — Installable Mac and Windows apps with their own window, because a Docker container cannot execute the project's host toolchain or the agent CLI on the machine. Docker stays dropped. A terminal-only launch and a separate browser tab are dropped. ([/.agents/teams/discovery/sessions/2026-11-14-local-e2e-platform.md](/.agents/teams/discovery/sessions/2026-11-14-local-e2e-platform.md))

2026-11-14 — This version's performance is this-run duration, and a user story appears on a test only when metadata already contains one. Both are Hall's defaults while questions 1–7 stay open. ([/.agents/teams/discovery/sessions/2026-11-14-local-e2e-platform.md](/.agents/teams/discovery/sessions/2026-11-14-local-e2e-platform.md))

2026-11-14 — Accept or discard before an agent change remains: outcomes stand (nothing sticks until Accept; Discard restores pre-request; after Accept, previous+new for that test). Signed ACCEPT_SNAPSHOT_AND_RERUN for TKT-8 means; no US-8 AC rewrite; optional run-suite.md one-liners. ([/.agents/teams/discovery/sessions/2026-11-14-tkt8-accept-discard.md](/.agents/teams/discovery/sessions/2026-11-14-tkt8-accept-discard.md))

2026-11-14 — The feature doc states behaviour, not language. The build is a Go binary. No application code until the user approves docs/features/local-runner/run-suite.md (linked from docs/features/local-runner/README.md). ([/.agents/teams/discovery/sessions/2026-11-14-local-e2e-platform.md](/.agents/teams/discovery/sessions/2026-11-14-local-e2e-platform.md))

2026-11-14 — ACCEPT_PROTOCOL: sign off TKT-5 as-is (`mellis:test` + extension points). Do not pick Playwright/Jest/TAP now. Optional doc note that `mellis:test` is interim; no US-5/story rewrite. ([/.agents/teams/discovery/sessions/2026-11-14-tkt5-reporter.md](/.agents/teams/discovery/sessions/2026-11-14-tkt5-reporter.md))

2026-11-14 — ACCEPT_CLI_DEFAULTS: sign off TKT-7 as-is (interim `claude` / `cursor-agent` + `-p <prompt>`, LookPath, product labels). No US-7 AC rewrite. Optional run-suite.md note that binaries/flags are interim. ([/.agents/teams/discovery/sessions/2026-11-14-tkt7-agent-cli.md](/.agents/teams/discovery/sessions/2026-11-14-tkt7-agent-cli.md))

2026-11-14 — ACCEPT_SNAPSHOT_AND_RERUN: sign off TKT-8 means (snapshot/restore/apply; full-suite rerun with previous+new for the named test). Outcomes over outputs; REQUIRE_CHANGE only if stakeholder bans transient writes or demands a product single-test filter. ([/.agents/teams/discovery/sessions/2026-11-14-tkt8-accept-discard.md](/.agents/teams/discovery/sessions/2026-11-14-tkt8-accept-discard.md))

## Worked / didn't

2026-11-14 — Calling out that a container cannot run the host toolchain or the installed CLI is the constraint the user confirmed. Shipping a container is an output, and it was dropped. ([/.agents/teams/discovery/sessions/2026-11-14-local-e2e-platform.md](/.agents/teams/discovery/sessions/2026-11-14-local-e2e-platform.md))

2026-11-14 — G1-first (hold the agent until a suite is readable) conflicted with Savoia's xyz, which needs an agent request inside 5 working days. The user set requirements instead. The xyz, the 50% XYZ, and G1–G4 were not confirmed by behaviour. ([/.agents/teams/discovery/sessions/2026-11-14-local-e2e-platform.md](/.agents/teams/discovery/sessions/2026-11-14-local-e2e-platform.md))

2026-11-14 — G4 (Mac and Windows each finish a run and an agent request from a Docker container) dropped with the container. The remaining behaviour is finishing them in the installable app's own window, with the suite command and agent CLIs running in the login-shell environment loaded on open. ([/.agents/teams/discovery/sessions/2026-11-14-local-e2e-platform.md](/.agents/teams/discovery/sessions/2026-11-14-local-e2e-platform.md))

2026-11-14 — Story-as-the-front-door was dropped. A story already in metadata is shown on the test, which is not evidence they choose the failure because of the story. ([/.agents/teams/discovery/sessions/2026-11-14-local-e2e-platform.md](/.agents/teams/discovery/sessions/2026-11-14-local-e2e-platform.md))

2026-11-14 — Outcome-vs-output sign-off landed three times: ACCEPT_PROTOCOL (TKT-5), ACCEPT_CLI_DEFAULTS (TKT-7), ACCEPT_SNAPSHOT_AND_RERUN (TKT-8). Decision-time tree state and named-test previous+new held; locking filter or forbidding transient CLI writes did not. Next learning: one real Accept+Discard, then one post-Accept rerun on a named project. ([/.agents/teams/discovery/sessions/2026-11-14-tkt5-reporter.md](/.agents/teams/discovery/sessions/2026-11-14-tkt5-reporter.md), [/.agents/teams/discovery/sessions/2026-11-14-tkt7-agent-cli.md](/.agents/teams/discovery/sessions/2026-11-14-tkt7-agent-cli.md), [/.agents/teams/discovery/sessions/2026-11-14-tkt8-accept-discard.md](/.agents/teams/discovery/sessions/2026-11-14-tkt8-accept-discard.md))

## Parked

2026-11-14 — Story-as-the-front-door, run history, flakiness, and traces. Revisit if a later session finds a story in metadata or a use beyond this-run duration. ([/.agents/teams/discovery/sessions/2026-11-14-local-e2e-platform.md](/.agents/teams/discovery/sessions/2026-11-14-local-e2e-platform.md))

2026-11-14 — Copy of the project (O48) or suggestion the developer applies (O56): revisit if Accept/Discard is struck or restore/trust fails on a real machine. ([/.agents/teams/discovery/sessions/2026-11-14-local-e2e-platform.md](/.agents/teams/discovery/sessions/2026-11-14-local-e2e-platform.md), [/.agents/teams/discovery/sessions/2026-11-14-tkt8-accept-discard.md](/.agents/teams/discovery/sessions/2026-11-14-tkt8-accept-discard.md))

2026-11-14 — Real project reporter (Playwright/Jest/TAP/etc.) until a project is named; then adapt via extension points. ([/.agents/teams/discovery/sessions/2026-11-14-tkt5-reporter.md](/.agents/teams/discovery/sessions/2026-11-14-tkt5-reporter.md))

2026-11-14 — Agent CLI machine override / teach-once picker, and hard-locking binaries after observed installs. Revisit if LookPath systematically misses a named CLI or product later locks. ([/.agents/teams/discovery/sessions/2026-11-14-tkt7-agent-cli.md](/.agents/teams/discovery/sessions/2026-11-14-tkt7-agent-cli.md))

2026-11-14 — REQUIRE_CHANGE / product single-test filter / ban on transient in-place CLI writes: parked unless a stakeholder raises that product rule. ([/.agents/teams/discovery/sessions/2026-11-14-tkt8-accept-discard.md](/.agents/teams/discovery/sessions/2026-11-14-tkt8-accept-discard.md))
