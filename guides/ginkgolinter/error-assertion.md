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
- `Expect(err).Should(BeNil())` — use `Expect(err).ShouldNot(HaveOccurred())`
- `Expect(err).ToNot(BeNil())` — use `Expect(err).To(HaveOccurred())`
- Checking error is nil with `BeNil()` — always prefer `HaveOccurred`
</patterns>

<related>
nil-assertion, succeed-matcher
