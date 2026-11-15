import './style.css';
import './app.css';
import {Events} from '@wailsio/runtime';
import {
  WindowKind,
  GetEnvironmentStatus,
  RetryLoadEnvironment,
  GetRunSettings,
  SaveRunSettings,
  StartSuite,
  StopSuite,
  GetSuiteState,
  ListAgentCLIs,
  AskAgent,
  StopAgent,
  GetAgentState,
  AcceptProposal,
  DiscardProposal,
  RerunTest,
} from '../bindings/mellis/app.js';

document.querySelector('#app').innerHTML = `
  <main class="shell">
    <header class="header">
      <div class="header-row">
        <div class="header-brand">
          <h1 class="brand">Mellis</h1>
          <p class="tagline">Local E2E runner</p>
        </div>
        <div class="app-menu">
          <button type="button" class="btn" id="btn-app-menu" aria-expanded="false" aria-controls="app-menu" aria-haspopup="true">
            Menu
          </button>
          <nav class="app-menu-panel" id="app-menu" hidden aria-label="Main app menu">
            <button type="button" class="app-menu-item" data-nav="run" id="menu-run">Run</button>
            <button type="button" class="app-menu-item" data-nav="settings" id="menu-settings">Settings</button>
          </nav>
        </div>
      </div>
    </header>

    <div class="page" id="page-run">
      <section class="run-panel" aria-labelledby="run-heading">
        <h2 id="run-heading">Run</h2>
        <p class="run-copy">
          Start the suite from this window. Mellis runs the suite command and project
          path from Settings, with the environment loaded on open. It does not start,
          stop, or restart your project's application — keep that running yourself.
        </p>
        <p class="run-status" id="window-kind" role="status">Loading…</p>
        <div class="run-actions">
          <button type="button" class="btn" id="btn-start-suite" disabled>Start suite</button>
          <button type="button" class="btn" id="btn-stop-suite" disabled>Stop run</button>
        </div>
        <p class="action-result" id="action-result" role="status" hidden></p>
        <div class="progress-panel" id="progress-panel" hidden>
          <p class="progress-counts" id="progress-counts" role="status"></p>
          <ul class="test-list" id="test-list" aria-live="polite"></ul>
        </div>
      </section>

      <section class="agent-panel" aria-labelledby="agent-heading">
        <h2 id="agent-heading">Ask an agent</h2>
        <p class="run-copy">
          From a result, ask an installed agent CLI to create a test, fix this test,
          or fix the application code this test is about. Choose the CLI before you send.
          Accept or discard a proposed change before it remains in the project.
        </p>
        <p class="agent-status" id="agent-status" role="status">Open a test result to ask an agent.</p>
        <div class="proposal-panel" id="proposal-panel" hidden>
          <h3 class="proposal-heading">Proposed change</h3>
          <ul class="proposal-list" id="proposal-list"></ul>
          <div class="run-actions">
            <button type="button" class="btn" id="btn-accept-proposal" disabled>Accept</button>
            <button type="button" class="btn" id="btn-discard-proposal" disabled>Discard</button>
            <button type="button" class="btn" id="btn-rerun-test" disabled>Run this test again</button>
          </div>
        </div>
        <div class="run-actions">
          <button type="button" class="btn" id="btn-stop-agent" disabled>Stop agent</button>
        </div>
      </section>
    </div>

    <div class="page" id="page-settings" hidden>
      <section class="settings-panel" aria-labelledby="settings-heading">
        <div class="page-heading-row">
          <h2 id="settings-heading">Settings</h2>
          <button type="button" class="btn" id="btn-back-to-run">Back to run</button>
        </div>
        <p class="run-copy">
          Environment, project path, and suite command live here. The run page starts
          the suite with these values.
        </p>
      </section>

      <section class="status-panel" aria-labelledby="env-heading">
        <h2 id="env-heading">Environment</h2>
        <p class="env-status" id="env-status" role="status">Loading environment…</p>
        <div class="env-actions">
          <button type="button" class="btn" id="btn-view-env">View environment</button>
          <button type="button" class="btn" id="btn-retry-env">Retry load</button>
        </div>
      </section>

      <section class="project-panel" aria-labelledby="project-heading">
        <h2 id="project-heading">Project and command</h2>
        <p class="run-copy">
          Point Mellis at the project and confirm the command that starts its E2E suite.
          Values from the last session are offered here for you to confirm or change.
        </p>
        <label class="field" for="project-path">
          <span class="field-label">Project path</span>
          <input type="text" id="project-path" class="field-input" spellcheck="false" autocomplete="off" placeholder="/path/to/project" />
        </label>
        <p class="session-path" id="session-path" role="status" hidden></p>
        <label class="field" for="suite-command">
          <span class="field-label">Suite command</span>
          <input type="text" id="suite-command" class="field-input" spellcheck="false" autocomplete="off" placeholder="npm run e2e" />
        </label>
        <p class="command-preview" id="command-preview" role="status">Command before run: (none yet)</p>
      </section>
    </div>
  </main>

  <dialog class="env-dialog" id="env-dialog" aria-labelledby="env-dialog-title">
    <form method="dialog" class="env-dialog-inner">
      <header class="env-dialog-header">
        <h2 id="env-dialog-title">Loaded environment</h2>
        <button type="submit" class="btn btn-quiet" value="close">Close</button>
      </header>
      <p class="env-dialog-meta" id="env-dialog-meta"></p>
      <pre class="env-list" id="env-list" tabindex="0"></pre>
      <footer class="env-dialog-footer">
        <button type="button" class="btn" id="btn-retry-in-dialog">Retry load</button>
      </footer>
    </form>
  </dialog>

  <dialog class="result-dialog" id="result-dialog" aria-labelledby="result-dialog-title">
    <form method="dialog" class="result-dialog-inner">
      <header class="result-dialog-header">
        <h2 id="result-dialog-title">Test result</h2>
        <button type="submit" class="btn btn-quiet" value="close">Close</button>
      </header>
      <p class="result-dialog-status" id="result-dialog-status" role="status"></p>
      <p class="result-dialog-meta" id="result-dialog-duration" hidden></p>
      <section class="result-dialog-section" id="result-dialog-story-section" hidden>
        <h3>User story</h3>
        <p id="result-dialog-story"></p>
      </section>
      <section class="result-dialog-section" id="result-dialog-failure-section" hidden>
        <h3>Failure</h3>
        <pre id="result-dialog-failure"></pre>
      </section>
      <section class="result-dialog-section agent-ask-section" id="agent-ask-section">
        <h3>Ask an agent</h3>
        <p class="agent-ask-copy">Choose an installed CLI, then send one action.</p>
        <fieldset class="agent-cli-fieldset">
          <legend class="field-label">Agent CLI</legend>
          <div class="agent-cli-list" id="agent-cli-list"></div>
        </fieldset>
        <p class="agent-ask-note" id="agent-ask-note" role="status" hidden></p>
        <div class="agent-ask-actions">
          <button type="button" class="btn" data-agent-action="create-test">Create a test</button>
          <button type="button" class="btn" data-agent-action="fix-test">Fix this test</button>
          <button type="button" class="btn" data-agent-action="fix-app">Fix application code</button>
        </div>
      </section>
    </form>
  </dialog>
`;

