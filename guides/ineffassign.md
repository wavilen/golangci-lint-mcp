# ineffassign

<instructions>
Ineffassign detects assignments that are never used — the assigned value is overwritten or goes out of scope before being read. These are often bugs where a calculation result is silently discarded.

Remove the unused assignment or use the value before it is overwritten.
</instructions>

<examples>
## Good
```go
func calc(a, b int) int {
    result := a * b
    return result
}
```
</examples>

<patterns>
- Remove variable reassignments where the previous value is never read before being overwritten
- Check error assignments that are silently discarded: `x, err = f(); x, err = g()`
- Eliminate loop variables overwritten each iteration without using the previous value
- Remove assignments in `if` init statements where the value is unused in the body
</patterns>

<related>
unused, wastedassign, govet, prealloc
</related>
