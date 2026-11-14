package main

import (
	"errors"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"testing"
	"time"

	"mellis/internal/shellenv"
	"mellis/internal/suiteparse"
)

func testPATH() string {
	if runtime.GOOS == "windows" {
		return `C:\Windows\System32`
	}
	return "/usr/bin:/bin"
}

func waitForSuite(t *testing.T, app *App) SuiteState {
	t.Helper()
	deadline := time.Now().Add(5 * time.Second)
	for time.Now().Before(deadline) {
		st := app.GetSuiteState()
		if st.Finished {
			return st
		}
		time.Sleep(10 * time.Millisecond)
	}
	t.Fatalf("suite did not finish; state=%+v", app.GetSuiteState())
	return SuiteState{}
}

func TestStartSuiteBlockedWhenLoadFailed(t *testing.T) {
	app := NewApp()
	app.loader = func(time.Duration) (shellenv.Vars, error) {
		return nil, errors.New("shell missing")
	}
	app.loadEnvironment()

	res := app.StartSuite(t.TempDir(), "echo suite")
	if !res.Blocked || res.Started {
		t.Fatalf("expected blocked, got %+v", res)
	}
	if !strings.Contains(res.Message, "Environment could not be loaded") {
		t.Fatalf("message=%q", res.Message)
	}

	agent := app.StartAgent("echo agent")
	if !agent.Blocked || agent.Started {
		t.Fatalf("expected agent blocked, got %+v", agent)
	}
}

func TestRetryRestoresSuiteAndAgent(t *testing.T) {
	app := NewApp()
	calls := 0
	path := testPATH()
	app.loader = func(time.Duration) (shellenv.Vars, error) {
		calls++
		if calls == 1 {
			return nil, errors.New("first fail")
		}
		return shellenv.Vars{
			"PATH": path,
			"HOME": "/tmp",
		}, nil
	}
	app.loadEnvironment()
	if app.EnvironmentReady() {
		t.Fatal("expected not ready after fail")
	}
	dir := t.TempDir()
	if !app.StartSuite(dir, "true").Blocked {
		t.Fatal("suite should be blocked")
	}

	st := app.RetryLoadEnvironment()
	if !st.Loaded {
		t.Fatalf("retry should succeed: %+v", st)
	}
	if !app.EnvironmentReady() {
		t.Fatal("expected ready after retry")
	}

	suiteCmd := "printf suite-ok"
	if runtime.GOOS == "windows" {
		suiteCmd = "echo suite-ok"
	}
	suite := app.StartSuite(dir, suiteCmd)
	if suite.Blocked || !suite.Started {
		t.Fatalf("suite after retry: %+v", suite)
	}
	if suite.UsedPATH != path {
		t.Fatalf("UsedPATH=%q", suite.UsedPATH)
	}
	final := waitForSuite(t, app)
	if !strings.Contains(final.Output, "suite-ok") {
		t.Fatalf("output=%q", final.Output)
	}

	agentCmd := "printf agent-ok"
	if runtime.GOOS == "windows" {
		agentCmd = "echo agent-ok"
	}
	agent := app.StartAgent(agentCmd)
	if agent.Blocked || !agent.Started {
		t.Fatalf("agent after retry: %+v", agent)
	}
	if !strings.Contains(agent.Output, "agent-ok") {
		t.Fatalf("output=%q", agent.Output)
	}
}

func TestStartUsesLoadedPATH(t *testing.T) {
	app := NewApp()
	marker := "/mellis-test-bin:" + testPATH()
	if runtime.GOOS == "windows" {
		marker = `C:\mellis-test-bin;` + testPATH()
	}
	app.loader = func(time.Duration) (shellenv.Vars, error) {
		return shellenv.Vars{
			"PATH": marker,
			"HOME": "/tmp",
		}, nil
	}
	app.loadEnvironment()

	cmd := `printf '%s' "$PATH"`
	if runtime.GOOS == "windows" {
		cmd = "echo %PATH%"
	}
	res := app.StartSuite(t.TempDir(), cmd)
	if !res.Started || res.Blocked {
		t.Fatalf("result=%+v", res)
	}
	final := waitForSuite(t, app)
	if !strings.Contains(final.Output, "mellis-test-bin") {
		t.Fatalf("child PATH=%q want marker containing mellis-test-bin", final.Output)
	}
	if res.UsedPATH != marker {
		t.Fatalf("UsedPATH=%q", res.UsedPATH)
	}
}

