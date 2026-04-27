# wrapcheck

<instructions>
Wrapcheck detects errors returned from external packages that are not wrapped with context. Unwrapped errors lose their stack context, making it difficult to trace where an error originated.

Wrap all errors from external packages with `errors.Wrap(err, "context")` from `github.com/pkg/errors` before returning. Prefer `Wrap` over `Wrapf` — dynamic format strings cause Sentry cardinality explosion. Use `Wrapf` only when values come from a small bounded set (e.g., iota constants).
</instructions>

<examples>
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
- Wrap errors from stdlib or third-party packages with `errors.Wrap`
- Add context when passing errors across package boundaries using `errors.Wrap`
- Wrap errors with call-site context before returning them across package boundaries
- Add call-site context to internal package errors before passing them to external callers
</patterns>

<related>
errcheck, err113, govet
</related>
