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
- Rename single-letter parameters and return values to descriptive names
- Expand short variable names to descriptive names for variables with long lifetimes
- Replace abbreviations like `c`, `r`, `p` with full descriptive names (e.g., `cfg`, `req`, `proc`)
</patterns>

<related>
godoclint, revive, gocritic
</related>
