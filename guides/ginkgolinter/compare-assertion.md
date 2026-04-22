# ginkgolinter: compare-assertion

<instructions>
Detects `Expect(x == y).Should(BeTrue())` or `Expect(x != y).Should(BeTrue())` in Ginkgo tests. Use `Expect(x).Should(Equal(y))` for equality and `Expect(x).ShouldNot(Equal(y))` for inequality. The `Equal` matcher uses `reflect.DeepEqual` and shows both values on failure.
</instructions>

<examples>
## Bad
```go
Expect(result == expected).Should(BeTrue())
Expect(name != "").Should(BeTrue())
```

## Good
```go
Expect(result).Should(Equal(expected))
Expect(name).ShouldNot(BeEmpty())
```
</examples>

<patterns>
- Use `Expect(a).Should(Equal(b))` instead of `Expect(a == b).Should(BeTrue())`
- Use `Expect(a).ShouldNot(Equal(b))` instead of `Expect(a != b).Should(BeTrue())`
- Use `Expect(a).To(BeNumerically(">", b))` instead of `Expect(a > b).Should(BeTrue())`
</patterns>

<related>
type-compare, nil-assertion
