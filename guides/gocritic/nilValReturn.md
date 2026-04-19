# gocritic: nilValReturn

<instructions>
Detects return statements that return a typed nil value where the caller will see a non-nil error interface. In Go, returning a typed nil pointer as an interface value results in a non-nil interface with a nil underlying value, which breaks `err != nil` checks in callers.

Return an explicit `nil` for the interface or return the concrete value directly without wrapping it in a nil check at the return site.
</instructions>

<examples>
## Bad
```go
func getResource() (*Resource, error) {
    var r *Resource // nil pointer
    if !found {
        return r, nil // typed nil returned as non-nil interface
    }
    return r, nil
}
```

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
- Returning typed nil pointer when no error occurred
- Returning nil struct pointer as interface value
- Helper functions returning `(SomeType, error)` where type is nil
- Factory methods that conditionally return nil concrete values
</patterns>

<related>
externalErrorReassign, uncheckedInlineErr
</related>
