# testifylint: require-error

<instructions>
Detects missing error assertion before using the result of a function that returns a value and an error. If `result, err := fn()` is called and `result` is used without checking `err`, a nil pointer dereference or incorrect behavior can occur. Add `require.NoError(t, err)` before accessing the result.
</instructions>

<examples>
## Bad
```go
result, err := strconv.Atoi(input)
assert.Equal(t, 42, result) // what if err != nil?
```

## Good
```go
result, err := strconv.Atoi(input)
require.NoError(t, err)
assert.Equal(t, 42, result)
```
</examples>

<patterns>
- `val, err := fn()` with `val` used but `err` unchecked — add `require.NoError`
- Multiple return values where error is ignored — always check first
- `result.Method()` after error-returning call — require.NoError before use
</patterns>

<related>
go-require, error-nil
