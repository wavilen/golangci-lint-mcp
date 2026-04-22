# revive: empty-block

<instructions>
Detects empty code blocks — `{}` with no statements inside. Empty blocks are dead code and usually indicate a missing implementation, a leftover from refactoring, or a logic error where code was accidentally deleted.

Remove the empty block entirely. If the block represents an intentional no-op, add a comment explaining why it is empty.
</instructions>

<examples>
## Bad
```go
switch status {
case http.StatusOK:
    handleSuccess()
case http.StatusNotFound:
    // forgot to implement
default:
}
```

## Good
```go
switch status {
case http.StatusOK:
    handleSuccess()
case http.StatusNotFound:
    handleNotFound()
}
```
</examples>

<patterns>
- Remove empty `default` cases in switch statements or add a comment explaining intent
- Remove empty `else` blocks left after removing dead code
- Remove empty `select` cases intended as placeholders
- Replace empty `for` loop bodies with the condition expression directly
- Remove empty blocks left after moving code elsewhere — remove the braces entirely
</patterns>

<related>
empty-lines, identical-branches
