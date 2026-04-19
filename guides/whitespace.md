# whitespace

<instructions>
Whitespace detects unnecessary blank lines in code. Extra whitespace between declarations, after opening braces, or before closing braces adds visual noise without improving readability.

Remove extraneous blank lines. Keep one blank line between function definitions and between logical groups of declarations, but remove double blanks and leading/trailing whitespace in blocks.
</instructions>

<examples>
## Bad
```go
func process() error {

    data, err := read()
    if err != nil {

        return err
    }

    return nil
}
```

## Good
```go
func process() error {
    data, err := read()
    if err != nil {
        return err
    }

    return nil
}
```
</examples>

<patterns>
- Blank lines after opening braces `{` or before closing braces `}`
- Multiple consecutive blank lines between functions
- Leading blank lines in file bodies
- Trailing whitespace at end of code blocks
</patterns>

<related>
wsl_v5, nlreturn, decorder
</related>
