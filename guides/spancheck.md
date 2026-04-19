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
- `tracer.Start()` without corresponding `span.End()`
- `span.End()` not deferred (missed on error paths)
- Setting span status after `span.End()` has no effect
- Missing error recording: `span.RecordError(err)` before `End()`
</patterns>

<related>
errcheck, contextcheck
</related>
