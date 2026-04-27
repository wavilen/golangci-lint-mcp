# ginkgolinter: async-assertion

<instructions>
Detects `Eventually` assertions without explicit timeout or with missing `Should`/`ShouldNot`. Use `Eventually(x).WithTimeout(duration).Should(matcher)` to specify a timeout. Without an explicit timeout, Gomega uses its default (1s), which may be too short for slow operations or too long for fast tests.
</instructions>

<examples>
## Good
```go
Eventually(func() int { return counter.Load() }).
    WithTimeout(5 * time.Second).
    WithPolling(100 * time.Millisecond).
    Should(Equal(10))
```
</examples>

<patterns>
- Add `WithTimeout` to `Eventually(fn).Should(...)` calls that lack an explicit timeout
- Add explicit `WithPolling` interval to `Eventually` assertions for predictable timing
- Add `WithTimeout` to `Consistently` assertions for clarity on maximum duration
</patterns>

<related>
ginkgolinter/async-intervals, ginkgolinter/expect-to
</related>
