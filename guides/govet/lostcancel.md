# govet: lostcancel

<instructions>
Reports cancellation functions returned by `context.WithCancel`, `context.WithTimeout`, or `context.WithDeadline` that are never called. The cancel function releases resources and signals goroutines to stop. Losing it prevents cleanup and leaks goroutines.

Always `defer cancel()` immediately after creating a cancellable context.
</instructions>

<examples>
## Bad
```go
ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
// cancel function discarded — context never cancelled
doWork(ctx)
```

## Good
```go
ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
defer cancel() // ensure resources are released
doWork(ctx)
```
</examples>

<patterns>
- Discarding cancel function with blank identifier `_`
- Storing cancel in a variable but never calling it
- Conditional cancel that skips the cancel path
</patterns>

<related>
defers, httpresponse
</related>
