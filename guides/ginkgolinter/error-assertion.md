# ginkgolinter: error-assertion

<instructions>
Detects `Expect(err).Should(BeNil())` instead of the idiomatic `Expect(err).ShouldNot(HaveOccurred())`. Using `HaveOccurred()` is the standard Gomega pattern for error checking and provides better failure messages that explicitly say "Expected no error but got ..." instead of just "Expected nil".
</instructions>

<examples>
## Bad
```go
Expect(err).Should(BeNil())
```

## Good
```go
Expect(err).ShouldNot(HaveOccurred())
```
</examples>

<patterns>
- Use `Expect(err).ShouldNot(HaveOccurred())` instead of `Expect(err).Should(BeNil())`
- Use `Expect(err).To(HaveOccurred())` instead of `Expect(err).ToNot(BeNil())`
- Prefer `HaveOccurred` over `BeNil()` for all error value assertions
</patterns>

<related>
nil-assertion, succeed-matcher
