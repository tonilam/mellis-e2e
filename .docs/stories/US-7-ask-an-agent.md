# US-7: Ask an installed agent

* **User value:** As a developer, I want to send a create or fix request to an agent CLI already installed on my machine, so that the change is made in the project I am looking at.
* **Source:** [docs/features/local-runner/run-suite.md](../../docs/features/local-runner/run-suite.md) — Ask an installed agent; Edge cases
* **Tickets:** [TKT-7](../tickets/TKT-7-ask-an-agent.md)

## Acceptance criteria

* From a result, the developer can ask for one of three actions: create a test, fix this test, or fix the application code this test is about.
* Before the request is sent, the developer chooses which installed agent CLI to call.
* When Claude Code CLI is installed and available in the loaded environment, it is one of the choices.
* When the Cursor CLI is installed and available in the loaded environment, it is one of the choices.
* The request includes the test name and the pass or fail.
* When the result has failure text, the request includes that text.
* When the test has a user story, the request includes that story.
* While the agent is working, the window shows that the agent is working.
* The developer can stop the agent request on its own. Stopping it leaves the reported test results in the window.
* When the chosen CLI cannot be started, the project path, the run, and the results still work.
* When the chosen CLI cannot be started, the agent actions say that the CLI is unavailable.
* When the chosen CLI is missing from the loaded environment, the window names that CLI.
