// Package cmdexec resolves command names against a loaded environment PATH.
package cmdexec

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

// FirstToken returns the first whitespace-separated token of a command line.
func FirstToken(command string) string {
	command = strings.TrimSpace(command)
	if command == "" {
		return ""
	}
	fields := strings.Fields(command)
	if len(fields) == 0 {
		return ""
	}
	return fields[0]
}

// LookPath searches pathEnv (colon/semicolon-separated) for an executable named name.
// Absolute or relative names that include a separator are checked directly.
func LookPath(name, pathEnv string) (string, error) {
	if name == "" {
		return "", fmt.Errorf("empty command")
	}
	if strings.ContainsRune(name, filepath.Separator) || (runtime.GOOS == "windows" && strings.ContainsRune(name, '/')) {
		return checkExecutable(name)
	}
	if runtime.GOOS == "windows" && filepath.VolumeName(name) != "" {
		return checkExecutable(name)
	}

	var lastErr error
	for _, dir := range filepath.SplitList(pathEnv) {
		if dir == "" {
			dir = "."
		}
		candidates := []string{filepath.Join(dir, name)}
		if runtime.GOOS == "windows" {
			exts := []string{"", ".com", ".exe", ".bat", ".cmd"}
			candidates = candidates[:0]
			for _, ext := range exts {
				candidates = append(candidates, filepath.Join(dir, name+ext))
			}
		}
		for _, c := range candidates {
			if path, err := checkExecutable(c); err == nil {
				return path, nil
			} else {
				lastErr = err
			}
		}
	}
	if lastErr == nil {
		lastErr = os.ErrNotExist
	}
	return "", fmt.Errorf("command %q not found in PATH: %w", name, lastErr)
}

func checkExecutable(path string) (string, error) {
	fi, err := os.Stat(path)
	if err != nil {
		return "", err
	}
	if fi.IsDir() {
		return "", fmt.Errorf("%s is a directory", path)
	}
	if runtime.GOOS != "windows" {
		if fi.Mode()&0o111 == 0 {
			return "", fmt.Errorf("%s is not executable", path)
		}
	}
	return path, nil
}
