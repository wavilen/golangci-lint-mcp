# govet: structtag

<instructions>
Reports malformed struct tags. Common issues include duplicate tag keys, invalid JSON/YAML/db tag syntax, missing tag values, and keys that don't follow the `key:"value"` format. These cause silent failures at runtime when marshaling or unmarshaling.

Fix tag format to follow the convention: `json:"fieldname,omitempty"`.
</instructions>

<examples>
## Bad
```go
type User struct {
    Name string `json:"name" json:"username"` // duplicate key
    Email string `json:"email, string"`       // space after comma is invalid
}
```

## Good
```go
type User struct {
    Name  string `json:"name"`
    Email string `json:"email"`
}
```
</examples>

<patterns>
- Remove duplicate keys in struct tags — use one `json` or `yaml` key per field
- Remove spaces from struct tag option values — use `json:"name,omitempty"` not `json:"name, omitempty"`
- Provide non-empty values for all struct tag keys — never `json:""`
- Fix mismatched quotes in struct tag strings — use backtick-delimited valid key-value pairs
</patterns>

<related>
composites, stdmethods
</related>
