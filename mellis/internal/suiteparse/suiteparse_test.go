package suiteparse

import (
	"testing"
)

func TestProgressHandlerParsesLifecycle(t *testing.T) {
	tr := NewTracker()
	lines := []string{
		"noise",
		"mellis:test waiting Login flow",
		"mellis:test running Login flow",
		"mellis:test passed Login flow",
		"mellis:test waiting Checkout",
		"mellis:test running Checkout",
		"mellis:test failed Checkout",
	}
	for _, line := range lines {
		tr.FeedLine(line)
	}
	snap := tr.Snapshot()
	if len(snap.Tests) != 2 {
		t.Fatalf("tests=%+v", snap.Tests)
	}
	if snap.Tests[0].Name != "Login flow" || snap.Tests[0].Status != StatusPassed {
		t.Fatalf("first=%+v", snap.Tests[0])
	}
	if snap.Tests[1].Name != "Checkout" || snap.Tests[1].Status != StatusFailed {
		t.Fatalf("second=%+v", snap.Tests[1])
	}
	if snap.Counts.Finished != 2 || snap.Counts.Running != 0 || snap.Counts.Waiting != 0 {
		t.Fatalf("counts=%+v", snap.Counts)
	}
}

func TestProgressHandlerOptionalFields(t *testing.T) {
	tr := NewTracker()
	tr.FeedLine("mellis:test waiting Checkout")
	tr.FeedLine("mellis:test running Checkout")
	tr.FeedLine("mellis:test failed Checkout duration=500ms story=As a shopper I can pay failure=Expected button visible")
	snap := tr.Snapshot()
	if len(snap.Tests) != 1 {
		t.Fatalf("tests=%+v", snap.Tests)
	}
	te := snap.Tests[0]
	if te.Status != StatusFailed || te.Name != "Checkout" {
		t.Fatalf("te=%+v", te)
	}
	if te.Duration != "500ms" {
		t.Fatalf("duration=%q", te.Duration)
	}
	if te.Story != "As a shopper I can pay" {
		t.Fatalf("story=%q", te.Story)
	}
	if te.Failure != "Expected button visible" {
		t.Fatalf("failure=%q", te.Failure)
	}
}

func TestProgressHandlerDurationOnly(t *testing.T) {
	tr := NewTracker()
	tr.FeedLine("mellis:test passed Fast duration=12ms")
	te := tr.Snapshot().Tests[0]
	if te.Duration != "12ms" || te.Story != "" || te.Failure != "" {
		t.Fatalf("te=%+v", te)
	}
}

func TestProgressHandlerNoDurationKeepsAbsent(t *testing.T) {
	tr := NewTracker()
	tr.FeedLine("mellis:test passed NoTiming")
	te := tr.Snapshot().Tests[0]
	if te.Duration != "" {
		t.Fatalf("expected no duration, got %q", te.Duration)
	}
	if te.Status != StatusPassed || te.Name != "NoTiming" {
		t.Fatalf("te=%+v", te)
	}
}

func TestProgressHandlerStoryAbsent(t *testing.T) {
	tr := NewTracker()
	tr.FeedLine("mellis:test passed Plain name only")
	te := tr.Snapshot().Tests[0]
	if te.Story != "" || te.Name != "Plain name only" {
		t.Fatalf("te=%+v", te)
	}
}

func TestOptionalFieldsDoNotWipeOnLaterStatus(t *testing.T) {
	tr := NewTracker()
	tr.FeedLine("mellis:test failed X duration=1s story=US-1 failure=boom")
	tr.FeedLine("mellis:test failed X") // same status, no fields — no wipe
	te := tr.Snapshot().Tests[0]
	if te.Duration != "1s" || te.Story != "US-1" || te.Failure != "boom" {
		t.Fatalf("fields wiped: %+v", te)
	}
}

func TestCountsWhileRunning(t *testing.T) {
	tr := NewTracker()
	tr.FeedLine("mellis:test waiting A")
	tr.FeedLine("mellis:test waiting B")
	tr.FeedLine("mellis:test waiting C")
	tr.FeedLine("mellis:test running A")
	tr.FeedLine("mellis:test passed A")
	tr.FeedLine("mellis:test running B")
	snap := tr.Snapshot()
	if snap.Counts.Finished != 1 || snap.Counts.Running != 1 || snap.Counts.Waiting != 1 {
		t.Fatalf("counts=%+v", snap.Counts)
	}
}

func TestAddHandlerExtensionPoint(t *testing.T) {
	tr := NewTracker()
	tr.AddHandler(LineHandlerFunc(func(line string) []Event {
		if line == "CUSTOM start Foo" {
			return []Event{{Name: "Foo", Status: StatusRunning}}
		}
		return nil
	}))
	tr.FeedLine("CUSTOM start Foo")
	snap := tr.Snapshot()
	if len(snap.Tests) != 1 || snap.Tests[0].Status != StatusRunning {
		t.Fatalf("snap=%+v", snap)
	}
}

// LineHandlerFunc adapts a function to LineHandler (test helper / extension).
type LineHandlerFunc func(line string) []Event

func (f LineHandlerFunc) HandleLine(line string) []Event { return f(line) }
