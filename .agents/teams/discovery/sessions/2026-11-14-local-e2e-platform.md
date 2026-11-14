---
date: 2026-11-14
request: A local E2E platform that scans a project, runs its suite, shows pass, fail, performance, and a user story when metadata has one, and calls a local agent CLI, running as a Docker container on Mac and Windows.
attendees: [david-kelley, alex-osborn, alberto-savoia, jeff-gothelf, rob-fitzpatrick, erika-hall]
status: decided
---
# Discovery: local E2E platform

**Request (verbatim):** check /Users/tonilam/Works/beezhub/.agents/teams/discovery in other repo, and bring the discovery team into this repo.

This is a new repo, we want to implement a local E2E platform that help developer to run existing E2E tests in a project. The user have to enter the project path and this application will scan all the E2E tests. The user can run the application to trigger the E2E test suite and monitor the progress. The application should show what test pass and what test failed, the performance of each test, and if the E2E test has meta data contains user story, the application should should the user story in each test. The application should be able to call local agent e.g. Claude Code CLI or Cursor CLI to create new tests or fix the codes of a failed test or fix a test depends of user request. There is no specific programming language we need to use, so please feel free to discuss with the discovery team and find what tech stack suit us. The application should be able to run as a docker container in Mac and Windows.

**Brief:** The repository was empty. Bringing the team is done; that sentence is out of the product. There is no application, no `docs/features/`, and no changelog entry before this session. The only actor in the request is a developer on their own machine. The intended flow is stated intent: enter a project path, scan existing E2E tests, trigger the suite, watch progress, read pass, fail, performance, and a user story when metadata contains one, then call a local agent CLI to create a test, fix the code of a failed test, or fix a test. Claude Code CLI and Cursor CLI are examples. Constraints stated: local, Docker on Mac and Windows, no mandated language. No sample project, no metadata example, and no user research. Evidence is the request. It is intent, and it describes no past run.

## Framing

Point of view, from Kelley: a developer on their own machine needs a run of the project's existing end-to-end tests they can read (pass, fail, how each test performed, the story when metadata has one), and a following moment in which they may ask the agent already installed on that machine to create or fix.

How Might We, kept: recognise the tests already in the project; watch a run until pass and fail are known; show performance in a way they can use on this run; show a user story when metadata has one; hand a create or fix to the installed agent and see it land on the same test; begin from a story; make a failure clear enough to choose a next step; trust that the run and the agent are on the project they chose. Round 2 added two: start the suite the way the project already runs it, and finish a run when they never ask an agent.

Business outcome, from Gothelf, as behaviour: the developer starts a run of a suite they already have and chooses a next action on a specific test; when metadata contains a story they can use it while choosing; after a create or fix, a later run shows the new test or a changed pass or fail; a developer on Mac and a developer on Windows each finish that run and that agent request while the app is a Docker container that can touch the project and the installed CLI. A container image and a spawned process are outputs.

## Idea board

Know which end-to-end tests this project already has: K1–K4, O1, O27–O29, O35, O36, O55, O57, K32–K35, K65. Side by side: mark a folder (K3), confirm each unsure test (O27), follow a note the project keeps (O55, O57).

Know what passed and failed while a run is going: K5–K8, O2–O5, O10, O11, O24, O25, O41, K38. The live board is one picture (K5 with O3).

Read a suite the project itself is running: O37, O54, O59, K36, K37.

See how each test performed on this run: K9–K12, O6–O9, O30, O33, O45, O51, K39–K41, K59, K63. Duration for this run is one picture (K9 with O6). A previous run on this machine (O9) and a history of status and duration (O45) sit side by side.

See the story that belongs with a test: K13–K16, O12–O15, O39, K57, K64.

Start from a story and reach its tests: K21–K24, O34, O40, O52, O53, O60, K42–K44, K62.

Ask the installed agent to create or fix, and decide what stays: K17–K20, O16–O23, O32, O38, O46–O48, O56, K45, K47, K49, K50, K58, K60, K61. Accept or discard before it stays is one picture (K20 with O22 and O46). A copy (O48), a preview (K20), and a suggestion the developer applies (O56) sit side by side.

Understand a failure well enough to choose: K25–K27, O11, O23, O31, O42–O44, K46, K48, K53, K54.

Trust that this is their project and their agent: K28–K31, O19, O26, K55, K56.

Finish a run when they do not ask an agent: O20, O21, K51, K52.

Counts: Kelley 31 ideas plus the Round 2 builds; Osborn O1–O87; Savoia S1–S16; Gothelf G1–G4; Fitzpatrick F1–F15; Hall H1–H6.

