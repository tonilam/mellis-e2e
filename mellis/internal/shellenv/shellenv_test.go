package shellenv

import (
	"runtime"
	"strings"
	"testing"
	"time"
)

func TestParseNullEnv(t *testing.T) {
	data := []byte("PATH=/usr/bin:/bin\x00HOME=/Users/dev\x00\x00")
	vars, err := ParseNullEnv(data)
	if err != nil {
		t.Fatal(err)
	}
	if vars["PATH"] != "/usr/bin:/bin" {
		t.Fatalf("PATH=%q", vars["PATH"])
	}
	if vars["HOME"] != "/Users/dev" {
		t.Fatalf("HOME=%q", vars["HOME"])
	}
}

func TestParseLineEnv(t *testing.T) {
	data := []byte("Path=C:\\Windows\\System32\r\nHOME=C:\\Users\\dev\r\n")
	vars, err := ParseLineEnv(data)
	if err != nil {
		t.Fatal(err)
	}
	if vars.PATH() != `C:\Windows\System32` {
		t.Fatalf("PATH=%q", vars.PATH())
	}
}

func TestLoadIncludesPATH(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("unix-oriented smoke; Windows covered by ParseLineEnv")
	}
	vars, err := Load(20 * time.Second)
	if err != nil {
		t.Fatalf("Load: %v", err)
	}
	path := vars.PATH()
	if path == "" {
		t.Fatal("expected non-empty PATH")
	}
	// Login shell on this Mac should include standard bins.
	if !strings.Contains(path, "/usr/bin") && !strings.Contains(path, "/bin") {
		t.Fatalf("PATH missing standard bins: %s", path)
	}
	env := vars.Environ()
	if len(env) < 5 {
		t.Fatalf("expected several env entries, got %d", len(env))
	}
}
