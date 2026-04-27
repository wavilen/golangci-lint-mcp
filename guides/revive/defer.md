# revive: defer

<instructions>
Detects common defer misuse: deferring in loops (defers accumulate until the function returns, consuming resources), deferring calls that reference loop variables (captured by reference), and deferring in hot paths where the overhead matters.

Move defer out of loops — call a helper function or manually manage cleanup. When defer must be used, be aware of its scope and resource implications.
</instructions>

<examples>
## Good
```go
func processFiles(paths []string) error {
    for _, p := range paths {
        err := processFile(p)
        if err != nil {
            return err
        }
    }
    return nil
}

func processFile(path string) error {
    f, err := os.Open(path)
    if err != nil {
        return fmt.Errorf("opening file: %w", err)
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
revive/deep-exit, gocritic/deferInLoop
</related>
