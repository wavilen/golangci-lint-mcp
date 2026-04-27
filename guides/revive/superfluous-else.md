# revive: superfluous-else

<instructions>
Detects `else` blocks that are unnecessary because the `if` block already returns, panics, or exits. When the `if` branch terminates execution, the code after it runs unconditionally — making the `else` keyword and its extra indentation redundant.

Remove the `else` keyword and unindent the code block. The control flow is clearer when the happy path is not nested.
</instructions>

<examples>
## Good
```go
func get(name string) string {
    if name == "" {
        return "default"
    }
    return name
}
```
</examples>

<patterns>
- Remove `else` after an `if` branch that returns, panics, or calls `log.Fatal`
- Remove indentation from continuation code after early-returning `else if` chains
- Remove `else` when the `if` block returns — the code below runs unconditionally
- Move main logic out of `else` blocks following error-check `if` statements
- Replace guard clause patterns obscured by `else` with direct continuation code
</patterns>

<related>
revive/early-return, revive/indent-error-flow, revive/unnecessary-if
</related>
