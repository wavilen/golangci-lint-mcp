# tparallel

<instructions>
Tparallel detects incorrect usage of `t.Parallel()` in tests. It checks that if a top-level test calls `t.Parallel()`, its subtests also call `t.Parallel()`, and vice versa — mixed parallel/non-parallel subtests cause confusing test behavior.

Ensure all subtests in a parallel test also call `t.Parallel()`, or none of them do.
</instructions>

<examples>
## Bad
```go
func TestSuite(t *testing.T) {
    t.Parallel()
    t.Run("A", func(t *testing.T) {
        // missing t.Parallel()
    })
    t.Run("B", func(t *testing.T) {
        t.Parallel()
    })
}
```

## Good
```go
func TestSuite(t *testing.T) {
    t.Parallel()
    t.Run("A", func(t *testing.T) {
        t.Parallel()
    })
    t.Run("B", func(t *testing.T) {
        t.Parallel()
    })
}
```
</examples>

<patterns>
- Parent test calls `t.Parallel()` but some subtests don't
- Subtest calls `t.Parallel()` but parent does not
- Only first or last subtest marked parallel in a group
- `t.Parallel()` in helper functions that may or may not run in parallel context
</patterns>

<related>
paralleltest, thelper, testpackage
</related>
