# TKT-7 SESSION_REPORT

**Ticket:** TKT-7 — Ask an installed agent  
**Status:** in-progress (AI-1..AI-12 checked; sign-off left for reviewer)  
**Story:** US-7

## AC verification

| AC | Result | Evidence |
| --- | --- | --- |
| AC-1 | Met | Result dialog: Create a test / Fix this test / Fix application code → `AskAgent` with action ids. `TestAskAgentSupportsThreeActions`. |
| AC-2 | Met | CLI radio list in result dialog; selection required before send (`selectedCLIId`). |
| AC-3 | Met | Known choice “Claude Code CLI” (`claude`) marked available via LookPath on loaded PATH. `TestListAgentCLIsMarksClaudeAndCursor`. |
| AC-4 | Met | Known choice “Cursor CLI” (`cursor-agent`) when available in loaded PATH. Same list API/test. |
| AC-5 | Met | Prompt always includes test name + Result status. `TestAskAgentCompletesAndIncludesContextInPrompt`, `BuildPrompt` tests. |
| AC-6 | Met | Failure included only when present. Prompt tests + cursor-agent ask test. |
| AC-7 | Met | Story included only when present. Same. |
| AC-8 | Met | Agent state Message “Agent working” + UI `agent-status` working state / `agent:state` events. Stop/running test. |
| AC-9 | Met | `StopAgent` keeps suite tests. `TestAskAgentRunsWithPromptAndKeepsResultsOnStop`. |
| AC-10 | Met | After missing/cannot-start agent, `StartSuite` still succeeds. Missing + cannot-start tests. |
| AC-11 | Met | Start failure → “Claude Code CLI is unavailable” + `Unavailable` state. `TestAskAgentCannotStartSaysUnavailable`. |
| AC-12 | Met | Missing → “CLI missing from the loaded environment: claude”; UI names missing binary on choices. `TestAskAgentMissingCLINamesBinary`. |

Verification: `cd mellis && go test ./...` (pass). App build: `CGO_LDFLAGS="-framework UniformTypeIdentifiers" wails build` (pass; UTType link needs that flag on this macOS SDK).

Accept/discard of proposed file changes is **not** implemented (TKT-8). Finished/stopped agent shows a placeholder that the request completed/stopped.

## Files changed

- `mellis/internal/agentcli/agentcli.go` — known CLIs, availability, prompt builder
- `mellis/internal/agentcli/agentcli_test.go` — unit tests
- `mellis/app.go` — `ListAgentCLIs`, `AskAgent`, `StopAgent`, `GetAgentState`, async agent watch/events
- `mellis/app_test.go` — ask/missing/unavailable/stop/keep-results/three-actions
- `mellis/frontend/src/main.js` — result ask UI, CLI pick, working/stop panel
- `mellis/frontend/src/app.css` — agent panel / ask styles
- `mellis/frontend/wailsjs/go/main/App.{js,d.ts}`, `models.ts` — bindings
- `.docs/tickets/TKT-7-ask-an-agent.md` — status in-progress; AI checkboxes
- `.docs/tickets/README.md` — TKT-7 in-progress

## Clarifications (documented assumptions)

**CLARIFICATION_NEEDED (product):** Exact Cursor agent CLI binary name and non-interactive flags are not specified in US-7 / run-suite.md.

**Assumptions for TKT-7:**

| Choice label | Binary on PATH | Invocation |
| --- | --- | --- |
| Claude Code CLI | `claude` | `claude -p <prompt>` |
| Cursor CLI | `cursor-agent` | `cursor-agent -p <prompt>` |

- Availability = `LookPath` against the loaded login-shell PATH (TKT-2).
- Prompt always has test name + pass/fail; failure/story only when present.
- Process runs in the saved project path with loaded env; stop uses the same process-group kill as the suite.
- No file snapshot / accept / discard (TKT-8).

## Changelog one-liner (do not write to docs/changelogs yet)

- From a result, ask create/fix via installed Claude Code or Cursor CLI with working/stop; name missing/unavailable CLIs without blocking path/run/results.
