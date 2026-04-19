# ginkgolinter: have-len-zero

<instructions>
Detects `Expect(x).Should(HaveLen(0))` which can be simplified to `Expect(x).Should(BeEmpty())`. The `BeEmpty` matcher is more semantic and works with any collection, string, or channel. It reads naturally: "the result should be empty" versus "the result should have length zero."
</instructions>

<examples>
## Bad
```go
Expect(results).Should(HaveLen(0))
```

## Good
```go
Expect(results).Should(BeEmpty())
```
</examples>

<patterns>
- `Expect(x).To(HaveLen(0))` — use `Expect(x).To(BeEmpty())`
- `Expect(x).ToNot(HaveLen(0))` — use `Expect(x).ToNot(BeEmpty())`
- `Expect(len(x)).To(Equal(0))` — use `Expect(x).To(BeEmpty())`
</patterns>

<related>
len-assertion, nil-assertion
