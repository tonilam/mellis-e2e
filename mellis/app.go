package main

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"sync"
	"time"

	"mellis/internal/agentcli"
	"mellis/internal/cmdexec"
	"mellis/internal/projectsnap"
	"mellis/internal/runsettings"
	"mellis/internal/shellenv"
	"mellis/internal/suiteparse"
	"mellis/internal/suiterun"

	wailsRuntime "github.com/wailsapp/wails/v2/pkg/runtime"
)

const (
	suiteStateEvent = "suite:state"
	agentStateEvent = "agent:state"
)

// Loader loads the login-shell environment. Tests can replace it.
type Loader func(timeout time.Duration) (shellenv.Vars, error)

// SettingsPath locates the persisted run settings file. Tests can replace it.
type SettingsPath = runsettings.PathFunc

// App is the Wails application backend.
type App struct {
	ctx          context.Context
	loader       Loader
	settingsPath SettingsPath

	mu       sync.RWMutex
	vars     shellenv.Vars
	loaded   bool
	loadErr  string
	settings runsettings.Settings

	suiteMu     sync.Mutex
	suiteCmd    *exec.Cmd
	suiteCancel context.CancelFunc
	tracker     *suiteparse.Tracker
	suiteState  SuiteState

	agentMu          sync.Mutex
	agentCmd         *exec.Cmd
	agentCancel      context.CancelFunc
	agentState       AgentState
	agentBeforeSnap  *projectsnap.Snap
	agentProposed    *projectsnap.Snap
	agentProjectPath string

	rerunMu       sync.Mutex
	rerunPrevious map[string]suiteparse.Test
}

// EnvStatus is the environment snapshot shown in the UI.
type EnvStatus struct {
	Loaded bool              `json:"loaded"`
	Error  string            `json:"error"`
	Path   string            `json:"path"`
	Vars   map[string]string `json:"vars"`
	Count  int               `json:"count"`
}

// RunSettings is the project path and suite command for this session.
type RunSettings struct {
	ProjectPath string `json:"projectPath"`
	Command     string `json:"command"`
}

// ActionResult is returned when starting a suite or agent command.
// Failed is true when the process started and exited non-zero; Message then
// carries the failure the command reported (stdout/stderr), not an invented
// "application not running" explanation. StartSuite never starts, stops, or
// restarts the project's application — it only runs the given suite command.
type ActionResult struct {
	Started  bool   `json:"started"`
	Blocked  bool   `json:"blocked"`
	Failed   bool   `json:"failed"`
	Message  string `json:"message"`
	UsedPATH string `json:"usedPath"`
	Output   string `json:"output"`
	ExitCode int    `json:"exitCode"`
}

// SuiteState is the live or final progress of the current suite run.
type SuiteState struct {
	Running  bool              `json:"running"`
	Finished bool              `json:"finished"`
	Stopped  bool              `json:"stopped"`
	Failed   bool              `json:"failed"`
	Message  string            `json:"message"`
	Output   string            `json:"output"`
	ExitCode int               `json:"exitCode"`
	Tests    []suiteparse.Test `json:"tests"`
	Counts   suiteparse.Counts `json:"counts"`
}

// AgentCLIChoice is one agent CLI the developer can pick before send.
type AgentCLIChoice = agentcli.CLI

// AskAgentRequest is the create/fix ask from a test result.
type AskAgentRequest = agentcli.Request

// ProposedChange is one file difference the agent produced (relative path).
type ProposedChange = projectsnap.Change

// AgentState is the live or final status of an agent request.
// When PendingDecision is true, ProposedChanges lists what accept would keep;
// the project tree itself already matches the pre-request snapshot.
type AgentState struct {
	Running         bool             `json:"running"`
	Finished        bool             `json:"finished"`
	Stopped         bool             `json:"stopped"`
	Failed          bool             `json:"failed"`
	Unavailable     bool             `json:"unavailable"`
	Message         string           `json:"message"`
	Output          string           `json:"output"`
	ExitCode        int              `json:"exitCode"`
	CLIID           string           `json:"cliId"`
	CLILabel        string           `json:"cliLabel"`
	CLIBinary       string           `json:"cliBinary"`
	Action          string           `json:"action"`
	TestName        string           `json:"testName"`
	PendingDecision bool             `json:"pendingDecision"`
	ProposedChanges []ProposedChange `json:"proposedChanges"`
	Accepted        bool             `json:"accepted"`
	Discarded       bool             `json:"discarded"`
	CanRerun        bool             `json:"canRerun"`
}

// NewApp creates a new App.
func NewApp() *App {
	return &App{
		loader:       shellenv.Load,
		settingsPath: runsettings.DefaultPath,
		tracker:      suiteparse.NewTracker(),
	}
}

func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
	a.loadPersistedSettings()
	a.loadEnvironment()
}

