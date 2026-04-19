# revive: defer

<instructions>
Detects common defer misuse: deferring in loops (defers accumulate until the function returns, consuming resources), deferring calls that reference loop variables (captured by reference), and deferring in hot paths where the overhead matters.

Move defer out of loops — call a helper function or manually manage cleanup. When defer must be used, be aware of its scope and resource implications.
</instructions>

<examples>
## Bad
```go
func processFiles(paths []string) error {
    for _, p := range paths {
        f, err := os.Open(p)
        if err != nil {
            return err
        }
        defer f.Close() // defers accumulate until function returns
        // process f...
    }
    return nil
}
```

## Good
```go
func processFiles(paths []string) error {
    for _, p := range paths {
        if err := processFile(p); err != nil {
            return err
        }
    }
    return nil
}

func processFile(path string) error {
    f, err := os.Open(path)
    if err != nil {
        return err
    }
    defer f.Close()
    // process f...
    return nil
}
```
</examples>

<patterns>
- `defer` inside `for` loops accumulating deferred calls
- Defer capturing loop variable by reference
- Deferred close/mutex unlock in hot-path functions
- Defer in `main` where direct cleanup would suffice
- Named return values interacting unexpectedly with deferred functions
</patterns>

<related>
deep-exit, gocritic/deferInLoop
