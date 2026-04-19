# gocritic: switchTrue

<instructions>
Detects `switch true` statements that should be written as `switch { ... }` (omitting the `true`). In Go, a `switch` without an expression is equivalent to `switch true`, making the explicit `true` redundant.

Remove the `true` from `switch true` — use bare `switch` instead.
</instructions>

<examples>
## Bad
```go
switch true {
case x > 0:
	positive()
case x < 0:
	negative()
default:
	zero()
}
```

## Good
```go
switch {
case x > 0:
	positive()
case x < 0:
	negative()
default:
	zero()
}
```
</examples>

<patterns>
- `switch true { ... }` → `switch { ... }`
- Any expression in switch that evaluates to a constant `true`
</patterns>

<related>
singleCaseSwitch, ifElseChain
</related>
