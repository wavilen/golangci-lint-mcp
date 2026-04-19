# revive: var-naming

<instructions>
Enforces Go variable naming conventions: use camelCase (not snake_case), avoid stuttering with the package name, and use short names for narrow scopes (e.g., loop variables) and descriptive names for wider scopes. Consistent naming improves readability across the codebase.

Rename variables to follow camelCase without underscores. Shorten overly verbose names in small scopes. Make package-level names descriptive.
</instructions>

<examples>
## Bad
```go
user_name := "Alice"
http_request := buildRequest()
for user_index := range users {
    process(users[user_index])
}
```

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
- Variable names using snake_case (e.g., `user_name`)
- Loop variables with overly long names (e.g., `userIndex` instead of `i`)
- Stuttering names like `userServiceService` or `httpHTTPHandler`
- Single-letter names at package scope where a descriptive name is needed
- ALL_CAPS variable names mimicking constants from other languages
</patterns>

<related>
var-declaration, receiver-naming, unexported-naming, package-naming
