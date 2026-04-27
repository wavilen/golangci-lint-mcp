# revive: file-length-limit

<instructions>
Enforces a maximum line count per source file. Very long files are hard to navigate, slow to review, and often indicate that the file has too many responsibilities. Splitting large files improves maintainability and reduces merge conflicts.

Split the file into smaller, focused files. Group related types and functions by responsibility or domain concept.
</instructions>

<examples>
## Good
```go
// user_handler.go     — HTTP route definitions (~100 lines)
// user_validation.go  — input validation logic (~80 lines)
// user_repository.go  — database access layer (~150 lines)
// user_response.go    — response formatting (~60 lines)
```
</examples>

<patterns>
- Separate god files into separate files grouped by responsibility
- Separate unrelated types or services into their own files
- Organize utility files by a clear principle — or split them by domain
- Separate generated code files that exceed size limits into smaller units
- Separate large test files by feature into separate test files
</patterns>

<related>
revive/line-length-limit, revive/function-length, revive/max-public-structs
</related>
