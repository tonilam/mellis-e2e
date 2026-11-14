# US-8: Accept or discard the change

* **User value:** As a developer, I want to accept or discard an agent's proposed change before it remains, so that my project changes only when I keep the edit, and I can rerun that test to see the new result.
* **Source:** [docs/features/local-runner/run-suite.md](../../docs/features/local-runner/run-suite.md) — Ask an installed agent
* **Tickets:** [TKT-8](../tickets/TKT-8-accept-or-discard.md)

## Acceptance criteria

* The agent shows a proposed change before that change remains in the project.
* The developer can accept the proposed change. After accept, the change is in the project.
* The developer can discard the proposed change. After discard, the project files match the tree from before the request.
* Nothing from the request remains in the project until the developer accepts.
* After an accepted change, the developer can run that same test again from the window.
* After that rerun, the window shows the new result beside the previous result for that test.