**Shortlist:**

1. **Read the suite.** On one named project, enter the path, find the tests already there, run them, watch, and read pass, fail, one performance reading, and the story when metadata has one. Why: until one real suite is readable, every other screen is a guess.
2. **Act from the failure.** From a result, ask the installed CLI to create a test, fix the test, or fix the application. The developer accepts or discards before a change stays. A copy (O48) and a suggestion they apply (O56) stay neighbours. Why: this is the job the request adds on top of runners that already show pass and fail.
3. **Watch, and leave the launch with the project.** The developer starts the suite the way that project already runs it. The app receives the results, shows the story, and is where the agent is called. Why: it may meet the need with the toolchain they already trust. It strains the sentence that says the application triggers the suite.
4. **Story is the front door.** Stories found in metadata are the list. The developer runs one story. A story with no test is a gap for the agent. Why: the story could be the unit of work. It closes if the first project has no story in test metadata.

## Hypotheses and riskiest assumptions

| ID | Hypothesis | XYZ | Risk | Value | Evidence |
|---|---|---|---|---|---|
| MEH | A developer who already runs E2E locally will move those runs here and send at least one create-or-fix through their installed CLI. | — | High | High | Request. No past episode (F1). |
| XYZ | At least 50% of weekly local-E2E developers who have an agent CLI will do that within 5 working days. | At least 50% of that group will start a suite here and send one create-or-fix. | High | Low for this sprint | The group has not been counted (Savoia). |
| xyz | The requester, on one project named in advance, moves their next runs onto a stand-in. | At least 2 of the next 3 runs within 5 working days, including 1 agent request. Fewer than 3 runs is inconclusive. Zero runs fails. | High | High | The only person in the brief. |
| G1 | They choose a next action on a specific test from a result that matches the suite they count, uses the performance reading they already use, and shows the story when metadata has it. | — | High | High | Request. Runner unknown (H1). |
| G2 | A later run shows a new test, or a changed pass or fail, after a request to the installed CLI. | — | High | High | Request. Which of the three jobs is real is unknown (H4). |
| G3 | They choose which failure to act on because of the user story. | — | High | Low until a story exists in the files | No metadata example (H3). |
| G4 | A developer on Mac and a developer on Windows each finish a run and an agent edit while the app is a Docker container that can touch the project and the host CLI. | — | High | High | Stated constraint. No observed container session. |
| H4 gate | They accept or discard an agent write before it stays. | — | High | High for direction 2 | Request is silent (T3). |

## Cheapest tests

Savoia set these pass lines before any result. One window of the next 3 runs can score only one of S12, S14, or S15. Gothelf wants the original xyz, agent included, as the single sprint test. That disagreement is in Dissent. Nothing below has been run.

| Assumption | Pretotype / experiment | Pass if (set in advance) | Cost |
|---|---|---|---|
| They will start runs here and stay for pass, fail, and one duration (direction 1, S12) | Mechanical Turk. One page of rows. A person runs the project's own command and types results while it goes. No scanner, no agent, no container. | On the project named beforehand, within 5 working days: at least 2 of the next 3 runs start by handing over the path; they stay until every test shows pass or fail and one duration; on at least 1 run they write a next action that names a test and its result or its duration. Fewer than 3 runs: inconclusive. Zero: fail. | About 3 hours to prepare the page, plus operator time on each run. |
| They will ask the installed CLI on a failure they watched, and accept or discard (direction 2, S13) | One-Night Stand, only after S12 or the S14 watch gate. One CLI call. A discard restores the tree before the sitting ends. A missing CLI is blocked, and blocked is not a pass. | Within 5 working days of a watched failure: one CLI request (create, fix the test, or fix the application); they accept or discard before anything remains. Discard plus a restored tree passes the gate. Accept passes only with a later run of that test inside 5 days. A patch applied with no decision fails this direction. | 4–6 hours once. |
| They keep starting the suite themselves and still watch it here (direction 3, S14) | Infiltrator. Same page, filled only from a run they start with the command they already use. Agent gate is S13, after the watch gate, in its own window. | They start at least 3 runs with their existing command, and for at least 2 of those 3 the page is open before the run ends and stays until pass, fail, and one duration are filled. The direction passes when the watch gate and the S13 gate both pass. If they ask the page to trigger the suite, the watch gate fails. | About 2 hours on top of the page. |
| Story metadata is how they choose a run (direction 4, S15) | Mechanical Turk list built only from metadata they point at. Stop inside an hour if they cannot point. Do not invent stories. | Fail immediately if they cannot point at metadata that contains a user story. Otherwise at least 2 of the next 3 runs are chosen from that list and only that story's tests run. A CLI create is scored only for a gap they acknowledge. | 1–2 hours if they point at the format. |
| Container, toolchain, and host CLI hold together (H5, G4) | One-day probe on the named project. Savoia: a fake container does not test this, so it is a build, and it stays out of S12–S15. | A container on an OS they have finishes one suite run and one call to the host CLI. | One day, and only if they hold all three. |
| What the files already contain (H1–H4, H6) | Hall: one session on the requester's machine. Contextual inquiry plus the test files. | A named path, the runner, the last slice, the result they opened, the story text or a recorded absence, and the last edit or a recorded absence. | Day 1, inside about 2 hours. Stop if they cannot open a project. |

