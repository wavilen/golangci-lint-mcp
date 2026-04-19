# errorlint: asserts

<instructions>
Detects type assertions on errors like `err.(*os.PathError)` or `switch err.(type)` that bypass the error wrapping chain. Wrapped errors will not match a direct type assertion. Use `errors.As(err, &target)` which walks the error chain via `Unwrap()` to find the matching type.
</instructions>

<examples>
## Bad
```go
pathErr, ok := err.(*os.PathError)
if ok {
    slog.Info("path error", "path", pathErr.Path)
}
```

## Good
```go
var pathErr *os.PathError
if errors.As(err, &pathErr) {
    slog.Info("path error", "path", pathErr.Path)
}
```
</examples>

<patterns>
- `err.(*SpecificType)` — use `errors.As(err, &target)`
- `switch err.(type)` on errors — use `errors.As` for each case
- `_, ok := err.(MyError)` — use `var e MyError; ok := errors.As(err, &e)`
</patterns>

<related>
errorf, comparison
