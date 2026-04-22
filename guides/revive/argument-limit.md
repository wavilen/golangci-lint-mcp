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
- Group related parameters into a struct when data-transfer or builder functions accumulate too many arguments
- Use the options pattern for constructor functions that set many optional fields
- Wrap API handler path/query parameters in a request struct to reduce argument count
- Separate functions where callers frequently mix up argument order into smaller focused functions
- Refactor parameter creep by periodically consolidating related arguments into structs
</patterns>

<related>
function-result-limit, function-length