// beforeClose stops any running suite or agent when the developer closes the window.
func (a *App) beforeClose(ctx context.Context) (prevent bool) {
	a.StopSuite()
	a.StopAgent()
	return false
}

func (a *App) loadPersistedSettings() {
	s, err := runsettings.Load(a.settingsPath)
	a.mu.Lock()
	defer a.mu.Unlock()
	if err != nil {
		a.settings = runsettings.Settings{}
		return
	}
	a.settings = s
}

func (a *App) loadEnvironment() {
	vars, err := a.loader(shellenv.DefaultTimeout)
	a.mu.Lock()
	defer a.mu.Unlock()
	if err != nil {
		a.vars = nil
		a.loaded = false
		a.loadErr = err.Error()
		return
	}
	a.vars = vars
	a.loaded = true
	a.loadErr = ""
}

// WindowKind reports where the UI runs. Used by the frontend so a reviewer
// can confirm the run is shown in the app window, not a browser tab.
func (a *App) WindowKind() string {
	return "app-window"
}

// GetEnvironmentStatus returns the loaded login-shell environment, or the
// failure message when the load did not succeed.
func (a *App) GetEnvironmentStatus() EnvStatus {
	a.mu.RLock()
	defer a.mu.RUnlock()
	st := EnvStatus{
		Loaded: a.loaded,
		Error:  a.loadErr,
	}
	if a.vars != nil {
		st.Path = a.vars.PATH()
		st.Count = len(a.vars)
		st.Vars = map[string]string(a.vars)
	}
	return st
}

// RetryLoadEnvironment reloads the developer's login-shell environment.
func (a *App) RetryLoadEnvironment() EnvStatus {
	a.loadEnvironment()
	return a.GetEnvironmentStatus()
}

// EnvironmentReady reports whether suite and agent actions may start.
func (a *App) EnvironmentReady() bool {
	a.mu.RLock()
	defer a.mu.RUnlock()
	return a.loaded
}

// GetRunSettings returns the session path and command (including remembered values).
func (a *App) GetRunSettings() RunSettings {
	a.mu.RLock()
	defer a.mu.RUnlock()
	return RunSettings{
		ProjectPath: a.settings.ProjectPath,
		Command:     a.settings.Command,
	}
}

// SaveRunSettings stores the path and command for this session and the next.
func (a *App) SaveRunSettings(projectPath, command string) RunSettings {
	s := runsettings.Settings{
		ProjectPath: strings.TrimSpace(projectPath),
		Command:     strings.TrimSpace(command),
	}
	_ = runsettings.Save(a.settingsPath, s)
	a.mu.Lock()
	a.settings = s
	a.mu.Unlock()
	return RunSettings{ProjectPath: s.ProjectPath, Command: s.Command}
}

// GetSuiteState returns the current suite progress snapshot.
func (a *App) GetSuiteState() SuiteState {
	a.suiteMu.Lock()
	defer a.suiteMu.Unlock()
	return a.copySuiteStateLocked()
}

// StartSuite runs the suite command in the project path with the login-shell
// environment loaded at open. It does not start, stop, or restart the project's
// application. When that application is not running, the command still runs; any
// failure shown is the one the command reported.
// When the environment has not loaded, the path is missing, the command is missing
// from PATH, or the command cannot be started, the start is blocked and no test
// results are invented.
// On success the command runs asynchronously; progress is available via
// GetSuiteState and the "suite:state" event.
func (a *App) StartSuite(projectPath, command string) ActionResult {
	a.rerunMu.Lock()
	a.rerunPrevious = nil
	a.rerunMu.Unlock()
	return a.startSuite(projectPath, command)
}

// StopSuite stops the running suite command. Results already reported stay.
func (a *App) StopSuite() ActionResult {
	a.suiteMu.Lock()
	cmd := a.suiteCmd
	cancel := a.suiteCancel
	running := a.suiteState.Running
	a.suiteMu.Unlock()

	if !running || cmd == nil {
		return ActionResult{
			Started: false,
			Blocked: false,
			Message: "No suite is running",
		}
	}

	if cancel != nil {
		cancel()
	}
	_ = suiterun.KillTree(cmd)

	// waitSuite goroutine marks finished/stopped and keeps reported tests.
	deadline := time.Now().Add(3 * time.Second)
	for time.Now().Before(deadline) {
		st := a.GetSuiteState()
		if st.Finished && st.Stopped {
			return ActionResult{
				Started:  true,
				Message:  "Suite stopped",
				Output:   st.Output,
				ExitCode: st.ExitCode,
				Failed:   st.Failed,
			}
		}
		time.Sleep(10 * time.Millisecond)
	}
	// Force local finished state if Wait has not returned yet.
	a.suiteMu.Lock()
	if a.suiteState.Running {
		a.suiteState.Running = false
		a.suiteState.Finished = true
		a.suiteState.Stopped = true
		a.suiteState.Message = "Suite stopped"
		a.publishSuiteStateLocked()
	}
	st := a.copySuiteStateLocked()
	a.suiteMu.Unlock()
	return ActionResult{
		Started:  true,
		Message:  st.Message,
		Output:   st.Output,
		ExitCode: st.ExitCode,
		Failed:   st.Failed,
	}
}

