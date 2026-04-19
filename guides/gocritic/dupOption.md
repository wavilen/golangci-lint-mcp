# gocritic: dupOption

<instructions>
Detects duplicate option patterns in functional option constructors or struct literal chains where the same option is applied more than once. The later application overwrites the earlier one, making the first one dead code.

Remove the duplicate option application. If both are needed with different values, rename or restructure to clarify intent.
</instructions>

<examples>
## Bad
```go
server := NewServer(
    WithPort(8080),
    WithTimeout(30*time.Second),
    WithPort(9090), // overwrites first WithPort
)
```

## Good
```go
server := NewServer(
    WithPort(9090),
    WithTimeout(30*time.Second),
)
```
</examples>

<patterns>
- Same functional option passed twice in constructor
- Duplicate key in struct literal or map literal
- Same `With*` option applied multiple times
- Duplicate field assignment in struct initialization
</patterns>

<related>
dupArg, dupSubExpr, dupBranchBody
</related>
