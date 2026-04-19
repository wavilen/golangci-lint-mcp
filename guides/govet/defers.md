# govet: defers

<instructions>
Reports `defer` calls inside loops. Each defer accumulates until the surrounding function returns, so deferring in a loop causes all deferred calls to pile up and run only at function exit — potentially exhausting resources.

Move the loop body into a separate function so each `defer` runs at the end of its iteration, or use explicit cleanup without `defer`.
</instructions>

<examples>
## Bad
```go
for _, path := range paths {
    f, err := os.Open(path)
    if err != nil {
        return err
    }
    defer f.Close() // all files close only when function returns
    process(f)
}
```

## Good
```go
for _, path := range paths {
    if err := processFile(path); err != nil {
        return err
    }
}

func processFile(path string) error {
    f, err := os.Open(path)
    if err != nil {
        return err
    }
    defer f.Close() // closes after each call
    return process(f)
}
```
</examples>

<patterns>
- `defer` inside `for` loop body
- `defer` inside `range` loop
- Deferred resource cleanup accumulating across iterations
</patterns>

<related>
lostcancel, loopclosure
</related>
