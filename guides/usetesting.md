# usetesting

<instructions>
Usetesting detects uses of `os.TempDir()` and `os.MkdirTemp` in tests instead of `t.TempDir()`. The testing method automatically cleans up when the test finishes, preventing temp file accumulation.

Use `t.TempDir()` in tests, which creates a temporary directory and removes it after the test completes.
</instructions>

<examples>
## Bad
```go
func TestParse(t *testing.T) {
    dir := os.MkdirTemp("", "test")
    defer os.RemoveAll(dir)
    // ...
}
```

## Good
```go
func TestParse(t *testing.T) {
    dir := t.TempDir()
    // automatically cleaned up
    // ...
}
```
</examples>

<patterns>
- Replace `os.MkdirTemp` in tests with `t.TempDir()`
- Use `t.TempDir()` + `os.Create` instead of `os.CreateTemp` in tests
- Replace manual `defer os.RemoveAll(dir)` with `t.TempDir()` auto-cleanup
- Replace deprecated `ioutil.TempDir` with `t.TempDir()`
</patterns>

<related>
thelper, testpackage, paralleltest
</related>
