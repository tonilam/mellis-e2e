package projectsnap

import (
	"os"
	"path/filepath"
	"runtime"
	"sort"
	"testing"
)

func TestCaptureRestoreAndDiff(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("path slash expectations")
	}
	proj := t.TempDir()
	mustWrite(t, filepath.Join(proj, "a.txt"), "one")
	mustWrite(t, filepath.Join(proj, "keep.txt"), "keep")
	mustWrite(t, filepath.Join(proj, "sub", "b.txt"), "two")

	snap, err := Capture(proj)
	if err != nil {
		t.Fatal(err)
	}
	defer snap.Close()

	mustWrite(t, filepath.Join(proj, "a.txt"), "changed")
	mustWrite(t, filepath.Join(proj, "new.txt"), "new")
	if err := os.Remove(filepath.Join(proj, "sub", "b.txt")); err != nil {
		t.Fatal(err)
	}

	changes, err := snap.DiffAgainst(proj)
	if err != nil {
		t.Fatal(err)
	}
	sort.Slice(changes, func(i, j int) bool { return changes[i].Path < changes[j].Path })
	want := []Change{
		{Path: "a.txt", Kind: KindModified},
		{Path: "new.txt", Kind: KindAdded},
		{Path: "sub/b.txt", Kind: KindRemoved},
	}
	if len(changes) != len(want) {
		t.Fatalf("changes=%+v want=%+v", changes, want)
	}
	for i := range want {
		if changes[i] != want[i] {
			t.Fatalf("changes[%d]=%+v want=%+v", i, changes[i], want[i])
		}
	}

	if err := snap.Restore(proj); err != nil {
		t.Fatal(err)
	}
	assertFile(t, filepath.Join(proj, "a.txt"), "one")
	assertFile(t, filepath.Join(proj, "keep.txt"), "keep")
	assertFile(t, filepath.Join(proj, "sub", "b.txt"), "two")
	if _, err := os.Stat(filepath.Join(proj, "new.txt")); !os.IsNotExist(err) {
		t.Fatalf("new.txt should be gone after restore, err=%v", err)
	}
}

func TestApplyProposal(t *testing.T) {
	proj := t.TempDir()
	mustWrite(t, filepath.Join(proj, "a.txt"), "before")

	snap, err := Capture(proj)
	if err != nil {
		t.Fatal(err)
	}
	defer snap.Close()

	mustWrite(t, filepath.Join(proj, "a.txt"), "after")
	proposed, err := Capture(proj)
	if err != nil {
		t.Fatal(err)
	}
	defer proposed.Close()

	if err := snap.Restore(proj); err != nil {
		t.Fatal(err)
	}
	assertFile(t, filepath.Join(proj, "a.txt"), "before")

	if err := Apply(proposed.Dir, proj); err != nil {
		t.Fatal(err)
	}
	assertFile(t, filepath.Join(proj, "a.txt"), "after")
}

func mustWrite(t *testing.T, path, body string) {
	t.Helper()
	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(path, []byte(body), 0o644); err != nil {
		t.Fatal(err)
	}
}

func assertFile(t *testing.T, path, want string) {
	t.Helper()
	b, err := os.ReadFile(path)
	if err != nil {
		t.Fatal(err)
	}
	if string(b) != want {
		t.Fatalf("%s=%q want %q", path, b, want)
	}
}
