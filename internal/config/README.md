# Overview of the files in the folder `internal/config/`

## Database Configurations and Setup `db.go`

Initializes a SQLite database using environment-based configuration with a fallback default, applies connection pooling settings, and ensures schema consistency via idempotent table creation.

Database Initialization (i.e. `InitDB`) and Schema Creation (i.e. `SetupDB`) are handled by separate functions to ensure clear separation of concerns and independent error handling.