## Research plan

Questions still open: what counts as an E2E test and what slice they run (H1); what they look at after pass and fail (H2); whether a user story is already in the files, and in what form (H3); which agent job they last did, and whether a write landed before anyone looked (H4); whether the container, the project's toolchain, and the host CLI all stay mandatory (H5); whether version one is one project (H6).

Already answered, from the request and the empty repo: the actor is a developer on their own machine; the app is meant to run as a Docker container on Mac and Windows; agent examples are Claude Code CLI and Cursor CLI; no runner, metadata format, or sample project was given; this repo has no product behaviour to read.

Method: one session, the requester, at the machine, about 2 hours. Fitzpatrick's rewrites below are the questions. A two-OS check is a second day, and only if they hold container plus toolchain plus host CLI together. No survey: the only recipient would be the person who asked, and closed options would be the board's guesses.

Stop rule, from Hall: stop at the end of that session, or at the end of day 2 when the two-OS check is required. If they cannot open a project, stop. Another conversation will not create a missing episode. Defaults if the session cannot answer: show pass, fail, and this run's duration; show a user story on the test when the file already contains one; the developer accepts a write before it lands; version one is the one project they open. Hall's default order is read the suite, understand the failure, then an agent write. Who starts the suite has no default until the session.

Interview guide: the questions in the next section.

## Dissent

**What the first five days are allowed to prove.** Savoia splits the original xyz. S12 scores starting and watching a run, and leaves the agent out. S13 scores the agent only after a failure they stayed for. Fitzpatrick wants the same split, so a single pass cannot hide which act happened. Gothelf holds the original xyz: at least 2 of the next 3 runs from the stand-in, including 1 agent request, inside 5 days. His reason: a first bet that withholds the agent makes that threshold impossible, so G1 and G2 are read from one window and are two kill rules. Kelley keeps Read the suite, Act from the failure, and Watch as separate moments a person can storyboard, and leaves the scoring to the others.

**Who starts the suite.** The request says the application triggers the suite. Osborn's watch family (O37, O54, O59) and shortlist 3 leave the launch with the project. Savoia: if they ask the page to trigger the suite, the watch gate fails, and that fact is recorded for the read-the-suite test. Hall sets no default. The session has to show how the last run was started.

**Accept before a write stays.** Kelley, Osborn, and Hall's stop-rule default keep an accept or discard before a change remains (K20, O22, O46). Gothelf puts that gate on the stand-in so it can be seen, and says it sits outside Savoia's original pass line. Fitzpatrick: it has not been observed, so the first real keep-or-discard is the evidence, and an explanation, a copy, or a suggestion they apply spends less reputation than a decision on their tree. Savoia: a discard that restores the tree passes the gate; a patch applied with no decision fails direction 2.

**The story as the front door.** The request supports showing a story when metadata has one. Gothelf and Hall park the shelf of stories (K22–K24, G3) until the files contain one and the developer uses it to choose. A story they only read on the row is a different, smaller feature.

**Whose yes counts.** Hall flags the xyz as sponsor bias: the person who asked for the product is the person who would perform the yes, and an agent request inside 5 days can be courtesy. Savoia still wants that person's time, because there is no one else in the brief. The 50% XYZ is parked by both. It names a group that has not been counted.

## Questions for the user

Checked against the Mom Test. Fitzpatrick's rewrites replace the drafts.

1. What happened that made you ask? Walk through the day, the project, and who was involved.
2. When did you last run E2E tests? Which project was it, what exactly did you run, and what did you use?
3. On that last run, what did you do to start it, and where did you watch it while it was going?
4. Once you knew what passed and what failed on that last run, what else did you check, and what did you do with it?
5. Last time you went looking for the story behind a test, where was it written? Point at that test, or at the place the story was.
6. Last time a test failed, what failed, what did you change, and what produced the change? Did the change stay in the project, and who looked at it first? The last time you added a test, what produced it?
7. Last time you ran tests, which machine were you on? The last time a tool had to see both the project and a program already installed on that machine, what did you set up, and what happened?

