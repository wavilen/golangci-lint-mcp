# wastedassign

<instructions>
Wastedassign detects assignments that are wasted — the value is assigned but never read before being reassigned or going out of scope. Unlike ineffassign, wastedassign specifically tracks assignment flow and finds values that are written but never consumed.

Use the assigned value, return it, or remove the assignment entirely.
</instructions>

<examples>
## Bad
```go
func compute(n int) int {
    result := n * 2
    if n > 10 {
        result = n * 3 // first assignment wasted
    }
    return result
}
```

## Good
```go
func compute(n int) int {
    if n > 10 {
        return n * 3
    }
    return n * 2
}
```
</examples>

<patterns>
- Restructure conditional branches to use the assigned value or assign only when needed
- Move loop-body assignments to after the loop or consume the value within the iteration
- Check or handle each error assignment before reassigning the error variable
</patterns>

<related>
ineffassign, unused, govet
