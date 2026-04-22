# containedctx

<instructions>
Containedctx detects structs that contain `context.Context` fields. Storing a context in a struct is an anti-pattern because contexts should flow through function call stacks, not be stored as dependencies. It makes lifetimes opaque and breaks cancellation propagation.

Pass the context as the first argument to methods instead of storing it on the struct.
</instructions>

<examples>
## Bad
```go
type Service struct {
    ctx context.Context
    db  *sql.DB
}

func (s *Service) Query() (*sql.Rows, error) {
    return s.db.QueryContext(s.ctx, "SELECT ...")
}
```

## Good
```go
type Service struct {
    db *sql.DB
}

func (s *Service) Query(ctx context.Context) (*sql.Rows, error) {
    return s.db.QueryContext(ctx, "SELECT ...")
}
```
</examples>

<patterns>
- Pass `context.Context` as the first function argument instead of storing it on a struct
- Pass `context.Context` as the first method argument rather than capturing it in a constructor
- Pass `context.Context` per operation in background workers rather than embedding it as a struct field
</patterns>

<related>
contextcheck, revive, gocritic
</related>
