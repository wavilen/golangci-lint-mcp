# revive: empty-lines

<instructions>
Detects unnecessary blank lines in specific contexts, such as multiple consecutive blank lines, blank lines at the start or end of a block, or blank lines between a function signature and its doc comment. Extra blank lines break visual flow and waste vertical space.

Remove the extra blank lines. Keep at most one blank line between declarations and none directly inside braces.
</instructions>

<examples>
## Bad
```go
func process() {


    // two blank lines above

    data := load()
    return data
}
```

## Good
```go
func process() {
    data := load()
    return data
}
```
</examples>

<patterns>
- Multiple consecutive blank lines between functions or declarations
- Blank line immediately after an opening brace
- Blank line immediately before a closing brace
- Blank lines between a doc comment and its declaration
- Extra blank lines introduced by automated formatting tools or merges
</patterns>

<related>
empty-block, comment-spacings
