# revive: enforce-map-style

<instructions>
Enforces a consistent style for map initialization. Choose between `make(map[K]V)` and `map[K]V{}` literal syntax and apply it consistently across the codebase. Mixed styles create visual inconsistency.

Configure the preferred style in your revive config and apply it uniformly. Use `make` when the map will be populated later or when the size is known; use literals when initializing with known entries.
</instructions>

<examples>
## Bad
```go
// Team uses make-style but this file uses literals
m := map[string]int{}
```

## Good
```go
// Consistent with project style: make
m := make(map[string]int)

// Or consistent with project style: literal
m := map[string]int{}
```
</examples>

<patterns>
- Mixed map initialization styles within the same package
- Inconsistent code review standards for map creation
- Auto-generated code using a different style than hand-written code
- Files authored by different team members with different habits
</patterns>

<related>
enforce-slice-style, enforce-switch-style
