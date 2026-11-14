---
id: TKT-6
type: new-story
status: draft
story: US-6
---

# TKT-6: Read the result

**Type:** new-story
**Story:** [US-6](../stories/US-6-read-the-result.md)
**Status:** draft

## Request

with the user story we created and our target for a shippable product, can you create all the tickets we needed?

This ticket makes US-6 true for the tests reported by TKT-5. The story file already exists. Do not add a second story. Performance is how long the test took in this run. A user story is shown only when the test metadata already contains one.

## Acceptance criteria

* AC-1: Each reported test shows whether it passed or failed.
* AC-2: When the output includes a duration, the test shows how long it took in this run.
* AC-3: When the output includes no duration, the test shows pass or fail and shows no duration.
* AC-4: When the test's metadata contains a user story, the row shows that story.
* AC-5: When the test's metadata contains a user story, opening the test shows the story text.
* AC-6: When the test's metadata has no user story, the row shows the test name and shows no story.
* AC-7: When a failed test's output includes failure text, the test shows that text.
* AC-8: The finished window remains usable when the developer does not ask an agent to change the project.
* AC-9: When the command finishes and reports no tests, the window says the run finished with no tests.

## Action items

- [ ] AI-1 (AC-1): Show passed or failed on each reported test.
- [ ] AI-2 (AC-2): When the output includes a duration, show how long that test took in this run.
- [ ] AI-3 (AC-3): When the output includes no duration, show pass or fail and show no duration.
- [ ] AI-4 (AC-4): When the test metadata contains a user story, show that story on the row.
- [ ] AI-5 (AC-5): When the test metadata contains a user story, opening the test shows the story text.
- [ ] AI-6 (AC-6): When the test metadata has no user story, show the test name and show no story.
- [ ] AI-7 (AC-7): When a failed test's output includes failure text, show that text.
- [ ] AI-8 (AC-8): The finished window remains usable when the developer does not ask an agent to change the project.
- [ ] AI-9 (AC-9): When the command finishes and reports no tests, the window says the run finished with no tests.

## Sign-off

- [ ] Every action item above is checked
- [ ] Every acceptance criterion above is met
- Date:
- Result:
