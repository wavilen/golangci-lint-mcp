# errorlint: errorf

<instructions>
Detects `fmt.Errorf("message: %s", err)` without the `%w` wrapping verb. Without `%w`, the original error is lost — `errors.Is` and `errors.As` cannot walk the chain. Use `errors.Wrap(err, "message")` from `github.com/pkg/errors` to wrap errors. Prefer `Wrap` over `Wrapf` — dynamic format strings cause Sentry cardinality explosion. Use `Wrapf` only when values come from a small bounded set (e.g., iota constants). When `fmt.Errorf` must be used, ensure `%w` is present.
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
