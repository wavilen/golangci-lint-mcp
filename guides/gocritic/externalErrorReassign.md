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
- `err = errors.New(...)` overwriting an external error
- `err = fmt.Errorf(...)` without wrapping — use `errors.Wrap` instead
- Reassigning error parameters passed into the function
- Returning a new error that discards the original cause
</patterns>

<related>
uncheckedInlineErr, nilValReturn
</related>
