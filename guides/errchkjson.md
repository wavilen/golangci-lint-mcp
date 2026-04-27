# errchkjson

<instructions>
Errchkjson reports types that cannot be safely marshaled to JSON. It detects non-JSON-serializable types passed to `json.Marshal` and `json.MarshalIndent`, such as channels, functions, or complex types without JSON tags.

Ensure all struct fields passed to JSON marshaling have serializable types and appropriate JSON tags.
</instructions>

<examples>
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
- Add `json:"-"` tags to channel or function fields, or avoid passing those structs to `json.Marshal`
- Add `json:"-"` tags to fields that cannot be serialized (channels, functions, sync types)
- Use string keys in maps passed to `json.Marshal`, or implement custom marshaling for non-string keys
- Ensure fields that must appear in JSON output are exported — unexported fields are silently omitted
</patterns>

<related>
errcheck, govet, musttag
</related>
