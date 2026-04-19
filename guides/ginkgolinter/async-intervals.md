# ginkgolinter: async-intervals

<instructions>
Detects `Eventually(x).WithPolling(interval)` where the polling interval is too aggressive (e.g., sub-millisecond). Very short polling intervals waste CPU and can cause flaky tests due to resource contention. Use reasonable intervals — typically 10ms minimum, with 100ms being a good default.
</instructions>

<examples>
## Bad
```go
Eventually(func() bool { return ready() }).
    WithPolling(time.Microsecond).
    Should(BeTrue())
```

## Good
```go
Eventually(func() bool { return ready() }).
    WithPolling(100 * time.Millisecond).
    WithTimeout(5 * time.Second).
    Should(BeTrue())
```
</examples>

<patterns>
- `WithPolling(time.Millisecond)` — too aggressive, use 10ms+ minimum
- No explicit polling interval on `Eventually` — add one for clarity
- `WithPolling(0)` — invalid, use a reasonable minimum
</patterns>

<related>
async-assertion, expect-to
