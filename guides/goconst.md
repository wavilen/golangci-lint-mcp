# goconst

<instructions>
Goconst finds repeated strings that could be replaced by constants. Magic strings duplicated across files hurt maintainability — changing them requires finding every occurrence.

Extract the repeated string into a `const` declaration. Use a typed constant if the string carries semantic meaning beyond its value.
</instructions>

<examples>
## Bad
```go
func role(r string) bool {
    return r == "admin" || r == "superuser"
}
func canEdit(r string) bool {
    return r == "admin"
}
```

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
- Error message strings used in multiple places
- Status or role strings like "active", "admin", "pending"
- HTTP method or content-type literals repeated across handlers
- SQL fragment strings duplicated in query builders
</patterns>

<related>
mnd, gochecknoglobals, dupl
</related>
