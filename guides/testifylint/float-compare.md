# testifylint: float-compare

<instructions>
Detects direct float equality assertions like `assert.Equal(t, 0.1+0.2, 0.3)` which fail due to floating-point precision. Use `assert.InDelta(t, expected, actual, tolerance)` or `assert.InEpsilon(t, expected, actual, relativeError)` for comparing floating-point numbers with a tolerance.
</instructions>

<examples>
## Bad
```go
assert.Equal(t, 0.3, result)
assert.True(t, math.Abs(result-0.3) < 1e-9)
```

## Good
```go
assert.InDelta(t, 0.3, result, 0.0001)
assert.InEpsilon(t, 0.3, result, 0.001)
```
</examples>

<patterns>
- Use `assert.InDelta` instead of `assert.Equal(t, floatA, floatB)`
- Use `assert.InDelta` instead of manual `math.Abs` comparison in assertions
- Use `assert.InDelta` or `assert.InEpsilon` for comparing computed float results
</patterns>

<related>
bool-compare, compares
