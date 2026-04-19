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
- Functions with many if-else branches handling different cases
- Complex switch statements with many case branches
- Functions combining validation, transformation, and error handling
- Legacy code that accumulated edge cases over time
- Business rule functions with many conditional outcomes
</patterns>

<related>
cognitive-complexity, function-length, max-control-nesting, gocyclo
