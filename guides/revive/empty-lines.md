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
- Remove multiple consecutive blank lines between functions or declarations to a single blank line
- Remove blank lines immediately after an opening brace
- Remove blank lines immediately before a closing brace
- Move doc comments directly above their declaration with no intervening blank lines
- Remove extra blank lines introduced by automated formatting tools or merges
</patterns>

<related>
empty-block, comment-spacings