// StartAgent starts an arbitrary agent command with the loaded environment.
// Kept for TKT-2 env probes. Prefer AskAgent for create/fix from a result.
// When the environment has not loaded, the start is blocked.
func (a *App) StartAgent(command string) ActionResult {
	return a.startWithEnv("agent", "", command, false)
}

// ListAgentCLIs returns known agent CLIs and whether each is available in the
// loaded environment PATH (Claude Code CLI and Cursor CLI).
func (a *App) ListAgentCLIs() []AgentCLIChoice {
	a.mu.RLock()
	loaded := a.loaded
	pathEnv := ""
	if a.vars != nil {
		pathEnv = a.vars.PATH()
	}
	a.mu.RUnlock()
	if !loaded {
		list := agentcli.Known()
		for i := range list {
			list[i].Available = false
		}
		return list
	}
	return agentcli.ListAgainstPATH(pathEnv)
}

// GetAgentState returns the current agent request snapshot.
func (a *App) GetAgentState() AgentState {
	a.agentMu.Lock()
	defer a.agentMu.Unlock()
	return a.agentState
}

// AskAgent sends a create/fix request to the chosen installed agent CLI.
// The request includes test name, pass/fail, failure text when present, and
// story when present. The CLI runs asynchronously in the project path with the
// login-shell environment. When the CLI is missing or cannot start, path/run/
// results keep working and the agent actions report unavailable or the missing name.
// Before the CLI starts, the project tree is snapshotted. After the CLI exits,
// the project is restored and the written tree is held as a proposal until
// AcceptProposal or DiscardProposal.
func (a *App) AskAgent(req AskAgentRequest) ActionResult {
	return a.askAgent(req)
}

// AcceptProposal keeps the agent's proposed change in the project.
func (a *App) AcceptProposal() ActionResult {
	return a.acceptProposal()
}

// DiscardProposal drops the proposal and leaves the project as it was before the request.
func (a *App) DiscardProposal() ActionResult {
	return a.discardProposal()
}

// RerunTest runs the suite again after an accepted change so the named test's
// new result can be shown beside its previous result.
//
// CLARIFICATION_NEEDED: product does not define a single-test filter flag.
// Assumption for TKT-8: rerun uses the same suite command and project path as
// StartSuite (full command); the window keeps the prior result for that test
// name and shows the new result beside it when the suite reports that test again.
func (a *App) RerunTest(testName string) ActionResult {
	return a.rerunTest(testName)
}

// StopAgent stops the running agent request. Reported suite results stay.
func (a *App) StopAgent() ActionResult {
	a.agentMu.Lock()
	cmd := a.agentCmd
	cancel := a.agentCancel
	running := a.agentState.Running
	a.agentMu.Unlock()

	if !running || cmd == nil {
		return ActionResult{
			Started: false,
			Blocked: false,
			Message: "No agent request is running",
		}
	}

	if cancel != nil {
		cancel()
	}
	_ = suiterun.KillTree(cmd)

	deadline := time.Now().Add(3 * time.Second)
	for time.Now().Before(deadline) {
		st := a.GetAgentState()
		if st.Finished && st.Stopped {
			return ActionResult{
				Started:  true,
				Message:  st.Message,
				Output:   st.Output,
				ExitCode: st.ExitCode,
				Failed:   st.Failed,
			}
		}
		time.Sleep(10 * time.Millisecond)
	}

	a.agentMu.Lock()
	if a.agentState.Running {
		a.agentState.Running = false
		a.agentState.Finished = true
		a.agentState.Stopped = true
		a.agentState.Message = "Agent request stopped"
		a.publishAgentStateLocked()
	}
	st := a.agentState
	a.agentMu.Unlock()
	return ActionResult{
		Started:  true,
		Message:  st.Message,
		Output:   st.Output,
		ExitCode: st.ExitCode,
		Failed:   st.Failed,
	}
}

