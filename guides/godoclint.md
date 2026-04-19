# godoclint

<instructions>
Godoclint checks that exported functions, types, and variables have properly formatted doc comments. Missing or malformed comments hurt API documentation and IDE tooltips.

Add a doc comment starting with the declared name immediately before each exported declaration.
</instructions>

<examples>
## Bad
```go
// This function parses config
func ParseConfig(path string) (*Config, error) {
    return nil, nil
}
```

## Good
```go
// ParseConfig reads and parses the configuration file at path.
// Returns an error if the file is missing or contains invalid YAML.
func ParseConfig(path string) (*Config, error) {
    return nil, nil
}
```
</examples>

<patterns>
- Exported functions without any doc comment
- Doc comments that don't start with the function or type name
- Exported types and constants missing documentation
- Packages without a package-level doc comment
</patterns>

<related>
godox, revive, errname
</related>
