# revive: var-naming

<instructions>
Enforces Go variable naming conventions: use camelCase (not snake_case), avoid stuttering with the package name, and use short names for narrow scopes (e.g., loop variables) and descriptive names for wider scopes. Consistent naming improves readability across the codebase.

Rename variables to follow camelCase without underscores. Shorten overly verbose names in small scopes. Make package-level names descriptive.
</instructions>

<examples>
## Good
```go
userName := "Alice"
req := buildRequest()
for i := range users {
    process(users[i])
}
```
</examples>

<patterns>
- Use camelCase for variable names instead of snake_case (e.g., `userName` not `user_name`)
- Use short names like `i` for loop variables instead of overly long names like `userIndex`
- Avoid stuttering names like `userServiceService` — simplify to `service`
- Use descriptive names at package scope instead of single-letter names
- Use camelCase for variables instead of ALL_CAPS (which is for constants in other languages)
</patterns>

<related>
revive/var-declaration, revive/receiver-naming, revive/unexported-naming, revive/package-naming
</related>
