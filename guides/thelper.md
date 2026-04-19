# thelper

<instructions>
Thelper detects test helpers that don't call `t.Helper()`. Without `t.Helper()`, when a helper function fails, the test output points to the helper's internal line instead of the caller's line, making debugging harder.

Add `t.Helper()` at the start of every test helper function that calls `t.Error`, `t.Fatal`, `t.Logf`, etc.
</instructions>

<examples>
## Bad
```go
func setupDB(t *testing.T) *sql.DB {
    db, err := sql.Open("sqlite3", ":memory:")
    if err != nil {
        t.Fatalf("failed to open db: %v", err)
    }
    return db
}
```

## Good
```go
func setupDB(t *testing.T) *sql.DB {
    t.Helper()
    db, err := sql.Open("sqlite3", ":memory:")
    if err != nil {
        t.Fatalf("failed to open db: %v", err)
    }
    return db
}
```
</examples>

<patterns>
- Helper functions using `t.Fatal`/`t.Error` without `t.Helper()`
- Table-driven test helpers that create subtests without `t.Helper()`
- Benchmark helpers missing `b.Helper()` (same issue)
- Helper calling another helper — both need `t.Helper()`
</patterns>

<related>
paralleltest, tparallel, testpackage
</related>