const envStatusEl = document.getElementById('env-status');
const windowKindEl = document.getElementById('window-kind');
const actionResultEl = document.getElementById('action-result');
const btnViewEnv = document.getElementById('btn-view-env');
const btnRetryEnv = document.getElementById('btn-retry-env');
const btnRetryInDialog = document.getElementById('btn-retry-in-dialog');
const btnStartSuite = document.getElementById('btn-start-suite');
const btnStopSuite = document.getElementById('btn-stop-suite');
const btnStopAgent = document.getElementById('btn-stop-agent');
const btnAppMenu = document.getElementById('btn-app-menu');
const appMenu = document.getElementById('app-menu');
const btnBackToRun = document.getElementById('btn-back-to-run');
const pageRun = document.getElementById('page-run');
const pageSettings = document.getElementById('page-settings');
const envDialog = document.getElementById('env-dialog');
const envDialogMeta = document.getElementById('env-dialog-meta');
const envList = document.getElementById('env-list');
const projectPathInput = document.getElementById('project-path');
const suiteCommandInput = document.getElementById('suite-command');
const sessionPathEl = document.getElementById('session-path');
const commandPreviewEl = document.getElementById('command-preview');
const progressPanel = document.getElementById('progress-panel');
const progressCounts = document.getElementById('progress-counts');
const testList = document.getElementById('test-list');
const resultDialog = document.getElementById('result-dialog');
const resultDialogTitle = document.getElementById('result-dialog-title');
const resultDialogStatus = document.getElementById('result-dialog-status');
const resultDialogDuration = document.getElementById('result-dialog-duration');
const resultDialogStorySection = document.getElementById('result-dialog-story-section');
const resultDialogStory = document.getElementById('result-dialog-story');
const resultDialogFailureSection = document.getElementById('result-dialog-failure-section');
const resultDialogFailure = document.getElementById('result-dialog-failure');
const agentCliList = document.getElementById('agent-cli-list');
const agentAskNote = document.getElementById('agent-ask-note');
const agentStatusEl = document.getElementById('agent-status');
const agentAskActions = document.querySelectorAll('[data-agent-action]');
const proposalPanel = document.getElementById('proposal-panel');
const proposalList = document.getElementById('proposal-list');
const btnAcceptProposal = document.getElementById('btn-accept-proposal');
const btnDiscardProposal = document.getElementById('btn-discard-proposal');
const btnRerunTest = document.getElementById('btn-rerun-test');

