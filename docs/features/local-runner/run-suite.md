# Run an existing E2E suite

## Overview

- A developer, on their own Mac or Windows machine, installs this application and opens it like any other app.
- The runner has its own window. The developer does not use a separate browser tab for the run.
- Their application is already running locally. This runner does not start it, stop it, or restart it.
- The suite starts by a command-line command in that project. The runner runs that command with the same toolchain the developer has in a new terminal.

## Preconditions

- The developer has installed the Mac app or the Windows app.
- The project's application is already running before they start the suite.
- The project has a command that starts the E2E suite.
- The toolchain those tests use is available in a new terminal on that machine (for example a language version manager or a virtualenv set up for the developer).

## Install and open

- On a Mac, the developer installs an app and opens it from the Dock, Launchpad, or Finder.
- On Windows, the developer installs an app and opens it from the Start menu.
- The runner shows its own window. It does not open a browser.
- The window has a main app menu. Settings is a separate page from the run page.
- On open, the runner loads the environment of the developer's login shell, including the PATH a new terminal would have. Suite commands and agent CLIs run with that environment.
- When that environment cannot be loaded, Settings says so and the runner does not start a suite or an agent request.
- The developer can open a view of the command environment the runner loaded, and can retry the load, from Settings.

## Choose the project and the command

- On Settings, the developer gives a project path.
- On Settings, the developer gives the command that starts the E2E suite, or confirms the command this runner remembered for that path.
- Settings shows the project path for as long as the session lasts, and shows the command before a run starts.
- The developer can change the command on Settings before the run.
- The runner remembers the path and the command on this machine for the next session.
- The run page does not show Environment, Project path, or Suite command controls.

## Run

- The developer starts the suite from the run page.
- The runner executes the command and project path from Settings, with the login-shell environment loaded at open.
- The runner does not start the project's application as part of this action.
- While the command is running, each test the output reports moves from waiting, to running, to passed or failed.
- Until a project reporter is mapped, the interim report language is `mellis:test <waiting|running|passed|failed> <name>[ duration=…][ story=…][ failure=…]` lines on stdout/stderr; unmatched lines are ignored and do not invent tests.
- The window shows how many tests are finished, how many are running, and how many are still waiting.
- The developer can stop the run. Results already reported stay in the window.
- When the command exits, the run is finished.
- When the run session shows a line longer than the window width, that line wraps so the full line stays within the window's max width without horizontal overflow.

## Results

- Each reported test shows whether it passed or failed.
- Each reported test shows how long it took in this run, when the output includes a duration.
- When the test's metadata contains a user story, the row shows that story, and opening the test shows the story text.
- When the test's metadata has no user story, the row shows the test name and no story.
- A failed test shows the failure text the command reported, when the output includes it.
- The window is complete at the end of the run if the developer never asks an agent to change the project.

## Ask an installed agent

- From a result, the developer can ask an agent CLI already installed on the machine to do one of three things: create a test, fix this test, or fix the application code this test is about.
- Claude Code CLI and the Cursor CLI are examples. The developer chooses which installed CLI to call before the request is sent.
- First shippable detection is interim: LookPath on the loaded PATH for `claude` and `cursor-agent`, invoked with `-p <prompt>`, shown as those product labels when present. Exact binary names and flags are not a product lock; revise when a real install differs.
- The request carries the test name, the pass or fail, the failure text when there is one, and the user story when the test has one.
- The window shows that the agent is working, and the developer can stop that request on its own.
- The agent proposes a change. The developer accepts it or discards it before anything remains in the project.
- After the agent request finishes, the project matches the tree from before the request until the developer accepts; accept leaves the proposed change; discard leaves that pre-request tree.
- A discarded change leaves the project as it was.
- After an accepted change, the developer can run that test again from the window. The new result is shown beside the previous result for that test.
- First shippable: “Run this test again” runs the same suite command; the window shows previous and new for that named test (a single-test filter is not required for this outcome).
- When the chosen CLI cannot be started, the project list, the run, and the results still work. The agent actions say that the CLI is unavailable.

## Edge cases

- The project path does not exist. The runner says so and does not start a command.
- The command cannot be started. The runner says so and does not invent test results.
- The command finishes and reports no tests. The window says the run finished with no tests.
- The project's application was not running. The runner still executes the command. Any failure is the failure the command reported. The runner does not start the application.
- A test result has no duration. The row shows pass or fail and shows no duration.
- The developer closes the window while a run is going. The runner stops the command.
- A command or an agent CLI is missing from the loaded environment. The runner names the missing command and does not invent a result.

## Out of scope

- Starting, stopping, or configuring the project's application, databases, or browsers, beyond what the E2E command itself does.
- Running this runner, the suite, or the agent CLI inside Docker.
- A history of earlier runs, a flakiness count, or a trace of each step.
- A home screen organised as a shelf of user stories. The story appears on the test that already carries it.
- Sharing one run across more than one machine or more than one developer.
- Editing project files by hand inside the runner. An agent change reaches the project only through accept or discard.

## Requirements

The requirement stage for this behaviour is [.docs/stories/](../../.docs/stories/README.md). Accept or discard before a change remains, and duration for this run only, are acceptance criteria there. The story text is whatever the test metadata already contains. The runner shows that text when it is present and shows no story when it is absent.
