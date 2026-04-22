# gocritic: typeSwitchVar

<instructions>
Detects type switches where the value is accessed through a separate type assertion instead of using the variable assigned by the type switch itself. When using `switch v := x.(type)`, the variable `v` already holds the correctly typed value in each case.

Use the type-switch variable directly instead of asserting `x.(Type)` again inside the case body.
</instructions>

<examples>
## Bad
```go
switch x.(type) {
case int:
	v := x.(int)
	slog.Info("result", "value", v+1)
}
```

## Good
```go
switch v := x.(type) {
case int:
	slog.Info("result", "value", v+1)
}
```
</examples>

<patterns>
- Use `switch v := x.(type)` and work with `v` inside cases — avoid redundant `x.(T)`
- Use the switch-assigned variable instead of re-asserting on the original interface value
- Replace `switch x.(type)` followed by `x.(T)` with `switch v := x.(type)` and use `v`
</patterns>

<related>
typeAssertChain, typeDefFirst
</related>
