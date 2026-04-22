# testifylint: len

<instructions>
Detects `assert.Equal(t, 5, len(x))` or `assert.NotEqual(t, 0, len(x))` where `assert.Len` or `assert.NotEmpty` should be used instead. The `Len` assertion provides a clearer failure message showing both the expected and actual lengths.
</instructions>

<examples>
## Bad
```go
assert.Equal(t, 3, len(items))
assert.NotEqual(t, 0, len(results))
```

## Good
```go
assert.Len(t, items, 3)
assert.NotEmpty(t, results)
```
</examples>

<patterns>
- Use `assert.Len(t, x, n)` instead of `assert.Equal(t, n, len(x))`
- Use `assert.Len(t, x, n)` instead of `assert.Equal(t, len(x), n)`
- Use `assert.NotEmpty(t, x)` instead of `assert.NotEqual(t, 0, len(x))`
</patterns>

<related>
empty, nil-compare
