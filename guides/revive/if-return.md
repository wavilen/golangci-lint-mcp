# revive: if-return

<instructions>
Detects redundant `if` statements that immediately return a boolean. Writing `if condition { return true } else { return false }` is unnecessary verbosity — you can return the condition directly.

Replace the if-else with a single `return condition`. For negated cases, use `return !condition`.
</instructions>

<examples>
## Bad
```go
func isValid() bool {
    if len(items) > 0 {
        return true
    } else {
        return false
    }
}
```

## Good
```go
func isValid() bool {
    return len(items) > 0
}
```
</examples>

<patterns>
- `if cond { return true } return false` patterns
- `if cond { return false } else { return true }` (should be `return !cond`)
- Boolean comparison functions wrapping a single expression
- Redundant guard checks before returning a boolean literal
- Switch statements with `return true`/`return false` in every case
</patterns>

<related>
early-return, bool-literal-in-expr, indent-error-flow
