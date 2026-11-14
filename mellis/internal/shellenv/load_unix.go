//go:build !windows

package shellenv

import (
	"bytes"
	"context"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"time"
)

// Load runs the developer's login shell and captures its environment,
// including the PATH a new terminal would have.
func Load(timeout time.Duration) (Vars, error) {
	if timeout <= 0 {
		timeout = DefaultTimeout
	}
	shell := loginShell()
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	// -l -i: login + interactive, so PATH matches a new terminal
	// (.zprofile / .bash_profile and .zshrc / .bashrc). Stdin is
	// /dev/null and the command is non-interactive work ending in exit.
	cmd := exec.CommandContext(ctx, shell, "-l", "-i", "-c", "env -0")
	cmd.Env = minimalBootstrapEnv()
	cmd.Stdin = nil
	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	err := cmd.Run()
	if ctx.Err() == context.DeadlineExceeded {
		return nil, fmt.Errorf("timed out loading login-shell environment from %s", shell)
	}
	if err != nil {
		msg := strings.TrimSpace(stderr.String())
		if msg == "" {
			msg = err.Error()
		}
		return nil, fmt.Errorf("login shell %s failed: %s", shell, msg)
	}

	vars, err := ParseNullEnv(stdout.Bytes())
	if err != nil {
		return nil, fmt.Errorf("parse environment from %s: %w", shell, err)
	}
	if vars.PATH() == "" {
		return nil, fmt.Errorf("login shell %s returned no PATH", shell)
	}
	return vars, nil
}

func loginShell() string {
	if s := os.Getenv("SHELL"); s != "" {
		return s
	}
	if _, err := os.Stat("/bin/zsh"); err == nil {
		return "/bin/zsh"
	}
	return "/bin/bash"
}

// minimalBootstrapEnv keeps launchd session vars (TMPDIR, SSH_AUTH_SOCK, …)
// but resets PATH so login/interactive profile scripts rebuild the toolchain
// PATH a new terminal would have, instead of inheriting a GUI-app PATH.
func minimalBootstrapEnv() []string {
	base := os.Environ()
	out := make([]string, 0, len(base)+2)
	seenPATH := false
	for _, kv := range base {
		k, _, _ := strings.Cut(kv, "=")
		switch k {
		case "PATH":
			out = append(out, "PATH=/usr/bin:/bin:/usr/sbin:/sbin")
			seenPATH = true
		case "TERM":
			if os.Getenv("TERM") == "" || os.Getenv("TERM") == "dumb" {
				out = append(out, "TERM=xterm-256color")
			} else {
				out = append(out, kv)
			}
		default:
			out = append(out, kv)
		}
	}
	if !seenPATH {
		out = append(out, "PATH=/usr/bin:/bin:/usr/sbin:/sbin")
	}
	if os.Getenv("HOME") == "" {
		if home, err := os.UserHomeDir(); err == nil {
			out = append(out, "HOME="+home)
		}
	}
	return out
}
