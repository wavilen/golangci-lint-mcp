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
- `switch x.(type)` followed by `x.(T)` inside case body
- Redundant type assertion after the switch already matched the type
- Using the original interface value instead of the switch-assigned variable
</patterns>

<related>
typeAssertChain, typeDefFirst
</related>