func TestStartSuiteBlockedWhenPathMissing(t *testing.T) {
	app := NewApp()
	app.loader = func(time.Duration) (shellenv.Vars, error) {
		return shellenv.Vars{"PATH": testPATH()}, nil
	}
	app.loadEnvironment()

	missing := filepath.Join(t.TempDir(), "no-such-project")
	res := app.StartSuite(missing, "true")
	if !res.Blocked || res.Started {
		t.Fatalf("expected blocked, got %+v", res)
	}
	if !strings.Contains(res.Message, "Project path does not exist") {
		t.Fatalf("message=%q", res.Message)
	}
	if res.Output != "" {
		t.Fatalf("invented output: %q", res.Output)
	}
}

func TestStartSuiteBlockedWhenCommandMissingFromEnv(t *testing.T) {
	app := NewApp()
	app.loader = func(time.Duration) (shellenv.Vars, error) {
		return shellenv.Vars{"PATH": t.TempDir()}, nil
	}
	app.loadEnvironment()

	res := app.StartSuite(t.TempDir(), "mellis-missing-cmd-xyz --flag")
	if !res.Blocked || res.Started {
		t.Fatalf("expected blocked, got %+v", res)
	}
	if !strings.Contains(res.Message, "Command missing from the loaded environment: mellis-missing-cmd-xyz") {
		t.Fatalf("message=%q", res.Message)
	}
	if res.Output != "" {
		t.Fatalf("invented output: %q", res.Output)
	}
}

func TestSaveRunSettingsPersistsAcrossLoad(t *testing.T) {
	dir := t.TempDir()
	pathFn := func() (string, error) {
		return filepath.Join(dir, "run-settings.json"), nil
	}

	app := NewApp()
	app.settingsPath = pathFn
	saved := app.SaveRunSettings("/proj/one", "npm run e2e")
	if saved.ProjectPath != "/proj/one" || saved.Command != "npm run e2e" {
		t.Fatalf("saved=%+v", saved)
	}

	app2 := NewApp()
	app2.settingsPath = pathFn
	app2.loadPersistedSettings()
	got := app2.GetRunSettings()
	if got.ProjectPath != "/proj/one" || got.Command != "npm run e2e" {
		t.Fatalf("got=%+v", got)
	}
}

func TestStartSuiteRunsInProjectDir(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("pwd check differs on windows")
	}
	app := NewApp()
	app.loader = func(time.Duration) (shellenv.Vars, error) {
		return shellenv.Vars{"PATH": testPATH()}, nil
	}
	app.loadEnvironment()

	dir := t.TempDir()
	res := app.StartSuite(dir, "pwd")
	if res.Blocked || !res.Started {
		t.Fatalf("result=%+v", res)
	}
	final := waitForSuite(t, app)
	got, err := filepath.EvalSymlinks(strings.TrimSpace(final.Output))
	if err != nil {
		got = strings.TrimSpace(final.Output)
	}
	want, err := filepath.EvalSymlinks(dir)
	if err != nil {
		want = dir
	}
	if got != want {
		t.Fatalf("cwd=%q want %q", got, want)
	}
}

func TestStartSuiteBlockedWhenPathIsFile(t *testing.T) {
	app := NewApp()
	app.loader = func(time.Duration) (shellenv.Vars, error) {
		return shellenv.Vars{"PATH": testPATH()}, nil
	}
	app.loadEnvironment()

	f := filepath.Join(t.TempDir(), "not-a-dir")
	if err := os.WriteFile(f, []byte("x"), 0o644); err != nil {
		t.Fatal(err)
	}
	res := app.StartSuite(f, "true")
	if !res.Blocked || res.Started {
		t.Fatalf("expected blocked, got %+v", res)
	}
	if !strings.Contains(res.Message, "Project path does not exist") {
		t.Fatalf("message=%q", res.Message)
	}
}

func TestStartSuiteSurfacesCommandFailure(t *testing.T) {
	app := NewApp()
	app.loader = func(time.Duration) (shellenv.Vars, error) {
		return shellenv.Vars{"PATH": testPATH()}, nil
	}
	app.loadEnvironment()

	dir := t.TempDir()
	cmd := `printf '%s\n' 'connection refused to app:3000'; exit 1`
	if runtime.GOOS == "windows" {
		cmd = "echo connection refused to app:3000 & exit /b 1"
	}
	res := app.StartSuite(dir, cmd)
	if res.Blocked || !res.Started {
		t.Fatalf("expected command to run, got %+v", res)
	}
	final := waitForSuite(t, app)
	if !final.Failed {
		t.Fatalf("expected Failed, got %+v", final)
	}
	if final.ExitCode == 0 {
		t.Fatalf("expected non-zero exit, got %+v", final)
	}
	if !strings.Contains(final.Message, "connection refused to app:3000") {
		t.Fatalf("failure on screen should be command report, got message=%q", final.Message)
	}
	if !strings.Contains(final.Output, "connection refused to app:3000") {
		t.Fatalf("output=%q", final.Output)
	}
}

