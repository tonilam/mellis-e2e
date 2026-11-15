---
id: TKT-9
type: new-story
status: signed-off
story: US-9
---

# TKT-9: Laravel login fixture

**Type:** new-story
**Story:** [US-9](../stories/US-9-laravel-login-fixture.md)
**Status:** signed-off

## Request

we need a basic Laravel project that could run in docker in order to test the application. The laravel docker image should include a mysql database, and have E2E playright test for the basic username password login workflow

## Acceptance criteria

* AC-1: A developer can start the Laravel application with Docker Compose and open a login page in a browser at a documented local URL.
* AC-2: The Docker Compose stack includes a MySQL service, and the Laravel application uses that MySQL database for user credentials.
* AC-3: A Playwright E2E test submits a valid username and password on the login page and the developer observes a successful authenticated state.
* AC-4: The Playwright suite command is documented and runs from the host against the already-running Docker application.

## Action items

- [x] AI-1 (AC-1): Scaffold a basic Laravel project with a username/password login page and Docker Compose so the app starts and the login URL is documented.
- [x] AI-2 (AC-2): Add a MySQL service to Docker Compose, wire Laravel to it, and seed a user whose credentials the login form accepts.
- [x] AI-3 (AC-3): Add a Playwright E2E test that logs in with the seeded username and password and asserts a successful authenticated state.
- [x] AI-4 (AC-4): Document the host Playwright suite command and verify it runs against the already-running Docker stack; record the pass/fail line in the ticket result.

## Sign-off

- [x] Every action item above is checked
- [x] Every acceptance criterion above is met
- Date: 2026-11-15
- Result: Verified `docker compose up --build -d` serves http://localhost:8000/login; MySQL holds seeded `demo@example.com`; `npm run test:e2e` → `1 passed (4.3s)`.
