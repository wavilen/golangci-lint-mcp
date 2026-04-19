# unparam

<instructions>
Unparam detects function parameters that are never used or always receive the same constant value. Unused parameters indicate dead code or a too-generic interface, while parameters that always get the same value suggest the function doesn't need that argument.

Remove unused parameters or replace them with the constant value inside the function.
</instructions>

<examples>
## Bad
```go
func greet(name string, verbose bool) string {
    // verbose is never used
    return "Hello, " + name
}
```

## Good
```go
func greet(name string) string {
    return "Hello, " + name
}
```
</examples>

<patterns>
- Function parameters that are never referenced in the function body
- Boolean parameters always passed as `true` or `false` by every caller
- Interface method parameters unused in specific implementations
- Context parameters accepted but never propagated
</patterns>

<related>
unused, ineffassign, govet
