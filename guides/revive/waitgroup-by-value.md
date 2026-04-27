# revive: waitgroup-by-value

<instructions>
Detects `sync.WaitGroup` passed by value to functions or goroutines. `WaitGroup` contains an internal counter and must not be copied after first use. Passing by value creates a copy with a separate counter, meaning `Add`, `Done`, and `Wait` operate on different instances — the caller's `Wait` never sees the goroutine's `Done`.

Change the function signature to accept `*sync.WaitGroup` (pointer). Always pass the address with `&wg`.
</instructions>

<examples>
## Good
```go
func worker(wg *sync.WaitGroup) {
    defer wg.Done()
    doWork()
}

func main() {
    var wg sync.WaitGroup
    wg.Add(1)
    go worker(&wg)
    wg.Wait()
}
```
</examples>

<patterns>
- Switch `sync.WaitGroup` function parameters to `*sync.WaitGroup` (pointer)
- Use `*sync.WaitGroup` by pointer in goroutine closures, never by value
- Avoid copying structs that embed `sync.WaitGroup` by value — use pointers
- Use pointer receivers for structs containing a `sync.WaitGroup` field
- Pass `&wg` to goroutine-launching helpers instead of copying by value
</patterns>

<related>
revive/use-waitgroup-go, revive/datarace
</related>
