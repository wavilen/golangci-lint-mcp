# embeddedstructfieldcheck

<instructions>
Embeddedstructfieldcheck detects embedded struct fields that are unexportated (lowercase) while the parent struct is exported. This can cause surprising API limitations — the embedded type's promoted methods are not accessible to external callers.

Either export the embedded type or use it as a named field instead of embedding. If the embedding is intentional and internal-only, keep the parent struct unexported too.
</instructions>

<examples>
## Bad
```go
type Server struct {
    *http.Client
    mu sync.Mutex
}
```

## Good
```go
type Server struct {
    client *http.Client
    mu     sync.Mutex
}
```
</examples>

<patterns>
- Use named fields instead of embedding unexported pointer types in exported structs
- Avoid embedding unexported interfaces in public API types — use them as named fields
- Prefer named `mu sync.Mutex` fields over embedding `sync.Mutex` in exported structs
- Guard public API surface by replacing convenience embeddings with explicit named fields
</patterns>

<related>
exhaustruct, tagliatelle, recvcheck
</related>
