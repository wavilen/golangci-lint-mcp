# testifylint: useless-assert

<instructions>
Detects assertions that always pass or always fail: `assert.True(t, true)`, `assert.Equal(t, 1, 1)`, `assert.False(t, false)`. These are no-ops that provide no test value and likely indicate a typo, copy-paste error, or incomplete test. Remove the assertion or replace it with a meaningful check.
</instructions>

<examples>
## Bad
```go
assert.True(t, true)
assert.Equal(t, 42, 42)
assert.False(t, false)
```

## Good
```go
assert.True(t, isValid)
assert.Equal(t, expected, actual)
assert.False(t, hasError)
```
</examples>

<patterns>
- Remove `assert.True(t, true)` or replace with an actual boolean expression
- Replace `assert.Equal(t, x, x)` with a comparison against an expected constant
- Use a real condition instead of `assert.False(t, false)`
</patterns>

<related>
blank-import, expected-actual
