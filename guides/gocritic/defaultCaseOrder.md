# gocritic: defaultCaseOrder

<instructions>
Detects `switch` and `select` statements where the `default` case is not the last case. Go convention places `default` at the bottom for readability — it signals the catch-all after all specific cases have been listed.

Move the `default` case to the end of the `switch` or `select` statement.
</instructions>

<examples>
## Bad
```go
switch v := x.(type) {
default:
	handleAny(v)
case int:
	handleInt(v)
case string:
	handleString(v)
}
```

## Good
```go
switch v := x.(type) {
case int:
	handleInt(v)
case string:
	handleString(v)
default:
	handleAny(v)
}
```
</examples>

<patterns>
- Move `default` to the last position in `switch` statements
- Move `default` to the end of `select` statements
- Reorder type switches — place `default` after all specific type cases
</patterns>

<related>
singleCaseSwitch, switchTrue, typeSwitchVar
</related>
