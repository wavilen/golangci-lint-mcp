# revive: cyclomatic

<instructions>
Measures cyclomatic complexity — the number of independent paths through a function's code. Each `if`, `for`, `case`, `&&`, and `||` adds one to the count. High cyclomatic complexity means the function is hard to test thoroughly (more test paths needed) and hard to understand.

Break the function into smaller functions with single responsibilities. Use early returns to reduce branching, and extract complex conditional logic into well-named helper functions.
</instructions>

<examples>
## Good
```go
func classify(x int) string {
    if x > 100 {
        return classifyLarge(x)
    }
    if x > 50 {
        return classifyMedium(x)
    }
    if x > 0 {
        return "small"
    }
    return "zero-or-negative"
}
```
</examples>

<patterns>
- Flatten functions with many if-else branches into smaller functions per case
- Extract complex switch statements with many branches into a dispatch table or helper methods
- Separate validation, transformation, and error handling into distinct functions
- Extract edge cases from legacy code into focused helper functions
- Separate business rule functions with many conditional outcomes into rule-specific functions
</patterns>

<related>
revive/cognitive-complexity, revive/function-length, revive/max-control-nesting, gocyclo
</related>
