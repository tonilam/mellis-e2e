---
id: TKT-7
type: new-story
status: signed-off
story: US-7
---

# TKT-7: Ask an installed agent

**Type:** new-story
**Story:** [US-7](../stories/US-7-ask-an-agent.md)
**Status:** signed-off

## Request

with the user story we created and our target for a shippable product, can you create all the tickets we needed?

This ticket makes US-7 true from a result produced by TKT-6. The story file already exists. Do not add a second story. The agent CLI runs with the environment from TKT-2. Keeping or discarding the change is TKT-8.

## Acceptance criteria

* AC-1: From a result, the developer can ask for one of three actions: create a test, fix this test, or fix the application code this test is about.
* AC-2: Before the request is sent, the developer chooses which installed agent CLI to call.
* AC-3: When Claude Code CLI is installed and available in the loaded environment, it is one of the choices.
* AC-4: When the Cursor CLI is installed and available in the loaded environment, it is one of the choices.
* AC-5: The request includes the test name and the pass or fail.
* AC-6: When the result has failure text, the request includes that text.
* AC-7: When the test has a user story, the request includes that story.
* AC-8: While the agent is working, the window shows that the agent is working.
* AC-9: The developer can stop the agent request on its own. Stopping it leaves the reported test results in the window.
* AC-10: When the chosen CLI cannot be started, the project path, the run, and the results still work.
* AC-11: When the chosen CLI cannot be started, the agent actions say that the CLI is unavailable.
* AC-12: When the chosen CLI is missing from the loaded environment, the window names that CLI.

## Action items

- [x] AI-1 (AC-1): From a result, the developer can ask to create a test, fix this test, or fix the application code this test is about.
- [x] AI-2 (AC-2): Before the request is sent, the developer chooses which installed agent CLI to call.
- [x] AI-3 (AC-3): When Claude Code CLI is installed and available in the loaded environment, it is one of the choices.
- [x] AI-4 (AC-4): When the Cursor CLI is installed and available in the loaded environment, it is one of the choices.
- [x] AI-5 (AC-5): The request includes the test name and the pass or fail.
- [x] AI-6 (AC-6): When the result has failure text, the request includes that text.
- [x] AI-7 (AC-7): When the test has a user story, the request includes that story.
- [x] AI-8 (AC-8): While the agent is working, the window shows that the agent is working.
- [x] AI-9 (AC-9): The developer can stop the agent request on its own, and the reported test results stay in the window.
- [x] AI-10 (AC-10): When the chosen CLI cannot be started, the project path, the run, and the results still work.
- [x] AI-11 (AC-11): When the chosen CLI cannot be started, the agent actions say that the CLI is unavailable.
- [x] AI-12 (AC-12): When the chosen CLI is missing from the loaded environment, the window names that CLI.

## Sign-off

- [x] Every action item above is checked
- [x] Every acceptance criterion above is met
- Date: 2026-11-14
- Result: Verified create/fix ask via Claude Code or Cursor CLI when on loaded PATH; working/stop keeps results; missing/unavailable named without blocking path/run (`go test ./...`). Discovery [ACCEPT_CLI_DEFAULTS](../../.agents/teams/discovery/sessions/2026-11-14-tkt7-agent-cli.md).
