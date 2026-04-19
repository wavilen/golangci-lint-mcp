# modernize: simplifyrange

<instructions>
Detects `for i := range x` where only the value is needed but the index is unused, or `for i, _ := range x` where the blank identifier can be omitted. Go allows `for i := range x` to iterate over indices only, and `for _, v := range x` to iterate over values only. The checker flags redundant or confusing range forms.

Simplify range expressions by removing unused variables or blank identifiers to make intent clear.
</instructions>

<examples>
## Bad
```go
for i, _ := range items {
    slog.Info("index", "i", i)
}

for _, _ = range items {
    count++
}
```

## Good
```go
for i := range items {
    slog.Info("index", "i", i)
}

for range items {
    count++
}
```
</examples>

<patterns>
- `for i, _ := range slice` where only the index is used — drop the blank
- `for _, _ := range slice` where neither value is used — use `for range`
- `for i := range n` where n is an integer — Go 1.22 range-over-int
</patterns>

<related>
loopvar, reloop, intrange
