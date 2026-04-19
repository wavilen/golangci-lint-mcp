# exhaustruct

<instructions>
Exhaustruct checks that struct literals initialize all exported fields. Omitting fields can lead to zero-value bugs where a struct is used with unintended default values.

Either initialize every exported field in the literal or explicitly use the field name with its zero value to signal intent.
</instructions>

<examples>
## Bad
```go
cfg := Config{
    Host: "localhost",
    Port: 8080,
    // Timeout is missing — defaults to 0, meaning no timeout
}
```

## Good
```go
cfg := Config{
    Host:    "localhost",
    Port:    8080,
    Timeout: 30 * time.Second,
}
```
</examples>

<patterns>
- Configuration structs with many fields where some are silently zero-valued
- DTOs constructed without required fields, causing downstream nil or zero checks
- API response structs missing optional but semantically important fields
- Test fixtures that omit fields, masking bugs that only appear in production
</patterns>

<related>
exhaustive, govet, revive
</related>
