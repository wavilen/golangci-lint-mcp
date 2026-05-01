# gocritic: exitAfterDefer

<instructions>
Detects `os.Exit` or `log.Fatal` calls in functions that contain deferred statements. `os.Exit` terminates immediately without running deferred cleanup — file handles leak, transactions remain open, and temporary files persist. Return an error instead and let the caller handle cleanup and exit.
</instructions>

<examples>
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
- Replace `log.Fatal` after defer with `log.Printf` + `return` to allow deferred cleanup
- Replace `os.Exit` in functions with deferred cleanup — return an error instead
- Replace `log.Fatalln` in request handlers with error return to execute deferred writes
- Avoid `runtime.Goexit` followed by `os.Exit` — use error propagation
</patterns>

<related>
gocritic/deferInLoop, gocritic/unnecessaryDefer, gocritic/badCall
</related>
