# testifylint: error-as

<instructions>
Detects direct usage of `errors.As(err, &target)` in tests when testify provides `assert.ErrorAs(t, err, &target)` for consistency with other test assertions. Using the testify wrapper keeps the assertion style uniform and adds test failure reporting automatically.
</instructions>

<examples>
## Bad
```go
var pathErr *os.PathError
if !errors.As(err, &pathErr) {
    t.Fatal("expected PathError")
}
```

## Good
```go
var pathErr *os.PathError
assert.ErrorAs(t, err, &pathErr, "expected PathError")
```
</examples>

<patterns>
- Use `assert.ErrorAs` instead of `errors.As(err, &target)` in test code
- Use `assert.ErrorAs` instead of manual type assertions on errors
- Use `assert.ErrorAs` instead of `assert.True(t, errors.As(err, &target))`
</patterns>

<related>
error-nil, require-error
