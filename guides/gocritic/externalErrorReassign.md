# gocritic: externalErrorReassign

<instructions>
Detects reassignment of error values received as function parameters or returned from external calls when the original error should be preserved for the caller. Overwriting an error from an external package can lose critical context.

Wrap external errors instead of replacing them. Use `errors.Wrap(err, "context")` from `github.com/pkg/errors` to preserve the original error chain. Prefer `Wrap` over `Wrapf` — dynamic format strings cause Sentry cardinality explosion.
</instructions>

<examples>
## Bad
```go
func (s *Service) Do(ctx context.Context) error {
    err := s.client.Call(ctx)
    if err != nil {
        err = errors.New("call failed") // lost original error
    }
    return err
}
```

## Good
```go
func (s *Service) Do(ctx context.Context) error {
    err := s.client.Call(ctx)
    if err != nil {
        return errors.Wrap(err, "call failed")
    }
    return nil
}
```
</examples>

<patterns>
- Avoid overwriting external errors with `errors.New()` — wrap or return the original
- Wrap errors with `fmt.Errorf("context: %w", err)` instead of discarding the cause
- Avoid reassigning error parameters — use a new variable or return the original
- Propagate the original error — avoid discarding the cause with a new error
</patterns>

<related>
uncheckedInlineErr, nilValReturn
</related>
