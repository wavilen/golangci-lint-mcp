# revive: bare-return

<instructions>
Detects bare `return` statements in functions with named return values. Bare returns save keystrokes but harm readability — the reader must look at the function signature to know what is being returned. Explicit returns make data flow obvious.

Replace bare `return` with explicit `return` listing the named return values.
</instructions>

<examples>
## Bad
```go
func divmod(a, b int) (quotient, remainder int) {
    quotient = a / b
    remainder = a % b
    return // what is returned?
}
```

## Good
```go
func divmod(a, b int) (quotient, remainder int) {
    quotient = a / b
    remainder = a % b
    return quotient, remainder
}
```
</examples>

<patterns>
- Short helper functions using named returns with bare return
- Functions with multiple return paths mixing bare and explicit returns
- Deferred functions modifying named return values with bare return
- Refactored functions that gained named returns but kept bare returns
</patterns>

<related>
early-return, if-return, nakedret