func TestStartSuiteOnlyRunsGivenCommand(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("shell script probe")
	}
	dir := t.TempDir()
	marker := filepath.Join(dir, "suite-ran")
	script := filepath.Join(dir, "suite.sh")
	body := "#!/bin/sh\nprintf 'ok' > \"$1\"\n"
	if err := os.WriteFile(script, []byte(body), 0o755); err != nil {
		t.Fatal(err)
	}

	app := NewApp()
	app.loader = func(time.Duration) (shellenv.Vars, error) {
		return shellenv.Vars{"PATH": testPATH()}, nil
	}
	app.loadEnvironment()

	// StartSuite must run only this command — no separate project-app launch.
	res := app.StartSuite(dir, script+" "+marker)
	if res.Blocked || !res.Started {
		t.Fatalf("result=%+v", res)
	}
	final := waitForSuite(t, app)
	if final.Failed {
		t.Fatalf("result=%+v", final)
	}
	data, err := os.ReadFile(marker)
	if err != nil {
		t.Fatalf("suite command did not run: %v", err)
	}
	if string(data) != "ok" {
		t.Fatalf("marker=%q", data)
	}
}

func TestStartSuiteStreamsProgressAndFinishes(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("shell progress script")
	}
	app := NewApp()
	app.loader = func(time.Duration) (shellenv.Vars, error) {
		return shellenv.Vars{"PATH": testPATH()}, nil
	}
	app.loadEnvironment()

	dir := t.TempDir()
	cmd := strings.Join([]string{
		`printf '%s\n' 'mellis:test waiting Alpha'`,
		`printf '%s\n' 'mellis:test waiting Beta'`,
		`printf '%s\n' 'mellis:test running Alpha'`,
		`printf '%s\n' 'mellis:test passed Alpha'`,
		`printf '%s\n' 'mellis:test running Beta'`,
		`printf '%s\n' 'mellis:test failed Beta'`,
	}, "; ")
	res := app.StartSuite(dir, cmd)
	if res.Blocked || !res.Started {
		t.Fatalf("start=%+v", res)
	}
	final := waitForSuite(t, app)
	if !final.Finished || final.Running || final.Stopped {
		t.Fatalf("final=%+v", final)
	}
	if len(final.Tests) != 2 {
		t.Fatalf("tests=%+v", final.Tests)
	}
	if final.Tests[0].Status != suiteparse.StatusPassed || final.Tests[1].Status != suiteparse.StatusFailed {
		t.Fatalf("tests=%+v", final.Tests)
	}
	if final.Counts.Finished != 2 || final.Counts.Running != 0 || final.Counts.Waiting != 0 {
		t.Fatalf("counts=%+v", final.Counts)
	}
}

