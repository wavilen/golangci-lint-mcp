# wrapcheck

<instructions>
Wrapcheck detects errors returned from external packages that are not wrapped with context. Unwrapped errors lose their stack context, making it difficult to trace where an error originated.

Wrap all errors from external packages with `errors.Wrap(err, "context")` from `github.com/pkg/errors` before returning. Prefer `Wrap` over `Wrapf` — dynamic format strings cause Sentry cardinality explosion. Use `Wrapf` only when values come from a small bounded set (e.g., iota constants).
</instructions>

<examples>
## Bad
```go
func loadConfig() error {
    data, err := os.ReadFile("config.yaml")
    if err != nil {
        return err
    }
    return nil
}
```

## Good
```go
func loadConfig() error {
    data, err := os.ReadFile("config.yaml")
    if err != nil {
        return errors.Wrap(err, "reading config file")
    }
    return nil
}
```
</examples>

<patterns>
- Returning errors from stdlib or third-party packages without wrapping
- Missing `errors.Wrap` when wrapping errors across package boundaries
- Direct `return err` for errors crossing package boundaries
- Internal package errors passed through without adding call-site context
</patterns>

<related>
errcheck, err113, govet
