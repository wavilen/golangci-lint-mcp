# testifylint: nil-compare

<instructions>
Detects `assert.Nil(t, x)` for non-error types where a more specific assertion is available. For errors, use `assert.NoError`. For collections, use `assert.Empty`. For pointers and interfaces where nil is semantically meaningful, `assert.Nil` is fine — this checker only flags cases where a better alternative exists.
</instructions>

<examples>
## Bad
```go
assert.Nil(t, err)
assert.Nil(t, []int{})
```

## Good
```go
assert.NoError(t, err)
assert.Empty(t, []int{})
```
</examples>

<patterns>
- `assert.Nil(t, err)` — use `assert.NoError(t, err)`
- `assert.Nil(t, emptySlice)` — use `assert.Empty(t, slice)`
- `assert.NotNil(t, err)` — use `assert.Error(t, err)`
</patterns>

<related>
error-nil, empty, len