let lastStatus = null;
let persistTimer = null;
let suiteRunning = false;
let agentRunning = false;
let agentCLIs = [];
let selectedCLIId = '';
let detailTest = null;
let lastAgentState = null;

function statusLabel(status) {
  switch (status) {
    case 'passed':
      return 'passed';
    case 'failed':
      return 'failed';
    case 'running':
      return 'running';
    case 'waiting':
      return 'waiting';
    default:
      return status || 'waiting';
  }
}

function setMenuOpen(open) {
  appMenu.hidden = !open;
  btnAppMenu.setAttribute('aria-expanded', open ? 'true' : 'false');
}

function showPage(page) {
  const next = page === 'settings' ? 'settings' : 'run';
  pageRun.hidden = next !== 'run';
  pageSettings.hidden = next !== 'settings';
  setMenuOpen(false);
}

function renderAgentCLIChoices() {
  agentCliList.innerHTML = '';
  if (!agentCLIs.length) {
    agentAskNote.hidden = false;
    agentAskNote.textContent = 'No agent CLIs are configured.';
    return;
  }

  for (const cli of agentCLIs) {
    const id = `agent-cli-${cli.id}`;
    const label = document.createElement('label');
    label.className = 'agent-cli-option';
    label.htmlFor = id;

    const input = document.createElement('input');
    input.type = 'radio';
    input.name = 'agent-cli';
    input.id = id;
    input.value = cli.id;
    input.disabled = !cli.available || !lastStatus?.loaded || agentRunning;
    if (cli.id === selectedCLIId && cli.available) {
      input.checked = true;
    }
    input.addEventListener('change', () => {
      if (input.checked) {
        selectedCLIId = cli.id;
        updateAgentAskNote();
      }
    });

    const text = document.createElement('span');
    text.textContent = cli.available
      ? cli.label
      : `${cli.label} (missing: ${cli.binary})`;

    label.appendChild(input);
    label.appendChild(text);
    agentCliList.appendChild(label);
  }

  const available = agentCLIs.filter((c) => c.available);
  if (!available.length) {
    selectedCLIId = '';
  } else if (!available.some((c) => c.id === selectedCLIId)) {
    selectedCLIId = available[0].id;
    const radio = document.getElementById(`agent-cli-${selectedCLIId}`);
    if (radio) {
      radio.checked = true;
    }
  }
  updateAgentAskNote();
  updateAgentAskButtons();
}

function updateAgentAskNote() {
  const missing = agentCLIs.filter((c) => !c.available);
  const parts = [];
  if (!lastStatus?.loaded) {
    parts.push('Environment is not loaded; agent actions are unavailable.');
  }
  for (const cli of missing) {
    parts.push(`${cli.label} is missing from the loaded environment (${cli.binary}).`);
  }
  if (parts.length) {
    agentAskNote.hidden = false;
    agentAskNote.textContent = parts.join(' ');
  } else {
    agentAskNote.hidden = true;
    agentAskNote.textContent = '';
  }
}

