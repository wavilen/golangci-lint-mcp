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
- Embedded unexported pointer types in exported structs
- Embedded unexported interfaces in public API types
- sync.Mutex or sync.WaitGroup embedded in exported structs
- Embedding for convenience that leaks into public API
</patterns>

<related>
exhaustruct, tagliatelle, recvcheck
</related>
