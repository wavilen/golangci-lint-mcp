# iface

<instructions>
Ifacet detects interfaces that define methods with pointer receiver parameters using the implementing type itself (e.g., `func(*T) Method()` in an interface that `T` cannot satisfy). This creates interfaces that are impossible to implement correctly.

Redesign the interface to accept the concrete type as a parameter, or use value receivers, or restructure the interface to avoid self-referential pointer receivers.
</instructions>

<examples>
## Bad
```go
type Modifier interface {
    Modify(*Modifier) // interface satisfied by *Modifier — confusing
}
```

## Good
```go
type Modifier interface {
    Modify() error
}

func (m *MyModifier) Modify() error {
    // implementation
    return nil
}
```
</examples>

<patterns>
- Avoid interface methods that accept the interface type itself as a parameter
- Resolve circular satisfaction by using value receivers or accepting concrete types
- Simplify interface embedding to avoid impossible-to-satisfy method sets
- Replace pointer receiver parameters in interface methods with value receivers or concrete types
</patterns>

<related>
interfacebloat, recvcheck, exhaustruct
</related>
