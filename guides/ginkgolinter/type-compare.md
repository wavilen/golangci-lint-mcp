# ginkgolinter: type-compare

<instructions>
Detects `Expect(x).Should(BeAssignableToTypeOf(MyType{}))` when simpler type matchers exist. Use `Expect(x).To(BeAnExistingType())` or use `HaveExistingField` patterns where appropriate. For checking struct types, consider using `BeAssignableToTypeOf` only when no more specific matcher is available.
</instructions>

<examples>
## Bad
```go
Expect(result).To(BeAssignableToTypeOf(MyType{}))
```

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
- `BeAssignableToTypeOf` with zero-value struct — sometimes unavoidable but check alternatives
- Type comparison with `reflect.TypeOf` — use Gomega matchers
- Checking interface implementation — use type assertion or `Satisfy`
</patterns>

<related>
compare-assertion, nil-assertion
