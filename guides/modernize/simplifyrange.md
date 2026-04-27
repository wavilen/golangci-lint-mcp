# modernize: simplifyrange

<instructions>
Detects `for i := range x` where only the value is needed but the index is unused, or `for i, _ := range x` where the blank identifier can be omitted. Go allows `for i := range x` to iterate over indices only, and `for _, v := range x` to iterate over values only. The checker flags redundant or confusing range forms.

Simplify range expressions by removing unused variables or blank identifiers to make intent clear.
</instructions>

<examples>
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
- Remove the blank identifier in `for i, _ := range slice` when only the index is used
- Use `for range` instead of `for _, _ := range slice` when neither value is used
- Use Go 1.22 range-over-int for `for i := range n` where n is an integer
</patterns>

<related>
modernize/loopvar, modernize/reloop, intrange
</related>
