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
- Assign derived contexts to a new variable instead of reassigning `ctx` in loops
- Use `taskCtx` or `loopCtx` to avoid shadowing the outer `ctx` variable
- Avoid accumulating context values by reusing the same variable name across iterations
- Set derived contexts on a new variable: `loopCtx := context.WithValue(ctx, ...)`
</patterns>

<related>
contextcheck, noctx, spancheck
