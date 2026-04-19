# revive: range-val-in-closure

<instructions>
Detects range loop variables captured by reference in closures launched inside the loop. Before Go 1.22, all iterations shared the same variable, so goroutines or deferred closures would see only the last value. Even with Go 1.22+, capturing by reference can be confusing.

Create a local copy of the loop variable inside the loop body, or pass it as a parameter to the goroutine function.
</instructions>

<examples>
## Bad
```go
for _, item := range items {
    go func() {
        process(item) // captures loop variable
    }()
}
```

## Good
```go
for _, item := range items {
    item := item // shadow with local copy
    go func() {
        process(item)
    }()
}
```
</examples>

<patterns>
- Goroutines launched inside range loops that close over the iteration variable
- Deferred function calls inside range loops referencing loop variables
- Closures stored in a slice inside a range loop
- Anonymous functions passed to `sync.Once` or similar inside loops
- WaitGroup callbacks capturing range variables
</patterns>

<related>
range-val-address, datarace
