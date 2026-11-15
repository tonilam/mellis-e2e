# US-5: Watch the run

* **User value:** As a developer, I want to see the suite move test by test and to stop it, so that I know what is still going and can end a run I no longer need.
* **Source:** [docs/features/local-runner/run-suite.md](../../docs/features/local-runner/run-suite.md) — Run; Edge cases
* **Tickets:** [TKT-5](../tickets/TKT-5-watch-the-run.md), [TKT-11](../tickets/TKT-11-wrap-run-session-output.md)

## Acceptance criteria

* While the command is running, each test the output reports moves from waiting, to running, to passed or failed.
* While the command is running, the window shows how many tests are finished, how many are running, and how many are still waiting.
* The developer can stop the run.
* After a stop, results already reported stay in the window.
* After a stop, the command is no longer running.
* When the command exits on its own, the run is finished.
* When the developer closes the window while a run is going, the runner stops the command.
* When the run session shows a line longer than the window width, that line wraps so the full line stays within the window's max width without horizontal overflow.
