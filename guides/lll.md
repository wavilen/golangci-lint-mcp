# lll

<instructions>
LLL (Line Length Linter) checks that lines don't exceed a maximum length (default 120 characters). Long lines reduce readability, especially in side-by-side diffs and split-screen editors.

Break long lines at natural boundaries: after operators, before function arguments, or by extracting long strings into constants or variables.
</instructions>

<examples>
## Bad
```go
func main() {
    http.HandleFunc("/api/v1/users", func(w http.ResponseWriter, r *http.Request) { fmt.Fprintf(w, `{"status":"ok","users":[{"id":1,"name":"Alice"},{"id":2,"name":"Bob"}]}`) })
}
```

## Good
```go
func main() {
    handler := func(w http.ResponseWriter, r *http.Request) {
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
- Long string literals that should be split or extracted into constants
- Chained method calls that should be formatted with line breaks between calls
- Function signatures with many parameters that need multi-line formatting
</patterns>

<related>
funlen, godoclint, revive
</related>
