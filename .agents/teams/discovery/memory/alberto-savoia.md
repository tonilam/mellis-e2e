---
owner: alberto-savoia
team: discovery
updated: 2026-11-14
---

# Memory: Alberto Savoia

> [!warning] Owner-only
> Only the `alberto-savoia` agent writes this file. Other members don't read it. The facilitator may read it but never writes it. Rules: [Memory rules](/.agents/teams/discovery/TEAM.md#memory-rules).

## Domain

- 2026-11-14 — Local runner: the developer already has the project's application running; this runner does not start, stop, or restart it. E2E tests start with a CLI command, and the runner starts the suite by invoking that command. ([/.agents/teams/discovery/sessions/2026-11-14-local-e2e-platform.md](/.agents/teams/discovery/sessions/2026-11-14-local-e2e-platform.md))
- 2026-11-14 — Show a user story on a test only when that test's metadata already contains one. Performance for this version is how long the test took in this run. Both are Hall's defaults; report questions 1–7 are still unanswered. ([/.agents/teams/discovery/sessions/2026-11-14-local-e2e-platform.md](/.agents/teams/discovery/sessions/2026-11-14-local-e2e-platform.md))
- 2026-11-14 — The 2026-11-14 choices are requirements the user set, not an observed run. The MEH and S12–S15 were not run, so there is still no Own Data. ([/.agents/teams/discovery/sessions/2026-11-14-local-e2e-platform.md](/.agents/teams/discovery/sessions/2026-11-14-local-e2e-platform.md))

## User decisions and preferences

- 2026-11-14 — Installable Mac app and installable Windows app, each with its own window. The developer opens it from the Dock, Launchpad, Finder, or the Start menu. On open it loads the developer's login-shell environment, and the suite command and agent CLIs run with that environment. Reason: a Docker container cannot execute the project's host toolchain or the agent CLI installed on the machine. ([/.agents/teams/discovery/sessions/2026-11-14-local-e2e-platform.md](/.agents/teams/discovery/sessions/2026-11-14-local-e2e-platform.md))
- 2026-11-14 — Dropped: Docker as the place the UI, the suite, or the agent CLI runs; a terminal-only launch; a separate browser tab; the runner booting the application under test. ([/.agents/teams/discovery/sessions/2026-11-14-local-e2e-platform.md](/.agents/teams/discovery/sessions/2026-11-14-local-e2e-platform.md))
- 2026-11-14 — Accept or discard before an agent change remains is a draft default in docs/features/local-runner/run-suite.md (linked from docs/features/local-runner/README.md) that the user can still strike. It was not observed. ([/.agents/teams/discovery/sessions/2026-11-14-local-e2e-platform.md](/.agents/teams/discovery/sessions/2026-11-14-local-e2e-platform.md))
- 2026-11-14 — The build is a Go binary. The feature doc states behaviour, not the language. No application code until the user approves docs/features/local-runner/run-suite.md. ([/.agents/teams/discovery/sessions/2026-11-14-local-e2e-platform.md](/.agents/teams/discovery/sessions/2026-11-14-local-e2e-platform.md))

## Worked / didn't

- 2026-11-14 — Pass lines for S12–S15 were set and not executed before close. Do not treat the packaging decision as evidence a developer will move runs here or call an agent. ([/.agents/teams/discovery/sessions/2026-11-14-local-e2e-platform.md](/.agents/teams/discovery/sessions/2026-11-14-local-e2e-platform.md))

## Parked

- 2026-11-14 — Parked for this version: story-as-the-front-door, run history, flakiness, and traces. ([/.agents/teams/discovery/sessions/2026-11-14-local-e2e-platform.md](/.agents/teams/discovery/sessions/2026-11-14-local-e2e-platform.md))
- 2026-11-14 — S12–S15 remain unused. One window of the next 3 runs on a project named beforehand can score only one of trigger-here, watch-their-own-run, or story-list. The agent one-night stand is only after a result they stayed for, with accept or discard. Revisit only if questions 1–7 are answered before build. ([/.agents/teams/discovery/sessions/2026-11-14-local-e2e-platform.md](/.agents/teams/discovery/sessions/2026-11-14-local-e2e-platform.md))
