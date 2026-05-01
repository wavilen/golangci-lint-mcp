# lll

<instructions>
Flags lines exceeding a configurable length limit (default 120 characters). Long lines break side-by-side diffs, get truncated by code review tools, and force horizontal scrolling that disrupts reading flow. Break lines at natural boundaries — after operators, before function arguments, or by extracting long strings into constants.
</instructions>

<examples>
## Good
```go
func main() {
    handler := func(w http.ResponseWriter, _ *http.Request) {
        response := `{"status":"ok","users":[` +
            `{"id":1,"name":"Alice"},` +
            `{"id":2,"name":"Bob"}]}`
        fmt.Fprint(w, response)
    }
    http.HandleFunc("/api/v1/users", handler)
}
```
</examples>

<patterns>
- Break long string literals with `+` concatenation at natural boundaries, or extract into `const` declarations
- Break chained `Builder` method calls (e.g., `strings.Builder`, `bytes.Buffer`) with a newline after each `.Method()` call
- Format function signatures with many parameters across multiple lines
</patterns>

<related>
funlen, godoclint, revive/line-length-limit
</related>
