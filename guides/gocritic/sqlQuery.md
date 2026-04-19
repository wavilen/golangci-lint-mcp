# gocritic: sqlQuery

<instructions>
Detects `database/sql` query execution where the result rows are not scanned or closed. Executing `db.Query` without iterating and closing the returned `*sql.Rows` leaks database connections. Use `db.Exec` for statements that don't return rows, or ensure `rows.Close()` is always called (typically via `defer`).

Use `db.Exec` for INSERT/UPDATE/DELETE. For `db.Query`, always defer `rows.Close()` and iterate with `rows.Next()`.
</instructions>

<examples>
## Bad
```go
_, err := db.Query("DELETE FROM users WHERE active = false")
// result rows discarded — use db.Exec instead
```

## Good
```go
_, err := db.Exec("DELETE FROM users WHERE active = false")
```
</examples>

<patterns>
- Using `db.Query()` for INSERT/UPDATE/DELETE statements
- Calling `db.Query()` without iterating `rows.Next()`
- Forgetting `defer rows.Close()` after `db.Query()`
- Discarding the `*sql.Rows` return value
</patterns>

<related>
uncheckedInlineErr, exitAfterDefer
</related>
