# nilnil

<instructions>
Detects functions that return `(nil, nil)` — a nil pointer alongside a nil error. This is ambiguous: the caller cannot distinguish "success with no result" from "not found" or "uninitialized." Return a meaningful sentinel error, a zero-value wrapper, or an explicit not-found error instead.
</instructions>

<examples>
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
</related>
