# TKT-8 SESSION_REPORT

**Ticket:** TKT-8 — Accept or discard the change  
**Status:** in-progress (AI-1..AI-6 checked; sign-off left for reviewer)  
**Story:** US-8

## AC verification

| AC | Result | Evidence |
| --- | --- | --- |
| AC-1 | Met | After AskAgent finishes, `PendingDecision` + `ProposedChanges` (file list) in agent panel. Project already restored. `TestAcceptProposalKeepsChangeDiscardRestores`, `TestDiscardProposalRestoresTree`. |
| AC-2 | Met | `AcceptProposal` applies held proposal onto project. File content matches agent write after accept. Same accept test. |
| AC-3 | Met | `DiscardProposal` leaves project matching pre-request tree (already restored at proposal time; discard drops proposal). `TestDiscardProposalRestoresTree`. |
| AC-4 | Met | After agent exit, runner restores from snapshot before showing proposal; reads before accept show original files / no new files. Accept/discard tests. |
| AC-5 | Met | After accept, `CanRerun` + UI “Run this test again” → `RerunTest`. `TestRerunTestShowsPreviousBesideNew`. |
| AC-6 | Met | Rerun attaches `HasPrevious` / `PreviousStatus` / `PreviousFailure` beside the new result for that test name. Same rerun test. |

Verification: `cd mellis && go test ./...` (pass). App build: `CGO_LDFLAGS="-framework UniformTypeIdentifiers" wails build` (pass).

## Files changed

- `mellis/internal/projectsnap/` — capture / restore / diff / apply tree snapshots
- `mellis/app.go` — snapshot around AskAgent; `AcceptProposal`, `DiscardProposal`, `RerunTest`; previous-result merge
- `mellis/internal/suiteparse/suiteparse.go` — previous result fields on `Test`
- `mellis/app_test.go` — accept / discard / rerun tests; TKT-7 completion expectations updated for proposal state
- `mellis/frontend/src/main.js` — proposal panel, accept/discard/rerun, previous+new row display
- `mellis/frontend/src/app.css` — proposal / previous styles
- `mellis/frontend/wailsjs/go/**` — bindings regenerated
- `.docs/tickets/TKT-8-accept-or-discard.md` — status in-progress; AI checkboxes
- `.docs/tickets/README.md` — TKT-8 in-progress

## Clarifications (documented assumptions)

**CLARIFICATION_NEEDED (product):** Agent CLIs write in-place in the project path. Exact binary flags remain the TKT-7 interim defaults.

**Assumptions for TKT-8:**

| Topic | Assumption |
| --- | --- |
| In-place CLI writes | Snapshot before ask → CLI writes → capture proposal → restore project → Accept applies / Discard drops |
| Until accept | After the request finishes, project matches pre-request tree; proposal held outside the project |
| Snapshot skip | `.git` skipped; snapshots live under system temp, not inside the project |
| Rerun filter | No product single-test filter; `RerunTest` runs the same suite command/path as StartSuite and shows previous+new for the named test |

## Changelog one-liner (do not write to docs/changelogs yet)

- Accept or discard agent proposals before they remain; after accept, rerun that test and show new result beside previous.
