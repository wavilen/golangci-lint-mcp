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
- Inconsistent casing between `json` and `yaml` tags on the same field
- snake_case in JSON tags when camelCase is configured
- PascalCase tags for protocols expecting lowercase
- Mixed conventions within a single struct definition
</patterns>

<related>
tagalign, exhaustruct, inamedparam
</related>
