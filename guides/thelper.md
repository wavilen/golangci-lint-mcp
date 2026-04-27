# thelper

<instructions>
Thelper detects test helpers that don't call `t.Helper()`. Without `t.Helper()`, when a helper function fails, the test output points to the helper's internal line instead of the caller's line, making debugging harder.

Add `t.Helper()` at the start of every test helper function that calls `t.Error`, `t.Fatal`, `t.Logf`, etc.
</instructions>

<examples>
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
- Add `t.Helper()` at the start of test helpers that call `t.Fatal`/`t.Error`
- Call `t.Helper()` in table-driven test helpers that create subtests
- Add `b.Helper()` in benchmark helpers that report failures
- Ensure `t.Helper()` is called in every helper in a chain — both caller and callee need it
</patterns>

<related>
paralleltest, tparallel, testpackage, usetesting, testifylint/suite-thelper
</related>
