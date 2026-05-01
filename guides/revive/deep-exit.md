# revive: deep-exit

<instructions>
Detects `os.Exit` and `log.Fatal` calls outside `main` or top-level handlers. These bypass deferred cleanup functions, skip `runtime` finalizers, and make the code untestable as a library. Return errors up the call stack and let `main` decide how to terminate.
</instructions>

<examples>
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
- Return errors from library functions instead of calling `log.Fatal` or `log.Fatalf`
- Propagate errors up the call stack instead of calling `os.Exit` in helper functions
- Replace `fmt.Fprintf(os.Stderr, ...); os.Exit(1)` patterns with error returns
- Return errors from test helper functions instead of calling `t.FailNow`
- Handle flag parsing errors with error returns rather than exiting directly
</patterns>

<related>
revive/defer, revive/error-return
</related>
