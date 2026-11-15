# Laravel login fixture

Basic Laravel app used to exercise Mellis against a real username/password E2E workflow.

## What you get

- Laravel app with a `/login` form (username + password)
- Docker Compose stack: `app` + MySQL
- Seeded user stored in MySQL
- Host-side Playwright suite for the login happy path

## Login URL

After Compose is up: [http://localhost:8000/login](http://localhost:8000/login)

## Seeded credentials

| Field | Value |
| --- | --- |
| Username | `demo@example.com` |
| Password | `password` |

Username is the user's email address.

## Start the application (Docker)

From this directory:

```bash
docker compose up --build -d
```

Wait until `app` is healthy, then open the login URL above.

Stop:

```bash
docker compose down
```

## Playwright suite command (host)

The suite runs on the host against the already-running Docker app. Do not start the app from Playwright.

```bash
npm install
npx playwright install chromium
npm run test:e2e
```

Suite command for Mellis:

```text
npm run test:e2e
```

Project path for Mellis: this directory (`fixtures/laravel-login`).

The Playwright config includes `e2e/mellis-reporter.ts`, which prints `mellis:test` progress lines so Mellis can show the login test while the suite runs.
