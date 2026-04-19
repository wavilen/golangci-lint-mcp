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
- HTTP handlers that store request context on a struct instead of passing it through
- Service objects that capture context in a constructor
- Background workers that embed context as a field rather than accepting it per operation
</patterns>

<related>
contextcheck, revive, gocritic
</related>
