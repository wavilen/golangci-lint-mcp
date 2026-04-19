# tagalign

<instructions>
Tagalign aligns struct tags for readability. Misaligned struct tags make it harder to scan field definitions, especially in structs with many fields.

Align struct tag key-value pairs into columns so each tag type (`json`, `yaml`, `db`, `validate`) starts at the same column across fields.
</instructions>

<examples>
## Bad
```go
type User struct {
    Name string `json:"name" validate:"required"`
    Age  int    `json:"age"`
    Email string `json:"email" validate:"required,email"`
}
```

## Good
```go
type User struct {
    Name  string `json:"name"  validate:"required"`
    Age   int    `json:"age"`
    Email string `json:"email" validate:"required,email"`
}
```
</examples>

<patterns>
- Struct tags with inconsistent spacing between fields
- Multiple tag keys not aligned across struct fields
- Missing tags on some fields causing alignment gaps
- Long field names pushing tags far right with no alignment
</patterns>

<related>
tagliatelle, exhaustruct, embeddedstructfieldcheck
</related>
