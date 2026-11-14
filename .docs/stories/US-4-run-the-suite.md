# US-4: Run the suite

* **User value:** As a developer, I want to start the project's end-to-end command from the runner, so that the tests run against the application I already have running.
* **Source:** [docs/features/local-runner/run-suite.md](../../docs/features/local-runner/run-suite.md) — Run; Edge cases
* **Tickets:** [TKT-4](../tickets/TKT-4-run-the-suite.md)

## Acceptance criteria

* The developer can start the suite from the window.
* The runner executes the command shown for the session, in the project path shown for the session, with the environment loaded on open.
* Starting the suite does not start the project's application.
* Starting the suite does not stop or restart the project's application.
* When the project's application is not running, the runner still executes the command.
* When the project's application is not running, a failure on screen is the failure the command reported.
* When the project's application is not running, the runner does not start that application.
