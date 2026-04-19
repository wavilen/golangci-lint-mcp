# varnamelen

<instructions>
Varnamelen checks that variable names are meaningful — rejecting single-letter or very short names that obscure intent. Short names make code harder to understand, especially for readers unfamiliar with the context.

Use descriptive names for variables with large or non-obvious scope. Short names are acceptable for small-scope variables like loop indices.
</instructions>

<examples>
## Bad
```go
func process(r io.Reader) ([]byte, error) {
    b, e := io.ReadAll(r)
    if e != nil {
        return nil, e
    }
    return b, nil
}
```

## Good
```go
func process(reader io.Reader) ([]byte, error) {
    data, err := io.ReadAll(reader)
    if err != nil {
        return nil, errors.Wrap(err, "reading input")
    }
    return data, nil
}
```
</examples>

<patterns>
- Single-letter variable names for function parameters or return values
- Short names for variables with long lifetimes (function scope or wider)
- Abbreviations like `c`, `r`, `p` for configuration, request, or processor objects
</patterns>

<related>
godoclint, revive, gocritic
</related>
