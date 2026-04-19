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
- `Expect(x).To(Equal(y))` — use `Expect(x).Should(Equal(y))`
- `Expect(x).ToNot(matcher)` — use `Expect(x).ShouldNot(matcher)`
- Mixing `.To` and `.Should` — standardize on `.Should`/`.ShouldNot`
</patterns>

<related>
async-assertion, compare-assertion
