# nakedret

<instructions>
Nakedret detects naked returns in functions longer than a configured limit. Naked returns (`return` with no value) rely on named return values and are easy to misunderstand, especially in longer functions.

Use explicit returns specifying the values being returned. Naked returns in short functions (under ~5 lines) are acceptable, but longer functions should always be explicit.
</instructions>

<examples>
## Bad
```go
func compute(a, b int) (sum, diff int) {
    sum = a + b
    diff = a - b
    if diff < 0 {
        diff = -diff
    }
    return
}
```

## Good
```go
func compute(a, b int) (sum, diff int) {
    sum = a + b
    diff = a - b
    if diff < 0 {
        diff = -diff
    }
    return sum, diff
}
```
</examples>

<patterns>
- Bare `return` at the end of multi-line functions
- Naked returns combined with defer that modifies named returns
- Early naked returns in conditional branches
- Named return values used only for naked returns, not for documentation
</patterns>

<related>
nonamedreturns, nlreturn, errname
</related>