function updateAgentAskButtons() {
  const canAsk =
    !!lastStatus?.loaded &&
    !!selectedCLIId &&
    !agentRunning &&
    !!detailTest;
  for (const btn of agentAskActions) {
    btn.disabled = !canAsk;
  }
}

function openTestDetail(te) {
  detailTest = te;
  resultDialogTitle.textContent = te.name || 'Test result';
  resultDialogStatus.textContent = statusLabel(te.status);
  resultDialogStatus.className = `result-dialog-status is-${te.status || 'waiting'}`;

  if (te.duration) {
    resultDialogDuration.hidden = false;
    resultDialogDuration.textContent = `Duration: ${te.duration}`;
  } else {
    resultDialogDuration.hidden = true;
    resultDialogDuration.textContent = '';
  }

  if (te.story) {
    resultDialogStorySection.hidden = false;
    resultDialogStory.textContent = te.story;
  } else {
    resultDialogStorySection.hidden = true;
    resultDialogStory.textContent = '';
  }

  if (te.failure) {
    resultDialogFailureSection.hidden = false;
    resultDialogFailure.textContent = te.failure;
  } else {
    resultDialogFailureSection.hidden = true;
    resultDialogFailure.textContent = '';
  }

  renderAgentCLIChoices();
  resultDialog.showModal();
}

function renderTestRow(te) {
  const li = document.createElement('li');
  li.className = `test-row is-${te.status || 'waiting'}`;
  li.tabIndex = 0;
  li.setAttribute('role', 'button');
  li.setAttribute('aria-label', `Open details for ${te.name}`);

  const main = document.createElement('div');
  main.className = 'test-row-main';
  const nameEl = document.createElement('span');
  nameEl.className = 'test-row-name';
  nameEl.textContent = te.name;
  const statusEl = document.createElement('span');
  statusEl.className = 'test-row-status';
  statusEl.textContent = statusLabel(te.status);
  main.appendChild(nameEl);
  main.appendChild(statusEl);
  li.appendChild(main);

  if (te.hasPrevious) {
    const prev = document.createElement('p');
    prev.className = 'test-row-previous';
    const prevDur = te.previousDuration ? ` · ${te.previousDuration}` : '';
    prev.textContent = `Previous: ${statusLabel(te.previousStatus)}${prevDur}`;
    if (te.previousFailure) {
      prev.textContent += ` — ${te.previousFailure}`;
    }
    const current = document.createElement('p');
    current.className = 'test-row-current';
    const curDur = te.duration ? ` · ${te.duration}` : '';
    current.textContent = `New: ${statusLabel(te.status)}${curDur}`;
    li.appendChild(prev);
    li.appendChild(current);
  } else if (te.duration) {
    const dur = document.createElement('span');
    dur.className = 'test-row-duration';
    dur.textContent = te.duration;
    li.appendChild(dur);
  }

  if (te.story) {
    const story = document.createElement('p');
    story.className = 'test-row-story';
    story.textContent = te.story;
    li.appendChild(story);
  }

  if (te.failure) {
    const fail = document.createElement('pre');
    fail.className = 'test-row-failure';
    fail.textContent = te.failure;
    li.appendChild(fail);
  }

  const open = () => openTestDetail(te);
  li.addEventListener('click', open);
  li.addEventListener('keydown', (ev) => {
    if (ev.key === 'Enter' || ev.key === ' ') {
      ev.preventDefault();
      open();
    }
  });
  return li;
}

function renderProposal(st) {
  const pending = !!st?.pendingDecision;
  const canRerun = !!st?.canRerun;
  const accepted = !!st?.accepted;
  const discarded = !!st?.discarded;
  const changes = st?.proposedChanges || [];

  if (!pending && !accepted && !discarded) {
    proposalPanel.hidden = true;
    proposalList.innerHTML = '';
    btnAcceptProposal.disabled = true;
    btnDiscardProposal.disabled = true;
    btnRerunTest.disabled = true;
    return;
  }

  proposalPanel.hidden = false;
  proposalList.innerHTML = '';
  if (pending || accepted) {
    if (!changes.length) {
      const li = document.createElement('li');
      li.className = 'proposal-item';
      li.textContent = 'No file changes in the proposal.';
      proposalList.appendChild(li);
    } else {
      for (const ch of changes) {
        const li = document.createElement('li');
        li.className = `proposal-item is-${ch.kind || 'modified'}`;
        li.textContent = `${ch.kind || 'modified'}: ${ch.path}`;
        proposalList.appendChild(li);
      }
    }
  } else if (discarded) {
    const li = document.createElement('li');
    li.className = 'proposal-item';
    li.textContent = 'Proposal discarded. Project matches the tree from before the request.';
    proposalList.appendChild(li);
  }

  btnAcceptProposal.disabled = !pending;
  btnDiscardProposal.disabled = !pending;
  btnRerunTest.disabled = !canRerun || suiteRunning;
}

