//go:build !windows

package shellenv_test

import (
	"os"
	"os/exec"
	"strings"
	"testing"
	"time"

	"mellis/internal/shellenv"
)

func TestLoadedPATHMatchesNewTerminal(t *testing.T) {
	vars, err := shellenv.Load(20 * time.Second)
	if err != nil {
		t.Fatal(err)
	}

	shell := os.Getenv("SHELL")
	if shell == "" {
		shell = "/bin/zsh"
	}
	cmd := exec.Command(shell, "-l", "-i", "-c", `printf %s "$PATH"`)
	cmd.Env = []string{
		"HOME=" + os.Getenv("HOME"),
		"USER=" + os.Getenv("USER"),
		"LOGNAME=" + os.Getenv("LOGNAME"),
		"TERM=xterm-256color",
		"PATH=/usr/bin:/bin:/usr/sbin:/sbin",
	}
	out, err := cmd.CombinedOutput()
	if err != nil {
		t.Fatalf("terminal shell: %v %s", err, out)
	}

	best := ""
	for _, line := range strings.Split(string(out), "\n") {
		line = strings.TrimSpace(line)
		if strings.Contains(line, "/bin") && strings.Contains(line, ":") && len(line) > len(best) {
			best = line
		}
	}
	if best == "" {
		t.Fatalf("could not parse terminal PATH from %q", out)
	}
	if vars.PATH() != best {
		t.Fatalf("PATH mismatch\nloaded=%q\nterm  =%q", vars.PATH(), best)
	}
}
