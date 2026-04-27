# ginkgolinter: nil-assertion

<instructions>
Detects `Expect(x == nil).Should(BeTrue())` or `Expect(x).Should(BeNil())` for error values. When checking errors, use `Expect(err).ShouldNot(HaveOccurred())` for semantic clarity. For non-error nil checks, `Expect(x).To(BeNil())` is acceptable but `BeTrue()` with a nil comparison should be avoided.
</instructions>

<examples>
## Good
```go
Expect(err).ShouldNot(HaveOccurred())
Expect(x).To(BeNil())
```
</examples>

<patterns>
- Use `Expect(x).To(BeNil())` instead of `Expect(x == nil).Should(BeTrue())`
- Use `Expect(x).ToNot(BeNil())` instead of `Expect(x != nil).Should(BeTrue())`
- Use `Expect(err).ShouldNot(HaveOccurred())` instead of `Expect(err).Should(BeNil())`
</patterns>

<related>
ginkgolinter/error-assertion, ginkgolinter/succeed-matcher
</related>
