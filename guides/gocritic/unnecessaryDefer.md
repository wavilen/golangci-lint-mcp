# gocritic: unnecessaryDefer

<instructions>
Detects `defer` statements in positions where they provide no benefit — specifically, `defer` as the last statement before a function return. When `defer` is the final statement, it runs immediately before the function returns anyway, so it's equivalent to an inline call but with unnecessary overhead.

Remove the `defer` keyword and call the function directly. `defer` is useful when cleanup must run regardless of which return path is taken, not at the end of a linear function.
</instructions>

<examples>
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
- Remove `defer` as the last statement before `return` — call directly
- Remove `defer f.Close()` at the end of a function with no branching — call directly
- Remove redundant `defer` after all error checks — the resource is about to go out of scope
- Remove duplicate `defer f.Close()` calls — only one is needed per resource
</patterns>

<related>
gocritic/deferInLoop, gocritic/exitAfterDefer, gocritic/badLock
</related>
