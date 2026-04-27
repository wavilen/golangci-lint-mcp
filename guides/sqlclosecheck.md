# sqlclosecheck

<instructions>
Sqlclosecheck detects missing `Close()` calls on `database/sql` types: `Rows`, `Stmt`, and `NamedStmt`. Unclosed database resources leak connections and exhaust the connection pool.

Always defer `Close()` immediately after successfully opening rows, statements, or named statements.
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
    // ...
}
```
</examples>

<patterns>
- Add `defer rows.Close()` immediately after `db.Query` or `db.QueryContext`
- Add `defer stmt.Close()` immediately after `db.Prepare` or `db.PrepareContext`
- Close the previous `*sql.Rows` before reassigning the variable to a new query result
- Avoid returning `*sql.Rows` from functions — process data inside and close rows locally
</patterns>

<related>
rowserrcheck, bodyclose, errcheck
</related>
