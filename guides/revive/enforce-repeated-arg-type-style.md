# revive: enforce-repeated-arg-type-style

<instructions>
Enforces a consistent style for function parameters and return values that share the same type. Go allows omitting repeated types (`x, y int`) but some teams prefer explicit types (`x int, y int`) for clarity.

Apply the project's chosen style consistently. Configure the preference in your revive config.
</instructions>

<examples>
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
- Use consistent shorthand or explicit type style for repeated parameter types within a function signature
- Use the project's convention for repeated argument types — don't mix styles
- Ensure auto-generated code to the project's chosen repeated-arg-type style
- Ensure consistent repeated-arg-type style in code reviews
</patterns>

<related>
revive/enforce-map-style, revive/enforce-slice-style
</related>
