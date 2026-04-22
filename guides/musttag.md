# musttag

<instructions>
Musttag enforces that struct fields passed to marshaling/unmarshaling functions have explicit struct tags. Without tags, field names default to the Go exported name, which may not match the expected JSON/YAML/XML schema.

Add the appropriate struct tag (`json`, `yaml`, `xml`, etc.) to every exported field in structs used with marshal/unmarshal functions.
</instructions>

<examples>
## Bad
```go
type User struct {
    Name  string
    Email string
}
data, _ := json.Marshal(User{Name: "alice"})
```

## Good
```go
type User struct {
    Name  string `json:"name"`
    Email string `json:"email"`
}
data, _ := json.Marshal(User{Name: "alice"})
```
</examples>

<patterns>
- Add `json` tags to structs passed to `json.Marshal`/`json.Unmarshal`
- Add `yaml` tags to structs used with `yaml.Marshal`/`yaml.Unmarshal`
- Ensure protobuf-generated structs carry proper struct tags
- Tag fields for third-party marshalers like `mapstructure` or `toml`
</patterns>

<related>
tagalign, tagliatelle
</related>
