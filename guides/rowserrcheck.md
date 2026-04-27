# rowserrcheck

<instructions>
Rowserrcheck verifies that `database/sql.Rows.Err()` is checked after iterating with `rows.Next()`. If a query fails mid-iteration, `Next()` returns false but the error is only available via `Err()` — silently returning incomplete results.

Always check `rows.Err()` after the `rows.Next()` loop, before processing the results.
</instructions>

<examples>
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
- Check `rows.Err()` immediately after every `rows.Next()` loop
- Add `defer rows.Close()` immediately after acquiring `*sql.Rows` from a query
- Check `rows.Err()` after early `break` from a `rows.Next()` loop
- Check the error from `db.Query`/`db.QueryContext` before iterating rows
</patterns>

<related>
sqlclosecheck, errcheck, bodyclose
</related>
