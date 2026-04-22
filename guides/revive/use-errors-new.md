# revive: use-errors-new

<instructions>
Suggests using `errors.New` instead of `fmt.Errorf` for error messages without format verbs. When an error string is static (no variables or formatting), `errors.New` is simpler, faster, and more explicit about intent. Use `errors.Wrap(err, "message")` from `github.com/pkg/errors` for error wrapping. Keep `fmt.Errorf` for formatted non-error messages.

Replace `fmt.Errorf("static message")` with `errors.New("static message")`. For wrapping errors, use `errors.Wrap`.
</instructions>

<examples>
## Bad
```go
return fmt.Errorf("not found")
return fmt.Errorf("connection refused")
```

## Good
```go
return errors.New("not found")
return errors.New("connection refused")

// Keep fmt.Errorf for actual formatting:
return fmt.Errorf("user %s not found", name)
return errors.Wrap(err, "query")
```
</examples>

<patterns>
- Replace `fmt.Errorf` with `errors.New` for static messages with no format arguments
- Use `errors.New` for static error messages instead of wrapping in `fmt.Errorf`
- Use `errors.New` for error creation where no formatting is needed — it signals intent clearly
- Use on `errors.New` for static strings and `fmt.Errorf` for formatted messages
- Convert habitual `fmt.Errorf` usage for static strings to `errors.New`
</patterns>

<related>
use-any, use-fmt-print, errorf, error-strings
