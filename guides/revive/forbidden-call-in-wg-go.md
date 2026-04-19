# revive: forbidden-call-in-wg-go

<instructions>
Detects disallowed function calls inside `sync.WaitGroup` goroutine handlers. Certain calls — like `runtime.Goexit()` — can prevent `wg.Done()` from being reached, causing deadlocks because the WaitGroup counter never returns to zero.

Ensure `wg.Done()` is deferred as the first statement in the goroutine so it always runs, even if the function panics or exits early.
</instructions>

<examples>
## Bad
```go
var wg sync.WaitGroup
wg.Add(1)
go func() {
    runtime.Goexit() // prevents wg.Done() from being reached
    wg.Done()
}()
wg.Wait()
```

## Good
```go
var wg sync.WaitGroup
wg.Add(1)
go func() {
    defer wg.Done()
    doWork()
}()
wg.Wait()
```
</examples>

<patterns>
- `runtime.Goexit()` inside a goroutine without deferred `wg.Done()`
- `os.Exit()` or `log.Fatal()` in a WaitGroup goroutine
- Panic without recovery preventing `wg.Done()` from executing
- `wg.Done()` placed after a call that might block forever
- Forgetting to defer `wg.Done()` when adding error handling paths
</patterns>

<related>
datarace, defer
