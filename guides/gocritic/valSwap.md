# gocritic: valSwap

<instructions>
Detects manual variable swaps that use a temporary variable instead of Go's parallel assignment. Swapping two values with `a, b = b, a` is more concise and idiomatic than the three-line temp-variable pattern.

Use parallel assignment `a, b = b, a` to swap two variables.
</instructions>

<examples>
## Bad
```go
tmp := a
a = b
b = tmp
```

## Good
```go
a, b = b, a
```
</examples>

<patterns>
- Three-line swap with `tmp := a; a = b; b = tmp`
- Swap in sorting or reversal algorithms
- Any pair exchange that could use parallel assignment
</patterns>

<related>
assignOp
</related>
