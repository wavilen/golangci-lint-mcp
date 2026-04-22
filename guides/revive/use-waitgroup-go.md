# revive: use-waitgroup-go

<instructions>
Detects `sync.WaitGroup.Add` calls placed outside the goroutine launch site where they could race with `WaitGroup.Wait`. The `Add` call should happen before the `go` statement (or inside the goroutine with proper synchronization) to ensure the counter is incremented before `Wait` can observe zero.

Call `wg.Add(1)` immediately before the `go` statement that launches the goroutine, not inside the goroutine itself.
</instructions>

<examples>
## Bad
```go
for _, item := range items {
    go func() {
        wg.Add(1) // too late — race with Wait
        defer wg.Done()
        process(item)
    }()
}
wg.Wait()
```

## Good
```go
for _, item := range items {
    wg.Add(1)
    go func() {
        defer wg.Done()
        process(item)
    }()
}
wg.Wait()
```
</examples>

<patterns>
- Call `wg.Add(1)` before the `go` statement, not inside the goroutine
- Move `wg.Add(1)` inside the loop body per-iteration rather than outside for loop-based spawning
- Ensure the WaitGroup counter is incremented consistently at every launch site
- Add `wg.Add` before every goroutine launch — never omit it
- Set the correct count in `wg.Add` for batch goroutine patterns
</patterns>

<related>
waitgroup-by-value, datarace
