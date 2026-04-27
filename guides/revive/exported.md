# revive: exported

<instructions>
Enforces that all exported (public) functions, types, constants, and variables have doc comments. Doc comments should start with the name of the identifier they describe and be complete sentences. This rule replaces the original `golint` check for documentation.

Add a doc comment immediately above the exported declaration. Start the comment with the name of the identifier: `// Name does X.`
</instructions>

<examples>
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
- Add doc comments to all exported types
- Document exported functions with comments starting with the function name
- Use the identifier name at the start of doc comments — e.g., `// Name does X.`
- Add comments to exported constants and variables
- Add a package-level comment starting with "Package {name}" for every package
</patterns>

<related>
revive/comments-density, godoclint, revive/comment-spacings
</related>
