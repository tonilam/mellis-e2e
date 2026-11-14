# TKT-5 SESSION_REPORT

**Ticket:** TKT-5 — Watch the run  
**Status:** in-progress (AI-1..AI-7 checked; sign-off left for reviewer)  
**Story:** US-5

## AC verification

| AC | Result | Evidence |
| --- | --- | --- |
| AC-1 | Met | Async stdout/stderr line feed → `suiteparse.Tracker`; UI rows move waiting→running→passed/failed. `TestStartSuiteStreamsProgressAndFinishes`. |
| AC-2 | Met | Counts panel: finished / running / waiting from tracker snapshot. Same test + UI `progress-counts`. |
| AC-3 | Met | `StopSuite` + Stop run button. `TestStopSuiteKeepsResultsAndStopsCommand`. |
| AC-4 | Met | After stop, `SuiteState.Tests` retained (passed Slow kept). Same stop test. |
| AC-5 | Met | After stop, `Running=false`, process group killed. Same stop test. |
| AC-6 | Met | Natural exit sets `Finished=true`, `Stopped=false`. Stream/finish test. |
| AC-7 | Met | `OnBeforeClose` → `StopSuite`. `TestBeforeCloseStopsSuite`. |

Verification command: `cd mellis && go test ./...` (pass). App builds: `wails build`.

## Files changed

- `mellis/app.go` — async `StartSuite`, `StopSuite`, `GetSuiteState`, stream + `suite:state` events
- `mellis/main.go` — `OnBeforeClose`
- `mellis/app_test.go` — wait helper + progress/stop/close tests
- `mellis/internal/suiteparse/` — framework-agnostic line parser + tracker + tests
- `mellis/internal/suiterun/` — process-group kill (unix/windows)
- `mellis/frontend/src/main.js` — progress list, counts, stop button, `EventsOn('suite:state')`
- `mellis/frontend/src/app.css` — progress styles
- `mellis/frontend/wailsjs/go/**` — bindings for new methods
- `.docs/tickets/TKT-5-watch-the-run.md` — status in-progress; AI checkboxes
- `.docs/tickets/README.md` — TKT-5 in-progress

## Clarifications (documented assumptions)

**CLARIFICATION_NEEDED (product):** Which real reporter formats (Playwright list, Jest, TAP, …) should ship first is not specified in US-5 / run-suite.md.

**Assumption for TKT-5:** Ship a minimal documented line protocol and extension points; do not invent multi-framework product scope.

Built-in format (case-sensitive):

```text
mellis:test <status> <name>
```

`<status>` ∈ `waiting` | `running` | `passed` | `failed`. Unmatched lines are ignored (no invented tests).

Extension points:

- `suiteparse.LineHandler` / `ProgressHandler` / `MultiHandler`
- `Tracker.AddHandler` for framework-specific adapters later

Suites that do not emit these lines still run; the window shows finished with empty progress (no invented rows). Mapping from Playwright/Jest output is a follow-up product decision.

## Changelog one-liner (do not write to docs/changelogs yet)

- Live suite progress (waiting→running→pass/fail), counts, stop (keep results), finish on exit, stop on window close.
