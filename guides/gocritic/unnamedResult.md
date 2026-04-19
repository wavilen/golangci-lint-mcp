# gocritic: unnamedResult

<instructions>
Detects named return types in function signatures that have unnamed results. Unnamed return values make it harder to understand what each value represents, especially in functions with multiple returns of the same type.

Name all return values in the function signature, or document them clearly in the function's doc comment.
</instructions>

<examples>
## Bad
```go
func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("division by zero")
	}
	return a / b, nil
}
```

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
- Multiple return values with no names: `(int, string, error)`
- Returns of same type that are easy to mix up
- Public API functions where return value names appear in documentation
</patterns>

<related>
tooManyResultsChecker, paramTypeCombine
</related>
