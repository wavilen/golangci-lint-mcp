# gocritic: valSwap

<instructions>
Detects manual variable swaps that use a temporary variable instead of Go's parallel assignment. Swapping two values with `a, b = b, a` is more concise and idiomatic than the three-line temp-variable pattern.

Use parallel assignment `a, b = b, a` to swap two variables.
</instructions>

<examples>
## Good
```go
a, b = b, a
```
</examples>

<patterns>
- Replace three-line swap `tmp := a; a = b; b = tmp` with parallel assignment `a, b = b, a`
- Use parallel assignment `a, b = b, a` for swaps in sorting or reversal algorithms
</patterns>

<related>
gocritic/assignOp
</related>
