# testifylint: suite-method-signature

<instructions>
Detects suite test methods with incorrect signatures. Testify suite methods must have no parameters — the suite provides `s.T()` for accessing `*testing.T`. A method like `func (s *MySuite) TestFoo(t *testing.T)` will not be called by the suite runner and the test will silently not run.
</instructions>

<examples>
## Good
```go
func (s *MySuite) TestFoo() {
    s.Equal(expected, actual)
}
```
</examples>

<patterns>
- Remove the `t *testing.T` parameter from `func (s *Suite) TestXxx(t *testing.T)`
- Remove the `ctx context.Context` parameter from `func (s *Suite) TestXxx(ctx context.Context)`
- Ensure suite methods have zero parameters — use `s.T()` to access `*testing.T`
</patterns>

<related>
testifylint/suite-thelper, testifylint/suite-dont-use-pkg
</related>
