# testifylint: go-require

<instructions>
Detects `assert.NoError(t, err)` or `assert.True(t, ...)` at the start of a test where a failure would make subsequent code meaningless. Use `require.NoError(t, err)` to stop test execution immediately. `require` assertions call `t.FailNow()` on failure, preventing nil pointer dereferences and confusing cascading failures.
</instructions>

<examples>
## Good
```go
result, err := Parse(input)
require.NoError(t, err)      // stops here if err != nil
assert.Equal(t, 5, result.ID)
```
</examples>

<patterns>
- Use `require.NoError` when `assert.NoError` is followed by using the result
- Use `require.True` instead of `assert.True(t, setupOK)` at test start
- Use `require` for setup or initialization checks that must pass before continuing
</patterns>

<related>
testifylint/require-error, testifylint/error-nil
</related>
