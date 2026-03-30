# Overview of the files in the folder `internal/model/`

## Entity

### Input Validation

The `binding:"required"` tag is a instruction for Gin's internal validator.

### Struct Tags `json:"..."`

Without these, Gin would use the capital-letter Go names in the API responses.

### Soft Deletion `ArchivedAt`

Instead of physically removing a row from the database, it is marked as deleted by storing a timestamp in that field. If the field is nil, the product is active; if it has a value, it is "soft deleted."

### Handling Nulls with Pointers `*time.Time`

In Go, a regular time.Time cannot be `null`; its default value is `0001-01-01`. By using a pointer, the field can be literally nil. This is to represent an "Optional" or "Nullable" database column.

### Clean JSON Output `omitempty`

This tag tells Go, if this pointer is nil, don't even include this key in the JSON response. This keeps the API payload small and clean by hiding fields that aren't relevant to that specific item.
