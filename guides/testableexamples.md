# testableexamples

<instructions>
Testableexamples checks that Go example functions (those starting with `func Example`) are actually testable — they must end with an `// Output:` comment. Without it, the example is displayed in documentation but never executed as a test.

Add an `// Output:` comment at the end of the example function body showing the expected stdout output.
</instructions>

<examples>
## Good
```go
func ExampleGreeting() {
    fmt.Println("hello")
    // Output: hello
}
```
</examples>

<patterns>
- Add `// Output:` comments to all example functions so they run as tests
- Ensure `// Output:` matches the actual stdout of the example
- Use `// Output:` (empty) only when the example genuinely prints nothing
- Avoid `os.Exit` or `panic` in examples — they prevent test execution
</patterns>

<related>
testpackage, godoclint, thelper
</related>
