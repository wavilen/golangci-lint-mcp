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
- `assert.Equal(t, n, len(x))` — use `assert.Len(t, x, n)`
- `assert.Equal(t, len(x), n)` — use `assert.Len(t, x, n)`
- `assert.NotEqual(t, 0, len(x))` — use `assert.NotEmpty(t, x)`
</patterns>

<related>
empty, nil-compare
