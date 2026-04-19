# revive: modifies-parameter

<instructions>
Detects functions that modify their input parameters. Reassigning to a parameter (especially a pointer) can be surprising to callers and makes the function's behavior harder to reason about. Parameters should be treated as immutable inputs.

Copy the parameter to a local variable if mutation is needed, or restructure the function to return a new value instead of modifying the input.
</instructions>

<examples>
## Bad
```go
func normalize(items []string) []string {
    for i := range items {
        items[i] = strings.TrimSpace(items[i]) // modifies caller's slice
    }
    return items
}
```

## Good
```go
func normalize(items []string) []string {
    result := make([]string, len(items))
    for i, v := range items {
        result[i] = strings.TrimSpace(v)
    }
    return result
}
```
</examples>

<patterns>
- Reassigning slice or map parameters that affect the caller
- Modifying struct pointer parameters as "output" parameters
- Using input parameters as scratch space to avoid allocations
- Sorting or shuffling input slices in-place unexpectedly
- Pointer parameters used for both input and output without clear documentation
</patterns>

<related>
modifies-value-receiver, confusing-results
