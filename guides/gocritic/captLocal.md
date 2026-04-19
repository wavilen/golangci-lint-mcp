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
- Short-lived local variables with `PascalCase` names
- Loop iteration variables with capitalized names
- Receiver names that are capitalized
- Block-scoped variables that look like exported symbols
</patterns>

<related>
docStub, exposedSyncMutex
</related>
