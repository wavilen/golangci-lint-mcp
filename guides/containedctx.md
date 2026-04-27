# containedctx

<instructions>
Containedctx detects structs that contain `context.Context` fields. Storing a context in a struct is an anti-pattern because contexts should flow through function call stacks, not be stored as dependencies. It makes lifetimes opaque and breaks cancellation propagation.

Pass the context as the first argument to methods instead of storing it on the struct.
</instructions>

<examples>
## Good
```go
type Service struct {
    db *sql.DB
}

func (s *Service) Query(ctx context.Context) (*sql.Rows, error) {
    rows, err := s.db.QueryContext(ctx, "SELECT ...")
    if err != nil {
        return nil, fmt.Errorf("executing query: %w", err)
    }
    return rows, nil
}
```
</examples>

<patterns>
- Pass `context.Context` as the first function argument instead of storing it on a struct
- Pass `context.Context` as the first method argument rather than capturing it in a constructor
- Pass `context.Context` per operation in background workers rather than embedding it as a struct field
</patterns>

<related>
contextcheck, staticcheck/SA1013, gocritic
</related>