func (a *App) askAgent(req AskAgentRequest) ActionResult {
	a.mu.RLock()
	loaded := a.loaded
	loadErr := a.loadErr
	vars := a.vars
	projectPath := a.settings.ProjectPath
	a.mu.RUnlock()

	if !loaded {
		msg := "Environment could not be loaded"
		if loadErr != "" {
			msg = msg + ": " + loadErr
		}
		return ActionResult{
			Started: false,
			Blocked: true,
			Message: "Cannot start agent: " + msg,
		}
	}

	cli, ok := agentcli.Find(req.CLIID)
	if !ok {
		return ActionResult{
			Started:  false,
			Blocked:  true,
			Message:  fmt.Sprintf("Unknown agent CLI: %s", req.CLIID),
			UsedPATH: vars.PATH(),
		}
	}
	if !agentcli.ValidAction(req.Action) {
		return ActionResult{
			Started:  false,
			Blocked:  true,
			Message:  fmt.Sprintf("Unknown agent action: %s", req.Action),
			UsedPATH: vars.PATH(),
		}
	}

	prompt, err := agentcli.BuildPrompt(req)
	if err != nil {
		return ActionResult{
			Started:  false,
			Blocked:  true,
			Message:  err.Error(),
			UsedPATH: vars.PATH(),
		}
	}

	resolved, lookErr := cmdexec.LookPath(cli.Binary, vars.PATH())
	if lookErr != nil {
		msg := fmt.Sprintf("CLI missing from the loaded environment: %s", cli.Binary)
		a.agentMu.Lock()
		a.clearProposalLocked()
		a.agentState = AgentState{
			Running:     false,
			Finished:    true,
			Unavailable: true,
			Message:     msg,
			CLIID:       cli.ID,
			CLILabel:    cli.Label,
			CLIBinary:   cli.Binary,
			Action:      req.Action,
			TestName:    strings.TrimSpace(req.TestName),
		}
		a.publishAgentStateLocked()
		a.agentMu.Unlock()
		return ActionResult{
			Started:  false,
			Blocked:  true,
			Message:  msg,
			UsedPATH: vars.PATH(),
		}
	}

	var beforeSnap *projectsnap.Snap
	if strings.TrimSpace(projectPath) != "" {
		snap, snapErr := projectsnap.Capture(projectPath)
		if snapErr != nil {
			return ActionResult{
				Started:  false,
				Blocked:  true,
				Message:  fmt.Sprintf("Cannot snapshot project before agent: %v", snapErr),
				UsedPATH: vars.PATH(),
			}
		}
		beforeSnap = snap
	}

	a.agentMu.Lock()
	if a.agentState.Running {
		a.agentMu.Unlock()
		if beforeSnap != nil {
			_ = beforeSnap.Close()
		}
		return ActionResult{
			Started:  false,
			Blocked:  true,
			Message:  "Cannot start agent: an agent request is already running",
			UsedPATH: vars.PATH(),
		}
	}
	if a.agentState.PendingDecision {
		a.agentMu.Unlock()
		if beforeSnap != nil {
			_ = beforeSnap.Close()
		}
		return ActionResult{
			Started:  false,
			Blocked:  true,
			Message:  "Cannot start agent: accept or discard the pending proposed change first",
			UsedPATH: vars.PATH(),
		}
	}

	a.clearProposalLocked()
	a.agentBeforeSnap = beforeSnap
	a.agentProjectPath = strings.TrimSpace(projectPath)

	ctx, cancel := context.WithCancel(context.Background())
	cmd := exec.CommandContext(ctx, resolved, "-p", prompt)
	cmd.Env = vars.Environ()
	if a.agentProjectPath != "" {
		cmd.Dir = a.agentProjectPath
	}
	suiterun.PrepareCmd(cmd)

	pr, pw := io.Pipe()
	cmd.Stdout = pw
	cmd.Stderr = pw

	a.agentCancel = cancel
	a.agentCmd = cmd
	a.agentState = AgentState{
		Running:   true,
		Finished:  false,
		Stopped:   false,
		Message:   "Agent working",
		CLIID:     cli.ID,
		CLILabel:  cli.Label,
		CLIBinary: cli.Binary,
		Action:    req.Action,
		TestName:  strings.TrimSpace(req.TestName),
	}
	a.publishAgentStateLocked()
	a.agentMu.Unlock()

	if err := cmd.Start(); err != nil {
		_ = pw.Close()
		_ = pr.Close()
		cancel()
		msg := fmt.Sprintf("%s is unavailable", cli.Label)
		a.agentMu.Lock()
		a.agentCmd = nil
		a.agentCancel = nil
		a.clearProposalLocked()
		a.agentState = AgentState{
			Running:     false,
			Finished:    true,
			Unavailable: true,
			Message:     msg,
			CLIID:       cli.ID,
			CLILabel:    cli.Label,
			CLIBinary:   cli.Binary,
			Action:      req.Action,
			TestName:    strings.TrimSpace(req.TestName),
		}
		a.publishAgentStateLocked()
		a.agentMu.Unlock()
		return ActionResult{
			Started:  false,
			Blocked:  true,
			ExitCode: -1,
			Message:  msg,
			UsedPATH: vars.PATH(),
		}
	}

	go a.watchAgent(ctx, cmd, pr, pw)

	return ActionResult{
		Started:  true,
		Blocked:  false,
		Message:  "Agent working",
		UsedPATH: vars.PATH(),
	}
}

