# revive: nested-structs

<instructions>
Flags struct types defined inline within another struct field. Nested struct definitions reduce readability and make it harder to reuse or test the inner type. Extract the nested struct into a named type at the package level.

Move the anonymous struct definition out of the enclosing struct and give it a descriptive name. Reference the new type in the field declaration.
</instructions>

<examples>
## Bad
```go
type Server struct {
    Config struct {
        Host string
        Port int
    }
    Logger *log.Logger
}
```

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
- Anonymous struct types embedded directly as fields
- Nested struct definitions for config or options objects
- Inline struct definitions in API request/response types
- One-off anonymous structs used instead of reusable named types
- Deeply nested struct literals that are hard to construct
</patterns>

<related>
struct-tag, optimize-operands-order
