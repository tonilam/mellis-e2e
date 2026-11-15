# Tickets

Development starts here. Run [ticket-start](../../.agents/workflows/ticket-start/SKILL.md). When the ticket you need is not listed, that workflow creates it before any code.

A ticket is one of: `new-story`, `update-story`, or `bug`. It links one user story, states pass/fail acceptance criteria, and lists the action items to check off. Sign-off waits until every action item is checked and every criterion is met.

These eight tickets are the first shippable runner, in build order. Each one delivers a story that already exists. Later tickets extend the product or fixtures. No application code starts while a ticket is `draft`.

| ID | Type | Story | Status | Title |
| --- | --- | --- | --- | --- |
| [TKT-1](TKT-1-open-the-app.md) | new-story | [US-1](../stories/US-1-open-the-app.md) | signed-off | Open the installed app |
| [TKT-2](TKT-2-load-the-shell-environment.md) | new-story | [US-2](../stories/US-2-load-the-shell-environment.md) | signed-off | Load the shell environment |
| [TKT-3](TKT-3-choose-the-project-and-command.md) | new-story | [US-3](../stories/US-3-choose-the-project-and-command.md) | signed-off | Choose the project and the command |
| [TKT-4](TKT-4-run-the-suite.md) | new-story | [US-4](../stories/US-4-run-the-suite.md) | signed-off | Run the suite |
| [TKT-5](TKT-5-watch-the-run.md) | new-story | [US-5](../stories/US-5-watch-the-run.md) | signed-off | Watch the run |
| [TKT-6](TKT-6-read-the-result.md) | new-story | [US-6](../stories/US-6-read-the-result.md) | signed-off | Read the result |
| [TKT-7](TKT-7-ask-an-agent.md) | new-story | [US-7](../stories/US-7-ask-an-agent.md) | signed-off | Ask an installed agent |
| [TKT-8](TKT-8-accept-or-discard.md) | new-story | [US-8](../stories/US-8-accept-or-discard.md) | signed-off | Accept or discard the change |
| [TKT-9](TKT-9-laravel-login-fixture.md) | new-story | [US-9](../stories/US-9-laravel-login-fixture.md) | signed-off | Laravel login fixture |
| [TKT-10](TKT-10-open-settings-from-the-menu.md) | new-story | [US-10](../stories/US-10-open-settings-from-the-menu.md) | signed-off | Open Settings from the main app menu |
