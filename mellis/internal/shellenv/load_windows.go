//go:build windows

package shellenv

import (
	"bytes"
	"context"
	"fmt"
	"os/exec"
	"strings"
	"time"
)

// Load captures the user+system environment a new cmd session would see.
// On Windows there is no Unix-style login shell; `cmd /c set` after a clean
// process start reflects the machine and user PATH from the registry.
func Load(timeout time.Duration) (Vars, error) {
	if timeout <= 0 {
		timeout = DefaultTimeout
	}
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	cmd := exec.CommandContext(ctx, "cmd", "/c", "set")
	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	err := cmd.Run()
	if ctx.Err() == context.DeadlineExceeded {
		return nil, fmt.Errorf("timed out loading Windows environment")
	}
	if err != nil {
		msg := strings.TrimSpace(stderr.String())
		if msg == "" {
			msg = err.Error()
		}
		return nil, fmt.Errorf("cmd /c set failed: %s", msg)
	}

	vars, err := ParseLineEnv(stdout.Bytes())
	if err != nil {
		return nil, fmt.Errorf("parse Windows environment: %w", err)
	}
	if vars.PATH() == "" {
		return nil, fmt.Errorf("Windows environment returned no PATH")
	}
	return vars, nil
}
