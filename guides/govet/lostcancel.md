# govet: lostcancel

<instructions>
Reports cancellation functions from `context.WithCancel`, `context.WithTimeout`, or `context.WithDeadline` that are never called. Without calling `cancel()`, context resources are never released and child goroutines have no stop signal — they leak indefinitely. Always `defer cancel()` immediately after creating a cancellable context.
</instructions>

<examples>
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
govet/defers, govet/httpresponse, contextcheck
</related>
