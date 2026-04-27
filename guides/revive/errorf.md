# revive: errorf

<instructions>
Detects `errors.New(fmt.Sprintf(...))` calls that should be written as `fmt.Errorf(...)`. Using `errors.New` with a `fmt.Sprintf` argument is redundant — `fmt.Errorf` formats a string and creates an error in one step.

Replace `errors.New(fmt.Sprintf(...))` with `fmt.Errorf(...)`. If wrapping an error, use `errors.Wrap(err, "context")` from `github.com/pkg/errors` instead of `fmt.Errorf("context: %w", err)`.
</instructions>

<examples>
## Good
```go
return fmt.Errorf("invalid id %d", id)
return errors.Wrap(err, "lookup failed")
```
</examples>

<patterns>
- Replace `errors.New(fmt.Sprintf(...))` with `fmt.Errorf(...)`
- Use `fmt.Errorf` with `%w` for error wrapping instead of `errors.Wrap` from `pkg/errors`
- Replace `errors.New(fmt.Sprintf(...))` with `fmt.Errorf(...)` for formatted error messages
- Use `fmt.Errorf` directly instead of building strings with `fmt.Sprintf` for error messages
- Simplify verbose error construction patterns from other languages to `fmt.Errorf`
</patterns>

<related>
revive/error-naming, revive/error-return, revive/error-strings
</related>
