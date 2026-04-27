# ginkgolinter: focus-container

<instructions>
Detects focused Ginkgo containers (`FDescribe`, `FContext`, `FIt`, `FWhen`) left in test code. Focus markers cause only the focused tests to run, skipping all other tests. This is useful during development but must be removed before committing. CI pipelines should detect and reject focused containers.
</instructions>

<examples>
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
- Remove the `F` prefix from `FDescribe`, `FContext`, `FIt`, `FWhen` focus containers
- Remove the `F` prefix from `FMeasure` before committing
- Remove all debug focus markers before committing — CI should reject focused tests
</patterns>

<related>
ginkgolinter/spec-pollution, ginkgolinter/expect-to
</related>
