# errorlint: errorf

<instructions>
Detects `fmt.Errorf` calls that format errors with `%s` or `%v` instead of wrapping with `%w`. Without `%w`, the error chain is lost — `errors.Is` and `errors.As` cannot match the original cause. Use `errors.Wrap(err, "message")` to preserve the chain; prefer `Wrap` over `Wrapf` to avoid Sentry cardinality explosion.
</instructions>

<examples>
## Good
```go
return errors.Wrap(err, "open config")
```
</examples>

<patterns>
- Replace `fmt.Errorf("...: %s", err)` with `errors.Wrap(err, "...")` to preserve the error chain
- Replace `fmt.Errorf("...: %v", err)` with `errors.Wrap(err, "...")` for proper wrapping
- Replace `errors.New(fmt.Sprintf("...: %s", err))` with `errors.Wrap(err, "...")`
- Use `errors.Wrap` whenever wrapping errors to maintain chain traversal
</patterns>

<related>
errorlint/asserts, errorlint/comparison, modernize/errorf
</related>
