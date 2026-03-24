# Overview of the files in the folder `internal/handler/`

## Database Configurations and Setup `product_handler.go`

### Strict JSON Decoding

Instead of using Gin’s default binding, the code uses `json.Decoder` with `DisallowUnknownFields()`. This ensures that any unexpected fields in the request body are rejected, improving API security and preventing silent bugs from malformed input.

### Use of DTO for Input Validation

The `UpdateProduct` struct is used as a DTO (Data Transfer Object) for request validation, separate from the database model. This prevents direct exposure of internal models and allows stricter control over input fields and validation rules.
