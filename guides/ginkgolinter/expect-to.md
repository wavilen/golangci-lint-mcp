# ginkgolinter: expect-to

<instructions>
Detects inconsistent use of `.To()` versus `.Should()` in Ginkgo assertions. Ginkgo style prefers `.Should()`/`.ShouldNot()` for consistency. Mixing `.To()` and `.Should()` in the same test suite makes code harder to read. Pick one convention and use it throughout.
</instructions>

<examples>
## Bad
```go
Expect(result).To(Equal(42))
Expect(err).ToNot(HaveOccurred())
```

## Good
```go
Expect(result).Should(Equal(42))
Expect(err).ShouldNot(HaveOccurred())
```
</examples>

<patterns>
- Use `Expect(x).Should(Equal(y))` instead of `Expect(x).To(Equal(y))`
- Use `Expect(x).ShouldNot(matcher)` instead of `Expect(x).ToNot(matcher)`
- Use `.Should`/`.ShouldNot` consistently — avoid mixing with `.To`/`.ToNot`
</patterns>

<related>
async-assertion, compare-assertion
