# testifylint: empty

<instructions>
Detects `assert.Len(t, x, 0)` or `assert.Equal(t, 0, len(x))` used to check if a collection is empty. Use `assert.Empty(t, x)` or `assert.NotEmpty(t, x)` for semantic clarity. The `Empty` assertion works with slices, maps, channels, strings, and any type with a `Len()` method.
</instructions>

<examples>
## Good
```go
assert.Empty(t, results)
assert.Empty(t, items)
```
</examples>

<patterns>
- Use `assert.Empty(t, x)` instead of `assert.Len(t, x, 0)`
- Use `assert.Empty(t, x)` instead of `assert.Equal(t, 0, len(x))`
- Use `assert.NotEmpty(t, x)` instead of `assert.NotEqual(t, 0, len(x))`
</patterns>

<related>
testifylint/len, testifylint/nil-compare
</related>
