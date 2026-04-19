# govet: errorsas

<instructions>
Reports `errors.As` calls where the second argument is not a pointer. `errors.As` requires a non-nil pointer target so it can assign the matched error value. Passing a value or nil target will never succeed.

Always pass a pointer to the target variable: `errors.As(err, &target)`.
</instructions>

<examples>
## Bad
```go
var target *MyError
if errors.As(err, target) { // target is nil pointer, not &target
    slog.Info("error target", "target", target)
}
```

## Good
```go
var target *MyError
if errors.As(err, &target) { // pointer to target variable
    slog.Info("error target", "target", target)
}
```
</examples>

<patterns>
- Passing a non-pointer to `errors.As` second argument
- Passing nil as `errors.As` target
- Passing value type instead of `&value` to `errors.As`
</patterns>

<related>
ifaceassert, nilfunc
</related>
