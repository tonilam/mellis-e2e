package runsettings

import (
	"os"
	"path/filepath"
	"testing"
)

func TestLoadMissingFileReturnsEmpty(t *testing.T) {
	dir := t.TempDir()
	pathFn := func() (string, error) {
		return filepath.Join(dir, "missing.json"), nil
	}
	s, err := Load(pathFn)
	if err != nil {
		t.Fatal(err)
	}
	if s.ProjectPath != "" || s.Command != "" {
		t.Fatalf("expected empty, got %+v", s)
	}
}

func TestSaveAndLoadRoundTrip(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "nested", "run-settings.json")
	pathFn := func() (string, error) {
		return path, nil
	}
	want := Settings{
		ProjectPath: "/Users/dev/my-app",
		Command:     "npm run e2e",
	}
	if err := Save(pathFn, want); err != nil {
		t.Fatal(err)
	}
	got, err := Load(pathFn)
	if err != nil {
		t.Fatal(err)
	}
	if got != want {
		t.Fatalf("got %+v want %+v", got, want)
	}
	if _, err := os.Stat(path); err != nil {
		t.Fatal(err)
	}
}