func (a *App) watchAgent(ctx context.Context, cmd *exec.Cmd, pr *io.PipeReader, pw *io.PipeWriter) {
	var output strings.Builder
	var outputMu sync.Mutex

	doneRead := make(chan struct{})
	go func() {
		defer close(doneRead)
		scanner := bufio.NewScanner(pr)
		scanner.Buffer(make([]byte, 0, 64*1024), 1024*1024)
		for scanner.Scan() {
			line := scanner.Text()
			outputMu.Lock()
			if output.Len() > 0 {
				output.WriteByte('\n')
			}
			output.WriteString(line)
			outputMu.Unlock()

			a.agentMu.Lock()
			outputMu.Lock()
			a.agentState.Output = output.String()
			outputMu.Unlock()
			a.agentState.Message = "Agent working"
			a.publishAgentStateLocked()
			a.agentMu.Unlock()
		}
	}()

	waitErr := cmd.Wait()
	_ = pw.Close()
	<-doneRead
	_ = pr.Close()

	stopped := false
	select {
	case <-ctx.Done():
		stopped = true
	default:
	}

	exitCode := 0
	failed := false
	msg := "Agent request completed"
	if waitErr != nil {
		if ee, ok := waitErr.(*exec.ExitError); ok {
			exitCode = ee.ExitCode()
			failed = !stopped
			outputMu.Lock()
			out := output.String()
			outputMu.Unlock()
			if stopped {
				msg = "Agent request stopped"
			} else if out != "" {
				msg = out
			} else {
				msg = fmt.Sprintf("agent exited with status %d", exitCode)
			}
		} else if stopped {
			msg = "Agent request stopped"
		} else {
			failed = true
			msg = fmt.Sprintf("agent failed: %v", waitErr)
			exitCode = -1
		}
	} else if stopped {
		msg = "Agent request stopped"
	}

	outputMu.Lock()
	out := output.String()
	outputMu.Unlock()

	a.agentMu.Lock()
	a.agentState.Running = false
	a.agentState.Finished = true
	a.agentState.Stopped = stopped
	a.agentState.Failed = failed
	a.agentState.ExitCode = exitCode
	a.agentState.Output = out
	a.agentCmd = nil
	a.agentCancel = nil

	if stopped || failed {
		a.restoreAndDropProposalLocked()
		if stopped {
			a.agentState.Message = "Agent request stopped"
		} else {
			a.agentState.Message = msg
		}
	} else {
		a.finalizeProposalLocked()
		if a.agentState.PendingDecision {
			a.agentState.Message = "Proposed change ready. Accept or discard."
		} else if a.agentProjectPath == "" {
			a.agentState.Message = msg
		} else {
			a.agentState.Message = msg
		}
	}
	a.publishAgentStateLocked()
	a.agentMu.Unlock()
}

func (a *App) finalizeProposalLocked() {
	projectPath := a.agentProjectPath
	before := a.agentBeforeSnap
	if projectPath == "" || before == nil {
		a.clearProposalLocked()
		return
	}

	changes, diffErr := before.DiffAgainst(projectPath)
	if diffErr != nil {
		a.restoreAndDropProposalLocked()
		a.agentState.Failed = true
		a.agentState.Message = fmt.Sprintf("Cannot read proposed change: %v", diffErr)
		return
	}

	proposed, capErr := projectsnap.Capture(projectPath)
	if capErr != nil {
		a.restoreAndDropProposalLocked()
		a.agentState.Failed = true
		a.agentState.Message = fmt.Sprintf("Cannot hold proposed change: %v", capErr)
		return
	}
	if err := before.Restore(projectPath); err != nil {
		_ = proposed.Close()
		a.agentState.Failed = true
		a.agentState.Message = fmt.Sprintf("Cannot restore project after agent: %v", err)
		a.clearProposalLocked()
		return
	}

	_ = before.Close()
	a.agentBeforeSnap = nil
	a.agentProposed = proposed
	a.agentState.PendingDecision = true
	a.agentState.ProposedChanges = changes
	a.agentState.Accepted = false
	a.agentState.Discarded = false
	a.agentState.CanRerun = false
}

func (a *App) restoreAndDropProposalLocked() {
	if a.agentBeforeSnap != nil && a.agentProjectPath != "" {
		_ = a.agentBeforeSnap.Restore(a.agentProjectPath)
	}
	a.clearProposalLocked()
}

