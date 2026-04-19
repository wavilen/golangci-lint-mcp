# testableexamples

<instructions>
Testableexamples checks that Go example functions (those starting with `func Example`) are actually testable — they must end with an `// Output:` comment. Without it, the example is displayed in documentation but never executed as a test.

Add an `// Output:` comment at the end of the example function body showing the expected stdout output.
</instructions>

<examples>
## Bad
```go
func ExampleGreeting() {
    fmt.Println("hello")
}
```

## Good
```go
func ExampleGreeting() {
    fmt.Println("hello")
    // Output: hello
}
```
</examples>

<patterns>
- Example functions without `// Output:` comment
- `// Output:` with wrong expected output (test fails silently)
- Empty `// Output:` when the example prints nothing
- Examples that call `os.Exit` or panic (not testable)
</patterns>

<related>
testpackage, godoclint, thelper
</related>
