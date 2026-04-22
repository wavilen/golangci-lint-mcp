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
- Replace bare `return` with explicit return values in short helper functions using named returns
- Use explicit return values consistently when functions mix bare and explicit returns
- Return explicit named values in deferred functions that modify named return values
- Replace bare returns with explicit returns when refactoring adds named return values
</patterns>

<related>
early-return, if-return, nakedret
