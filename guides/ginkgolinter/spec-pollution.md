# ginkgolinter: spec-pollution

<instructions>
Detects mutable state shared across Ginkgo specs that can cause test pollution. Modifying package-level variables or shared struct fields in `BeforeEach`/`It` without resetting in `AfterEach` causes later specs to see dirty state. Use local variables inside `It` blocks or ensure proper cleanup in `AfterEach`.
</instructions>

<examples>
## Good
```go
var _ = Describe("cache", func() {
    var cache *Cache

    BeforeEach(func() {
        cache = NewCache() // fresh instance per spec
    })
    It("is empty initially", func() {
        Expect(cache.Get("key")).To(BeEmpty())
    })
})
```
</examples>

<patterns>
- Use local `var` in `Describe` instead of package-level `var` modified in `BeforeEach`
- Add cleanup in `AfterEach` for any shared state modified during specs
- Separate each spec — avoid `BeforeEach` depending on previous spec's state
</patterns>

<related>
ginkgolinter/focus-container, ginkgolinter/nil-assertion
</related>
