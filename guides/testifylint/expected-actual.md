# testifylint: expected-actual

<instructions>
Detects `assert.Equal(t, actual, expected)` where arguments are in the wrong order. Testify convention is `assert.Equal(t, expected, actual)` — expected value first. Correct ordering produces useful diff output showing what was expected versus what was received.
</instructions>

<examples>
## Good
```go
assert.Equal(t, expected, result)
assert.Equal(t, 42, calculate())
```
</examples>

<patterns>
- Move the expected value to the first position: `assert.Equal(t, constant, computed)` instead of `(t, computed, constant)`
- Move the expected value to the first position: `assert.Equal(t, expected, result)` instead of `(t, result, expected)`
- Move literal or expected values to the first argument position
</patterns>

<related>
testifylint/formatter, testifylint/contains-unnecessary-format, testifylint/bool-compare, testifylint/nil-compare
</related>