func TestStopSuiteKeepsResultsAndStopsCommand(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("shell sleep script")
	}
	app := NewApp()
	app.loader = func(time.Duration) (shellenv.Vars, error) {
		return shellenv.Vars{"PATH": testPATH()}, nil
	}
	app.loadEnvironment()

	dir := t.TempDir()
	marker := filepath.Join(dir, "still-running")
	cmd := strings.Join([]string{
		`printf '%s\n' 'mellis:test waiting Slow'`,
		`printf '%s\n' 'mellis:test running Slow'`,
		`printf '%s\n' 'mellis:test passed Slow'`,
		`printf '%s\n' 'mellis:test waiting Hang'`,
		`printf '%s\n' 'mellis:test running Hang'`,
		`touch '` + marker + `'`,
		`sleep 30`,
		`printf '%s\n' 'mellis:test passed Hang'`,
	}, "; ")
	res := app.StartSuite(dir, cmd)
	if res.Blocked || !res.Started {
		t.Fatalf("start=%+v", res)
	}

	deadline := time.Now().Add(3 * time.Second)
	for time.Now().Before(deadline) {
		if _, err := os.Stat(marker); err == nil {
			break
		}
		time.Sleep(20 * time.Millisecond)
	}
	if _, err := os.Stat(marker); err != nil {
		t.Fatal("suite did not reach sleep")
	}

	// Wait until progress includes the passed test.
	deadline = time.Now().Add(2 * time.Second)
	var mid SuiteState
	for time.Now().Before(deadline) {
		mid = app.GetSuiteState()
		if mid.Running && len(mid.Tests) >= 1 && mid.Tests[0].Status == suiteparse.StatusPassed {
			break
		}
		time.Sleep(20 * time.Millisecond)
	}
	if !mid.Running {
		t.Fatalf("expected running mid-state, got %+v", mid)
	}

	stop := app.StopSuite()
	if stop.Blocked {
		t.Fatalf("stop=%+v", stop)
	}
	final := waitForSuite(t, app)
	if !final.Finished || !final.Stopped || final.Running {
		t.Fatalf("final=%+v", final)
	}
	if len(final.Tests) < 1 || final.Tests[0].Name != "Slow" || final.Tests[0].Status != suiteparse.StatusPassed {
		t.Fatalf("results should remain after stop: %+v", final.Tests)
	}
	// Hang should still be running or waiting in the snapshot — not passed after stop.
	foundHang := false
	for _, te := range final.Tests {
		if te.Name == "Hang" {
			foundHang = true
			if te.Status == suiteparse.StatusPassed {
				t.Fatalf("Hang should not have finished after stop: %+v", final.Tests)
			}
		}
	}
	if !foundHang {
		t.Fatalf("expected Hang row kept: %+v", final.Tests)
	}
}

func TestBeforeCloseStopsSuite(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("shell sleep script")
	}
	app := NewApp()
	app.loader = func(time.Duration) (shellenv.Vars, error) {
		return shellenv.Vars{"PATH": testPATH()}, nil
	}
	app.loadEnvironment()

	dir := t.TempDir()
	res := app.StartSuite(dir, "sleep 30")
	if res.Blocked || !res.Started {
		t.Fatalf("start=%+v", res)
	}
	deadline := time.Now().Add(2 * time.Second)
	for time.Now().Before(deadline) {
		if app.GetSuiteState().Running {
			break
		}
		time.Sleep(10 * time.Millisecond)
	}
	if !app.GetSuiteState().Running {
		t.Fatal("suite not running")
	}

	prevent := app.beforeClose(nil)
	if prevent {
		t.Fatal("beforeClose should allow window close")
	}
	final := waitForSuite(t, app)
	if !final.Finished || final.Running {
		t.Fatalf("final=%+v", final)
	}
	if !final.Stopped {
		t.Fatalf("expected stopped after window close, got %+v", final)
	}
}

func TestStartSuiteReportsResultFields(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("shell progress script")
	}
	app := NewApp()
	app.loader = func(time.Duration) (shellenv.Vars, error) {
		return shellenv.Vars{"PATH": testPATH()}, nil
	}
	app.loadEnvironment()

	dir := t.TempDir()
	cmd := strings.Join([]string{
		`printf '%s\n' 'mellis:test waiting WithStory'`,
		`printf '%s\n' 'mellis:test running WithStory'`,
		`printf '%s\n' 'mellis:test failed WithStory duration=250ms story=As a developer I see the story failure=Expected visible button'`,
		`printf '%s\n' 'mellis:test waiting NoMeta'`,
		`printf '%s\n' 'mellis:test running NoMeta'`,
		`printf '%s\n' 'mellis:test passed NoMeta'`,
	}, "; ")
	res := app.StartSuite(dir, cmd)
	if res.Blocked || !res.Started {
		t.Fatalf("start=%+v", res)
	}
	final := waitForSuite(t, app)
	if len(final.Tests) != 2 {
		t.Fatalf("tests=%+v", final.Tests)
	}
	withStory := final.Tests[0]
	if withStory.Status != suiteparse.StatusFailed {
		t.Fatalf("WithStory status=%v", withStory.Status)
	}
	if withStory.Duration != "250ms" {
		t.Fatalf("duration=%q", withStory.Duration)
	}
	if withStory.Story != "As a developer I see the story" {
		t.Fatalf("story=%q", withStory.Story)
	}
	if withStory.Failure != "Expected visible button" {
		t.Fatalf("failure=%q", withStory.Failure)
	}
	noMeta := final.Tests[1]
	if noMeta.Name != "NoMeta" || noMeta.Status != suiteparse.StatusPassed {
		t.Fatalf("NoMeta=%+v", noMeta)
	}
	if noMeta.Duration != "" || noMeta.Story != "" || noMeta.Failure != "" {
		t.Fatalf("NoMeta should omit optional fields: %+v", noMeta)
	}
	// Finished state remains readable without starting an agent.
	if !final.Finished || final.Running {
		t.Fatalf("final=%+v", final)
	}
}