function applyEnvStatus(st) {
  lastStatus = st;
  const suiteBtn = btnStartSuite;

  if (st.loaded) {
    envStatusEl.textContent = 'Login-shell environment loaded.';
    envStatusEl.classList.remove('is-error');
    envStatusEl.classList.add('is-ok');
    suiteBtn.disabled = suiteRunning;
  } else {
    const detail = st.error ? `: ${st.error}` : '';
    envStatusEl.textContent = `Environment could not be loaded${detail}`;
    envStatusEl.classList.add('is-error');
    envStatusEl.classList.remove('is-ok');
    suiteBtn.disabled = true;
  }
  fillEnvDialog(st);
  return refreshAgentCLIs();
}

function fillEnvDialog(st) {
  if (!st || !st.loaded) {
    envDialogMeta.textContent = st?.error
      ? `Load failed: ${st.error}`
      : 'No environment loaded.';
    envList.textContent = '';
    return;
  }
  envDialogMeta.textContent = `${st.count} variables · PATH length ${st.path?.length ?? 0}`;
  const vars = st.vars || {};
  const keys = Object.keys(vars).sort();
  envList.textContent = keys.map((k) => `${k}=${vars[k]}`).join('\n');
}

function showSessionPath(path) {
  const trimmed = (path || '').trim();
  if (!trimmed) {
    sessionPathEl.hidden = true;
    sessionPathEl.textContent = '';
    return;
  }
  sessionPathEl.hidden = false;
  sessionPathEl.textContent = `Session project path: ${trimmed}`;
}

function showCommandPreview(command) {
  const trimmed = (command || '').trim();
  commandPreviewEl.textContent = trimmed
    ? `Command before run: ${trimmed}`
    : 'Command before run: (none yet)';
}

function schedulePersist() {
  clearTimeout(persistTimer);
  persistTimer = setTimeout(() => {
    SaveRunSettings(projectPathInput.value, suiteCommandInput.value).catch((err) => {
      console.error('save settings', err);
    });
  }, 300);
}

function onPathOrCommandInput() {
  showSessionPath(projectPathInput.value);
  showCommandPreview(suiteCommandInput.value);
  schedulePersist();
}

function showActionResult(res) {
  actionResultEl.hidden = false;
  actionResultEl.classList.remove('is-blocked', 'is-failed', 'is-ok');
  if (res.blocked) {
    actionResultEl.textContent = res.message;
    actionResultEl.classList.add('is-blocked');
    return;
  }
  if (res.failed) {
    actionResultEl.classList.add('is-failed');
    const detail = res.message || res.output || `Command exited with status ${res.exitCode}`;
    actionResultEl.textContent = `Suite failed.\n${detail}`;
    return;
  }
  actionResultEl.classList.add('is-ok');
  const out = res.output ? `\n${res.output}` : '';
  actionResultEl.textContent = `${res.message || 'Suite finished.'}${out}`;
}

