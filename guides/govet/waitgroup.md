# govet: waitgroup

<instructions>
Reports incorrect usage of `sync.WaitGroup`, most commonly passing it by value instead of pointer. When a `WaitGroup` is copied, the counter is duplicated and the copy's `Done()` calls don't affect the original, causing `Wait()` to block forever or return prematurely.

Pass `*sync.WaitGroup` to functions, and call `wg.Add(1)` before launching the goroutine.
</instructions>

<examples>
## Bad
```go
func worker(wg sync.WaitGroup) { // passed by value — copy of counter
    defer wg.Done()
}
```

## Good
```go
func worker(wg *sync.WaitGroup) { // pointer — shared counter
    defer wg.Done()
}
```
</examples>

<patterns>
- Value receiver on function accepting `sync.WaitGroup`
- Passing `sync.WaitGroup` by value to goroutine closure
- `wg.Add(1)` called inside goroutine instead of before `go` statement
- Copying `sync.WaitGroup` into struct by value
</patterns>

<related>
copylocks, atomic
</related>
