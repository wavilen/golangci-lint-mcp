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
- `var x = value` inside functions where `:=` would work
- Declarations with explicit type and immediate assignment (`var s string = "hello"`)
- Package-level `var` with complex initialization that could be a function
- Multiple `var` declarations in a block that could use individual `:=`
- Generated code using `var` consistently instead of `:=`
</patterns>

<related>
var-naming, unnecessary-stmt
