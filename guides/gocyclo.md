# gocyclo

<instructions>
Gocyclo computes cyclomatic complexity — the number of independent paths through the code. High complexity means more test cases and harder reasoning. Default threshold is 10.

Simplify by extracting branches into helpers, using early returns, and replacing conditionals with table-driven logic.
</instructions>

<examples>
## Good
```go
var statusRanges = []struct {
    min, max int
    label    string
}{
    {200, 300, "success"},
    {300, 400, "redirect"},
    {400, 500, "client error"},
    {500, 600, "server error"},
}

func classify(code int) string {
    for _, r := range statusRanges {
        if code >= r.min && code < r.max {
            return r.label
        }
    }
    return "unknown"
}
```
</examples>

<patterns>
- Replace sequential if/else-if chains with table-driven dispatch
- Convert long conditional chains into maps or lookup tables
- Simplify routing logic with map-based dispatch
</patterns>

<related>
cyclop, gocognit, maintidx, funlen, revive/cyclomatic
</related>
