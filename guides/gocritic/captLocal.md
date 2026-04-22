# gocritic: captLocal

<instructions>
Detects local variables with names that start with a capital letter. In Go, capitalization controls visibility — capitalized identifiers are exported. Local variables are never exported, so capitalized names mislead readers into thinking the variable has broader scope.

Use camelCase starting with a lowercase letter for all local variables.
</instructions>

<examples>
## Bad
```go
func calculate() int {
	Result := 42
	return Result
}
```

## Good
```go
func calculate() int {
	result := 42
	return result
}
```
</examples>

<patterns>
- Rename short-lived `PascalCase` locals to `camelCase`
- Rename loop iteration variables with capitalized names to lowercase
- Use lowercase receiver names — avoid capitalized receiver names
- Rename block-scoped variables that look like exported symbols
</patterns>

<related>
docStub, exposedSyncMutex
</related>
