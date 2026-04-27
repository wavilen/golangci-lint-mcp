# errorlint: asserts

<instructions>
Detects type assertions on errors like `err.(*os.PathError)` or `switch err.(type)` that bypass the error wrapping chain. Wrapped errors will not match a direct type assertion. Use `errors.As(err, &target)` which walks the error chain via `Unwrap()` to find the matching type.
</instructions>

<examples>
## Good
```go
var pathErr *os.PathError
if errors.As(err, &pathErr) {
    slog.Info("path error", "path", pathErr.Path)
}
```
</examples>

<patterns>
- Use `errors.As(err, &target)` instead of direct type assertions like `err.(*SpecificType)`
- Replace `switch err.(type)` on errors with `errors.As` checks for each case
- Use `var e MyError; ok := errors.As(err, &e)` instead of `_, ok := err.(MyError)`
</patterns>

<related>
errorlint/errorf, errorlint/comparison, govet/errorsas
</related>
