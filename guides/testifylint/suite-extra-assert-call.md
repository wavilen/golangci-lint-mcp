# testifylint: suite-extra-assert-call

<instructions>
Detects `s.Assert().Equal(t, ...)` where the extra `Assert()` wrapper is redundant inside suite methods. Suite types already provide assertion methods directly through `s.Equal`, `s.True`, etc. The `Assert()` method returns a standard `*Assertions` object, but calling it adds unnecessary indirection.
</instructions>

<examples>
## Bad
```go
func (s *MySuite) TestFeature() {
    s.Assert().Equal(expected, actual)
    s.Assert().True(ok)
}
```

## Good
```go
func (s *MySuite) TestFeature() {
    s.Equal(expected, actual)
    s.True(ok)
}
```
</examples>

<patterns>
- `s.Assert().Xxx(...)` — use `s.Xxx(...)` directly
- `s.Require().Equal(...)` is fine for require-style, but `s.Assert()` is unnecessary
- Any `s.Assert()` call in suite method — replace with direct suite method
</patterns>

<related>
suite-dont-use-pkg, suite-method-signature
