# revive: var-declaration

<instructions>
Detects variable declarations using the `var` keyword where a short variable declaration (`:=`) would be more idiomatic. When a variable is declared and immediately assigned, `:=` is more concise and preferred in Go. Use `var` only when the zero value is intentional or when declaring package-level variables.

Replace `var x = expr` with `x := expr` inside functions. Keep `var` for zero-value declarations or package-level variables.
</instructions>

<examples>
## Bad
```go
func process() {
    var name = "Alice"
    var count = len(items)
    var err = doSomething()
}
```

## Good
```go
func process() {
    name := "Alice"
    count := len(items)
    err := doSomething()
}
```
</examples>

<patterns>
- Use `:=` instead of `var x = value` inside functions
- Replace `var s string = "hello"` with `s := "hello"` for declarations with immediate assignment
- Extract complex package-level `var` initialization into an `init()` or dedicated function
- Replace multiple `var` declarations in a block with individual `:=` assignments
- Use `:=` in generated code instead of `var` for consistency
</patterns>

<related>
var-naming, unnecessary-stmt
