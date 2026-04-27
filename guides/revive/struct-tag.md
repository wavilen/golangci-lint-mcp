# revive: struct-tag

<instructions>
Checks that struct field tags follow conventions for common tag formats: `json`, `xml`, `yaml`, `db`, `protobuf`. It validates tag syntax, detects missing tags on exported fields, and flags inconsistencies like mixing `omitempty` usage. Well-formed tags ensure correct serialization behavior.

Ensure every exported field has appropriate tags, the tag key matches the expected format, and tag options are consistent across the struct.
</instructions>

<examples>
## Good
```go
type User struct {
    Name  string `json:"name" xml:"name"`
    Email string `json:"email" xml:"email"`
}
```
</examples>

<patterns>
- Ensure tag values are properly quoted within backtick-enclosed strings
- Remove struct tags from unexported fields — most encoders ignore them
- Use consistent field naming conventions across `json`, `xml`, `yaml`, and other tag formats
- Replace typos in tag keys (e.g., `jason` → `json`)
- Add `omitempty` to pointer or slice fields where it would reduce serialization noise
</patterns>

<related>
revive/nested-structs, revive/enforce-map-style, revive/enforce-slice-style
</related>
