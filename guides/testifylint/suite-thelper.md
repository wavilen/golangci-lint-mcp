# testifylint: suite-thelper

<instructions>
Detects incorrect usage of `t.Helper()` or missing helper patterns in testify suite methods. Suite methods should not manually call `s.T().Helper()` since the suite framework already manages helper marking. If you write helper methods on a suite, mark them with proper helper delegation.
</instructions>

<examples>
## Bad
```go
func (s *MySuite) TestFeature() {
    s.T().Helper()
    s.Equal(expected, actual)
}
```

## Good
```go
func (s *MySuite) TestFeature() {
    s.Equal(expected, actual)
}

func (s *MySuite) helperMethod() {
    s.T().Helper() // only in actual helper methods, not test methods
    s.Equal(1, 2)
}
```
</examples>

<patterns>
- Remove `s.T().Helper()` from test methods — not needed in suite test functions
- Add `s.T().Helper()` in custom suite helper methods for correct line reporting
- Remove `Helper()` calls from `SetupSuite`/`TearDownSuite` — unnecessary there
</patterns>

<related>
suite-method-signature, suite-extra-assert-call