func (a *App) clearProposalLocked() {
	if a.agentBeforeSnap != nil {
		_ = a.agentBeforeSnap.Close()
		a.agentBeforeSnap = nil
	}
	if a.agentProposed != nil {
		_ = a.agentProposed.Close()
		a.agentProposed = nil
	}
	a.agentProjectPath = ""
	a.agentState.PendingDecision = false
	a.agentState.ProposedChanges = nil
	a.agentState.Accepted = false
	a.agentState.Discarded = false
	a.agentState.CanRerun = false
}

func (a *App) acceptProposal() ActionResult {
	a.agentMu.Lock()
	defer a.agentMu.Unlock()
	if !a.agentState.PendingDecision || a.agentProposed == nil {
		return ActionResult{
			Started: false,
			Blocked: true,
			Message: "No proposed change to accept",
		}
	}
	projectPath := a.agentProjectPath
	if projectPath == "" {
		a.clearProposalLocked()
		return ActionResult{
			Started: false,
			Blocked: true,
			Message: "No project path for proposed change",
		}
	}
	if err := projectsnap.Apply(a.agentProposed.Dir, projectPath); err != nil {
		return ActionResult{
			Started: false,
			Blocked: true,
			Message: fmt.Sprintf("Cannot accept proposed change: %v", err),
		}
	}
	testName := a.agentState.TestName
	_ = a.agentProposed.Close()
	a.agentProposed = nil
	a.agentProjectPath = projectPath
	a.agentState.PendingDecision = false
	a.agentState.Accepted = true
	a.agentState.Discarded = false
	a.agentState.CanRerun = strings.TrimSpace(testName) != ""
	a.agentState.Message = "Proposed change accepted"
	a.publishAgentStateLocked()
	return ActionResult{
		Started: true,
		Message: a.agentState.Message,
	}
}

func (a *App) discardProposal() ActionResult {
	a.agentMu.Lock()
	defer a.agentMu.Unlock()
	if !a.agentState.PendingDecision {
		return ActionResult{
			Started: false,
			Blocked: true,
			Message: "No proposed change to discard",
		}
	}
	// Project was already restored when the proposal was prepared.
	a.clearProposalLocked()
	a.agentState.Finished = true
	a.agentState.Discarded = true
	a.agentState.Accepted = false
	a.agentState.CanRerun = false
	a.agentState.Message = "Proposed change discarded"
	a.publishAgentStateLocked()
	return ActionResult{
		Started: true,
		Message: a.agentState.Message,
	}
}

func (a *App) rerunTest(testName string) ActionResult {
	testName = strings.TrimSpace(testName)
	if testName == "" {
		return ActionResult{
			Started: false,
			Blocked: true,
			Message: "Cannot rerun: test name is empty",
		}
	}

	a.agentMu.Lock()
	canRerun := a.agentState.CanRerun && a.agentState.Accepted
	agentTest := a.agentState.TestName
	a.agentMu.Unlock()
	if !canRerun {
		return ActionResult{
			Started: false,
			Blocked: true,
			Message: "Cannot rerun: accept a proposed change first",
		}
	}
	if agentTest != "" && agentTest != testName {
		return ActionResult{
			Started: false,
			Blocked: true,
			Message: fmt.Sprintf("Cannot rerun: expected test %q", agentTest),
		}
	}

	a.suiteMu.Lock()
	var prev *suiteparse.Test
	for i := range a.suiteState.Tests {
		if a.suiteState.Tests[i].Name == testName {
			tcopy := a.suiteState.Tests[i]
			prev = &tcopy
			break
		}
	}
	a.suiteMu.Unlock()
	if prev == nil {
		// Still allow rerun when the suite list was cleared; previous is empty.
		prev = &suiteparse.Test{Name: testName}
	}

	a.rerunMu.Lock()
	a.rerunPrevious = map[string]suiteparse.Test{testName: *prev}
	a.rerunMu.Unlock()

	a.mu.RLock()
	projectPath := a.settings.ProjectPath
	command := a.settings.Command
	a.mu.RUnlock()

	return a.startSuite(projectPath, command)
}

func (a *App) publishAgentStateLocked() {
	if a.ctx == nil {
		return
	}
	wailsRuntime.EventsEmit(a.ctx, agentStateEvent, a.agentState)
}

func (a *App) attachPreviousResults(tests []suiteparse.Test) []suiteparse.Test {
	a.rerunMu.Lock()
	prev := a.rerunPrevious
	a.rerunMu.Unlock()
	if len(prev) == 0 {
		return tests
	}
	out := make([]suiteparse.Test, len(tests))
	copy(out, tests)
	for i := range out {
		if p, ok := prev[out[i].Name]; ok {
			out[i].HasPrevious = true
			out[i].PreviousStatus = p.Status
			out[i].PreviousDuration = p.Duration
			out[i].PreviousFailure = p.Failure
		}
	}
	return out
}

