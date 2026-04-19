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
- `wg.Add` called inside the goroutine instead of before the `go` statement
- Loop-based goroutine spawning where `Add` is outside the loop but should be per-iteration
- `WaitGroup` counter managed inconsistently across different launch sites
- Missing `wg.Add` entirely (goroutines launched without tracking)
- `Add` with the wrong count in batch goroutine patterns
</patterns>

<related>
waitgroup-by-value, datarace
