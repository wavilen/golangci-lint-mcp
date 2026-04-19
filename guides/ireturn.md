# ireturn

<instructions>
Ireturn rejects functions that return interfaces instead of concrete types. Returning interfaces forces the caller to depend on an abstraction they didn't choose, limiting flexibility and making the API harder to mock in tests.

Return concrete types and let the caller decide whether to accept them as interfaces.
</instructions>

<examples>
## Bad
```go
func NewStore(path string) (io.Closer, error) {
    f, err := os.Open(path)
    if err != nil {
        return nil, err
    }
    return f, nil
}
```

## Good
```go
func NewStore(path string) (*os.File, error) {
    f, err := os.Open(path)
    if err != nil {
        return nil, err
    }
    return f, nil
}
```
</examples>

<patterns>
- Factory functions returning interface types instead of concrete implementations
- Constructor-style functions that unnecessarily hide the concrete type
- Functions returning `error` interface wrapped in another interface
- Providers that return `io.Reader` when a concrete `*bytes.Buffer` would suffice
</patterns>

<related>
interfacebloat, revive, gocritic
</related>
