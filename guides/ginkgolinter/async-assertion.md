# ginkgolinter: async-assertion

<instructions>
Detects `Eventually` assertions without explicit timeout or with missing `Should`/`ShouldNot`. Use `Eventually(x).WithTimeout(duration).Should(matcher)` to specify a timeout. Without an explicit timeout, Gomega uses its default (1s), which may be too short for slow operations or too long for fast tests.
</instructions>

<examples>
## Bad
```go
Eventually(func() int { return counter.Load() }).Should(Equal(10))
```

## Good
```go
Eventually(func() int { return counter.Load() }).
    WithTimeout(5 * time.Second).
    WithPolling(100 * time.Millisecond).
    Should(Equal(10))
```
</examples>

<patterns>
- `Eventually(fn).Should(...)` without timeout — add `WithTimeout`
- Missing `WithPolling` interval — add explicit polling interval
- `Consistently` without timeout — add `WithTimeout` for clarity
</patterns>

<related>
async-intervals, expect-to
