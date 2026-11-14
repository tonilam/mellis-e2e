# US-6: Read the result

* **User value:** As a developer, I want each test to show whether it passed, how long it took, its failure text, and its user story when the test already has one, so that I can decide what to do next from this run alone.
* **Source:** [docs/features/local-runner/run-suite.md](../../docs/features/local-runner/run-suite.md) — Results; Edge cases
* **Tickets:** [TKT-6](../tickets/TKT-6-read-the-result.md)

## Acceptance criteria

* Each reported test shows whether it passed or failed.
* When the output includes a duration, the test shows how long it took in this run.
* When the output includes no duration, the test shows pass or fail and shows no duration.
* When the test's metadata contains a user story, the row shows that story.
* When the test's metadata contains a user story, opening the test shows the story text.
* When the test's metadata has no user story, the row shows the test name and shows no story.
* When a failed test's output includes failure text, the test shows that text.
* The finished window remains usable when the developer does not ask an agent to change the project.
* When the command finishes and reports no tests, the window says the run finished with no tests.
