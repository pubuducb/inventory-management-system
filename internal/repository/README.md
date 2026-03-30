# Overview of the files in the folder `internal/repository/`

## Database Configurations and Setup

Performs database operations by executing SQL queries. The use of parameterized queries ensures safe updates and reduces the risk of SQL injection.

During update and archive (delete) operations, the repository validates whether any rows were affected; if not, it treats the operation as an error.

### Parameterized Queries

An SQL statement that uses placeholders instead of directly embedding user input into the query string (like in `repo.db.QueryRow(query, product.Name, product.Price)`). The database engine receives the SQL structure (the code) first, parses it, and builds an execution plan before it ever sees the user data. The database does not combine them like a string.
