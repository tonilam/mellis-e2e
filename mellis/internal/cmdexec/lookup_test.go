package cmdexec

import (
	"os"
	"path/filepath"
	"runtime"
	"testing"
)

func TestFirstToken(t *testing.T) {
	if got := FirstToken("  npm run e2e "); got != "npm" {
		t.Fatalf("got %q", got)
	}
	if got := FirstToken(""); got != "" {
		t.Fatalf("got %q", got)
	}
}

func TestLookPathFindsBinary(t *testing.T) {
	dir := t.TempDir()
	name := "mellis-probe"
	path := filepath.Join(dir, name)
	if runtime.GOOS == "windows" {
		path += ".exe"
	}
	if err := os.WriteFile(path, []byte("#!/bin/sh\necho ok\n"), 0o755); err != nil {
		t.Fatal(err)
	}
	found, err := LookPath(name, dir)
	if err != nil {
		t.Fatal(err)
	}
	if found != path {
		t.Fatalf("found=%q want %q", found, path)
	}
}

func TestLookPathMissing(t *testing.T) {
	_, err := LookPath("mellis-definitely-missing-xyz", t.TempDir())
	if err == nil {
		t.Fatal("expected error")
	}
}
