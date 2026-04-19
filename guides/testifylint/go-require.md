# testifylint: go-require

<instructions>
Detects `assert.NoError(t, err)` or `assert.True(t, ...)` at the start of a test where a failure would make subsequent code meaningless. Use `require.NoError(t, err)` to stop test execution immediately. `require` assertions call `t.FailNow()` on failure, preventing nil pointer dereferences and confusing cascading failures.
</instructions>

<examples>
## Bad
```go
result, err := Parse(input)
assert.NoError(t, err)       // test continues even if err != nil
assert.Equal(t, 5, result.ID) // panics or gives wrong error
```

## Good
```go
result, err := Parse(input)
require.NoError(t, err)      // stops here if err != nil
assert.Equal(t, 5, result.ID)
```
</examples>

<patterns>
- `assert.NoError` followed by using the result — use `require.NoError`
- `assert.True(t, setupOK)` at test start — use `require.True`
- Setup or initialization check that must pass — always use `require`
</patterns>

<related>
require-error, error-nil