func TestStartSuiteEmptyRunFinishesWithNoTests(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("shell echo script")
	}
	app := NewApp()
	app.loader = func(time.Duration) (shellenv.Vars, error) {
		return shellenv.Vars{"PATH": testPATH()}, nil
	}
	app.loadEnvironment()

	dir := t.TempDir()
	res := app.StartSuite(dir, "printf 'no protocol lines\\n'")
	if res.Blocked || !res.Started {
		t.Fatalf("start=%+v", res)
	}
	final := waitForSuite(t, app)
	if !final.Finished || final.Running {
		t.Fatalf("final=%+v", final)
	}
	if len(final.Tests) != 0 {
		t.Fatalf("expected no tests, got %+v", final.Tests)
	}
}

func waitForAgent(t *testing.T, app *App) AgentState {
	t.Helper()
	deadline := time.Now().Add(5 * time.Second)
	for time.Now().Before(deadline) {
		st := app.GetAgentState()
		if st.Finished {
			return st
		}
		time.Sleep(10 * time.Millisecond)
	}
	t.Fatalf("agent did not finish; state=%+v", app.GetAgentState())
	return AgentState{}
}

func writeFakeCLI(t *testing.T, dir, name, script string) {
	t.Helper()
	if runtime.GOOS == "windows" {
		t.Skip("unix fake CLI")
	}
	path := filepath.Join(dir, name)
	if err := os.WriteFile(path, []byte(script), 0o755); err != nil {
		t.Fatal(err)
	}
}

func TestListAgentCLIsMarksClaudeAndCursor(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("unix fake CLI")
	}
	binDir := t.TempDir()
	writeFakeCLI(t, binDir, "claude", "#!/bin/sh\necho ok\n")
	app := NewApp()
	app.loader = func(time.Duration) (shellenv.Vars, error) {
		return shellenv.Vars{"PATH": binDir + ":" + testPATH()}, nil
	}
	app.loadEnvironment()

	list := app.ListAgentCLIs()
	var claude, cursor bool
	for _, c := range list {
		if c.ID == "claude" && c.Available && c.Label == "Claude Code CLI" {
			claude = true
		}
		if c.ID == "cursor" && !c.Available && c.Label == "Cursor CLI" && c.Binary == "cursor-agent" {
			cursor = true
		}
	}
	if !claude || !cursor {
		t.Fatalf("list=%+v", list)
	}
}

func TestAskAgentMissingCLINamesBinary(t *testing.T) {
	app := NewApp()
	emptyBin := t.TempDir()
	app.loader = func(time.Duration) (shellenv.Vars, error) {
		return shellenv.Vars{"PATH": emptyBin + ":" + testPATH()}, nil
	}
	app.loadEnvironment()
	app.SaveRunSettings(t.TempDir(), "true")

	res := app.AskAgent(AskAgentRequest{
		CLIID:    "claude",
		Action:   "fix-test",
		TestName: "Login",
		Status:   "failed",
		Failure:  "boom",
		Story:    "As a user",
	})
	if !res.Blocked || res.Started {
		t.Fatalf("expected blocked missing CLI, got %+v", res)
	}
	if !strings.Contains(res.Message, "CLI missing from the loaded environment: claude") {
		t.Fatalf("message=%q", res.Message)
	}
	st := app.GetAgentState()
	if !st.Unavailable || !strings.Contains(st.Message, "claude") {
		t.Fatalf("state=%+v", st)
	}

	// Path / suite still work when agent CLI is missing.
	suite := app.StartSuite(app.GetRunSettings().ProjectPath, "true")
	if suite.Blocked || !suite.Started {
		t.Fatalf("suite should still start: %+v", suite)
	}
	_ = waitForSuite(t, app)
}

