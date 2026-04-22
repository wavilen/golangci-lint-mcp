# revive: unnecessary-if

<instructions>
Detects `if` statements that are unnecessary because the condition is a constant boolean (`true` or `false`) or the if/else branches are identical. An always-true condition means the `if` wrapper is noise; always-false means the code is dead; identical branches mean the condition doesn't matter.

Remove the `if` and keep only the relevant branch, or rewrite the logic to eliminate the tautological condition.
</instructions>

<examples>
## Bad
```go
if true {
    doWork()
}
if err != nil && err != io.EOF {
} else {
    doWork() // identical in both branches
}
```

## Good
```go
doWork()
```
</examples>

<patterns>
- Remove `if true { ... }` wrappers around code that always executes
- Remove `if false { ... }` dead code blocks entirely
- Combine if/else branches with identical bodies into a single code path
- Remove conditions that are always true or false due to type constraints
- Simplify `if flag == true` to `if flag`
</patterns>

<related>
unnecessary-stmt, unnecessary-format, bool-literal-in-expr
