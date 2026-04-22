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
- Remove unused function parameters or replace with blank identifier `_`
- Remove boolean parameters that always receive the same constant value
- Eliminate unused interface method parameters or provide a narrower interface
- Propagate `context.Context` to downstream calls or remove it from the signature
</patterns>

<related>
unused, ineffassign, govet