## Facilitator note: stack

The team stayed at product level, by its protocol. Every member recorded the same three failure conditions for a later technology choice: it has to run the target project's own toolchain; it has to call an agent CLI already installed on the developer's machine; it has to run as a Docker container on Mac and on Windows, with the run and the agent able to touch the project and those CLIs.

Docker Desktop on Mac and on Windows runs a Linux container. A Linux process cannot execute a Mac or Windows binary. The recommendation below is the facilitator's, so those three conditions can be true together.

- **In the container:** TypeScript on Node. One process serves a local web UI and the API that holds the run. Official Node images publish linux/amd64 and linux/arm64, which covers Intel and Apple Silicon Macs and typical Windows Docker Desktop.
- **On the host:** a small companion, one binary for Mac and one for Windows, written in Go so the host needs no extra runtime. The developer starts the companion. It launches the container with the project directory mounted, runs the suite with the project's own command, and runs the installed CLI (`claude`, the Cursor CLI, or whichever they name). The container reaches the companion through Docker Desktop's host address.
- **Left unset until a project is named:** which reporter the companion reads. The image does not treat one framework as the definition of an E2E test.

Savoia's cheap tests leave this container out on purpose. The one-day probe (H5) is the first time the companion has to finish one real run and one real CLI call.

### 2026-11-14 follow-up: native binary

The user asked: if Docker cannot run the machine's E2E tests or the agent CLI, is a native Windows or Mac executable the better application?

Facilitator answer, pending their confirmation: yes. The companion above is already the program that can see the project, the shell PATH, and the installed CLIs. Putting the UI in that same binary removes the container, the mount, and Docker Desktop from the path that runs tests. Docker stays available later as an optional clean-room package. It is the wrong place to launch the project's toolchain or `claude` / the Cursor CLI.

Revised shape: one Go program, cross-compiled for macOS (arm64 and amd64) and Windows (amd64). The developer starts it from the terminal they already use for that project, so it inherits the same PATH as their tests. It serves the local web UI, runs the suite with the project's own command, and runs the installed agent CLI. A double-clicked app is a weak fit, because Finder and the Start menu do not load nvm, asdf, pyenv, or a project virtualenv.

The user confirmed this packaging on 2026-11-14 and added two requirements, verbatim: "yes, we will require the developer already have a running local application in their project and E2E test could be started in CLI."

### 2026-11-14 follow-up: installable app with its own window

The user asked: "could the runner be a native application with UI that could be install as a windows or mac app?"

Yes. The same Go program is packaged as an installable Mac app and an installable Windows app, with the UI in the app's own window (the system webview, via Wails). Opening it from the Dock or the Start menu does not inherit a terminal's PATH. On launch the app reads the developer's login-shell environment and runs the suite command and the agent CLIs with that environment. A terminal-only launch is no longer the product.

## Outcome

Decided for this session. These are requirements the requester set. They are not a past episode of someone running a suite. Questions 1–7 are still unanswered.

**Chosen**
- The application is an installable Mac app and an installable Windows app, with its own window. The developer opens it from the Dock, Launchpad, Finder, or the Start menu. On open it loads the developer's login-shell environment, and suite commands and agent CLIs run with that environment. Reason: the user asked for an installed native app with a UI, and a double-clicked app still has to see the toolchain a new terminal would see.
- The developer already has the project's application running locally. This runner does not start, stop, or restart that application.
- The project's E2E tests can be started with a CLI command. The runner starts the suite by invoking that command.

**Dropped**
- Docker as the place the UI, the suite, or the agent CLI runs.
- A terminal-only launch, and a separate browser tab, as the way the developer uses the runner.
- The runner booting the application under test.
- Story-as-the-front-door, run history, flakiness, and traces. A user story is shown on a test when that test's metadata already contains one. Performance for this version is how long the test took in this run. These follow Hall's defaults because the questions are still open.

**Enters the draft feature doc**
- [docs/features/local-runner/run-suite.md](/docs/features/local-runner/run-suite.md), linked from [docs/features/local-runner/README.md](/docs/features/local-runner/README.md).
- Accept or discard before an agent change remains is in that draft as a default, still visible for the user to strike.
- The build is the Go binary in the facilitator note. The feature doc states the behaviour, not the language.

**Still open**
- Which project, which command, and what a user-story annotation looks like in their files.
- Questions 1–7 above.
- The user has not yet approved the feature doc. No application code until they do.
