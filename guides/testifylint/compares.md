# testifylint: compares

<instructions>
Detects `assert.True(t, a > b)` or `assert.False(t, a < b)` where dedicated comparison assertions provide better failure output. Use `assert.Greater`, `assert.GreaterOrEqual`, `assert.Less`, or `assert.LessOrEqual` instead. These show both operands in the failure message.
</instructions>

<examples>
## Good
```go
assert.Greater(t, score, passingGrade)
assert.GreaterOrEqual(t, age, minimumAge)
```
</examples>

<patterns>
- Use `assert.Greater(t, a, b)` instead of `assert.True(t, a > b)`
- Use `assert.GreaterOrEqual(t, a, b)` instead of `assert.True(t, a >= b)`
- Use `assert.GreaterOrEqual(t, a, b)` instead of `assert.False(t, a < b)`
- Use `assert.Less(t, a, b)` instead of `assert.True(t, a < b)`
</patterns>

<related>
testifylint/bool-compare, testifylint/float-compare
</related>