function renderSuiteState(st) {
  if (!st) {
    return;
  }
  suiteRunning = !!st.running;
  btnStopSuite.disabled = !suiteRunning;
  if (lastStatus?.loaded) {
    btnStartSuite.disabled = suiteRunning;
  }
  if (lastAgentState) {
    btnRerunTest.disabled = !lastAgentState.canRerun || suiteRunning;
  }

  const tests = st.tests || [];
  const counts = st.counts || {finished: 0, running: 0, waiting: 0};
  const hasProgress = tests.length > 0 || st.running || st.finished;

  if (!hasProgress) {
    progressPanel.hidden = true;
    return;
  }

  progressPanel.hidden = false;
  progressCounts.textContent =
    `Finished: ${counts.finished ?? 0} · Running: ${counts.running ?? 0} · Waiting: ${counts.waiting ?? 0}`;

  testList.innerHTML = '';
  for (const te of tests) {
    testList.appendChild(renderTestRow(te));
  }

  actionResultEl.hidden = false;
  actionResultEl.classList.remove('is-blocked', 'is-failed', 'is-ok');
  if (st.running) {
    actionResultEl.textContent = 'Suite running…';
    return;
  }
  if (st.finished) {
    if (st.stopped) {
      actionResultEl.classList.add('is-ok');
      actionResultEl.textContent = 'Run stopped. Reported results are kept above.';
      return;
    }
    if (tests.length === 0) {
      actionResultEl.classList.add('is-ok');
      actionResultEl.textContent = 'The run finished with no tests.';
      return;
    }
    if (st.failed) {
      actionResultEl.classList.add('is-failed');
      const detail = st.message || st.output || `Command exited with status ${st.exitCode}`;
      actionResultEl.textContent = `Suite finished with failures.\n${detail}`;
      return;
    }
    actionResultEl.classList.add('is-ok');
    actionResultEl.textContent = 'Run finished.';
  }
}

function renderAgentState(st) {
  if (!st) {
    return;
  }
  lastAgentState = st;
  agentRunning = !!st.running;
  btnStopAgent.disabled = !agentRunning;
  updateAgentAskButtons();
  renderProposal(st);
  if (resultDialog.open) {
    renderAgentCLIChoices();
  }

  agentStatusEl.classList.remove('is-blocked', 'is-failed', 'is-ok', 'is-working');

  if (st.running) {
    agentStatusEl.classList.add('is-working');
    agentStatusEl.textContent = `Agent working${st.cliLabel ? ` (${st.cliLabel})` : ''}…`;
    return;
  }

  if (st.unavailable) {
    agentStatusEl.classList.add('is-blocked');
    const label = st.cliLabel || 'CLI';
    const named = st.cliBinary
      ? `${label} is unavailable (${st.cliBinary}).`
      : `${label} is unavailable.`;
    agentStatusEl.textContent = st.message || named;
    return;
  }

  if (st.pendingDecision) {
    agentStatusEl.classList.add('is-ok');
    agentStatusEl.textContent =
      st.message || 'Proposed change ready. Accept or discard before it remains in the project.';
    return;
  }

  if (st.finished) {
    if (st.stopped) {
      agentStatusEl.classList.add('is-ok');
      agentStatusEl.textContent =
        st.message ||
        'Agent request stopped. Reported test results are kept above. Project matches the tree from before the request.';
      return;
    }
    if (st.failed) {
      agentStatusEl.classList.add('is-failed');
      agentStatusEl.textContent = st.message || 'Agent request failed.';
      return;
    }
    agentStatusEl.classList.add('is-ok');
    agentStatusEl.textContent = st.message || 'Agent request finished.';
    return;
  }

  if (!st.message) {
    agentStatusEl.textContent = 'Open a test result to ask an agent.';
  }
}

async function refreshAgentCLIs() {
  try {
    agentCLIs = (await ListAgentCLIs()) || [];
  } catch (err) {
    console.error('list agent clis', err);
    agentCLIs = [];
  }
  if (resultDialog.open) {
    renderAgentCLIChoices();
  } else {
    updateAgentAskNote();
    updateAgentAskButtons();
  }
  return agentCLIs;
}

async function refreshEnv() {
  const st = await GetEnvironmentStatus();
  await applyEnvStatus(st);
  return st;
}

async function retryEnv() {
  envStatusEl.textContent = 'Retrying environment load…';
  const st = await RetryLoadEnvironment();
  await applyEnvStatus(st);
  return st;
}

async function loadRunSettings() {
  const settings = await GetRunSettings();
  projectPathInput.value = settings.projectPath || '';
  suiteCommandInput.value = settings.command || '';
  showSessionPath(projectPathInput.value);
  showCommandPreview(suiteCommandInput.value);
}

