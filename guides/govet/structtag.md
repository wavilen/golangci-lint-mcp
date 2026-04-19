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
- Duplicate keys in struct tags (`json:"x" json:"y"`)
- Invalid option syntax (spaces in options)
- Missing value in tag (`json:""`)
- Mismatched quotes in tag string
</patterns>

<related>
composites, stdmethods
</related>
