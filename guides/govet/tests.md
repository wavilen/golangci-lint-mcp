# govet: tests

<instructions>
Reports malformed test, benchmark, fuzz, and example function signatures. Functions named `TestXxx` must take `*testing.T`, benchmarks `BenchmarkXxx` must take `*testing.B`, and examples `ExampleXxx` must take no parameters. Wrong signatures mean the test runner ignores them silently.

Fix the function name and signature to match the expected pattern.
</instructions>

<examples>
## Bad
```go
func TestMyFunc(t testing.T) { // missing pointer — *testing.T required
    // ...
}
```

## Good
```go
func TestMyFunc(t *testing.T) {
    // ...
}
```
</examples>

<patterns>
- `TestXxx` without `*testing.T` parameter
- `BenchmarkXxx` without `*testing.B` parameter
- `ExampleXxx` with parameters (should have none)
- Test function with wrong name pattern (lowercase first letter)
</patterns>

<related>
testinggoroutine, unreachable
</related>
