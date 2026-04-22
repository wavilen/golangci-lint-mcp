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
- Store and `defer cancel()` immediately — never discard with `_`
- Call the cancel function before exiting — never store it without invoking
- Ensure `cancel()` runs on all code paths — avoid conditional skips
</patterns>

<related>
defers, httpresponse
</related>
