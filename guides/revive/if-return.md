# revive: if-return

<instructions>
Detects redundant `if` statements that immediately return a boolean. Writing `if condition { return true } else { return false }` is unnecessary verbosity — you can return the condition directly.

Replace the if-else with a single `return condition`. For negated cases, use `return !condition`.
</instructions>

<examples>
## Good
```go
func isValid() bool {
    return len(items) > 0
}
```
</examples>

<patterns>
- Replace `if cond { return true } return false` with `return cond`
- Simplify `if cond { return false } else { return true }` to `return !cond`
- Return boolean expressions directly instead of wrapping them in an if-return
- Remove redundant guard checks before returning a boolean literal
- Replace switch statements returning `true`/`false` per case with a direct return expression
</patterns>

<related>
revive/early-return, revive/bool-literal-in-expr, revive/indent-error-flow
</related>
