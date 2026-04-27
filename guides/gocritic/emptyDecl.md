# gocritic: emptyDecl

<instructions>
Detects empty declarations — files with only package declarations and imports but no code, or declarations with empty bodies that serve no purpose. Empty files and declarations add noise and can confuse readers about intended functionality.

Remove the empty file or declaration. If the file exists for future implementation, add a comment explaining the intent.
</instructions>

<examples>
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
- Remove files with only `package` and `import` declarations — no executable code
- Replace empty interface declarations with `any` when used as type constraints
- Remove empty struct types with no fields and no methods — use an established empty type
- Remove placeholder files committed without any content
</patterns>

<related>
gocritic/commentedOutCode, gocritic/codegenComment
</related>
