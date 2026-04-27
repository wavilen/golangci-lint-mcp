# godoclint

<instructions>
Godoclint checks that exported functions, types, and variables have properly formatted doc comments. Missing or malformed comments hurt API documentation and IDE tooltips.

Add a doc comment starting with the declared name immediately before each exported declaration.
</instructions>

<examples>
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
- Add doc comments to all exported functions
- Start doc comments with the name of the function or type they describe
- Add doc comments to all exported types and constants
- Add a package-level doc comment to every package
</patterns>

<related>
godox, revive, errname
</related>
