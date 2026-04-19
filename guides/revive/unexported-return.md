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
- Exported function returning a struct defined with a lowercase name
- Exported method on an exported type returning an unexported type
- Factory functions returning unexported implementation types
- API surface where the return type is invisible to calling packages
- Tests in another package unable to assert on the return type
</patterns>

<related>
unexported-naming, receiver-naming, var-naming
