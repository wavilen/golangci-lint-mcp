# errorlint: errorf

<instructions>
Detects `fmt.Errorf("message: %s", err)` without the `%w` wrapping verb. Without `%w`, the original error is lost — `errors.Is` and `errors.As` cannot walk the chain. Use `errors.Wrap(err, "message")` from `github.com/pkg/errors` to wrap errors. Prefer `Wrap` over `Wrapf` — dynamic format strings cause Sentry cardinality explosion. Use `Wrapf` only when values come from a small bounded set (e.g., iota constants). When `fmt.Errorf` must be used, ensure `%w` is present.
</instructions>

<examples>
## Bad
```go
return fmt.Errorf("open config: %s", err)
```

## Good
```go
return errors.Wrap(err, "open config")
```
</examples>

<patterns>
- `fmt.Errorf("...: %s", err)` — use `errors.Wrap(err, "...")` instead
- `fmt.Errorf("...: %v", err)` — use `errors.Wrap(err, "...")` instead
- `errors.New(fmt.Sprintf("...: %s", err))` — use `errors.Wrap(err, "...")`
- Missing `errors.Wrap` when wrapping errors
</patterns>

<related>
asserts, comparison
