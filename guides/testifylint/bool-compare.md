# testifylint: bool-compare

<instructions>
Detects `assert.True(t, a == b)` or `assert.False(t, a != b)` where a dedicated equality assertion is more readable and produces better failure messages. Use `assert.Equal` for equality and `assert.NotEqual` for inequality. The dedicated assertions show both values in the output, making test failures easier to diagnose.
</instructions>

<examples>
## Good
```go
assert.Equal(t, expected, result)
assert.Equal(t, expected, result)
```
</examples>

<patterns>
- Use `assert.Equal(t, y, x)` instead of `assert.True(t, x == y)`
- Use `assert.NotEqual(t, y, x)` instead of `assert.False(t, x == y)`
- Use `assert.NotEqual(t, y, x)` instead of `assert.True(t, x != y)`
</patterns>

<related>
testifylint/compares, testifylint/float-compare
</related>
