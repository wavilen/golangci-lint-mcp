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
- Add `t.Parallel()` to all subtests when the parent test is parallel
- Call `t.Parallel()` in the parent test if any subtest uses `t.Parallel()`
- Mark all subtests in a group as parallel — avoid mixing parallel and sequential
- Move `t.Parallel()` out of helper functions that may run in non-parallel contexts
</patterns>

<related>
paralleltest, thelper, testpackage
</related>
