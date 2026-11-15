# US-9: Run a Laravel login fixture for E2E

* **User value:** As a developer, I want a basic Laravel app with MySQL in Docker and a Playwright login suite, so that I can exercise the runner against a real username/password workflow.
* **Source:** Request via ticket-start — sample project for testing the application
* **Tickets:** [TKT-9](../tickets/TKT-9-laravel-login-fixture.md)

## Acceptance criteria

* AC-1: A developer can start the Laravel application with Docker Compose and open a login page in a browser at a documented local URL.
* AC-2: The Docker Compose stack includes a MySQL service, and the Laravel application uses that MySQL database for user credentials.
* AC-3: A Playwright E2E test submits a valid username and password on the login page and the developer observes a successful authenticated state.
* AC-4: The Playwright suite command is documented and runs from the host against the already-running Docker application.
