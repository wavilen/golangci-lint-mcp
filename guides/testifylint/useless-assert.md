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
- `assert.True(t, true)` — remove or replace with actual boolean expression
- `assert.Equal(t, x, x)` — compare against an expected constant, not itself
- `assert.False(t, false)` — use a real condition
</patterns>

<related>
blank-import, expected-actual
