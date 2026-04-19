# revive: max-control-nesting

<instructions>
Enforces a maximum nesting depth for control flow constructs (if, for, switch, select). Deep nesting makes code hard to follow because the reader must mentally track many context switches and scope levels.

Use early returns, extract nested logic into helper functions, or restructure with guard clauses to flatten the code.
</instructions>

<examples>
## Bad
```go
func handle(req Request) error {
    if req.Valid {
        if req.Auth {
            if req.HasPermission {
                if req.RateLimitOk {
                    return process(req)
                }
            }
        }
    }
    return errors.New("rejected")
}
```

## Good
```go
func handle(req Request) error {
    if !req.Valid {
        return errors.New("invalid")
    }
    if !req.Auth {
        return errors.New("unauthorized")
    }
    if !req.HasPermission {
        return errors.New("forbidden")
    }
    if !req.RateLimitOk {
        return errors.New("rate limited")
    }
    return process(req)
}
```
</examples>

<patterns>
- Nested if-statements where each level adds a guard condition
- Loops containing nested conditionals with further nesting
- Switch statements inside loops inside other conditionals
- Error handling paths that nest deeper than the happy path
- Callback or handler functions with many levels of validation
</patterns>

<related>
cognitive-complexity, cyclomatic, early-return, indent-error-flow
