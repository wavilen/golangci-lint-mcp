# gocritic: unnecessaryDefer

<instructions>
Detects `defer` statements in positions where they provide no benefit — specifically, `defer` as the last statement before a function return. When `defer` is the final statement, it runs immediately before the function returns anyway, so it's equivalent to an inline call but with unnecessary overhead.

Remove the `defer` keyword and call the function directly. `defer` is useful when cleanup must run regardless of which return path is taken, not at the end of a linear function.
</instructions>

<examples>
## Bad
```go
func process() error {
    f, err := os.Open("data.txt")
    if err != nil {
        return err
    }
    defer f.Close() // fine — may return before end

    err = scan(f)
    defer f.Close() // unnecessary — this is the last statement
    return err
}
```

## Good
```go
func process() error {
    f, err := os.Open("data.txt")
    if err != nil {
        return err
    }
    defer f.Close()

    err = scan(f)
    return err
}
```
</examples>

<patterns>
- `defer` as the last statement before `return`
- `defer f.Close()` at the end of a function with no branching
- Redundant `defer` after all error checks have passed
- Multiple `defer f.Close()` calls for the same resource
</patterns>

<related>
deferInLoop, exitAfterDefer, badLock
</related>
