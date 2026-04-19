# ginkgolinter: nil-assertion

<instructions>
Detects `Expect(x == nil).Should(BeTrue())` or `Expect(x).Should(BeNil())` for error values. When checking errors, use `Expect(err).ShouldNot(HaveOccurred())` for semantic clarity. For non-error nil checks, `Expect(x).To(BeNil())` is acceptable but `BeTrue()` with a nil comparison should be avoided.
</instructions>

<examples>
## Bad
```go
Expect(err == nil).Should(BeTrue())
Expect(x).Should(BeNil())
```

## Good
```go
Expect(err).ShouldNot(HaveOccurred())
Expect(x).To(BeNil())
```
</examples>

<patterns>
- `Expect(x == nil).Should(BeTrue())` — use `Expect(x).To(BeNil())`
- `Expect(x != nil).Should(BeTrue())` — use `Expect(x).ToNot(BeNil())`
- `Expect(err).Should(BeNil())` — use `Expect(err).ShouldNot(HaveOccurred())`
</patterns>

<related>
error-assertion, succeed-matcher
