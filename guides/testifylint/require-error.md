# testifylint: require-error

<instructions>
Detects missing error assertion before using the result of a function that returns a value and an error. If `result, err := fn()` is called and `result` is used without checking `err`, a nil pointer dereference or incorrect behavior can occur. Add `require.NoError(t, err)` before accessing the result.
</instructions>

<examples>
## Good
```go
result, err := strconv.Atoi(input)
require.NoError(t, err)
assert.Equal(t, 42, result)
```
</examples>

<patterns>
- Add `require.NoError` when `val, err := fn()` uses `val` but leaves `err` unchecked
- Check error return values first when multiple return values include an error
- Add `require.NoError` before using results from error-returning calls
</patterns>

<related>
testifylint/go-require, testifylint/error-nil
</related>
