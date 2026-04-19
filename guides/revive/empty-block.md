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
- Empty `default` cases in switch statements
- Empty `else` blocks left after removing dead code
- Empty `select` cases intended as placeholders
- Empty `for` loop bodies where the condition handles everything
- Empty blocks after code was moved elsewhere but braces remained
</patterns>

<related>
empty-lines, identical-branches
