// Package agentcli names known agent CLIs and builds the ask prompt.
//
// Assumptions (CLARIFICATION_NEEDED until product confirms):
//   - Claude Code CLI binary is "claude"; non-interactive flag is -p/--print.
//   - Cursor CLI binary is "cursor-agent"; non-interactive flag is -p.
//
package agentcli

import (
	"fmt"
	"strings"

	"mellis/internal/cmdexec"
)

// Known action identifiers sent from the UI.
const (
	ActionCreateTest = "create-test"
	ActionFixTest    = "fix-test"
	ActionFixApp     = "fix-app"
)

// CLI describes one installed-agent choice.
type CLI struct {
	ID        string `json:"id"`
	Label     string `json:"label"`
	Binary    string `json:"binary"`
	Available bool   `json:"available"`
}

// Request is the ask payload from a test result.
type Request struct {
	CLIID    string `json:"cliId"`
	Action   string `json:"action"`
	TestName string `json:"testName"`
	Status   string `json:"status"`
	Failure  string `json:"failure"`
	Story    string `json:"story"`
}

// Known returns the fixed set of agent CLIs Mellis can call.
func Known() []CLI {
	return []CLI{
		{ID: "claude", Label: "Claude Code CLI", Binary: "claude"},
		{ID: "cursor", Label: "Cursor CLI", Binary: "cursor-agent"},
	}
}

// ListAgainstPATH marks each known CLI available when its binary resolves in pathEnv.
func ListAgainstPATH(pathEnv string) []CLI {
	out := Known()
	for i := range out {
		_, err := cmdexec.LookPath(out[i].Binary, pathEnv)
		out[i].Available = err == nil
	}
	return out
}

// Find returns the known CLI with the given id, or false when unknown.
func Find(id string) (CLI, bool) {
	id = strings.TrimSpace(id)
	for _, c := range Known() {
		if c.ID == id {
			return c, true
		}
	}
	return CLI{}, false
}

// ActionLabel is the human wording for an ask action.
func ActionLabel(action string) string {
	switch strings.TrimSpace(action) {
	case ActionCreateTest:
		return "create a test"
	case ActionFixTest:
		return "fix this test"
	case ActionFixApp:
		return "fix the application code this test is about"
	default:
		return ""
	}
}

// ValidAction reports whether action is one of the three supported asks.
func ValidAction(action string) bool {
	return ActionLabel(action) != ""
}

// BuildPrompt composes the request text for the agent CLI.
// It always includes the test name and pass/fail. Failure and story are
// included only when present.
func BuildPrompt(req Request) (string, error) {
	label := ActionLabel(req.Action)
	if label == "" {
		return "", fmt.Errorf("unknown action %q", req.Action)
	}
	name := strings.TrimSpace(req.TestName)
	if name == "" {
		return "", fmt.Errorf("test name is required")
	}
	status := strings.TrimSpace(req.Status)
	if status == "" {
		return "", fmt.Errorf("test status is required")
	}

	var b strings.Builder
	b.WriteString("In the current project, please ")
	b.WriteString(label)
	b.WriteString(".\n\n")
	b.WriteString("Test name: ")
	b.WriteString(name)
	b.WriteByte('\n')
	b.WriteString("Result: ")
	b.WriteString(status)
	b.WriteByte('\n')
	if fail := strings.TrimSpace(req.Failure); fail != "" {
		b.WriteString("Failure:\n")
		b.WriteString(fail)
		b.WriteByte('\n')
	}
	if story := strings.TrimSpace(req.Story); story != "" {
		b.WriteString("User story:\n")
		b.WriteString(story)
		b.WriteByte('\n')
	}
	b.WriteString("\nPropose the change. Do not explain beyond what is needed for the change.\n")
	return b.String(), nil
}

// CommandLine builds the argv-style command string for logging/tests.
// The process is started with binary + "-p" + prompt (no shell).
func CommandLine(binary, prompt string) string {
	return binary + " -p " + prompt
}
