# revive: modifies-parameter

<instructions>
Detects functions that modify their input parameters. Reassigning to a parameter (especially a pointer) can be surprising to callers and makes the function's behavior harder to reason about. Parameters should be treated as immutable inputs.

Copy the parameter to a local variable if mutation is needed, or restructure the function to return a new value instead of modifying the input.
</instructions>

<examples>
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
- Use a local copy of slice or map parameters before modifying to avoid affecting the caller
- Return new values instead of modifying struct pointer parameters as "output" parameters
- Use a local variable for mutation instead of using input parameters as scratch space
- Use a copy of input slices before sorting or shuffling to avoid unexpected in-place mutation
- Document pointer parameters used for both input and output, or restructure to use separate types
</patterns>

<related>
revive/modifies-value-receiver, revive/confusing-results
</related>
