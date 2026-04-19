# gochecksumtype

<instructions>
Gochecksumtype ensures that interface types with a finite set of implementations have exhaustive type switches. It detects when a sum-type interface (marked with a `//sumtype:enclosed` comment) is switched on without covering all implementations.

Add a `default` case or handle the missing implementation type in the switch.
</instructions>

<examples>
## Bad
```go
//sumtype:enclosed
type Shape interface{ shape() }

type Circle struct{}
func (Circle) shape() {}

type Square struct{}
func (Square) shape() {}

func area(s Shape) float64 {
    switch s.(type) {
    case Circle:
        return 3.14
    // Square not handled
    }
    return 0
}
```

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
- Missing implementation in a type switch over a sum-type interface
- New struct added to a sealed interface but not added to existing switches
- No `default` case as a safety net for future additions
</patterns>

<related>
exhaustive, govet, staticcheck
