# fatcontext

<instructions>
Fatcontext detects contexts that get reassigned inside loops, leading to accumulated values (deadlines, cancel functions, values) that grow each iteration. This causes unexpected timeouts and memory growth.

Assign the derived context to a new variable instead of shadowing the loop variable.
</instructions>

<examples>
## Bad
```go
for _, item := range items {
    ctx, cancel := context.WithTimeout(ctx, time.Second)
    defer cancel()
    process(ctx, item)
}
```

## Good
```go
for _, item := range items {
    taskCtx, cancel := context.WithTimeout(ctx, time.Second)
    defer cancel()
    process(taskCtx, item)
}
```
</examples>

<patterns>
- Reassigning `ctx` inside a for loop with WithTimeout/WithDeadline
- Shadowing the outer ctx variable in a range loop
- Accumulating context values by reusing the same ctx variable name
- Using `ctx = context.WithValue(ctx, ...)` in a loop
</patterns>

<related>
contextcheck, noctx, spancheck
