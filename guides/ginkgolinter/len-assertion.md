# ginkgolinter: len-assertion

<instructions>
Detects `Expect(len(x)).Should(Equal(5))` in Ginkgo tests. This uses `len()` inside `Expect` which produces an unhelpful failure message. Use `Expect(x).Should(HaveLen(5))` instead — the Gomega `HaveLen` matcher works with any collection type and shows both expected and actual length in the failure output.
</instructions>

<examples>
## Good
```go
Expect(items).Should(HaveLen(5))
```
</examples>

<patterns>
- Use `Expect(x).Should(HaveLen(n))` instead of `Expect(len(x)).Should(Equal(n))`
- Use `Expect(x).ToNot(BeEmpty())` instead of `Expect(len(x)).To(BeNumerically(">", 0))`
- Use `Expect(x).To(BeEmpty())` instead of `Expect(len(x)).To(Equal(0))`
</patterns>

<related>
ginkgolinter/have-len-zero, ginkgolinter/nil-assertion
</related>
