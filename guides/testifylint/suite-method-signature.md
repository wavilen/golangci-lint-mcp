# testifylint: suite-method-signature

<instructions>
Detects suite test methods with incorrect signatures. Testify suite methods must have no parameters — the suite provides `s.T()` for accessing `*testing.T`. A method like `func (s *MySuite) TestFoo(t *testing.T)` will not be called by the suite runner and the test will silently not run.
</instructions>

<examples>
## Bad
```go
func (s *MySuite) TestFoo(t *testing.T) {
    assert.Equal(t, expected, actual)
}
```

## Good
```go
func (s *MySuite) TestFoo() {
    s.Equal(expected, actual)
}
```
</examples>

<patterns>
- `func (s *Suite) TestXxx(t *testing.T)` — remove the parameter
- `func (s *Suite) TestXxx(ctx context.Context)` — remove the parameter
- Any suite method accepting parameters — suite methods must have zero params
</patterns>

<related>
suite-thelper, suite-dont-use-pkg
