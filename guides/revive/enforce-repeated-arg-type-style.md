# revive: enforce-repeated-arg-type-style

<instructions>
Enforces a consistent style for function parameters and return values that share the same type. Go allows omitting repeated types (`x, y int`) but some teams prefer explicit types (`x int, y int`) for clarity.

Apply the project's chosen style consistently. Configure the preference in your revive config.
</instructions>

<examples>
## Bad
```go
// Team prefers explicit types but this uses shorthand
func move(x, y int) (dx, dy int) {
    return x, y
}
```

## Good
```go
// Consistent with project style: explicit
func move(x int, y int) (dx int, dy int) {
    return x, y
}

// Or consistent with project style: shorthand
func move(x, y int) (dx, dy int) {
    return x, y
}
```
</examples>

<patterns>
- Function signatures mixing shorthand and explicit type styles
- New team members using their preferred style without checking conventions
- Auto-generated code not matching the project's chosen style
- Code review feedback inconsistency on this point
</patterns>

<related>
enforce-map-style, enforce-slice-style
