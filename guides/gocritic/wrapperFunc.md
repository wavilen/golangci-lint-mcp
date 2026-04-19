# gocritic: wrapperFunc

<instructions>
Detects functions that simply delegate to another function with no additional logic. These wrapper functions add indirection without value — they wrap a single call and return its result. Inlining the call or using the target function directly improves readability.

Remove the wrapper and call the target function directly, or add meaningful logic to justify the wrapper.
</instructions>

<examples>
## Bad
```go
func handleError(err error) {
	log.Fatal(err)
}
```

## Good
```go
// Use log.Fatal directly at call sites, or add real value:
func handleError(err error) {
	log.Printf("operation failed: %v", err)
	os.Exit(1)
}
```
</examples>

<patterns>
- Function body is a single call to another function
- Wrapper adds no logic, validation, or transformation
- Method that just forwards to another method on an embedded field
</patterns>

<related>
unlambda, deferUnlambda
</related>