func (a *App) startSuite(projectPath, command string) ActionResult {
	projectPath = strings.TrimSpace(projectPath)
	command = strings.TrimSpace(command)

	a.mu.RLock()
	loaded := a.loaded
	loadErr := a.loadErr
	vars := a.vars
	a.mu.RUnlock()

	if !loaded {
		msg := "Environment could not be loaded"
		if loadErr != "" {
			msg = msg + ": " + loadErr
		}
		return ActionResult{
			Started: false,
			Blocked: true,
			Message: "Cannot start suite: " + msg,
		}
	}

	if projectPath == "" {
		return ActionResult{
			Started:  false,
			Blocked:  true,
			Message:  "Cannot start suite: project path is empty",
			UsedPATH: vars.PATH(),
		}
	}

	info, err := os.Stat(projectPath)
	if err != nil {
		if os.IsNotExist(err) {
			return ActionResult{
				Started:  false,
				Blocked:  true,
				Message:  fmt.Sprintf("Project path does not exist: %s", projectPath),
				UsedPATH: vars.PATH(),
			}
		}
		return ActionResult{
			Started:  false,
			Blocked:  true,
			Message:  fmt.Sprintf("Cannot start suite: cannot use project path %s: %v", projectPath, err),
			UsedPATH: vars.PATH(),
		}
	}
	if !info.IsDir() {
		return ActionResult{
			Started:  false,
			Blocked:  true,
			Message:  fmt.Sprintf("Project path does not exist: %s", projectPath),
			UsedPATH: vars.PATH(),
		}
	}

	if command == "" {
		return ActionResult{
			Started:  false,
			Blocked:  true,
			Message:  "Cannot start suite: no command given",
			UsedPATH: vars.PATH(),
		}
	}

	name := cmdexec.FirstToken(command)
	if _, err := cmdexec.LookPath(name, vars.PATH()); err != nil {
		return ActionResult{
			Started:  false,
			Blocked:  true,
			Message:  fmt.Sprintf("Command missing from the loaded environment: %s", name),
			UsedPATH: vars.PATH(),
		}
	}

	a.suiteMu.Lock()
	if a.suiteState.Running {
		a.suiteMu.Unlock()
		return ActionResult{
			Started:  false,
			Blocked:  true,
			Message:  "Cannot start suite: a suite is already running",
			UsedPATH: vars.PATH(),
		}
	}

	ctx, cancel := context.WithCancel(context.Background())
	cmd := a.commandWithEnv(vars, command)
	cmd.Dir = projectPath
	suiterun.PrepareCmd(cmd)

	pr, pw := io.Pipe()
	cmd.Stdout = pw
	cmd.Stderr = pw

	a.tracker.Reset()
	a.suiteCancel = cancel
	a.suiteCmd = cmd
	a.suiteState = SuiteState{
		Running:  true,
		Finished: false,
		Stopped:  false,
		Message:  "Suite running",
		Tests:    nil,
		Counts:   suiteparse.Counts{},
	}
	a.publishSuiteStateLocked()
	a.suiteMu.Unlock()

	if err := cmd.Start(); err != nil {
		_ = pw.Close()
		_ = pr.Close()
		cancel()
		a.suiteMu.Lock()
		a.suiteCmd = nil
		a.suiteCancel = nil
		a.suiteState = SuiteState{
			Running:  false,
			Finished: true,
			Message:  fmt.Sprintf("Cannot start command: %v", err),
		}
		a.publishSuiteStateLocked()
		a.suiteMu.Unlock()
		return ActionResult{
			Started:  false,
			Blocked:  true,
			ExitCode: -1,
			Message:  fmt.Sprintf("Cannot start command: %v", err),
			UsedPATH: vars.PATH(),
		}
	}

	go a.watchSuite(ctx, cmd, pr, pw, vars.PATH())

	return ActionResult{
		Started:  true,
		Blocked:  false,
		Message:  "Suite started",
		UsedPATH: vars.PATH(),
	}
}

