# revive: exported

<instructions>
Enforces that all exported (public) functions, types, constants, and variables have doc comments. Doc comments should start with the name of the identifier they describe and be complete sentences. This rule replaces the original `golint` check for documentation.

Add a doc comment immediately above the exported declaration. Start the comment with the name of the identifier: `// Name does X.`
</instructions>

<examples>
## Bad
```go
type Config struct {
    Timeout time.Duration
}

func NewClient(c Config) *Client {
    return &Client{cfg: c}
}
```

## Good
```go
// Config holds settings for creating a new Client.
type Config struct {
    // Timeout is the maximum duration for a request.
    Timeout time.Duration
}

// NewClient creates and returns a Client with the given configuration.
func NewClient(c Config) *Client {
    return &Client{cfg: c}
}
```
</examples>

<patterns>
- Exported types without doc comments
- Exported functions missing documentation
- Doc comments not starting with the identifier name
- Exported constants or variables without comments
- Package-level comments missing or malformed
</patterns>

<related>
comments-density, godoclint, comment-spacings
