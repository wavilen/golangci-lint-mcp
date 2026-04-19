# rowserrcheck

<instructions>
Rowserrcheck verifies that `database/sql.Rows.Err()` is checked after iterating with `rows.Next()`. If a query fails mid-iteration, `Next()` returns false but the error is only available via `Err()` — silently returning incomplete results.

Always check `rows.Err()` after the `rows.Next()` loop, before processing the results.
</instructions>

<examples>
## Bad
```go
rows, _ := db.Query("SELECT id FROM users")
for rows.Next() {
    var id int
    _ = rows.Scan(&id)
}
```

## Good
```go
rows, err := db.Query("SELECT id FROM users")
if err != nil {
    return err
}
defer rows.Close()
for rows.Next() {
    var id int
    if err := rows.Scan(&id); err != nil {
        return err
    }
}
if err := rows.Err(); err != nil {
    return err
}
```
</examples>

<patterns>
- Missing `rows.Err()` check after `rows.Next()` loop
- Missing `rows.Close()` in defer after query
- Early loop breaks without checking `rows.Err()`
- Query errors masked by ignoring the returned error
</patterns>

<related>
sqlclosecheck, errcheck, bodyclose
</related>
