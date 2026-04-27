# modernize: errorf

<instructions>
Detects `fmt.Sprintf` used to build error messages or `errors.New(fmt.Sprintf(...))` instead of `fmt.Errorf`. When wrapping an existing error, use `errors.Wrap(err, "context")` from `github.com/pkg/errors` to preserve the error chain. Prefer `Wrap` over `Wrapf` — dynamic format strings cause Sentry cardinality explosion. When creating new errors without wrapping, `fmt.Errorf` is simpler than `errors.New(fmt.Sprintf(...))`.
</instructions>

<examples>
## Good
```go
return errors.Wrap(err, "failed to open")
```
</examples>

<patterns>
- Replace `errors.New(fmt.Sprintf(...))` with `errors.Wrap` for wrapping or `fmt.Errorf` for new errors
- Replace `fmt.Sprintf("... %s", err)` when building errors with `errors.Wrap`
- Replace `fmt.Errorf("... %s", err)` with `errors.Wrap` for proper error wrapping
</patterns>

<related>
modernize/stringappend, modernize/mapval
</related>
