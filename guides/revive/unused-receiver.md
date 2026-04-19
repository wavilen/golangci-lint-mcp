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
- Methods that return constants unrelated to the receiver
- Interface implementations where the method logic doesn't access receiver state
- Methods added for future use with placeholder bodies
- Utility methods attached to a type for namespace reasons only
- Factory methods on a type that create unrelated objects
</patterns>

<related>
unused-parameter, receiver-naming, unexported-return