function sendAgentAsk(action) {
  if (!detailTest || !selectedCLIId) {
    return;
  }
  const req = {
    cliId: selectedCLIId,
    action,
    testName: detailTest.name || '',
    status: detailTest.status || '',
    failure: detailTest.failure || '',
    story: detailTest.story || '',
  };
  agentStatusEl.classList.remove('is-blocked', 'is-failed', 'is-ok');
  agentStatusEl.classList.add('is-working');
  agentStatusEl.textContent = 'Agent working…';
  agentRunning = true;
  btnStopAgent.disabled = false;
  updateAgentAskButtons();
  if (resultDialog.open) {
    resultDialog.close();
  }
  AskAgent(req)
    .then((res) => {
      if (res.blocked || !res.started) {
        agentRunning = false;
        btnStopAgent.disabled = true;
        agentStatusEl.classList.remove('is-working', 'is-ok', 'is-failed');
        agentStatusEl.classList.add('is-blocked');
        agentStatusEl.textContent = res.message || 'Agent CLI is unavailable.';
        updateAgentAskButtons();
        return GetAgentState().then(renderAgentState);
      }
      return GetAgentState().then(renderAgentState);
    })
    .catch((err) => {
      agentRunning = false;
      btnStopAgent.disabled = true;
      agentStatusEl.classList.remove('is-working');
      agentStatusEl.classList.add('is-blocked');
      agentStatusEl.textContent = String(err);
      updateAgentAskButtons();
    });
}

btnAppMenu.addEventListener('click', (ev) => {
  ev.stopPropagation();
  setMenuOpen(appMenu.hidden);
});

appMenu.addEventListener('click', (ev) => {
  const target = ev.target.closest('[data-nav]');
  if (!target) {
    return;
  }
  showPage(target.getAttribute('data-nav'));
});

btnBackToRun.addEventListener('click', () => {
  showPage('run');
});

document.addEventListener('click', (ev) => {
  if (!appMenu.hidden && !ev.target.closest('.app-menu')) {
    setMenuOpen(false);
  }
});

document.addEventListener('keydown', (ev) => {
  if (ev.key === 'Escape' && !appMenu.hidden) {
    setMenuOpen(false);
  }
});

projectPathInput.addEventListener('input', onPathOrCommandInput);
suiteCommandInput.addEventListener('input', onPathOrCommandInput);

btnViewEnv.addEventListener('click', () => {
  fillEnvDialog(lastStatus);
  envDialog.showModal();
});

btnRetryEnv.addEventListener('click', () => {
  retryEnv().catch((err) => {
    envStatusEl.textContent = `Retry failed: ${err}`;
    envStatusEl.classList.add('is-error');
  });
});

btnRetryInDialog.addEventListener('click', () => {
  retryEnv()
    .then((st) => {
      fillEnvDialog(st);
    })
    .catch((err) => {
      envDialogMeta.textContent = `Retry failed: ${err}`;
    });
});

btnStartSuite.addEventListener('click', () => {
  const projectPath = projectPathInput.value.trim();
  const command = suiteCommandInput.value.trim();
  showSessionPath(projectPath);
  showCommandPreview(command);
  btnStartSuite.disabled = true;
  actionResultEl.hidden = false;
  actionResultEl.classList.remove('is-blocked', 'is-failed', 'is-ok');
  actionResultEl.textContent = 'Starting suite…';
  progressPanel.hidden = true;
  testList.innerHTML = '';
  SaveRunSettings(projectPath, command)
    .then(() => StartSuite(projectPath, command))
    .then((res) => {
      if (res.blocked || !res.started) {
        showActionResult(res);
        if (lastStatus?.loaded) {
          btnStartSuite.disabled = false;
        }
        btnStopSuite.disabled = true;
        suiteRunning = false;
        return;
      }
      suiteRunning = true;
      btnStopSuite.disabled = false;
      actionResultEl.textContent = 'Suite running…';
      return GetSuiteState().then(renderSuiteState);
    })
    .catch((err) => {
      actionResultEl.hidden = false;
      actionResultEl.classList.remove('is-ok', 'is-failed');
      actionResultEl.classList.add('is-blocked');
      actionResultEl.textContent = String(err);
      suiteRunning = false;
      btnStopSuite.disabled = true;
      if (lastStatus?.loaded) {
        btnStartSuite.disabled = false;
      }
    });
});

