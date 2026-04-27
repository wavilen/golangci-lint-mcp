# grouper: const

<instructions>
Detects multiple consecutive `const` declarations that can be grouped into a single `const` block. Scattered `const` statements are harder to read and maintain. Group related constants together using the `const ( ... )` block form, especially when they share a type or are logically related.
</instructions>

<examples>
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
- Group sequential `const` declarations into a single `const ( ... )` block
- Group related constants declared separately for readability
- Use the grouped `const ( ... )` form instead of single-const-per-line declarations
</patterns>

<related>
grouper/var, grouper/type, grouper/import
</related>
