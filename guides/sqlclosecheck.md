# sqlclosecheck

<instructions>
Sqlclosecheck detects missing `Close()` calls on `database/sql` types: `Rows`, `Stmt`, and `NamedStmt`. Unclosed database resources leak connections and exhaust the connection pool.

Always defer `Close()` immediately after successfully opening rows, statements, or named statements.
</instructions>

<examples>
## Bad
```go
rows, err := db.Query("SELECT id FROM users")
if err != nil {
    return err
}
for rows.Next() {
    // ...
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
    // ...
}
```
</examples>

<patterns>
- `db.Query`/`db.QueryContext` without deferred `rows.Close()`
- `db.Prepare` without deferred `stmt.Close()`
- Reassigning rows variable before closing previous result
- Returning rows from function (transfers close responsibility)
</patterns>

<related>
rowserrcheck, bodyclose, errcheck
</related>
