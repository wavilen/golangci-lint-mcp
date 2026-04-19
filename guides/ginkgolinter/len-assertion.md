# ginkgolinter: len-assertion

<instructions>
Detects `Expect(len(x)).Should(Equal(5))` in Ginkgo tests. This uses `len()` inside `Expect` which produces an unhelpful failure message. Use `Expect(x).Should(HaveLen(5))` instead — the Gomega `HaveLen` matcher works with any collection type and shows both expected and actual length in the failure output.
</instructions>

<examples>
## Bad
```go
Expect(len(items)).Should(Equal(5))
```

## Good
```go
Expect(items).Should(HaveLen(5))
```
</examples>

<patterns>
- `Expect(len(x)).Should(Equal(n))` — use `Expect(x).Should(HaveLen(n))`
- `Expect(len(x)).To(BeNumerically(">", 0))` — use `Expect(x).ToNot(BeEmpty())`
- `Expect(len(x)).To(Equal(0))` — use `Expect(x).To(BeEmpty())`
</patterns>

<related>
have-len-zero, nil-assertion
