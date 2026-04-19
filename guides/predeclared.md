# predeclared

<instructions>
Predeclared detects when Go's predeclared identifiers (like `true`, `false`, `nil`, `iota`, `error`, `string`, `len`, `cap`, `copy`, etc.) are shadowed by local variable, type, or constant declarations. Shadowing built-in names causes confusion and subtle bugs.

Rename the shadowing identifier to something distinct from the predeclared name.
</instructions>

<examples>
## Bad
```go
type error struct {
    msg string
}
```

## Good
```go
type appError struct {
    msg string
}
```
</examples>

<patterns>
- `type error struct{}` shadowing the built-in `error` interface
- `var copy = ...` shadowing the built-in `copy` function
- `type string = myString` shadowing the built-in `string` type
- Local variables named `true`, `nil`, `len` hiding predeclared identifiers
</patterns>

<related>
govet, revive, staticcheck
