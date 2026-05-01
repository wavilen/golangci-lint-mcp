# cyclop

<instructions>
Measures cyclomatic complexity — the count of independent branching paths through a function. Each additional branch multiplies the test matrix (N+1 test cases per path), making behavior hard to reason about and changes risky to validate. Flatten branching by extracting paths into standalone functions or replacing conditionals with table-driven dispatch.
</instructions>

<examples>
## Good
```go
var handlers = map[string]func(Order) error{
    "created":  handleCreated,
    "approved": handleApproved,
    "shipped":  handleShipped,
}

func Process(status string, o Order) error {
    h, ok := handlers[status]
    if !ok {
        return fmt.Errorf("unknown: %s", status)
    }
    return h(o)
}
```
</examples>

<patterns>
- Extract each branching path into its own function to reduce the caller to a flat dispatch
- Replace conditional chains with `map[string]func` lookup tables for O(1) dispatch
- Decompose functions combining lookup, transform, and store into three separate functions
- Replace `switch` statements on type with interface method dispatch — one method per case, no branches
- Replace nested `if/else` guards with early returns to eliminate indent-based paths
</patterns>

<related>
gocyclo, gocognit, maintidx, funlen, nestif
</related>
