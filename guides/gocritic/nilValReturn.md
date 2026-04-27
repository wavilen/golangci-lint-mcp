# gocritic: nilValReturn

<instructions>
Detects return statements that return a typed nil value where the caller will see a non-nil error interface. In Go, returning a typed nil pointer as an interface value results in a non-nil interface with a nil underlying value, which breaks `err != nil` checks in callers.

Return an explicit `nil` for the interface or return the concrete value directly without wrapping it in a nil check at the return site.
</instructions>

<examples>
## Good
```go
func getResource() (*Resource, error) {
    if !found {
        return nil, nil // explicit nil
    }
    return &Resource{}, nil
}
```
</examples>

<patterns>
- Avoid returning typed nil pointer when no error occurred — return an explicitly initialized value
- Avoid returning nil struct pointer as interface value — return a concrete non-nil value
- Initialize helper function return types properly — avoid nil `(SomeType, error)` where type is nil
- Initialize factory method return values — avoid conditionally returning nil concrete values
</patterns>

<related>
gocritic/externalErrorReassign, gocritic/uncheckedInlineErr
</related>
