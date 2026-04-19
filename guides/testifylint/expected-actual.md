# testifylint: expected-actual

<instructions>
Detects `assert.Equal(t, actual, expected)` where arguments are in the wrong order. Testify convention is `assert.Equal(t, expected, actual)` — expected value first. Correct ordering produces useful diff output showing what was expected versus what was received.
</instructions>

<examples>
## Bad
```go
assert.Equal(t, result, expected)
assert.Equal(t, calculate(), 42)
```

## Good
```go
assert.Equal(t, expected, result)
assert.Equal(t, 42, calculate())
```
</examples>

<patterns>
- `assert.Equal(t, computed, constant)` — swap to `assert.Equal(t, constant, computed)`
- `assert.Equal(t, result, expected)` — swap to `assert.Equal(t, expected, result)`
- Literal or expected value in second position — move to first
</patterns>

<related>
formatter, contains-unnecessary-format
