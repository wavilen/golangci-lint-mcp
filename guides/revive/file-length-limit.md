# revive: file-length-limit

<instructions>
Enforces a maximum line count per source file. Very long files are hard to navigate, slow to review, and often indicate that the file has too many responsibilities. Splitting large files improves maintainability and reduces merge conflicts.

Split the file into smaller, focused files. Group related types and functions by responsibility or domain concept.
</instructions>

<examples>
## Bad
```go
// user_handler.go — 1200 lines mixing HTTP handling,
// validation, database queries, and response formatting
```

## Good
```go
// user_handler.go     — HTTP route definitions (~100 lines)
// user_validation.go  — input validation logic (~80 lines)
// user_repository.go  — database access layer (~150 lines)
// user_response.go    — response formatting (~60 lines)
```
</examples>

<patterns>
- God files accumulating all related functionality in one place
- Single files containing multiple unrelated types or services
- Utility files growing over time with no clear organization principle
- Generated code files exceeding reasonable size limits
- Test files with many test cases that could be split by feature
</patterns>

<related>
line-length-limit, function-length, max-public-structs
