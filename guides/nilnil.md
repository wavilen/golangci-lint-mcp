# nilnil

<instructions>
Nilnil detects functions that return both a nil pointer and a nil error, which is ambiguous — the caller cannot distinguish "success with nil result" from "no result available." This pattern violates Go's convention that `(nil, nil)` should mean success with no value.

Return a meaningful zero-value, a sentinel error, or a wrapper type instead.
</instructions>

<examples>
## Bad
```go
func FindUser(id int) (*User, error) {
    row := db.QueryRow("SELECT ...", id)
    var u User
    if err := row.Scan(&u.ID, &u.Name); err != nil {
        return nil, nil // ambiguous: found nothing? errored?
    }
    return &u, nil
}
```

## Good
```go
func FindUser(id int) (*User, error) {
    row := db.QueryRow("SELECT ...", id)
    var u User
    if err := row.Scan(&u.ID, &u.Name); err != nil {
        if errors.Is(err, sql.ErrNoRows) {
            return nil, fmt.Errorf("user %d not found", id)
        }
        return nil, err
    }
    return &u, nil
}
```
</examples>

<patterns>
- `(nil, nil)` return when a resource is not found
- Functions returning `(pointer, error)` where both are nil on empty result
- Repository/cache lookup methods that return nil-nil for missing items
</patterns>

<related>
nilerr, errcheck, govet
