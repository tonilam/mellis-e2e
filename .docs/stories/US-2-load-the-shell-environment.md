# US-2: Load the shell environment

* **User value:** As a developer, I want the installed app to use the same toolchain as a new terminal, so that the suite command and the agent CLI resolve the way they do when I type them.
* **Source:** [docs/features/local-runner/run-suite.md](../../docs/features/local-runner/run-suite.md) — Install and open
* **Tickets:** [TKT-2](../tickets/TKT-2-load-the-shell-environment.md)

## Acceptance criteria

* On open, the runner loads the developer's login-shell environment, including the PATH a new terminal would have.
* A suite command started later in the session runs with that loaded environment.
* An agent CLI started later in the session runs with that loaded environment.
* The developer can open a view of the environment the runner loaded.
* The developer can retry the load from that view.
* When the environment cannot be loaded, the window says so.
* When the environment cannot be loaded, the runner does not start a suite and does not start an agent request.
* After a failed load, a successful retry makes the suite and agent actions available again.
