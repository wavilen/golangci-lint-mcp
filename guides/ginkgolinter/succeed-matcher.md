# ginkgolinter: succeed-matcher

<instructions>
Detects `Expect(err).Should(BeNil())` when calling a function that returns only an error. Use `Expect(fn()).Should(Succeed())` for cleaner single-return error checking. The `Succeed()` matcher is specifically designed for functions that return only an error, making the intent clear.
</instructions>

<examples>
## Bad
```go
err := os.Remove(tmpFile)
Expect(err).Should(BeNil())
```

## Good
```go
Expect(os.Remove(tmpFile)).Should(Succeed())
```
</examples>

<patterns>
- `err := fn(); Expect(err).To(BeNil())` where fn returns only error — use `Expect(fn()).To(Succeed())`
- `Expect(err).ShouldNot(HaveOccurred())` for single-return — can simplify to `Expect(fn()).To(Succeed())`
- Two-line error check for single-return function — collapse to one line
</patterns>

<related>
error-assertion, nil-assertion
