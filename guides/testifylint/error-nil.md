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
- `assert.Nil(t, err)` — use `assert.NoError(t, err)`
- `assert.NotNil(t, err)` — use `assert.Error(t, err)`
- `require.Nil(t, err)` — use `require.NoError(t, err)`
</patterns>

<related>
error-as, require-error
