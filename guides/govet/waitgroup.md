# govet: waitgroup

<instructions>
Reports incorrect usage of `sync.WaitGroup`, most commonly passing it by value instead of pointer. When a `WaitGroup` is copied, the counter is duplicated and the copy's `Done()` calls don't affect the original, causing `Wait()` to block forever or return prematurely.

Pass `*sync.WaitGroup` to functions, and call `wg.Add(1)` before launching the goroutine.
</instructions>

<examples>
## Good
```go
func worker(wg *sync.WaitGroup) { // pointer — shared counter
    defer wg.Done()
}
```
</examples>

<patterns>
- Use pointer parameters for functions accepting `sync.WaitGroup` — never value receivers
- Pass `*sync.WaitGroup` to goroutine closures — never copy by value
- Call `wg.Add(1)` before the `go` statement, not inside the goroutine
- Store `*sync.WaitGroup` in structs instead of copying `sync.WaitGroup` by value
</patterns>

<related>
govet/copylocks, govet/atomic, revive/waitgroup-by-value
</related>
