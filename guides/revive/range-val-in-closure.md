# revive: range-val-in-closure

<instructions>
Detects range loop variables captured by reference in closures launched inside the loop. Before Go 1.22, all iterations shared the same variable, so goroutines or deferred closures would see only the last value. Even with Go 1.22+, capturing by reference can be confusing.

Create a local copy of the loop variable inside the loop body, or pass it as a parameter to the goroutine function.
</instructions>

<examples>
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
- Use the iteration variable before launching goroutines inside range loops — use `item := item` or pass as parameter
- Use a local copy of loop variables in deferred function calls inside range loops
- Use loop values by copy before storing closures in a slice inside a range loop
- Pass loop variables as arguments to anonymous functions instead of closing over them
- Use range variables before passing to WaitGroup callbacks
</patterns>

<related>
revive/range-val-address, revive/datarace
</related>
