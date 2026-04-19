# revive: enforce-slice-style

<instructions>
Enforces a consistent style for slice initialization. Choose between `make([]T, 0)`, `make([]T, 0, n)`, and `[]T{}` literal syntax and apply it consistently. Mixed styles create visual inconsistency across the codebase.

Configure the preferred style in your revive config and apply it everywhere. Use `make` with a capacity hint when the approximate size is known for better performance.
</instructions>

<examples>
## Bad
```go
// Team uses make-style but this file uses literals
items := []string{}
```

## Good
```go
// Consistent with project style: make
items := make([]string, 0)

// Or with capacity hint
items := make([]string, 0, len(input))
```
</examples>

<patterns>
- Mixed slice initialization styles within the same package
- Empty slice literals `[]T{}` vs `make([]T, 0)` inconsistency
- Some files using capacity hints while others do not
- Auto-generated code using a different style than hand-written code
</patterns>

<related>
enforce-map-style, enforce-switch-style
