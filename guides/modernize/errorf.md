# modernize: errorf

<instructions>
Detects `fmt.Sprintf` used to build error messages or `errors.New(fmt.Sprintf(...))` instead of `fmt.Errorf`. When wrapping an existing error, use `errors.Wrap(err, "context")` from `github.com/pkg/errors` to preserve the error chain. Prefer `Wrap` over `Wrapf` — dynamic format strings cause Sentry cardinality explosion. When creating new errors without wrapping, `fmt.Errorf` is simpler than `errors.New(fmt.Sprintf(...))`.
</instructions>

<examples>
## Bad
```go
return errors.New(fmt.Sprintf("failed to open %s: %s", path, err))
```

## Good
```go
return errors.Wrap(err, "failed to open")
```
</examples>

<patterns>
- `errors.New(fmt.Sprintf(...))` — use `errors.Wrap` for wrapping or `fmt.Errorf` for new errors
- `fmt.Sprintf("... %s", err)` when building errors — use `errors.Wrap` instead
- `fmt.Errorf("... %s", err)` — use `errors.Wrap` for proper error wrapping
</patterns>

<related>
stringappend, mapval
