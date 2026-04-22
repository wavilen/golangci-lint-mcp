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
- Rename identifiers with accented characters to ASCII equivalents (e.g., résumé → resume)
- Replace Greek letter identifiers with descriptive ASCII names (e.g., π → pi)
- Rename CJK character identifiers to English equivalents
- Replace non-ASCII whitespace with standard spaces or underscores
</patterns>

<related>
godox, godot, goheader
</related>
