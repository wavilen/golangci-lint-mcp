# gocognit

<instructions>
Gocognit measures cognitive complexity — how hard code is to read, penalizing nesting, flow interruptions, and boolean operators. The primary fix is SRP decomposition: split functions that mix concerns into single-responsibility helpers so each piece is independently understandable.
</instructions>

<examples>
## Good
```go
func HandleCreateOrder(w http.ResponseWriter, r *http.Request) {
    req, err := validateOrder(r)
    if err != nil {
        writeResponse(w, 400, err)
        return
    }
    order, err := processOrder(req)
    if err != nil {
        writeResponse(w, 500, err)
        return
    }
    writeResponse(w, 201, order)
}
```
</examples>

<patterns>
- Decompose god handlers into separate functions for validation, logic, and formatting
- Extract nested responsibilities into single-purpose helper functions
- Separate service methods into distinct lookup, transform, and persist steps
- Replace switch with interface dispatch (OCP)
- Prefer small focused interfaces over one fat interface (ISP)
</patterns>

<related>
gocyclo, cyclop, maintidx, nestif, revive/cognitive-complexity
</related>
