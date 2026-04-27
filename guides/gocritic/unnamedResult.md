# gocritic: unnamedResult

<instructions>
Detects named return types in function signatures that have unnamed results. Unnamed return values make it harder to understand what each value represents, especially in functions with multiple returns of the same type.

Name all return values in the function signature, or document them clearly in the function's doc comment.
</instructions>

<examples>
## Good
```go
func divide(a, b float64) (quotient float64, err error) {
	if b == 0 {
		return 0, errors.New("division by zero")
	}
	return a / b, nil
}
```
</examples>

<patterns>
- Add names to all return values when returning multiple: `(int, string, error)` → `(count int, name string, err error)`
- Add names to return values of the same type to prevent order mix-ups
- Add names to return values in public API functions — names appear in documentation
</patterns>

<related>
gocritic/tooManyResultsChecker, gocritic/paramTypeCombine
</related>
