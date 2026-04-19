# gocognit

<instructions>
Gocognit measures cognitive complexity — how hard code is to read, penalizing nesting, flow interruptions, and boolean operators. The primary fix is SRP decomposition: split functions that mix concerns into single-responsibility helpers so each piece is independently understandable.
</instructions>

<examples>
## Bad
```go
func HandleCreateOrder(w http.ResponseWriter, r *http.Request) {
    var req OrderRequest
    if json.NewDecoder(r.Body).Decode(&req) != nil {
        w.WriteHeader(400)
        return
    }
    if req.Customer == "" {
        w.WriteHeader(400)
        return
    }
    total := 0.0
    for _, i := range req.Items {
        total += float64(i.Qty) * i.Price
    }
    order := Order{Customer: req.Customer, Total: total * 1.08}
    db.Save(&order)
    w.WriteHeader(201)
    json.NewEncoder(w).Encode(order)
}
```

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
- God handlers mixing input validation, business logic, and response formatting
- Functions where multiple responsibilities create deeply nested control flow
- Service methods combining lookup, transformation, and persistence
- Replace switch with interface dispatch (OCP)
- Prefer small focused interfaces over one fat interface (ISP)
</patterns>

<related>
gocyclo, cyclop, maintidx, nestif
</related>
