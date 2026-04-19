# gocritic: sloppyLen

<instructions>
Detects unnecessary `len()` comparisons when the intent is just to check for emptiness. Using `len(s) == 0` or `len(s) > 0` works, but the idiomatic Go approach is to rely on the zero value directly: `s == ""` for strings, `s == nil` or direct truthiness for slices and maps.

Use the idiomatic emptiness check: compare to empty string, use `len()` only when the actual length matters.
</instructions>

<examples>
## Bad
```go
if len(name) == 0 {
    return errors.New("name required")
}
```

## Good
```go
if name == "" {
    return errors.New("name required")
}
```
</examples>

<patterns>
- `len(s) == 0` instead of `s == ""` for strings
- `len(s) > 0` instead of `s != ""` for strings
- `len(m) == 0` for maps when checking nil
- `len(slice) != 0` when a nil check suffices
</patterns>

<related>
sloppyReassign, sloppyTypeAssert, offBy1
</related>
