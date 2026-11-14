---
name: discovery
description: Discovery team for new-feature requests at the analysis stage. Six independent subagents in three pillars - brainstorm, explore, research - convened and facilitated by the analyst. Produces a discovery report that feeds the PRD/FRD. No design, no code.
facilitator: analyst
members:
  - david-kelley
  - alex-osborn
  - alberto-savoia
  - jeff-gothelf
  - rob-fitzpatrick
  - erika-hall
---

# Team: Discovery

Brainstorm, explore and research a new-feature request before it becomes a requirement. The analyst ([.agents/roles/product.md](/.agents/roles/product.md), [.cursor/agents/analyst.md](/.cursor/agents/analyst.md)) convenes the team and facilitates every session. Members never talk to each other directly. Everything goes through the facilitator.

Brought from the BeezHub discovery team on 2026-11-14. Member briefs and the session agenda are the same practice. Memory and sessions start empty: BeezHub domain memory was not copied, because this is a different product.

## Members

| Pillar | Member | Asks | Brings |
|---|---|---|---|
| Brainstorm | [David Kelley](/.agents/teams/discovery/members/david-kelley.md) | What's the real problem, and for whom? | Point of view, How Might We framing, judgment-free ideation, clusters, storyboards |
| Brainstorm | [Alex Osborn](/.agents/teams/discovery/members/alex-osborn.md) | What else could it be? | Quantity, Osborn's four rules, SCAMPER, combinations |
| Explore | [Alberto Savoia](/.agents/teams/discovery/members/alberto-savoia.md) | Should we build it at all? | XYZ hypotheses, pretotypes (fake door, Mechanical Turk, ...), skin in the game |
| Explore | [Jeff Gothelf](/.agents/teams/discovery/members/jeff-gothelf.md) | Which outcome, and what's the smallest bet? | Lean UX canvas, hypothesis statements, risk × value, in-sprint experiments |
| Research | [Rob Fitzpatrick](/.agents/teams/discovery/members/rob-fitzpatrick.md) | What have people actually done? | Evidence audit, Mom Test interview guides, commitment signals, question rewrites |
| Research | [Erika Hall](/.agents/teams/discovery/members/erika-hall.md) | What do we need to know, and how much is enough? | Research questions, method choice, bias check, stop rule |

> [!note] Personas
> Each member applies the published methods of the person it is named after. It is not that person: it doesn't claim to be them or put words in their mouth.

## Paths

Team root: [/.agents/teams/discovery/](/.agents/teams/discovery/). Use root-absolute paths (`/.agents/teams/discovery/...`) when briefing members. The repo root is this repository.

- `members/<slug>.md`: each member's brief (its subagent instructions).
- `memory/<slug>.md`: each member's memory. Only the owner writes it.
- `sessions/<YYYY-MM-DD>-<slug>.md`: discovery reports. Only the facilitator writes them.

## Memory ownership

- Each memory file has one owner (its `owner` property). Only that member's agent writes it.
- Nobody else writes it: not other members, not the facilitator, not the main session. To change what a member remembers, tell the member and let it decide.
- Members don't read each other's memory. They hear from teammates only through the facilitator, which keeps six independent perspectives instead of one.
- The facilitator may read any memory file.

## When to convene

- New capability: the full team, before the product design doc is drafted.
- Small extension of an existing flow: one member per pillar. The facilitator says which and why.
- Bugfix: never. Restoring intended behaviour needs no ideation.

## Session agenda (facilitator)

Each member is its own subagent. Never play a member yourself or run two members in one agent, because that is one perspective wearing six hats.

### 0. Prep

Ground the request first, so the team works from facts rather than the raw ask:

- Read the request, verbatim.
- Check `docs/changelogs/` for recent related changes and `docs/features/<area>/` (per [.agents/roles/product.md](/.agents/roles/product.md)) for existing behaviour this touches.
- Note the domain summary: actors, key entities, the existing flow, with file paths.
- Note constraints the user stated, and anything already decided.
- Note known unknowns.

Then write the session brief:

- The request, verbatim.
- Domain summary: actors, key entities, the existing flow this touches, with file paths.
- Constraints the user stated, and anything already decided.
- Known unknowns.
- Evidence available: docs, tickets, data, notes from the user. If there is no user research, say so plainly.

### 1. Round 1: diverge (parallel, independent)

Spawn every attendee in one message, in parallel, with the brief (template below). Nobody sees anyone else's Round 1, so no one anchors on the first idea voiced. Keep each agent's ID for later rounds.

### 2. Synthesis

Build the board from all replies without judging it:

- Keep every idea under its member ID (`K3`, `O12`, ...). Cluster by user need, starting from Kelley's clusters when present. Merge exact duplicates; keep near-duplicates side by side.
- Collect the hypotheses (Gothelf, Savoia) and the evidence audit (Fitzpatrick, Hall) into one assumption list.
- Shortlist 2–4 directions for Round 2, one line on why for each. Favour breadth over favourites; a wild idea can make the list.
- Draft your shape-changing questions for the user.

### 3. Round 2: build and test (parallel)

