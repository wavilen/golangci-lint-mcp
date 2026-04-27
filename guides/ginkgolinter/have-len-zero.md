# ginkgolinter: have-len-zero

<instructions>
Detects `Expect(x).Should(HaveLen(0))` which can be simplified to `Expect(x).Should(BeEmpty())`. The `BeEmpty` matcher is more semantic and works with any collection, string, or channel. It reads naturally: "the result should be empty" versus "the result should have length zero."
</instructions>

<examples>
## Good
```go
Expect(results).Should(BeEmpty())
```
</examples>

<patterns>
- Use `Expect(x).To(BeEmpty())` instead of `Expect(x).To(HaveLen(0))`
- Use `Expect(x).ToNot(BeEmpty())` instead of `Expect(x).ToNot(HaveLen(0))`
- Use `Expect(x).To(BeEmpty())` instead of `Expect(len(x)).To(Equal(0))`
</patterns>

<related>
ginkgolinter/len-assertion, ginkgolinter/nil-assertion
</related>
