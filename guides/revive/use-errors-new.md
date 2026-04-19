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
- `fmt.Errorf` called with a plain string and no arguments
- Static error messages wrapped in `fmt.Errorf` unnecessarily
- Error creation where `errors.New` is clearer about intent
- Consistent use of `fmt.Errorf` across a codebase even for static strings
- Migration from older code where `fmt.Errorf` was used habitually
</patterns>

<related>
use-any, use-fmt-print, errorf, error-strings
