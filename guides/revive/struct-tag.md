# revive: struct-tag

<instructions>
Checks that struct field tags follow conventions for common tag formats: `json`, `xml`, `yaml`, `db`, `protobuf`. It validates tag syntax, detects missing tags on exported fields, and flags inconsistencies like mixing `omitempty` usage. Well-formed tags ensure correct serialization behavior.

Ensure every exported field has appropriate tags, the tag key matches the expected format, and tag options are consistent across the struct.
</instructions>

<examples>
## Bad
```go
type User struct {
    Name  string `json:name`        // missing quotes
    Email string `json:"email" xml:"email"` // inconsistent naming
    age   int    `json:"age"`       // unexported field with tag (ignored)
}
```

## Good
```go
type User struct {
    Name  string `json:"name" xml:"name"`
    Email string `json:"email" xml:"email"`
}
```
</examples>

<patterns>
- Missing quotes around tag values
- Unexported fields with struct tags (tags are ignored by most encoders)
- Inconsistent field naming conventions between different tag formats
- Typos in tag keys (e.g., `jason` instead of `json`)
- Missing `omitempty` on pointer or slice fields where it would reduce output noise
</patterns>

<related>
nested-structs, enforce-map-style, enforce-slice-style
