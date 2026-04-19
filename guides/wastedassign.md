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
- Variable reassigned in conditional branches where some paths waste the initial value
- Assignment in a loop body overwritten each iteration without being read
- Error variable reassigned without checking the previous value
</patterns>

<related>
ineffassign, unused, govet
