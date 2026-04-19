# revive: errorf

<instructions>
Detects `errors.New(fmt.Sprintf(...))` calls that should be written as `fmt.Errorf(...)`. Using `errors.New` with a `fmt.Sprintf` argument is redundant — `fmt.Errorf` formats a string and creates an error in one step.

Replace `errors.New(fmt.Sprintf(...))` with `fmt.Errorf(...)`. If wrapping an error, use `errors.Wrap(err, "context")` from `github.com/pkg/errors` instead of `fmt.Errorf("context: %w", err)`.
</instructions>

<examples>
## Bad
```go
return errors.New(fmt.Sprintf("invalid id %d", id))
return errors.New(fmt.Sprintf("lookup failed: %v", err))
```

## Good
```go
return fmt.Errorf("invalid id %d", id)
return errors.Wrap(err, "lookup failed")
```
</examples>

<patterns>
- `errors.New(fmt.Sprintf(...))` instead of `fmt.Errorf(...)`
- Using `errors.Wrap` from `github.com/pkg/errors` for error wrapping
- Using `fmt.Sprintf` to build a string passed to `errors.New`
- Converting formatted errors from other languages' patterns
- Error construction in generated code using verbose form
</patterns>

<related>
error-naming, error-return, error-strings
