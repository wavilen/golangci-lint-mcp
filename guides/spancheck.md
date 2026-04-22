# spancheck

<instructions>
Spancheck verifies that OpenTelemetry spans are completed properly. Every span created via `tracer.Start()` must have `span.End()` called, typically via `defer`. Missing `End()` calls leak spans and produce incomplete traces.

Always call `span.End()` immediately after creating the span, usually with `defer`.
</instructions>

<examples>
## Bad
```go
ctx, span := tracer.Start(ctx, "operation")
result := doWork(ctx)
return result
```

## Good
```go
ctx, span := tracer.Start(ctx, "operation")
defer span.End()
result := doWork(ctx)
return result
```
</examples>

<patterns>
- Call `defer span.End()` immediately after `tracer.Start()`
- Call `span.End()` via `defer` to ensure it runs on all code paths, including error returns
- Set span status and attributes before calling `span.End()`
- Call `span.RecordError(err)` before `span.End()` when an error occurs
</patterns>

<related>
errcheck, contextcheck
</related>
