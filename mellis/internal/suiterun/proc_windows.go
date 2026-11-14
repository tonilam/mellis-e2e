//go:build windows

package suiterun

import (
	"fmt"
	"os/exec"
)

// PrepareCmd is a no-op on Windows; KillTree uses taskkill /T.
func PrepareCmd(cmd *exec.Cmd) {}

// KillTree stops the process and its children.
func KillTree(cmd *exec.Cmd) error {
	if cmd == nil || cmd.Process == nil {
		return nil
	}
	kill := exec.Command("taskkill", "/T", "/F", "/PID", fmt.Sprintf("%d", cmd.Process.Pid))
	if err := kill.Run(); err != nil {
		return cmd.Process.Kill()
	}
	return nil
}
