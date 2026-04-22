# revive: unused-receiver

<instructions>
Detects method receivers that are never referenced in the method body. An unused receiver often means the method doesn't actually depend on the receiver's state and could be a standalone function. Alternatively, it may indicate an incomplete implementation.

Convert the method to a package-level function if it doesn't need the receiver. Otherwise, prefix the receiver name with `_` to signal intentional non-use for interface satisfaction.
</instructions>

<examples>
## Bad
```go
func (s *Server) Version() string {
    return "1.0.0" // s never used
}
```

## Good
```go
func Version() string {
    return "1.0.0"
}

// Or if satisfying an interface:
func (_ *Server) Version() string {
    return "1.0.0"
}
```
</examples>

<patterns>
- Convert methods returning constants unrelated to the receiver into package-level functions
- Move interface implementation methods that don't access receiver state to standalone functions, or prefix the receiver with `_`
- Implement placeholder methods properly or prefix receiver with `_` until ready
- Convert utility methods attached to a type only for namespace into package-level functions
- Move factory methods that create unrelated objects to package-level functions
</patterns>

<related>
unused-parameter, receiver-naming, unexported-return