btnStopSuite.addEventListener('click', () => {
  btnStopSuite.disabled = true;
  StopSuite()
    .then(() => GetSuiteState())
    .then(renderSuiteState)
    .catch((err) => {
      actionResultEl.hidden = false;
      actionResultEl.classList.add('is-blocked');
      actionResultEl.textContent = String(err);
    });
});

btnStopAgent.addEventListener('click', () => {
  btnStopAgent.disabled = true;
  StopAgent()
    .then(() => GetAgentState())
    .then(renderAgentState)
    .catch((err) => {
      agentStatusEl.classList.add('is-blocked');
      agentStatusEl.textContent = String(err);
    });
});

btnAcceptProposal.addEventListener('click', () => {
  btnAcceptProposal.disabled = true;
  btnDiscardProposal.disabled = true;
  AcceptProposal()
    .then((res) => {
      if (res.blocked) {
        agentStatusEl.classList.add('is-blocked');
        agentStatusEl.textContent = res.message || 'Could not accept.';
      }
      return GetAgentState();
    })
    .then(renderAgentState)
    .catch((err) => {
      agentStatusEl.classList.add('is-blocked');
      agentStatusEl.textContent = String(err);
      btnAcceptProposal.disabled = false;
      btnDiscardProposal.disabled = false;
    });
});

btnDiscardProposal.addEventListener('click', () => {
  btnAcceptProposal.disabled = true;
  btnDiscardProposal.disabled = true;
  DiscardProposal()
    .then((res) => {
      if (res.blocked) {
        agentStatusEl.classList.add('is-blocked');
        agentStatusEl.textContent = res.message || 'Could not discard.';
      }
      return GetAgentState();
    })
    .then(renderAgentState)
    .catch((err) => {
      agentStatusEl.classList.add('is-blocked');
      agentStatusEl.textContent = String(err);
      btnAcceptProposal.disabled = false;
      btnDiscardProposal.disabled = false;
    });
});

btnRerunTest.addEventListener('click', () => {
  const testName = lastAgentState?.testName || '';
  if (!testName) {
    return;
  }
  btnRerunTest.disabled = true;
  actionResultEl.hidden = false;
  actionResultEl.classList.remove('is-blocked', 'is-failed', 'is-ok');
  actionResultEl.textContent = `Running ${testName} again…`;
  RerunTest(testName)
    .then((res) => {
      if (res.blocked || !res.started) {
        showActionResult(res);
        btnRerunTest.disabled = !lastAgentState?.canRerun;
        return;
      }
      suiteRunning = true;
      btnStopSuite.disabled = false;
      if (lastStatus?.loaded) {
        btnStartSuite.disabled = true;
      }
      return GetSuiteState().then(renderSuiteState);
    })
    .catch((err) => {
      actionResultEl.classList.add('is-blocked');
      actionResultEl.textContent = String(err);
      btnRerunTest.disabled = !lastAgentState?.canRerun;
    });
});

for (const btn of agentAskActions) {
  btn.addEventListener('click', (ev) => {
    ev.preventDefault();
    sendAgentAsk(btn.getAttribute('data-agent-action'));
  });
}

Events.On('suite:state', (ev) => {
  const st = ev?.data ?? ev;
  renderSuiteState(st);
});

Events.On('agent:state', (ev) => {
  renderAgentState(ev?.data ?? ev);
});

WindowKind()
  .then((kind) => {
    windowKindEl.textContent =
      kind === 'app-window'
        ? 'Showing the run in the app window.'
        : `Unexpected window kind: ${kind}`;
  })
  .catch((err) => {
    windowKindEl.textContent = `Could not confirm window kind: ${err}`;
    console.error(err);
  });

Promise.all([refreshEnv(), loadRunSettings()])
  .then(() => Promise.all([GetSuiteState(), GetAgentState()]))
  .then(([suiteSt, agentSt]) => {
    if (suiteSt && (suiteSt.running || suiteSt.finished || (suiteSt.tests && suiteSt.tests.length))) {
      renderSuiteState(suiteSt);
    }
    if (agentSt) {
      renderAgentState(agentSt);
    }
  })
  .catch((err) => {
    envStatusEl.textContent = `Environment could not be loaded: ${err}`;
    envStatusEl.classList.add('is-error');
    btnStartSuite.disabled = true;
  });
