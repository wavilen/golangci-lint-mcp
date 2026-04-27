# testifylint: suite-dont-use-pkg

<instructions>
Detects standalone `assert.Equal(t, ...)` or `require.NoError(t, ...)` inside suite test methods instead of using the suite's built-in assertion methods. Suite types embed `suite.Suite` which provides `s.Equal`, `s.NoError`, etc. Using these keeps assertions consistent and leverages the suite's failure handling.
</instructions>

<examples>
## Good
```go
func (s *MySuite) TestSomething() {
    s.Equal(expected, actual)
    s.Require().NoError(err)
}
```
</examples>

<patterns>
- Use `s.Xxx(...)` instead of `assert.Xxx(s.T(), ...)` in suite methods
- Use `s.Require().Xxx(...)` instead of `require.Xxx(s.T(), ...)` in suite methods
- Use suite methods consistently instead of mixing package-level assertions
</patterns>

<related>
testifylint/suite-extra-assert-call, testifylint/suite-method-signature
</related>
