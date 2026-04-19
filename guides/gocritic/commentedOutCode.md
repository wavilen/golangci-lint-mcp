# gocritic: commentedOutCode

<instructions>
Detects commented-out code that should be removed. Commented-out code clutters the codebase, confuses readers about intent, and decays without testing. Use version control to preserve history instead of commenting out code.

Remove commented-out code. If code might be needed later, rely on git history.
</instructions>

<examples>
## Bad
```go
// result := process(data)
// if result.Err != nil {
//     return result.Err
// }
```

## Good
```go
// Delete the commented code entirely.
// Git preserves the history if you need it back.
```
</examples>
