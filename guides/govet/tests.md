# govet: tests

<instructions>
Reports malformed test, benchmark, fuzz, and example function signatures. Functions named `TestXxx` must take `*testing.T`, benchmarks `BenchmarkXxx` must take `*testing.B`, and examples `ExampleXxx` must take no parameters. Wrong signatures mean the test runner ignores them silently.

Fix the function name and signature to match the expected pattern.
</instructions>

<examples>
## Good
```go
func TestMyFunc(_ *testing.T) {
    // ...
}
```
</examples>

<patterns>
- Add `*testing.T` as the parameter for `TestXxx` functions — the test runner ignores wrong signatures
- Add `*testing.B` as the parameter for `BenchmarkXxx` functions
- Remove parameters from `ExampleXxx` functions — examples must take no arguments
- Capitalize the first letter after `Test`/`Benchmark`/`Example` — lowercase names are ignored
</patterns>

<related>
govet/testinggoroutine, govet/unreachable
</related>
