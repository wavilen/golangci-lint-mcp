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
- Use `Expect(fn()).To(Succeed())` instead of `err := fn(); Expect(err).To(BeNil())` for single-return functions
- Simplify `Expect(err).ShouldNot(HaveOccurred())` for single-return functions to `Expect(fn()).To(Succeed())`
- Simplify two-line error checks for single-return functions into `Expect(fn()).To(Succeed())`
</patterns>

<related>
error-assertion, nil-assertion
