// Package suiteparse turns suite command output lines into per-test progress.
//
// Framework-agnostic design:
//   - LineHandler converts one stdout/stderr line into zero or more Events.
//   - Tracker applies Events to ordered test rows and Counts.
//   - ProgressHandler is the built-in line format for TKT-5 / TKT-6.
//   - MultiHandler and Tracker.AddHandler are extension points for
//     Playwright/Jest/etc. reporters without changing the UI model.
//
// Built-in line format (case-sensitive prefix):
//
//	mellis:test <status> <name>[ duration=<d>][ story=<s>][ failure=<f>]
//
// where <status> is waiting, running, passed, or failed, and <name> is the
// text after status until the first optional field (or end of line).
// Optional fields use known keys duration, story, and failure; each value
// runs until the next known key= or end of line. Unmatched lines are ignored.
//
// CLARIFICATION_NEEDED: real suite metadata encoding (Playwright/Jest/etc.)
// is not specified; these optional fields are the interim wire shape until
// a project reporter is mapped.
package suiteparse

import (
	"regexp"
	"strings"
	"sync"
)

// Status is the lifecycle of one reported test.
type Status string

const (
	StatusWaiting Status = "waiting"
	StatusRunning Status = "running"
	StatusPassed  Status = "passed"
	StatusFailed  Status = "failed"
)

// Event is one status change for a named test, with optional result fields.
type Event struct {
	Name        string
	Status      Status
	Duration    string // set only when HasDuration
	Story       string // set only when HasStory
	Failure     string // set only when HasFailure
	HasDuration bool
	HasStory    bool
	HasFailure  bool
}

// LineHandler turns one output line into zero or more events.
type LineHandler interface {
	HandleLine(line string) []Event
}

// Test is one row shown while a run is in progress or after it stops.
// Previous* fields hold the prior result for that test after a rerun (TKT-8).
type Test struct {
	Name             string `json:"name"`
	Status           Status `json:"status"`
	Duration         string `json:"duration,omitempty"`
	Story            string `json:"story,omitempty"`
	Failure          string `json:"failure,omitempty"`
	HasPrevious      bool   `json:"hasPrevious,omitempty"`
	PreviousStatus   Status `json:"previousStatus,omitempty"`
	PreviousDuration string `json:"previousDuration,omitempty"`
	PreviousFailure  string `json:"previousFailure,omitempty"`
}

// Counts are the three tallies shown while the command is running.
type Counts struct {
	Finished int `json:"finished"`
	Running  int `json:"running"`
	Waiting  int `json:"waiting"`
}

// Snapshot is the UI-facing progress model.
type Snapshot struct {
	Tests  []Test `json:"tests"`
	Counts Counts `json:"counts"`
}

// Tracker holds ordered tests and applies handler events.
type Tracker struct {
	mu       sync.Mutex
	handlers []LineHandler
	order    []string
	byName   map[string]Test
}

// NewTracker returns a tracker with the built-in ProgressHandler.
func NewTracker(extra ...LineHandler) *Tracker {
	handlers := make([]LineHandler, 0, 1+len(extra))
	handlers = append(handlers, ProgressHandler{})
	handlers = append(handlers, extra...)
	return &Tracker{
		handlers: handlers,
		byName:   make(map[string]Test),
	}
}

// AddHandler registers another line handler (framework extension point).
func (t *Tracker) AddHandler(h LineHandler) {
	if h == nil {
		return
	}
	t.mu.Lock()
	defer t.mu.Unlock()
	t.handlers = append(t.handlers, h)
}

// Reset clears all tests.
func (t *Tracker) Reset() {
	t.mu.Lock()
	defer t.mu.Unlock()
	t.order = nil
	t.byName = make(map[string]Test)
}

// FeedLine runs all handlers on one line. Returns true when state changed.
func (t *Tracker) FeedLine(line string) bool {
	t.mu.Lock()
	handlers := append([]LineHandler(nil), t.handlers...)
	t.mu.Unlock()

	var events []Event
	for _, h := range handlers {
		events = append(events, h.HandleLine(line)...)
	}
	if len(events) == 0 {
		return false
	}

	t.mu.Lock()
	defer t.mu.Unlock()
	changed := false
	for _, ev := range events {
		if applyEvent(t, ev) {
			changed = true
		}
	}
	return changed
}

// Apply applies a single event (for tests and non-line sources).
func (t *Tracker) Apply(ev Event) bool {
	t.mu.Lock()
	defer t.mu.Unlock()
	return applyEvent(t, ev)
}

