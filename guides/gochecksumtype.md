# gochecksumtype

<instructions>
Gochecksumtype ensures that interface types with a finite set of implementations have exhaustive type switches. It detects when a sum-type interface (marked with a `//sumtype:enclosed` comment) is switched on without covering all implementations.

Add a `default` case or handle the missing implementation type in the switch.
</instructions>

<examples>
## Good
```go
func area(s Shape) float64 {
    switch v := s.(type) {
    case Circle:
        return 3.14
    case Square:
        return 1.0
    default:
        panic(fmt.Sprintf("unhandled shape: %T", v))
    }
}
```
</examples>

<patterns>
- Handle all implementations in type switches over sum-type interfaces
- Add new struct types to every existing type switch when extending a sealed interface
- Include a `default` case as a safety net for future additions to the sum type
</patterns>

<related>
exhaustive, govet, staticcheck
</related>
