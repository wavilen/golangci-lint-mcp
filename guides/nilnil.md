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
- Return a sentinel error or `ErrNotFound` instead of `(nil, nil)` when a resource is missing
- Return a descriptive error instead of `(nil, nil)` for empty or not-found results
- Define a clear not-found error and return it from repository/cache lookups instead of nil-nil
</patterns>

<related>
nilerr, errcheck, govet
