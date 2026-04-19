# gocritic: exitAfterDefer

<instructions>
Detects calls to `os.Exit` or `log.Fatal` inside functions that contain deferred statements. `os.Exit` terminates the program immediately without running deferred functions, so any cleanup in `defer` statements is silently skipped.

Return an error instead of calling `os.Exit`, and let the caller decide how to handle the failure. If immediate exit is truly needed, move cleanup before the exit call.
</instructions>

<examples>
## Bad
```go
func process(path string) {
    f, err := os.Open(path)
    if err != nil {
        log.Fatal(err) // defers won't run
    }
    defer f.Close()
    // ...
}
```

## Good
```go
func process(path string) error {
    f, err := os.Open(path)
    if err != nil {
        return errors.Wrap(err, "opening file")
    }
    defer f.Close()
    // ...
    return nil
}
```
</examples>

<patterns>
- `log.Fatal` or `log.Fatalf` after defer statements
- `os.Exit` in functions with deferred cleanup
- `log.Fatalln` in request handlers with deferred response writes
- Calling `runtime.Goexit` followed by exit
</patterns>

<related>
deferInLoop, unnecessaryDefer, badCall
</related>
