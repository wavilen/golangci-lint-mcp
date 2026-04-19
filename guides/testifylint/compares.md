# testifylint: compares

<instructions>
Detects `assert.True(t, a > b)` or `assert.False(t, a < b)` where dedicated comparison assertions provide better failure output. Use `assert.Greater`, `assert.GreaterOrEqual`, `assert.Less`, or `assert.LessOrEqual` instead. These show both operands in the failure message.
</instructions>

<examples>
## Bad
```go
assert.True(t, score > passingGrade)
assert.False(t, age < minimumAge)
```

## Good
```go
assert.Greater(t, score, passingGrade)
assert.GreaterOrEqual(t, age, minimumAge)
```
</examples>

<patterns>
- `assert.True(t, a > b)` — use `assert.Greater(t, a, b)`
- `assert.True(t, a >= b)` — use `assert.GreaterOrEqual(t, a, b)`
- `assert.False(t, a < b)` — use `assert.GreaterOrEqual(t, a, b)`
- `assert.True(t, a < b)` — use `assert.Less(t, a, b)`
</patterns>

<related>
bool-compare, float-compare