func TestAskAgentRunsWithPromptAndKeepsResultsOnStop(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("unix fake CLI")
	}
	binDir := t.TempDir()
	writeFakeCLI(t, binDir, "claude", "#!/bin/sh\necho \"PROMPT:$2\"; sleep 30\n")
	proj := t.TempDir()

	app := NewApp()
	app.loader = func(time.Duration) (shellenv.Vars, error) {
		return shellenv.Vars{"PATH": binDir + ":" + testPATH()}, nil
	}
	app.loadEnvironment()
	app.SaveRunSettings(proj, "true")

	// Seed suite results that must survive stopping the agent.
	suiteCmd := strings.Join([]string{
		`printf '%s\n' 'mellis:test waiting KeepMe'`,
		`printf '%s\n' 'mellis:test running KeepMe'`,
		`printf '%s\n' 'mellis:test failed KeepMe failure=Expected button story=As a developer'`,
	}, "; ")
	if res := app.StartSuite(proj, suiteCmd); res.Blocked || !res.Started {
		t.Fatalf("suite=%+v", res)
	}
	suiteFinal := waitForSuite(t, app)
	if len(suiteFinal.Tests) != 1 {
		t.Fatalf("tests=%+v", suiteFinal.Tests)
	}

	ask := app.AskAgent(AskAgentRequest{
		CLIID:    "claude",
		Action:   "fix-app",
		TestName: "KeepMe",
		Status:   "failed",
		Failure:  "Expected button",
		Story:    "As a developer",
	})
	if ask.Blocked || !ask.Started {
		t.Fatalf("ask=%+v", ask)
	}
	deadline := time.Now().Add(2 * time.Second)
	for time.Now().Before(deadline) {
		if app.GetAgentState().Running {
			break
		}
		time.Sleep(10 * time.Millisecond)
	}
	if !app.GetAgentState().Running {
		t.Fatalf("agent not running: %+v", app.GetAgentState())
	}
	if app.GetAgentState().Message != "Agent working" {
		t.Fatalf("want working message, got %+v", app.GetAgentState())
	}

	stop := app.StopAgent()
	if !stop.Started {
		t.Fatalf("stop=%+v", stop)
	}
	agentFinal := waitForAgent(t, app)
	if !agentFinal.Stopped || agentFinal.Running {
		t.Fatalf("agentFinal=%+v", agentFinal)
	}
	if agentFinal.Message != "Agent request stopped" {
		t.Fatalf("message=%q", agentFinal.Message)
	}

	after := app.GetSuiteState()
	if len(after.Tests) != 1 || after.Tests[0].Name != "KeepMe" {
		t.Fatalf("results cleared after agent stop: %+v", after.Tests)
	}
}

func TestAskAgentCompletesAndIncludesContextInPrompt(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("unix fake CLI")
	}
	binDir := t.TempDir()
	writeFakeCLI(t, binDir, "cursor-agent", "#!/bin/sh\nprintf '%s\\n' \"$2\"\n")
	proj := t.TempDir()

	app := NewApp()
	app.loader = func(time.Duration) (shellenv.Vars, error) {
		return shellenv.Vars{"PATH": binDir + ":" + testPATH()}, nil
	}
	app.loadEnvironment()
	app.SaveRunSettings(proj, "true")

	res := app.AskAgent(AskAgentRequest{
		CLIID:    "cursor",
		Action:   "create-test",
		TestName: "Signup",
		Status:   "passed",
		Story:    "As a visitor I can sign up",
	})
	if res.Blocked || !res.Started {
		t.Fatalf("ask=%+v", res)
	}
	final := waitForAgent(t, app)
	if final.Failed || final.Stopped || !final.Finished {
		t.Fatalf("final=%+v", final)
	}
	if !final.PendingDecision {
		t.Fatalf("want pending proposal, got %+v", final)
	}
	if final.Message != "Proposed change ready. Accept or discard." {
		t.Fatalf("message=%q", final.Message)
	}
	out := final.Output
	for _, want := range []string{
		"create a test",
		"Test name: Signup",
		"Result: passed",
		"User story:\nAs a visitor I can sign up",
	} {
		if !strings.Contains(out, want) {
			t.Fatalf("output missing %q:\n%s", want, out)
		}
	}
	if strings.Contains(out, "Failure:") {
		t.Fatalf("unexpected failure section:\n%s", out)
	}
}

