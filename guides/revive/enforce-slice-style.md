# revive: enforce-slice-style

<instructions>
Enforces a consistent style for slice initialization. Choose between `make([]T, 0)`, `make([]T, 0, n)`, and `[]T{}` literal syntax and apply it consistently. Mixed styles create visual inconsistency across the codebase.

Configure the preferred style in your revive config and apply it everywhere. Use `make` with a capacity hint when the approximate size is known for better performance.
</instructions>

<examples>
## Good
```go
// Consistent with project style: make
items := make([]string, 0)

// Or with capacity hint
items := make([]string, 0, len(input))
```
</examples>

<patterns>
- Use consistent slice initialization style within the same package — either `make` or literals
- Use between empty slice literals `[]T{}` and `make([]T, 0)` consistently across the codebase
- Use capacity hints uniformly — either use them everywhere or nowhere for similar patterns
- Ensure auto-generated code to the slice initialization style used by hand-written code
</patterns>

<related>
revive/enforce-map-style, revive/enforce-switch-style
</related>
