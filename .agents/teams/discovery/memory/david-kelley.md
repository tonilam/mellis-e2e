---
owner: david-kelley
team: discovery
updated: 2026-11-14
---

# Memory: David Kelley

> [!warning] Owner-only
> Only the `david-kelley` agent writes this file. Other members don't read it. The facilitator may read it but never writes it. Rules: [Memory rules](/.agents/teams/discovery/TEAM.md#memory-rules).

## Domain

- 2026-11-14 — Local runner for one developer on their own machine. It runs end-to-end tests that already live in a project. Session decisions are requirements the requester set, not a past episode. Questions 1–7 in the report are still unanswered. ([/.agents/teams/discovery/sessions/2026-11-14-local-e2e-platform.md](/.agents/teams/discovery/sessions/2026-11-14-local-e2e-platform.md))
- 2026-11-14 — The developer already has the project's application running locally. The runner does not start, stop, or restart that application. ([/.agents/teams/discovery/sessions/2026-11-14-local-e2e-platform.md](/.agents/teams/discovery/sessions/2026-11-14-local-e2e-platform.md))
- 2026-11-14 — The project's end-to-end tests start with a CLI command. The runner starts the suite by invoking that command. Which project and which command are still open. ([/.agents/teams/discovery/sessions/2026-11-14-local-e2e-platform.md](/.agents/teams/discovery/sessions/2026-11-14-local-e2e-platform.md))
- 2026-11-14 — A user story is shown on a test when that test's metadata already contains one. The annotation format is unknown. Performance for this version is how long the test took in this run. ([/.agents/teams/discovery/sessions/2026-11-14-local-e2e-platform.md](/.agents/teams/discovery/sessions/2026-11-14-local-e2e-platform.md))

## User decisions and preferences

- 2026-11-14 — The runner is an installable Mac app and an installable Windows app with its own window. The developer opens it from the Dock, Launchpad, Finder, or the Start menu. On open it loads the developer's login-shell environment, and the suite command and agent CLIs run with that environment. Docker stays dropped because a container cannot execute the project's host toolchain or the agent CLI installed on the machine. ([/.agents/teams/discovery/sessions/2026-11-14-local-e2e-platform.md](/.agents/teams/discovery/sessions/2026-11-14-local-e2e-platform.md))
- 2026-11-14 — Dropped: Docker as the place the UI, the suite, or the agent CLI runs; a terminal-only launch; a separate browser tab; the runner booting the application under test. ([/.agents/teams/discovery/sessions/2026-11-14-local-e2e-platform.md](/.agents/teams/discovery/sessions/2026-11-14-local-e2e-platform.md))
- 2026-11-14 — Dropped for this version: story-as-the-front-door, run history, flakiness, and traces. Story-on-the-test and this-run duration follow Hall's defaults because the questions are still open. ([/.agents/teams/discovery/sessions/2026-11-14-local-e2e-platform.md](/.agents/teams/discovery/sessions/2026-11-14-local-e2e-platform.md))
- 2026-11-14 — Accept or discard before an agent change remains is in the draft feature doc as a default the user can still strike. The build is a Go binary. The doc states behaviour, not the language. No application code until the user approves it: [docs/features/local-runner/run-suite.md](/docs/features/local-runner/run-suite.md). ([/.agents/teams/discovery/sessions/2026-11-14-local-e2e-platform.md](/.agents/teams/discovery/sessions/2026-11-14-local-e2e-platform.md))

## Worked / didn't

- 2026-11-14 — Naming the product conditions a stack would fail (the project's own toolchain, the agent CLI on the machine, and a container that must touch both) is what made Docker fail in front of the user. They chose an installable windowed app that loads the login-shell environment on open. Use that move again. ([/.agents/teams/discovery/sessions/2026-11-14-local-e2e-platform.md](/.agents/teams/discovery/sessions/2026-11-14-local-e2e-platform.md))
- 2026-11-14 — Keeping Read the suite, Act from the failure, and Watch as separate moments left the trigger decision visible. The user chose the runner invoking the project's command. Leaving the launch with the project was not chosen. ([/.agents/teams/discovery/sessions/2026-11-14-local-e2e-platform.md](/.agents/teams/discovery/sessions/2026-11-14-local-e2e-platform.md))
- 2026-11-14 — One blended moment for the result and the agent ask was too tight. The framing that held is an ordered pair: read the run, then an optional agent ask. ([/.agents/teams/discovery/sessions/2026-11-14-local-e2e-platform.md](/.agents/teams/discovery/sessions/2026-11-14-local-e2e-platform.md))

## Parked

- 2026-11-14 — Story as the front door (a shelf of stories, run one story, a story with no test is a gap) waits until metadata contains a story and the developer uses it to choose. A story read on the test row is the smaller feature that was kept. ([/.agents/teams/discovery/sessions/2026-11-14-local-e2e-platform.md](/.agents/teams/discovery/sessions/2026-11-14-local-e2e-platform.md))
- 2026-11-14 — Run history, a previous run on this machine, flakiness, and traces. This version stops at the duration of the current run. ([/.agents/teams/discovery/sessions/2026-11-14-local-e2e-platform.md](/.agents/teams/discovery/sessions/2026-11-14-local-e2e-platform.md))
- 2026-11-14 — A copy the developer brings back, and a suggestion they apply, sit beside accept-or-discard. The draft default is accept or discard before a change remains, and the user can still strike it. ([/.agents/teams/discovery/sessions/2026-11-14-local-e2e-platform.md](/.agents/teams/discovery/sessions/2026-11-14-local-e2e-platform.md))
