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
- Structs passed to `json.Marshal`/`json.Unmarshal` without `json` tags
- Structs used with `yaml.Marshal` missing `yaml` tags
- Protobuf-generated structs without proper tags
- Third-party marshalers like `mapstructure` or `toml`
</patterns>

<related>
tagalign, tagliatelle
</related>
