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
- Replace `switch true { ... }` with `switch { ... }`
- Replace any constant-true expression in `switch` with a bare `switch`
</patterns>

<related>
singleCaseSwitch, ifElseChain
</related>
