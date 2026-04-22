# revive: unexported-return

<instructions>
Detects exported functions or methods that return unexported types. This forces callers outside the package to use the type opaquely, preventing them from declaring variables, using the type in their own signatures, or accessing fields. It also breaks documentation generation.

Either export the return type so callers can use it, or return an interface that exposes the needed behavior. If the type is intentionally opaque, document this decision.
</instructions>

<examples>
## Bad
```go
type result struct {
    Value int
}

func Process() result { // returns unexported type
    return result{Value: 42}
}
```

## Good
```go
type Result struct {
    Value int
}

func Process() Result {
    return Result{Value: 42}
}
```
</examples>

<patterns>
- Export the return type when an exported function returns an unexported struct
- Export the return type when an exported method returns an unexported type
- Return an interface from factory functions instead of an unexported implementation type
- Export types used in the API surface so calling packages can reference them
- Export return types needed for assertions in external test packages
</patterns>

<related>
unexported-naming, receiver-naming, var-naming
