# testifylint: suite-broken-parallel

<instructions>
Detects `t.Parallel()` called inside testify suite test methods. The suite runner manages test execution itself and calling `t.Parallel()` inside suite methods breaks the suite's setup/teardown lifecycle. Remove parallel calls from suite methods; parallelism should be managed at the suite level instead.
</instructions>

<examples>
## Good
```go
func (s *MySuite) TestFeature() {
    // test code — no Parallel() in suite methods
}
```
</examples>

<patterns>
- Remove `s.T().Parallel()` from suite methods — the suite runner manages execution
- Remove `t.Parallel()` from suite setup/teardown methods
- Use parallel suite runner configuration instead of `t.Parallel()` in individual suite tests
</patterns>

<related>
testifylint/suite-dont-use-pkg, testifylint/suite-extra-assert-call
</related>
