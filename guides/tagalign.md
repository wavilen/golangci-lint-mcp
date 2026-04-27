# tagalign

<instructions>
Tagalign aligns struct tags for readability. Misaligned struct tags make it harder to scan field definitions, especially in structs with many fields.

Align struct tag key-value pairs into columns so each tag type (`json`, `yaml`, `db`, `validate`) starts at the same column across fields.
</instructions>

<examples>
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
- Align struct tags with consistent spacing across fields
- Align each tag key (`json`, `yaml`, `db`) to the same column across fields
- Add missing struct tags to maintain consistent alignment
- Wrap or shorten long field names to keep tag columns aligned
</patterns>

<related>
tagliatelle, exhaustruct, embeddedstructfieldcheck
</related>
