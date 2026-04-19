# revive: redundant-test-main-exit

<instructions>
Detects `os.Exit` calls in `TestMain` functions that are unnecessary. Since Go 1.15, `TestMain` functions no longer need to call `os.Exit` — the testing framework handles exit codes automatically based on whether `m.Run()` reported failures. Calling `os.Exit(0)` is now redundant, and `os.Exit(non-zero)` can skip cleanup.

Remove the `os.Exit` call. Just call `m.Run()` and let the test framework handle the exit status.
</instructions>

<examples>
## Bad
```go
func TestMain(m *testing.M) {
    setup()
    code := m.Run()
    teardown()
    os.Exit(code)
}
```

## Good
```go
func TestMain(m *testing.M) {
    setup()
    defer teardown()
    m.Run()
}
```
</examples>

<patterns>
- `os.Exit(m.Run())` pattern in TestMain functions
- `os.Exit(0)` at the end of TestMain
- Direct `os.Exit` calls that bypass deferred cleanup
- Legacy TestMain implementations from pre-Go 1.15
- TestMain with manual exit code propagation
</patterns>

<related>
redundant-build-tag, redundant-import-alias, deep-exit
