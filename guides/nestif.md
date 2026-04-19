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
- Nested null/nil checks that can be combined with early returns
- Deeply nested permission or authorization logic
- Multi-level conditionals in data processing pipelines
- Callback-based code where nesting accumulates across async boundaries
</patterns>

<related>
gocognit, cyclop, gocyclo, funlen
</related>
