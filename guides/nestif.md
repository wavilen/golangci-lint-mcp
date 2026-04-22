# nestif

<instructions>
Nestif checks the depth of nested if statements. Deep nesting makes code hard to follow because the reader must track multiple conditionals simultaneously. Default threshold is 3 levels.

Flatten nesting with guard clauses (early returns), extract inner logic into helper functions, or restructure with switch statements.
</instructions>

<examples>
## Bad
```go
func grant(user *User, res *Resource) bool {
    if user != nil {
        if res != nil {
            if user.Active {
                if res.Owner == user.ID {
                    return true
                }
            }
        }
    }
    return false
}
```

## Good
```go
func grant(user *User, res *Resource) bool {
    if user == nil || res == nil || !user.Active {
        return false
    }
    return res.Owner == user.ID
}
```
</examples>

<patterns>
- Flatten nested nil checks using early returns to reduce indentation
- Simplify nested permission logic with guard clauses and early returns
- Flatten multi-level conditionals in pipelines using early returns or helper functions
- Extract nested callback logic into named functions to reduce indentation
</patterns>

<related>
gocognit, cyclop, gocyclo, funlen
</related>
