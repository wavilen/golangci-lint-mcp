# govet: errorsas

<instructions>
Reports `errors.As` calls where the second argument is not a pointer. `errors.As` requires a non-nil pointer target so it can assign the matched error value. Passing a value or nil target will never succeed.

Always pass a pointer to the target variable: `errors.As(err, &target)`.
</instructions>

<examples>
## Good
```go
var target *MyError
if errors.As(err, &target) { // pointer to target variable
    slog.Info("error target", "target", target)
}
```
</examples>

<patterns>
- Pass a pointer as the second argument to `errors.As` — never a value type
- Pass a non-nil pointer as the `errors.As` target
- Use `&value` (not bare `value`) for the `errors.As` target argument
</patterns>

<related>
govet/ifaceassert, govet/nilfunc, errorlint/asserts
</related>
