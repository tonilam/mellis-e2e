---
owner: rob-fitzpatrick
team: discovery
updated: 2026-11-14
---

# Memory: Rob Fitzpatrick

> [!warning] Owner-only
> Only the `rob-fitzpatrick` agent writes this file. Other members don't read it. The facilitator may read it but never writes it. Rules: [Memory rules](/.agents/teams/discovery/TEAM.md#memory-rules).

## Domain

2026-11-14 — A developer on their own machine already has the project's application running. This runner does not start, stop, or restart it. The project's E2E tests start with a CLI command, and the runner starts the suite by invoking that command. ([/.agents/teams/discovery/sessions/2026-11-14-local-e2e-platform.md](/.agents/teams/discovery/sessions/2026-11-14-local-e2e-platform.md))
2026-11-14 — The runner is an installable Mac app and an installable Windows app with its own window. The developer opens it from the Dock, Launchpad, Finder, or the Start menu. On open it loads the developer's login-shell environment, and the suite command and agent CLIs run with that environment. A Docker container cannot execute the project's host toolchain or the agent CLI installed on the machine, so Docker stays dropped. ([/.agents/teams/discovery/sessions/2026-11-14-local-e2e-platform.md](/.agents/teams/discovery/sessions/2026-11-14-local-e2e-platform.md))
2026-11-14 — For this version, a user story is shown on a test when that test's metadata already contains one, and performance is how long the test took in this run. Both are Hall's defaults while questions 1–7 are open, not observed behaviour. ([/.agents/teams/discovery/sessions/2026-11-14-local-e2e-platform.md](/.agents/teams/discovery/sessions/2026-11-14-local-e2e-platform.md))

## User decisions and preferences

2026-11-14 — The close is requirements the requester set, not a past episode of someone running a suite. Questions 1–7 in the report are still unanswered. ([/.agents/teams/discovery/sessions/2026-11-14-local-e2e-platform.md](/.agents/teams/discovery/sessions/2026-11-14-local-e2e-platform.md))
2026-11-14 — The user corrected the launch the same day. The runner is an installable Mac app and an installable Windows app with its own window, opened from the Dock, Launchpad, Finder, or the Start menu. On open it loads the login-shell environment; the suite command and agent CLIs run with that environment. Dropped: a terminal-only launch, a separate browser tab, and Docker as the place the UI, the suite, or the agent CLI runs. The runner still does not boot the application under test. Reason Docker stays dropped: the run has to see the host toolchain and the installed agent CLI. ([/.agents/teams/discovery/sessions/2026-11-14-local-e2e-platform.md](/.agents/teams/discovery/sessions/2026-11-14-local-e2e-platform.md))
2026-11-14 — Dropped for this version: story-as-the-front-door, run history, flakiness, and traces. ([/.agents/teams/discovery/sessions/2026-11-14-local-e2e-platform.md](/.agents/teams/discovery/sessions/2026-11-14-local-e2e-platform.md))
2026-11-14 — Accept or discard before an agent change remains is a draft default in docs/features/local-runner/run-suite.md (linked from docs/features/local-runner/README.md). The user can still strike it. It was not observed. ([/.agents/teams/discovery/sessions/2026-11-14-local-e2e-platform.md](/.agents/teams/discovery/sessions/2026-11-14-local-e2e-platform.md))
2026-11-14 — The build is a Go binary. The feature doc states behaviour, not the language. No application code until the user approves the doc. ([/.agents/teams/discovery/sessions/2026-11-14-local-e2e-platform.md](/.agents/teams/discovery/sessions/2026-11-14-local-e2e-platform.md))

## Worked / didn't

2026-11-14 — Worked: keeping the request labeled as intent. The outcome still calls the choices requirements, and questions 1–7 stayed unanswered. ([/.agents/teams/discovery/sessions/2026-11-14-local-e2e-platform.md](/.agents/teams/discovery/sessions/2026-11-14-local-e2e-platform.md))
2026-11-14 — Worked: splitting a run from an agent request, with Savoia. The decision starts the suite by command and leaves accept-or-discard as a strikeable default, so one yes did not cover both acts. ([/.agents/teams/discovery/sessions/2026-11-14-local-e2e-platform.md](/.agents/teams/discovery/sessions/2026-11-14-local-e2e-platform.md))
2026-11-14 — Worked: naming the host-toolchain and installed-CLI consequence without picking a stack. That was the user's reason for dropping Docker. ([/.agents/teams/discovery/sessions/2026-11-14-local-e2e-platform.md](/.agents/teams/discovery/sessions/2026-11-14-local-e2e-platform.md))
2026-11-14 — Didn't: the interview was not run. A packaging yes from the requester is not evidence of a past run. Hall had flagged that xyz yes as sponsor bias. ([/.agents/teams/discovery/sessions/2026-11-14-local-e2e-platform.md](/.agents/teams/discovery/sessions/2026-11-14-local-e2e-platform.md))

## Parked

2026-11-14 — Parked: story-as-the-front-door, run history, flakiness, traces, and the uncounted 50% XYZ. Revisit the story shelf only if a named project's files already contain a user story and the developer uses it to choose. ([/.agents/teams/discovery/sessions/2026-11-14-local-e2e-platform.md](/.agents/teams/discovery/sessions/2026-11-14-local-e2e-platform.md))
2026-11-14 — Parked as unconfirmed: which project, which command, and what a user-story annotation looks like in the files. ([/.agents/teams/discovery/sessions/2026-11-14-local-e2e-platform.md](/.agents/teams/discovery/sessions/2026-11-14-local-e2e-platform.md))
