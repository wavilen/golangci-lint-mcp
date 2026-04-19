# revive: superfluous-else

<instructions>
Detects `else` blocks that are unnecessary because the `if` block already returns, panics, or exits. When the `if` branch terminates execution, the code after it runs unconditionally — making the `else` keyword and its extra indentation redundant.

Remove the `else` keyword and unindent the code block. The control flow is clearer when the happy path is not nested.
</instructions>

<examples>
## Bad
```go
func get(name string) string {
    if name == "" {
        return "default"
    } else {
        return name
    }
}
```

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
- `if/else` where the `if` branch returns, panics, or calls `log.Fatal`
- `else if` chains where earlier branches all terminate
- Nested if-else where the else is a simple continuation after a return
- Error-check blocks followed by else with the main logic
- Guard clause patterns obscured by unnecessary else blocks
</patterns>

<related>
early-return, indent-error-flow, unnecessary-if
