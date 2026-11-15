# Stories

Requirement stage for the local E2E runner. Each story is one capability a developer can complete, with pass/fail acceptance criteria.

Behaviour summary: [docs/features/local-runner/run-suite.md](../../docs/features/local-runner/run-suite.md). Discovery: [.agents/teams/discovery/sessions/2026-11-14-local-e2e-platform.md](../../.agents/teams/discovery/sessions/2026-11-14-local-e2e-platform.md).

| ID | Story |
| --- | --- |
| [US-1](US-1-open-the-app.md) | Install the Mac or Windows app and open it in its own window |
| [US-2](US-2-load-the-shell-environment.md) | Load the login-shell environment on open |
| [US-3](US-3-choose-the-project-and-command.md) | Choose the project path and the E2E command |
| [US-4](US-4-run-the-suite.md) | Run the suite against an application that is already running |
| [US-5](US-5-watch-the-run.md) | Watch progress, stop the run, and stop it by closing the window |
| [US-6](US-6-read-the-result.md) | Read pass, fail, duration, failure text, and a user story |
| [US-7](US-7-ask-an-agent.md) | Ask an installed agent CLI to create or fix |
| [US-8](US-8-accept-or-discard.md) | Accept or discard the change, then rerun that test |
| [US-9](US-9-laravel-login-fixture.md) | Run a Laravel login fixture for E2E |
| [US-10](US-10-open-settings-from-the-menu.md) | Open Settings from the main app menu |

## Out of scope

These have no story in this stage.

- Starting, stopping, or configuring the project's application, databases, or browsers, beyond what the E2E command itself does.
- Running the runner, the suite, or the agent CLI inside Docker.
- A history of earlier runs, a flakiness count, or a trace of each step.
- A home screen organised as a shelf of user stories.
- Sharing one run across more than one machine or more than one developer.
- Editing project files by hand inside the runner.
