# revive: enforce-map-style

<instructions>
Enforces a consistent style for map initialization. Choose between `make(map[K]V)` and `map[K]V{}` literal syntax and apply it consistently across the codebase. Mixed styles create visual inconsistency.

Configure the preferred style in your revive config and apply it uniformly. Use `make` when the map will be populated later or when the size is known; use literals when initializing with known entries.
</instructions>

<examples>
## Good
```go
// Consistent with project style: make
m := make(map[string]int)

// Or consistent with project style: literal
m := map[string]int{}
```
</examples>

<patterns>
- Use either `make` or map literals consistently within the same package
- Ensure code review standards agree on a single map creation style
- Ensure auto-generated code to the same map style used by hand-written code
- Use the prevailing convention in files by different authors
</patterns>

<related>
revive/enforce-slice-style, revive/enforce-switch-style
</related>
