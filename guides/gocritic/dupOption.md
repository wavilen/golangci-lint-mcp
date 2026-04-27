# gocritic: dupOption

<instructions>
Detects duplicate option patterns in functional option constructors or struct literal chains where the same option is applied more than once. The later application overwrites the earlier one, making the first one dead code.

Remove the duplicate option application. If both are needed with different values, rename or restructure to clarify intent.
</instructions>

<examples>
## Good
```go
server := NewServer(
    WithPort(9090),
    WithTimeout(30*time.Second),
)
```
</examples>

<patterns>
- Remove duplicate functional options passed to constructors
- Remove duplicate keys in struct or map literals
- Remove duplicate `With*` options — pass each option only once
- Remove duplicate field assignments in struct initialization
</patterns>

<related>
gocritic/dupArg, gocritic/dupSubExpr, gocritic/dupBranchBody
</related>