Send the board, the shortlist and your draft questions to every attendee with `SendMessage`, which keeps its context. If an agent can't be reached, spawn a fresh one with the brief and the board. Each member works Round 2 of its brief.

### 4. Converge

Write the discovery report (format below). Keep the dissent: two members disagreeing is information for the user, so don't average it away. Then hand off to the analyst ([.agents/roles/product.md](/.agents/roles/product.md)): ask the user the shape-changing questions with Fitzpatrick's rewrites applied, and present the shortlist with its cheapest tests, before drafting the PRD/FRD in `docs/features/<area>/`.

### 5. Close

Once the user has decided, fill in the report's Outcome. Send each attendee the outcome: what the user chose and why, what was dropped, and what goes into the PRD/FRD. Ask each one to update its own memory and reply with the lines it changed. Always run close, even when a session is abandoned, so the team still learns from it.

One extra round is allowed if the user's answers reopen the problem. Needing more than that means the request is really several requests: tell the user and split it.

## Briefing a member

Agent tool, `subagent_type: general-purpose`. Round 1 prompt:

```
You are <Name>, a member of the discovery team (<pillar> pillar).
Read these first, in order:
1. /.agents/teams/discovery/TEAM.md, sections "Member protocol" and "Memory rules"
2. /.agents/teams/discovery/members/<slug>.md (your brief)
3. /.agents/teams/discovery/memory/<slug>.md (your memory; only you write it)

Session: <YYYY-MM-DD>-<slug>
<session brief>

Do Round 1 of your brief and reply in its output format. Don't update your memory yet; that happens at close.
```

Round 2, via `SendMessage`: `Round 2. Board: <board>. Shortlist: <shortlist>. Draft questions to the user: <questions>. Do Round 2 of your brief.`

Close, via `SendMessage`: `Session closed. Outcome: <decisions and why, what was dropped, what goes into the PRD/FRD>. Report: <path>. Update your memory now per Memory rules, then reply with the lines you added, changed or removed.`

## Member protocol

Every member follows these.

- **Start:** read your brief, then your memory. When a memory entry changes your contribution, say so.
- **Label every claim** `[evidence: <source>]` (a code path, doc, ticket, data from the brief, a cited web page) or `[assumption]`. Never invent user quotes, interview answers, metrics or experiment results. You design research and tests; the results come from real people.
- **Product level only:** what users can do and why it matters. Architecture, schemas and technology choices belong to a later architecture pass, after this session. If the requester asks for a stack, stay in your lens and put the product requirements that would make a stack fail under "Notes outside my lens". Do not pick a framework.
- **Read anything** in the product repos and docs. In the team folder, read only TEAM.md, your own brief, your own memory, and sessions the brief links.
- **Write nothing** except your own memory file. No code, no docs, no other team files.
- **ID your items** with your letter: K (Kelley), O (Osborn), S (Savoia), G (Gothelf), F (Fitzpatrick), H (Hall). `K3` is Kelley's third idea.
- **Stay in your lens.** Something outside it goes in one line under "Notes outside my lens".
- **Web research** is allowed for comparable products and published practice. Cite the source and label it as other people's data, not our users.
- **Reply** in the output format of your brief, and nothing else.

## Memory rules

Every member follows these. Update your memory only at close, when the facilitator sends the outcome.

- Keep durable, reusable learning only: domain facts, the user's decisions and standing preferences (with the reason), techniques that landed or flopped here, and parked ideas worth revisiting. No transcripts.
- One line per entry, with an absolute date and a link to the session: `2026-09-27 — <learning> ([/.agents/teams/discovery/sessions/2026-09-27-<slug>.md](/.agents/teams/discovery/sessions/2026-09-27-<slug>.md))`.
- When an entry turns out wrong, correct or delete it. Don't stack contradictions.
- Keep the file under about 150 lines, merging entries before adding new ones.
- Set `updated` in the frontmatter to today's date.

## Discovery report

The facilitator writes one per session at `sessions/<YYYY-MM-DD>-<slug>.md`:

```
---
date: <YYYY-MM-DD>
request: <one line>
attendees: [<slugs>]
status: open | decided | abandoned
---
# Discovery: <title>

**Request (verbatim):** ...
**Brief:** domain summary with paths, constraints, evidence available.

## Framing
Point of view · chosen How Might We questions · business and user outcomes

## Idea board
<cluster (user need)>: K1 ..., O4 ..., (n ideas in total, by member)
**Shortlist:** <direction>: why

## Hypotheses and riskiest assumptions
| ID | Hypothesis | XYZ | Risk | Value | Evidence |

## Cheapest tests
| Assumption | Pretotype / experiment | Pass if (set in advance) | Cost |

## Research plan
Research questions · already answered (with source) · method, who, how many, time box · interview guide · stop rule

## Dissent
Who disagreed, on what, and each side's reason.

## Questions for the user
Shape-changing only, checked against the Mom Test.

## Outcome
The user's decisions and why · what enters the PRD/FRD · what was parked and why · doc link (`docs/features/<area>/...`)
```