func TestAskAgentCannotStartSaysUnavailable(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("unix fake CLI")
	}
	binDir := t.TempDir()
	// LookPath finds +x file; Start fails (not a valid executable / no shebang).
	bad := filepath.Join(binDir, "claude")
	if err := os.WriteFile(bad, []byte("not-an-executable"), 0o755); err != nil {
		t.Fatal(err)
	}
	proj := t.TempDir()
	app := NewApp()
	app.loader = func(time.Duration) (shellenv.Vars, error) {
		return shellenv.Vars{"PATH": binDir + ":" + testPATH()}, nil
	}
	app.loadEnvironment()
	app.SaveRunSettings(proj, "true")

	res := app.AskAgent(AskAgentRequest{
		CLIID:    "claude",
		Action:   "fix-test",
		TestName: "T",
		Status:   "failed",
	})
	if !res.Blocked || res.Started {
		t.Fatalf("expected unavailable start, got %+v", res)
	}
	if !strings.Contains(res.Message, "Claude Code CLI is unavailable") {
		t.Fatalf("message=%q", res.Message)
	}
	st := app.GetAgentState()
	if !st.Unavailable {
		t.Fatalf("state=%+v", st)
	}

	// Project path, run, and results still work.
	suite := app.StartSuite(proj, "true")
	if suite.Blocked || !suite.Started {
		t.Fatalf("suite should still start: %+v", suite)
	}
	_ = waitForSuite(t, app)
}

func TestAskAgentSupportsThreeActions(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("unix fake CLI")
	}
	binDir := t.TempDir()
	writeFakeCLI(t, binDir, "claude", "#!/bin/sh\nprintf '%s\\n' \"$2\"\n")
	app := NewApp()
	app.loader = func(time.Duration) (shellenv.Vars, error) {
		return shellenv.Vars{"PATH": binDir}, nil
	}
	app.loadEnvironment()
	app.SaveRunSettings(t.TempDir(), "true")

	wants := map[string]string{
		"create-test": "create a test",
		"fix-test":    "fix this test",
		"fix-app":     "fix the application code this test is about",
	}
	for action, phrase := range wants {
		res := app.AskAgent(AskAgentRequest{
			CLIID:    "claude",
			Action:   action,
			TestName: "T",
			Status:   "failed",
		})
		if res.Blocked || !res.Started {
			t.Fatalf("action %s: %+v", action, res)
		}
		final := waitForAgent(t, app)
		if !strings.Contains(final.Output, phrase) {
			t.Fatalf("action %s output missing %q:\n%s", action, phrase, final.Output)
		}
		if !final.PendingDecision {
			t.Fatalf("action %s: want pending proposal, got %+v", action, final)
		}
		disc := app.DiscardProposal()
		if disc.Blocked || !disc.Started {
			t.Fatalf("action %s discard: %+v", action, disc)
		}
	}
}

func TestAcceptProposalKeepsChangeDiscardRestores(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("unix fake CLI")
	}
	binDir := t.TempDir()
	writeFakeCLI(t, binDir, "claude", `#!/bin/sh
printf '%s\n' "$2"
printf 'patched\n' > changed.txt
`)
	proj := t.TempDir()
	if err := os.WriteFile(filepath.Join(proj, "changed.txt"), []byte("original\n"), 0o644); err != nil {
		t.Fatal(err)
	}
	mustWriteExtra := filepath.Join(proj, "keep.txt")
	if err := os.WriteFile(mustWriteExtra, []byte("keep\n"), 0o644); err != nil {
		t.Fatal(err)
	}

	app := NewApp()
	app.loader = func(time.Duration) (shellenv.Vars, error) {
		return shellenv.Vars{"PATH": binDir + ":" + testPATH()}, nil
	}
	app.loadEnvironment()
	app.SaveRunSettings(proj, "true")

	ask := app.AskAgent(AskAgentRequest{
		CLIID:    "claude",
		Action:   "fix-app",
		TestName: "Login",
		Status:   "failed",
		Failure:  "boom",
	})
	if ask.Blocked || !ask.Started {
		t.Fatalf("ask=%+v", ask)
	}
	final := waitForAgent(t, app)
	if !final.PendingDecision || final.Failed {
		t.Fatalf("final=%+v", final)
	}
	if len(final.ProposedChanges) == 0 {
		t.Fatalf("expected proposed file changes, got %+v", final.ProposedChanges)
	}
	// Nothing remains in the project until accept.
	body, err := os.ReadFile(filepath.Join(proj, "changed.txt"))
	if err != nil {
		t.Fatal(err)
	}
	if string(body) != "original\n" {
		t.Fatalf("project changed before accept: %q", body)
	}

	acc := app.AcceptProposal()
	if acc.Blocked || !acc.Started {
		t.Fatalf("accept=%+v", acc)
	}
	st := app.GetAgentState()
	if !st.Accepted || !st.CanRerun || st.PendingDecision {
		t.Fatalf("after accept=%+v", st)
	}
	body, err = os.ReadFile(filepath.Join(proj, "changed.txt"))
	if err != nil {
		t.Fatal(err)
	}
	if string(body) != "patched\n" {
		t.Fatalf("after accept file=%q", body)
	}
}

func TestDiscardProposalRestoresTree(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("unix fake CLI")
	}
	binDir := t.TempDir()
	writeFakeCLI(t, binDir, "claude", `#!/bin/sh
printf 'x\n' > newfile.txt
`)
	proj := t.TempDir()
	if err := os.WriteFile(filepath.Join(proj, "old.txt"), []byte("old\n"), 0o644); err != nil {
		t.Fatal(err)
	}

	app := NewApp()
	app.loader = func(time.Duration) (shellenv.Vars, error) {
		return shellenv.Vars{"PATH": binDir + ":" + testPATH()}, nil
	}
	app.loadEnvironment()
	app.SaveRunSettings(proj, "true")

	if res := app.AskAgent(AskAgentRequest{
		CLIID: "claude", Action: "create-test", TestName: "T", Status: "failed",
	}); res.Blocked || !res.Started {
		t.Fatalf("ask=%+v", res)
	}
	final := waitForAgent(t, app)
	if !final.PendingDecision {
		t.Fatalf("final=%+v", final)
	}
	if _, err := os.Stat(filepath.Join(proj, "newfile.txt")); !os.IsNotExist(err) {
		t.Fatalf("newfile should not remain before accept: %v", err)
	}

	disc := app.DiscardProposal()
	if disc.Blocked || !disc.Started {
		t.Fatalf("discard=%+v", disc)
	}
	st := app.GetAgentState()
	if !st.Discarded || st.PendingDecision || st.CanRerun {
		t.Fatalf("after discard=%+v", st)
	}
	if _, err := os.Stat(filepath.Join(proj, "newfile.txt")); !os.IsNotExist(err) {
		t.Fatalf("newfile after discard: %v", err)
	}
	body, err := os.ReadFile(filepath.Join(proj, "old.txt"))
	if err != nil || string(body) != "old\n" {
		t.Fatalf("old.txt=%q err=%v", body, err)
	}
}

func TestRerunTestShowsPreviousBesideNew(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("unix fake CLI")
	}
	binDir := t.TempDir()
	writeFakeCLI(t, binDir, "claude", `#!/bin/sh
printf 'ok\n' > fix.txt
`)
	proj := t.TempDir()
	app := NewApp()
	app.loader = func(time.Duration) (shellenv.Vars, error) {
		return shellenv.Vars{"PATH": binDir + ":" + testPATH()}, nil
	}
	app.loadEnvironment()

	suiteCmd := strings.Join([]string{
		`printf '%s\n' 'mellis:test waiting Login'`,
		`printf '%s\n' 'mellis:test failed Login failure=Expected button duration=12ms'`,
	}, "; ")
	app.SaveRunSettings(proj, suiteCmd)
	if res := app.StartSuite(proj, suiteCmd); res.Blocked || !res.Started {
		t.Fatalf("suite=%+v", res)
	}
	first := waitForSuite(t, app)
	if len(first.Tests) != 1 || first.Tests[0].Status != suiteparse.StatusFailed {
		t.Fatalf("first=%+v", first.Tests)
	}

	if res := app.AskAgent(AskAgentRequest{
		CLIID: "claude", Action: "fix-test", TestName: "Login", Status: "failed", Failure: "Expected button",
	}); res.Blocked || !res.Started {
		t.Fatalf("ask=%+v", res)
	}
	_ = waitForAgent(t, app)
	if res := app.AcceptProposal(); res.Blocked || !res.Started {
		t.Fatalf("accept=%+v", res)
	}

	rerunCmd := strings.Join([]string{
		`printf '%s\n' 'mellis:test waiting Login'`,
		`printf '%s\n' 'mellis:test passed Login duration=9ms'`,
	}, "; ")
	app.SaveRunSettings(proj, rerunCmd)

	rerun := app.RerunTest("Login")
	if rerun.Blocked || !rerun.Started {
		t.Fatalf("rerun=%+v", rerun)
	}
	second := waitForSuite(t, app)
	if len(second.Tests) != 1 {
		t.Fatalf("second tests=%+v", second.Tests)
	}
	te := second.Tests[0]
	if te.Name != "Login" || te.Status != suiteparse.StatusPassed {
		t.Fatalf("new result=%+v", te)
	}
	if !te.HasPrevious || te.PreviousStatus != suiteparse.StatusFailed {
		t.Fatalf("previous missing: %+v", te)
	}
	if te.PreviousFailure != "Expected button" {
		t.Fatalf("previous failure=%q", te.PreviousFailure)
	}
}
