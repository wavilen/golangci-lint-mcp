# contextcheck

<instructions>
Contextcheck verifies that `context.Context` is properly propagated through function call chains. It flags places where a context is created but not passed down, or where a function accepting a context receives `context.Background()` instead of the caller's context.

Pass the parent context as the first argument to every function that needs it.
</instructions>

<examples>
## Good
```go
func (s *Service) Fetch(ctx context.Context) error {
    req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
    if err != nil {
        return err
    }
    resp, err := http.DefaultClient.Do(req)
    return err
}
```
</examples>

<patterns>
- Replace `http.NewRequest` with `http.NewRequestWithContext(ctx, ...)` to propagate context
- Use `context.WithCancel(parent)` in goroutines instead of discarding the parent context
- Pass `context.Context` to all helper functions that accept it
- Avoid `context.Background()` when the caller's context should propagate through the call chain
</patterns>

<related>
errcheck, govet, noctx, bodyclose, govet/lostcancel
</related>