func (a *App) watchSuite(ctx context.Context, cmd *exec.Cmd, pr *io.PipeReader, pw *io.PipeWriter, usedPATH string) {
	var output strings.Builder
	var outputMu sync.Mutex

	doneRead := make(chan struct{})
	go func() {
		defer close(doneRead)
		scanner := bufio.NewScanner(pr)
		// Suites can emit long failure blobs on one line.
		scanner.Buffer(make([]byte, 0, 64*1024), 1024*1024)
		for scanner.Scan() {
			line := scanner.Text()
			outputMu.Lock()
			if output.Len() > 0 {
				output.WriteByte('\n')
			}
			output.WriteString(line)
			outputMu.Unlock()

			if a.tracker.FeedLine(line) {
				a.suiteMu.Lock()
				snap := a.tracker.Snapshot()
				a.suiteState.Tests = a.attachPreviousResults(snap.Tests)
				a.suiteState.Counts = snap.Counts
				outputMu.Lock()
				a.suiteState.Output = output.String()
				outputMu.Unlock()
				a.publishSuiteStateLocked()
				a.suiteMu.Unlock()
			}
		}
	}()

	waitErr := cmd.Wait()
	_ = pw.Close()
	<-doneRead
	_ = pr.Close()

	stopped := false
	select {
	case <-ctx.Done():
		stopped = true
	default:
	}

	exitCode := 0
	failed := false
	msg := "Suite finished"
	if waitErr != nil {
		if ee, ok := waitErr.(*exec.ExitError); ok {
			exitCode = ee.ExitCode()
			failed = true
			outputMu.Lock()
			out := output.String()
			outputMu.Unlock()
			if stopped {
				msg = "Suite stopped"
			} else if out != "" {
				msg = out
			} else {
				msg = fmt.Sprintf("suite exited with status %d", exitCode)
			}
		} else if stopped {
			msg = "Suite stopped"
			failed = false
		} else {
			failed = true
			msg = fmt.Sprintf("suite failed: %v", waitErr)
			exitCode = -1
		}
	} else if stopped {
		msg = "Suite stopped"
	}

	outputMu.Lock()
	out := output.String()
	outputMu.Unlock()

	a.suiteMu.Lock()
	snap := a.tracker.Snapshot()
	a.suiteState.Running = false
	a.suiteState.Finished = true
	a.suiteState.Stopped = stopped
	a.suiteState.Failed = failed && !stopped
	a.suiteState.ExitCode = exitCode
	a.suiteState.Message = msg
	a.suiteState.Output = out
	a.suiteState.Tests = a.attachPreviousResults(snap.Tests)
	a.suiteState.Counts = snap.Counts
	a.suiteCmd = nil
	a.suiteCancel = nil
	a.publishSuiteStateLocked()
	a.suiteMu.Unlock()

	_ = usedPATH
}

func (a *App) publishSuiteStateLocked() {
	if a.ctx == nil {
		return
	}
	st := a.copySuiteStateLocked()
	wailsRuntime.EventsEmit(a.ctx, suiteStateEvent, st)
}

func (a *App) copySuiteStateLocked() SuiteState {
	st := a.suiteState
	if st.Tests != nil {
		tests := make([]suiteparse.Test, len(st.Tests))
		copy(tests, st.Tests)
		st.Tests = a.attachPreviousResults(tests)
	}
	return st
}

func (a *App) startWithEnv(kind, projectPath, command string, requireProject bool) ActionResult {
	a.mu.RLock()
	loaded := a.loaded
	loadErr := a.loadErr
	vars := a.vars
	a.mu.RUnlock()

	if !loaded {
		msg := "Environment could not be loaded"
		if loadErr != "" {
			msg = msg + ": " + loadErr
		}
		return ActionResult{
			Started: false,
			Blocked: true,
			Message: fmt.Sprintf("Cannot start %s: %s", kind, msg),
		}
	}

	command = strings.TrimSpace(command)
	if command == "" {
		return ActionResult{
			Started:  false,
			Blocked:  false,
			Message:  fmt.Sprintf("No %s command given", kind),
			UsedPATH: vars.PATH(),
		}
	}

	if requireProject {
		return a.startSuite(projectPath, command)
	}
	return a.runCommand(kind, "", command, vars)
}

func (a *App) runCommand(kind, projectPath, command string, vars shellenv.Vars) ActionResult {
	// Only the developer-supplied command runs. There is no project-app start,
	// stop, or restart step before or after this call.
	cmd := a.commandWithEnv(vars, command)
	if projectPath != "" {
		cmd.Dir = projectPath
	}
	out, err := cmd.CombinedOutput()
	output := strings.TrimSpace(string(out))
	result := ActionResult{
		Started:  true,
		Blocked:  false,
		UsedPATH: vars.PATH(),
		Output:   output,
	}
	if err != nil {
		if ee, ok := err.(*exec.ExitError); ok {
			result.Failed = true
			result.ExitCode = ee.ExitCode()
			// Surface the command's own report as the on-screen failure (AC-6).
			if output != "" {
				result.Message = output
			} else {
				result.Message = fmt.Sprintf("%s exited with status %d", kind, ee.ExitCode())
			}
			return result
		}
		// Could not start the process at all — no invented results.
		result.Started = false
		result.Blocked = true
		result.ExitCode = -1
		result.Output = ""
		result.Message = fmt.Sprintf("Cannot start command: %v", err)
		return result
	}
	result.Message = fmt.Sprintf("%s finished", kind)
	return result
}

func (a *App) commandWithEnv(vars shellenv.Vars, command string) *exec.Cmd {
	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/c", command)
	} else {
		shell := "/bin/sh"
		cmd = exec.Command(shell, "-c", command)
	}
	cmd.Env = vars.Environ()
	return cmd
}
