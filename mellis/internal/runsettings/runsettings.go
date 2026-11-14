// Package runsettings persists the project path and suite command on this machine.
package runsettings

import (
	"encoding/json"
	"errors"
	"os"
	"path/filepath"
)

const (
	appDirName  = "mellis"
	fileName    = "run-settings.json"
)

// Settings is the remembered project path and suite command.
type Settings struct {
	ProjectPath string `json:"projectPath"`
	Command     string `json:"command"`
}

// PathFunc locates the config file. Tests replace it.
type PathFunc func() (string, error)

// DefaultPath returns the config file under the user config directory.
func DefaultPath() (string, error) {
	dir, err := os.UserConfigDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(dir, appDirName, fileName), nil
}

// Load reads settings from disk. Missing file yields empty settings and no error.
func Load(pathFn PathFunc) (Settings, error) {
	if pathFn == nil {
		pathFn = DefaultPath
	}
	path, err := pathFn()
	if err != nil {
		return Settings{}, err
	}
	data, err := os.ReadFile(path)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return Settings{}, nil
		}
		return Settings{}, err
	}
	var s Settings
	if err := json.Unmarshal(data, &s); err != nil {
		return Settings{}, err
	}
	return s, nil
}

// Save writes settings to disk, creating the parent directory when needed.
func Save(pathFn PathFunc, s Settings) error {
	if pathFn == nil {
		pathFn = DefaultPath
	}
	path, err := pathFn()
	if err != nil {
		return err
	}
	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		return err
	}
	data, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(path, data, 0o644)
}
