# gocritic: sqlQuery

<instructions>
Detects `database/sql` query execution where the result rows are not scanned or closed. Executing `db.Query` without iterating and closing the returned `*sql.Rows` leaks database connections. Use `db.Exec` for statements that don't return rows, or ensure `rows.Close()` is always called (typically via `defer`).

Use `db.Exec` for INSERT/UPDATE/DELETE. For `db.Query`, always defer `rows.Close()` and iterate with `rows.Next()`.
</instructions>

<examples>
## Good
```go
_, err := db.Exec("DELETE FROM users WHERE active = false")
```
</examples>

<patterns>
- Use `db.Exec()` for INSERT/UPDATE/DELETE — reserve `db.Query()` for SELECT
- Always iterate `rows.Next()` after `db.Query()` — or use `db.QueryRow()` for single rows
- Add `defer rows.Close()` immediately after `db.Query()`
- Always capture and close the `*sql.Rows` return value from `db.Query()`
</patterns>

<related>
gocritic/uncheckedInlineErr, gocritic/exitAfterDefer
</related>
