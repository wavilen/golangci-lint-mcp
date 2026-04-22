# testifylint: error-nil

<instructions>
Detects `assert.Nil(t, err)` or `assert.NotNil(t, err)` for error values. Use `assert.NoError(t, err)` and `assert.Error(t, err)` for semantic clarity. The `NoError`/`Error` assertions produce clearer failure messages and communicate intent: you are checking for an error condition, not for nil-ness.
</instructions>

<examples>
## Bad
```go
assert.Nil(t, err)
assert.NotNil(t, err)
```

## Good
```go
assert.NoError(t, err)
assert.Error(t, err)
```
</examples>

<patterns>
- Use `assert.NoError(t, err)` instead of `assert.Nil(t, err)`
- Use `assert.Error(t, err)` instead of `assert.NotNil(t, err)`
- Use `require.NoError(t, err)` instead of `require.Nil(t, err)`
</patterns>

<related>
error-as, require-error
