# gocritic: deferInLoop

<instructions>
Detects `defer` statements inside loops. Deferred functions only execute when the surrounding function returns, not at the end of each loop iteration. In a loop, this causes resources to accumulate until the function exits, leading to resource leaks or exhaustion.

Move the loop body into a separate function so `defer` runs at the end of each iteration, or use explicit cleanup without `defer`.
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
        defer f.Close() // all files close only when function returns
        process(f)
    }
    return nil
}
```

## Good
```go
func processFiles(paths []string) error {
    for _, p := range paths {
        if err := processOne(p); err != nil {
            return err
        }
    }
    return nil
}

func processOne(path string) error {
    f, err := os.Open(path)
    if err != nil {
        return err
    }
    defer f.Close()
    process(f)
    return nil
}
```
</examples>

<patterns>
- Move `defer f.Close()` outside loops — wrap loop body in a helper function
- Move `defer mu.Unlock()` outside loop iterations — use a helper or anonymous function
- Move `defer rows.Close()` into a per-iteration helper function for database loops
- Wrap deferred cleanup in a per-iteration function to avoid accumulation in `range`/`for` blocks
</patterns>

<related>
unnecessaryDefer, exitAfterDefer, badLock
</related>
