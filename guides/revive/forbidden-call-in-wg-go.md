# revive: forbidden-call-in-wg-go

<instructions>
Detects `runtime.Goexit` and other disallowed calls inside `sync.WaitGroup` goroutine handlers. These calls prevent `wg.Done()` from being reached, leaving the WaitGroup counter above zero and causing `wg.Wait()` to deadlock. Defer `wg.Done()` as the first statement in the goroutine so it always runs, even on panic.
</instructions>

<examples>
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
- Use `defer wg.Done()` immediately in goroutines to ensure it runs even after `runtime.Goexit()`
- Avoid `os.Exit()` or `log.Fatal()` in WaitGroup goroutines — they prevent `wg.Done()` from running
- Add recovery from panics in WaitGroup goroutines to ensure `wg.Done()` executes
- Move `wg.Done()` before any call that might block forever
- Use `defer wg.Done()` as the first statement when adding error handling paths to goroutines
</patterns>

<related>
revive/datarace, revive/defer
</related>
