# testifylint: suite-broken-parallel

<instructions>
Detects `t.Parallel()` called inside testify suite test methods. The suite runner manages test execution itself and calling `t.Parallel()` inside suite methods breaks the suite's setup/teardown lifecycle. Remove parallel calls from suite methods; parallelism should be managed at the suite level instead.
</instructions>

<examples>
## Bad
```go
func (s *MySuite) TestFeature() {
    s.T().Parallel()
    // test code
}
```

## Good
```go
func (s *MySuite) TestFeature() {
    // test code — no Parallel() in suite methods
}
```
</examples>

<patterns>
- `s.T().Parallel()` inside suite method — remove it
- `t.Parallel()` in suite setup/teardown — remove it
- Running suite tests in parallel — use parallel suite runner configuration
</patterns>

<related>
suite-dont-use-pkg, suite-extra-assert-call
