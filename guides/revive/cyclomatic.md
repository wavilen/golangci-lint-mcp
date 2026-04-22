# revive: cyclomatic

<instructions>
Measures cyclomatic complexity — the number of independent paths through a function's code. Each `if`, `for`, `case`, `&&`, and `||` adds one to the count. High cyclomatic complexity means the function is hard to test thoroughly (more test paths needed) and hard to understand.

Break the function into smaller functions with single responsibilities. Use early returns to reduce branching, and extract complex conditional logic into well-named helper functions.
</instructions>

<examples>
## Bad
```go
func classify(x int) string {
    if x > 100 {
        if x%2 == 0 {
            return "big-even"
        } else if x%3 == 0 {
            return "big-multiple-of-3"
        }
        return "big-odd"
    } else if x > 50 {
        if x%2 == 0 {
            return "medium-even"
        }
        return "medium-odd"
    } else if x > 0 {
        return "small"
    }
    return "zero-or-negative"
}
```

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
cognitive-complexity, function-length, max-control-nesting, gocyclo
