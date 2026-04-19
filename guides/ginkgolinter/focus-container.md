# ginkgolinter: focus-container

<instructions>
Detects focused Ginkgo containers (`FDescribe`, `FContext`, `FIt`, `FWhen`) left in test code. Focus markers cause only the focused tests to run, skipping all other tests. This is useful during development but must be removed before committing. CI pipelines should detect and reject focused containers.
</instructions>

<examples>
## Bad
```go
var _ = Describe("parser", func() {
    FIt("handles edge case", func() {  // skips all other tests!
        Expect(Parse("")).To(HaveOccurred())
    })
})
```

## Good
```go
var _ = Describe("parser", func() {
    It("handles edge case", func() {
        Expect(Parse("")).To(HaveOccurred())
    })
})
```
</examples>

<patterns>
- `FDescribe`, `FContext`, `FIt`, `FWhen` — remove the `F` prefix
- `FMeasure` — remove the `F` prefix
- Leftover debug focus markers — always remove before committing
</patterns>

<related>
spec-pollution, expect-to
