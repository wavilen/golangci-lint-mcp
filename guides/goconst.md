# goconst

<instructions>
Goconst finds repeated strings that could be replaced by constants. Magic strings duplicated across files hurt maintainability — changing them requires finding every occurrence.

Extract the repeated string into a `const` declaration. Use a typed constant if the string carries semantic meaning beyond its value.
</instructions>

<examples>
## Good
```go
const RoleAdmin = "admin"

func role(r string) bool {
    return r == RoleAdmin || r == "superuser"
}
func canEdit(r string) bool {
    return r == RoleAdmin
}
```
</examples>

<patterns>
- Extract repeated error message strings into named constants
- Define constants for repeated status or role strings
- Replace repeated HTTP method and content-type literals with constants
- Extract duplicated SQL fragment strings into constants
</patterns>

<related>
mnd, gochecknoglobals, dupl
</related>
