# revive: nested-structs

<instructions>
Flags struct types defined inline within another struct field. Nested struct definitions reduce readability and make it harder to reuse or test the inner type. Extract the nested struct into a named type at the package level.

Move the anonymous struct definition out of the enclosing struct and give it a descriptive name. Reference the new type in the field declaration.
</instructions>

<examples>
## Good
```go
type ServerConfig struct {
    Host string
    Port int
}

type Server struct {
    Config ServerConfig
    Logger *log.Logger
}
```
</examples>

<patterns>
- Extract anonymous struct types embedded as fields into named top-level types
- Move nested config or options struct definitions to package-level named types
- Define named types for inline struct definitions in API request/response types
- Replace one-off anonymous structs with reusable named types
- Flatten deeply nested struct literals into composed named types
</patterns>

<related>
revive/struct-tag, revive/optimize-operands-order
</related>
