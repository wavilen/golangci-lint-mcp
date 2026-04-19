# revive: deep-exit

<instructions>
Detects calls to `os.Exit`, `log.Fatal`, and similar functions deep in the call stack. These bypass deferred functions, skip cleanup, and make the program impossible to test or reuse as a library. Only `main` should decide to terminate the process.

Return errors up the call stack instead. Let `main` or the top-level handler decide how to exit.
</instructions>

<examples>
## Bad
```go
func processFile(path string) {
    data, err := os.ReadFile(path)
    if err != nil {
        log.Fatalf("failed to read %s: %v", path, err) // kills the process
    }
    handle(data)
}
```

## Good
```go
func processFile(path string) error {
    data, err := os.ReadFile(path)
    if err != nil {
        return errors.Wrap(err, "reading file")
    }
    return handle(data)
}
```
</examples>

<patterns>
- `log.Fatal` or `log.Fatalf` in library or utility functions
- `os.Exit` in helper functions instead of returning errors
- `fmt.Fprintf(os.Stderr, ...); os.Exit(1)` patterns in packages other than main
- Test helper functions calling `t.FailNow` when they should return errors
- CLI flag parsing logic that exits on unknown flags
</patterns>

<related>
defer, error-return
