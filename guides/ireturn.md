# ireturn

<instructions>
Ireturn rejects functions that return interfaces instead of concrete types. Returning interfaces forces the caller to depend on an abstraction they didn't choose, limiting flexibility and making the API harder to mock in tests.

Return concrete types and let the caller decide whether to accept them as interfaces.
</instructions>

<examples>
## Good
```go
func NewStore(path string) (*os.File, error) {
    f, err := os.Open(path)
    if err != nil {
        return nil, fmt.Errorf("opening store: %w", err)
    }
    return f, nil
}
```
</examples>

<patterns>
- Return concrete types from factory functions and let callers abstract as needed
- Declare the concrete return type in constructors instead of hiding it behind an interface
- Return `error` directly rather than wrapping it in another interface
- Return `*bytes.Buffer` instead of `io.Reader` when the concrete type suffices
</patterns>

<related>
interfacebloat, revive/max-public-structs, gocritic/unnamedResult
</related>
