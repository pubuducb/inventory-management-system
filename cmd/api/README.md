# Overview of the files in the folder `cmd/api/`

## Entry Point of the API `main.go`

### Database Abstraction

Database initialization and configuration have been moved to `internal/configs/db.go`. This decouples the database logic from `main.go`, ensuring that database-specific changes do not require modifications to the entry point.

### Routing and Handler Isolation

Handler initialization and router setup are now encapsulated within `internal/server/server.go`. This allows the API to scale with new routes and handlers without cluttering main.go. This way, **the `main` package doesn’t need to change as the API scales**.

### Graceful Shutdown

Implemented a graceful shutdown mechanism (via `waitForShutdown()`) to ensure the application handles termination signals safely, allowing active connections to complete and resources to close properly.
