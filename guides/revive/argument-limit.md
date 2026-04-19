# revive: argument-limit

<instructions>
Flags functions that accept more than a configured number of parameters (default is typically 8). Functions with many arguments are hard to read, hard to call correctly, and suggest the function does too much.

Refactor by grouping related parameters into a struct, using the options pattern, or splitting the function into smaller focused ones.
</instructions>

<examples>
## Bad
```go
func CreateUser(name, email, phone, address, city, state, zip, country, role string, age int) error {
    // 10 parameters — easy to mix up order
    return nil
}
```

## Good
```go
type CreateUserParams struct {
    Name    string
    Email   string
    Phone   string
    Address Address
    Role    string
    Age     int
}

func CreateUser(p CreateUserParams) error {
    return nil
}
```
</examples>

<patterns>
- Data-transfer or builder functions accumulating parameters over time
- Constructor functions that set many optional fields
- API handler functions with many path/query parameters
- Functions where callers frequently mix up argument order
- Gradual parameter creep during development
</patterns>

<related>
function-result-limit, function-length
