# tagliatelle

<instructions>
Tagliatelle enforces naming conventions in struct tags — ensuring `json`, `yaml`, `xml`, and other tagged field names follow a consistent casing style (camelCase, snake_case, PascalCase, etc.).

Configure the required case in `.golangci.yml` under `linters.settings.tagliatelle.case`. Rename struct tag values to match the configured convention.
</instructions>

<examples>
## Bad
```go
type Config struct {
    MaxRetries int `json:"max_retries"`
    Timeout    int `json:"Timeout"`
}
```

## Good
```go
type Config struct {
    MaxRetries int `json:"maxRetries"`
    Timeout    int `json:"timeout"`
}
```
</examples>

<patterns>
- Apply the configured casing convention consistently across all tag types
- Convert snake_case JSON tags to the configured casing (e.g., camelCase)
- Lowercase PascalCase tags to match protocol conventions
- Standardize all struct tags to use one consistent casing convention
</patterns>

<related>
tagalign, exhaustruct, inamedparam
</related>
