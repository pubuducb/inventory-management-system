# Overview of the files in the folder `internal/repository/`

## Database Configurations and Setup `product_repo.go`

Performs database operations by executing SQL queries. The use of parameterized queries ensures safe updates and reduces the risk of SQL injection.

During update and archive (delete) operations, the repository validates whether any rows were affected; if not, it treats the operation as an error.
