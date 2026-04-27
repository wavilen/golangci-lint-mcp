# revive: max-control-nesting

<instructions>
Enforces a maximum nesting depth for control flow constructs (if, for, switch, select). Deep nesting makes code hard to follow because the reader must mentally track many context switches and scope levels.

Use early returns, extract nested logic into helper functions, or restructure with guard clauses to flatten the code.
</instructions>

<examples>
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
- Replace nested guard conditions with sequential early returns
- Extract nested loop conditionals into separate helper functions
- Move switch-inside-loop logic into a dedicated function to reduce nesting
- Convert error handling to use guard clauses instead of nesting deeper
- Flatten callback or handler validation into sequential checks with early returns
</patterns>

<related>
revive/cognitive-complexity, revive/cyclomatic, revive/early-return, revive/indent-error-flow
</related>
