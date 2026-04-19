# gocritic: emptyDecl

<instructions>
Detects empty declarations — files with only package declarations and imports but no code, or declarations with empty bodies that serve no purpose. Empty files and declarations add noise and can confuse readers about intended functionality.

Remove the empty file or declaration. If the file exists for future implementation, add a comment explaining the intent.
</instructions>

<examples>
## Bad
```go
package handlers

// No types, functions, or variables declared
```

## Good
```go
package handlers

// Handler processes incoming HTTP requests.
type Handler struct {
    Service *Service
}
```
</examples>

<patterns>
- Files with only `package` and `import` declarations
- Empty interface declarations used as type constraints
- Empty struct types with no fields and no methods
- Placeholder files committed without any content
</patterns>

<related>
commentedOutCode, codegenComment
</related>
