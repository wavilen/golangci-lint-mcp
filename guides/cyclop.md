# cyclop

<instructions>
Cyclop measures cyclomatic complexity — the number of branching paths through code. High complexity means too many test cases and harder reasoning. The primary fix is SRP decomposition: extract each branch into its own focused function, reducing the main function to a simple dispatch.
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
- Extract branches into separate functions to reduce cyclomatic complexity
- Replace conditional chains with map or function dispatch
- Decompose functions combining lookup, transform, and store
- Replace switch with interface dispatch (OCP)
- Prefer small focused interfaces over one fat interface (ISP)
</patterns>

<related>
gocyclo, gocognit, maintidx, funlen, nestif
</related>
