package agentcli

import (
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"testing"
)

func TestListAgainstPATHMarksAvailable(t *testing.T) {
	dir := t.TempDir()
	bin := "claude"
	if runtime.GOOS == "windows" {
		bin = "claude.exe"
	}
	path := filepath.Join(dir, bin)
	if err := os.WriteFile(path, []byte("#!/bin/sh\necho ok\n"), 0o755); err != nil {
		t.Fatal(err)
	}
	list := ListAgainstPATH(dir)
	var claude, cursor CLI
	for _, c := range list {
		switch c.ID {
		case "claude":
			claude = c
		case "cursor":
			cursor = c
		}
	}
	if !claude.Available {
		t.Fatalf("claude should be available in %s: %+v", dir, list)
	}
	if cursor.Available {
		t.Fatalf("cursor-agent should be missing: %+v", list)
	}
	if claude.Label != "Claude Code CLI" || cursor.Label != "Cursor CLI" {
		t.Fatalf("labels=%q / %q", claude.Label, cursor.Label)
	}
}

func TestBuildPromptIncludesRequiredAndOptional(t *testing.T) {
	p, err := BuildPrompt(Request{
		Action:   ActionFixTest,
		TestName: "Login fails",
		Status:   "failed",
		Failure:  "Expected visible button",
		Story:    "As a user I can log in",
	})
	if err != nil {
		t.Fatal(err)
	}
	for _, want := range []string{
		"fix this test",
		"Test name: Login fails",
		"Result: failed",
		"Failure:\nExpected visible button",
		"User story:\nAs a user I can log in",
	} {
		if !strings.Contains(p, want) {
			t.Fatalf("prompt missing %q:\n%s", want, p)
		}
	}
}

func TestBuildPromptOmitsEmptyOptional(t *testing.T) {
	p, err := BuildPrompt(Request{
		Action:   ActionCreateTest,
		TestName: "New flow",
		Status:   "passed",
	})
	if err != nil {
		t.Fatal(err)
	}
	if strings.Contains(p, "Failure:") || strings.Contains(p, "User story:") {
		t.Fatalf("unexpected optional sections:\n%s", p)
	}
	if !strings.Contains(p, "create a test") {
		t.Fatalf("prompt=%s", p)
	}
}

func TestBuildPromptFixApp(t *testing.T) {
	p, err := BuildPrompt(Request{
		Action:   ActionFixApp,
		TestName: "T",
		Status:   "failed",
	})
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(p, "fix the application code this test is about") {
		t.Fatalf("prompt=%s", p)
	}
}

func TestFindAndValidAction(t *testing.T) {
	c, ok := Find("claude")
	if !ok || c.Binary != "claude" {
		t.Fatalf("Find claude: %+v %v", c, ok)
	}
	if _, ok := Find("nope"); ok {
		t.Fatal("expected unknown")
	}
	if !ValidAction(ActionCreateTest) || ValidAction("other") {
		t.Fatal("ValidAction")
	}
}
