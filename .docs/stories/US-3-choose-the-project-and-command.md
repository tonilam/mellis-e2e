# US-3: Choose the project and the command

* **User value:** As a developer, I want to point the runner at a project and confirm the command that starts its end-to-end suite, so that the run uses the command that project already has.
* **Source:** [docs/features/local-runner/run-suite.md](../../docs/features/local-runner/run-suite.md) — Choose the project and the command; Edge cases
* **Tickets:** [TKT-3](../tickets/TKT-3-choose-the-project-and-command.md)

## Acceptance criteria

* The developer can enter a project path.
* The window shows that path for as long as the session lasts.
* The developer can enter the command that starts the E2E suite.
* The window shows that command before a run starts.
* The developer can change the command before the run starts.
* When this machine already has a path and command from an earlier session, the next session offers that same path and command for the developer to confirm or change.
* When the project path does not exist, the window says so and the runner does not start a command.
* When the command cannot be started, the window says so and the runner shows no invented test results.
* When the command is missing from the loaded environment, the window names that command and the runner shows no invented test results.
