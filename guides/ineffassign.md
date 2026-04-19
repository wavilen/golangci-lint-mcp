# ineffassign

<instructions>
Ineffassign detects assignments that are never used — the assigned value is overwritten or goes out of scope before being read. These are often bugs where a calculation result is silently discarded.

Remove the unused assignment or use the value before it is overwritten.
</instructions>

<examples>
## Bad
```go
func calc(a, b int) int {
    result := a + b
    result = a * b // previous assignment never used
    return result
}
```

## Good
```go
func calc(a, b int) int {
    result := a * b
    return result
}
```
</examples>

<patterns>
- Variable reassigned before the previous value is read
- Error assigned but never checked: `x, err = f(); x, err = g()`
- Loop variables overwritten each iteration without using the previous value
- Assignment in `if` init statement where the value is unused in body
</patterns>

<related>
unused, wastedassign, govet
