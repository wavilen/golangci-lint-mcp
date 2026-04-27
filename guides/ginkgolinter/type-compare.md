# ginkgolinter: type-compare

<instructions>
Detects `Expect(x).Should(BeAssignableToTypeOf(MyType{}))` when simpler type matchers exist. Use `Expect(x).To(BeAnExistingType())` or use `HaveExistingField` patterns where appropriate. For checking struct types, consider using `BeAssignableToTypeOf` only when no more specific matcher is available.
</instructions>

<examples>
## Good
```go
// When checking interface implementation:
_, ok := result.(MyInterface)
Expect(ok).To(BeTrue())

// When type assertion is needed:
Expect(result).To(BeAssignableToTypeOf(MyType{}))
```
</examples>

<patterns>
- Use `BeAssignableToTypeOf` only when no more specific matcher exists — check alternatives first
- Replace `reflect.TypeOf` comparisons with Gomega matchers for readability
- Use type assertion or `Satisfy` for checking interface implementation
</patterns>

<related>
ginkgolinter/compare-assertion, ginkgolinter/nil-assertion
</related>
