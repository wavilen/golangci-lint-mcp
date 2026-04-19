# revive: waitgroup-by-value

<instructions>
Detects `sync.WaitGroup` passed by value to functions or goroutines. `WaitGroup` contains an internal counter and must not be copied after first use. Passing by value creates a copy with a separate counter, meaning `Add`, `Done`, and `Wait` operate on different instances — the caller's `Wait` never sees the goroutine's `Done`.

Change the function signature to accept `*sync.WaitGroup` (pointer). Always pass the address with `&wg`.
</instructions>

<examples>
## Bad
```go
func worker(wg sync.WaitGroup) {
    defer wg.Done()
    doWork()
}

func main() {
    var wg sync.WaitGroup
    wg.Add(1)
    go worker(wg) // copied — counter not shared
    wg.Wait()     // hangs forever
}
```

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
- Function parameters of type `sync.WaitGroup` (not pointer)
- Goroutine closures capturing a WaitGroup by value
- Structs embedding `sync.WaitGroup` by value and being copied
- Method receivers using value type for structs containing WaitGroup
- Passing WaitGroup to goroutine-launching helpers without a pointer
</patterns>

<related>
use-waitgroup-go, datarace
