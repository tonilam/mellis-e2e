//go:build !windows

package suiterun

import (
	"os/exec"
	"syscall"
)

// PrepareCmd puts the process in its own group so Stop can kill children.
func PrepareCmd(cmd *exec.Cmd) {
	cmd.SysProcAttr = &syscall.SysProcAttr{Setpgid: true}
}

// KillTree stops the process group started by cmd.
func KillTree(cmd *exec.Cmd) error {
	if cmd == nil || cmd.Process == nil {
		return nil
	}
	pgid, err := syscall.Getpgid(cmd.Process.Pid)
	if err != nil {
		return cmd.Process.Kill()
	}
	// Negative pgid signals the whole group.
	if err := syscall.Kill(-pgid, syscall.SIGTERM); err != nil {
		_ = syscall.Kill(-pgid, syscall.SIGKILL)
		return err
	}
	return nil
}