func applyEvent(t *Tracker, ev Event) bool {
	name := strings.TrimSpace(ev.Name)
	if name == "" || !validStatus(ev.Status) {
		return false
	}
	prev, ok := t.byName[name]
	if !ok {
		te := Test{Name: name, Status: ev.Status}
		mergeOptional(&te, ev)
		t.order = append(t.order, name)
		t.byName[name] = te
		return true
	}
	next := prev
	changed := false
	if prev.Status != ev.Status {
		next.Status = ev.Status
		changed = true
	}
	if mergeOptional(&next, ev) {
		changed = true
	}
	if !changed {
		return false
	}
	t.byName[name] = next
	return true
}

func mergeOptional(te *Test, ev Event) bool {
	changed := false
	if ev.HasDuration && te.Duration != ev.Duration {
		te.Duration = ev.Duration
		changed = true
	}
	if ev.HasStory && te.Story != ev.Story {
		te.Story = ev.Story
		changed = true
	}
	if ev.HasFailure && te.Failure != ev.Failure {
		te.Failure = ev.Failure
		changed = true
	}
	return changed
}

// Snapshot returns a copy of the current progress.
func (t *Tracker) Snapshot() Snapshot {
	t.mu.Lock()
	defer t.mu.Unlock()
	tests := make([]Test, 0, len(t.order))
	var c Counts
	for _, name := range t.order {
		te := t.byName[name]
		tests = append(tests, te)
		switch te.Status {
		case StatusWaiting:
			c.Waiting++
		case StatusRunning:
			c.Running++
		case StatusPassed, StatusFailed:
			c.Finished++
		}
	}
	return Snapshot{Tests: tests, Counts: c}
}

func validStatus(s Status) bool {
	switch s {
	case StatusWaiting, StatusRunning, StatusPassed, StatusFailed:
		return true
	default:
		return false
	}
}

// ProgressHandler recognizes "mellis:test <status> <name>[ fields…]" lines.
type ProgressHandler struct{}

const progressPrefix = "mellis:test "

// known optional field keys on mellis:test lines.
var optionalFieldRE = regexp.MustCompile(`\s+(duration|story|failure)=`)

// HandleLine implements LineHandler.
func (ProgressHandler) HandleLine(line string) []Event {
	line = strings.TrimSpace(line)
	if !strings.HasPrefix(line, progressPrefix) {
		return nil
	}
	rest := strings.TrimSpace(strings.TrimPrefix(line, progressPrefix))
	statusStr, afterStatus, ok := strings.Cut(rest, " ")
	if !ok {
		return nil
	}
	st := Status(strings.ToLower(strings.TrimSpace(statusStr)))
	if !validStatus(st) {
		return nil
	}

	name, duration, story, failure, hasDur, hasStory, hasFail := splitNameAndFields(strings.TrimSpace(afterStatus))
	if name == "" {
		return nil
	}
	return []Event{{
		Name:        name,
		Status:      st,
		Duration:    duration,
		Story:       story,
		Failure:     failure,
		HasDuration: hasDur,
		HasStory:    hasStory,
		HasFailure:  hasFail,
	}}
}

func splitNameAndFields(afterStatus string) (name, duration, story, failure string, hasDur, hasStory, hasFail bool) {
	loc := optionalFieldRE.FindStringIndex(afterStatus)
	if loc == nil {
		return strings.TrimSpace(afterStatus), "", "", "", false, false, false
	}
	name = strings.TrimSpace(afterStatus[:loc[0]])
	fields := afterStatus[loc[0]:]
	duration, story, failure, hasDur, hasStory, hasFail = parseOptionalFields(fields)
	return name, duration, story, failure, hasDur, hasStory, hasFail
}

func parseOptionalFields(fields string) (duration, story, failure string, hasDur, hasStory, hasFail bool) {
	matches := optionalFieldRE.FindAllStringSubmatchIndex(fields, -1)
	for i, m := range matches {
		key := fields[m[2]:m[3]]
		valueStart := m[1]
		valueEnd := len(fields)
		if i+1 < len(matches) {
			valueEnd = matches[i+1][0]
		}
		value := strings.TrimSpace(fields[valueStart:valueEnd])
		switch key {
		case "duration":
			duration = value
			hasDur = true
		case "story":
			story = value
			hasStory = true
		case "failure":
			failure = value
			hasFail = true
		}
	}
	return duration, story, failure, hasDur, hasStory, hasFail
}

// MultiHandler runs several handlers and concatenates their events.
type MultiHandler struct {
	Handlers []LineHandler
}

// HandleLine implements LineHandler.
func (m MultiHandler) HandleLine(line string) []Event {
	var out []Event
	for _, h := range m.Handlers {
		if h == nil {
			continue
		}
		out = append(out, h.HandleLine(line)...)
	}
	return out
}
