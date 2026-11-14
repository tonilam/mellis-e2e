// Package shellenv loads the developer's login-shell environment so suite
// and agent processes see the same PATH as a new terminal.
package shellenv

import (
	"fmt"
	"sort"
	"strings"
	"time"
)

const DefaultTimeout = 15 * time.Second

// Vars is a snapshot of environment variables from a login shell.
type Vars map[string]string

// Environ returns KEY=VALUE pairs suitable for exec.Cmd.Env.
func (v Vars) Environ() []string {
	if len(v) == 0 {
		return nil
	}
	out := make([]string, 0, len(v))
	for k, val := range v {
		out = append(out, k+"="+val)
	}
	sort.Strings(out)
	return out
}

// PATH returns the PATH value, or empty if unset.
func (v Vars) PATH() string {
	if v == nil {
		return ""
	}
	if p, ok := v["PATH"]; ok {
		return p
	}
	// Windows may use Path.
	if p, ok := v["Path"]; ok {
		return p
	}
	return ""
}

// SortedKeys returns variable names in alphabetical order.
func (v Vars) SortedKeys() []string {
	keys := make([]string, 0, len(v))
	for k := range v {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	return keys
}

// ParseNullEnv parses null-separated KEY=VALUE pairs from env -0 output.
func ParseNullEnv(data []byte) (Vars, error) {
	parts := strings.Split(string(data), "\x00")
	out := make(Vars)
	for _, part := range parts {
		if part == "" {
			continue
		}
		k, val, ok := strings.Cut(part, "=")
		if !ok || k == "" {
			continue
		}
		out[k] = val
	}
	if len(out) == 0 {
		return nil, fmt.Errorf("no environment variables in shell output")
	}
	return out, nil
}

// ParseLineEnv parses newline-separated KEY=VALUE pairs (e.g. Windows `set`).
func ParseLineEnv(data []byte) (Vars, error) {
	lines := strings.Split(string(data), "\n")
	out := make(Vars)
	for _, line := range lines {
		line = strings.TrimRight(line, "\r")
		if line == "" {
			continue
		}
		k, val, ok := strings.Cut(line, "=")
		if !ok || k == "" {
			continue
		}
		out[k] = val
	}
	if len(out) == 0 {
		return nil, fmt.Errorf("no environment variables in shell output")
	}
	return out, nil
}
