# testifylint: suite-dont-use-pkg

<instructions>
Detects standalone `assert.Equal(t, ...)` or `require.NoError(t, ...)` inside suite test methods instead of using the suite's built-in assertion methods. Suite types embed `suite.Suite` which provides `s.Equal`, `s.NoError`, etc. Using these keeps assertions consistent and leverages the suite's failure handling.
</instructions>

<examples>
## Bad
```go
func (s *MySuite) TestSomething() {
    assert.Equal(s.T(), expected, actual)
    require.NoError(s.T(), err)
}
```

## Good
```go
func (s *MySuite) TestSomething() {
    s.Equal(expected, actual)
    s.Require().NoError(err)
}
```
</examples>

<patterns>
- `assert.Xxx(s.T(), ...)` in suite method — use `s.Xxx(...)`
- `require.Xxx(s.T(), ...)` in suite method — use `s.Require().Xxx(...)`
- Mixing package-level assertions with suite methods — use suite methods consistently
</patterns>

<related>
suite-extra-assert-call, suite-method-signature
