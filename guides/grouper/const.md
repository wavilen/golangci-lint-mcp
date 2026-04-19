# grouper: const

<instructions>
Detects multiple consecutive `const` declarations that can be grouped into a single `const` block. Scattered `const` statements are harder to read and maintain. Group related constants together using the `const ( ... )` block form, especially when they share a type or are logically related.
</instructions>

<examples>
## Bad
```go
const StatusOK = 200
const StatusNotFound = 404
const StatusError = 500
```

## Good
```go
const (
    StatusOK      = 200
    StatusNotFound = 404
    StatusError   = 500
)
```
</examples>

<patterns>
- Multiple `const` declarations in sequence — group into `const ( ... )`
- Related constants declared separately — group for readability
- Single const per line in a block of declarations — use grouped form
</patterns>

<related>
var, type, import
