# errchkjson

<instructions>
Errchkjson reports types that cannot be safely marshaled to JSON. It detects non-JSON-serializable types passed to `json.Marshal` and `json.MarshalIndent`, such as channels, functions, or complex types without JSON tags.

Ensure all struct fields passed to JSON marshaling have serializable types and appropriate JSON tags.
</instructions>

<examples>
## Bad
```go
type Config struct {
    Name    string
    Handler http.HandlerFunc
    Ch      chan int
}
data, _ := json.Marshal(Config{Name: "test"})
```

## Good
```go
type Config struct {
    Name string `json:"name"`
}
data, err := json.Marshal(Config{Name: "test"})
if err != nil {
    return err
}
```
</examples>

<patterns>
- Structs with channel or function fields passed to `json.Marshal`
- Missing `json:"-"` tags on non-serializable fields
- Maps with non-string keys used in JSON marshaling
- Unexported fields silently omitted from JSON output
</patterns>

<related>
errcheck, govet, musttag
