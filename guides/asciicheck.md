# asciicheck

<instructions>
Asciicheck detects non-ASCII identifiers in Go code. Identifiers containing Unicode characters reduce readability and can cause issues in some editors and toolchains.

Rename identifiers to use only ASCII letters, digits, and underscores. Keep Unicode in string literals and comments where it belongs.
</instructions>

<examples>
## Bad
```go
func calculateTotal(prix float64) float64 {
    réduction := 0.1
    return prix * (1 - réduction)
}
```

## Good
```go
func calculateTotal(price float64) float64 {
    discount := 0.1
    return price * (1 - discount)
}
```
</examples>

<patterns>
- Function or variable names with accented characters (résumé, café, naïve)
- Greek letters used as identifiers (π, Δ, α)
- CJK characters in variable or function names
- Non-ASCII whitespace in identifier boundaries
</patterns>

<related>
godox, godot, goheader
</related>
