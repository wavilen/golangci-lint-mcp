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
- Move `defer` out of loops into a helper function to prevent deferred calls from accumulating
- Extract loop body into a separate function when defer captures a loop variable by reference
- Call cleanup directly instead of deferring in hot-path functions
- Use direct cleanup in `main` where defer adds unnecessary overhead
- Ensure deferred functions that modify named return values use explicit returns
</patterns>

<related>
deep-exit, gocritic/deferInLoop
