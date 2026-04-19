# gocritic: rangeAppendAll

<instructions>
Detects `append(all, items...)` patterns inside range loops that build a slice by repeatedly appending. This pattern can be replaced with a more efficient approach. More importantly, it catches the common bug of `append`ing from the same slice being ranged over, which modifies the slice during iteration.

Use a separate slice for accumulation, or pre-allocate and copy. Never append from the slice you are ranging over.
</instructions>

<examples>
## Bad
```go
var all []int
for _, chunk := range chunks {
    all = append(all, chunk...) // repeated allocations
}
```

## Good
```go
total := 0
for _, chunk := range chunks {
    total += len(chunk)
}
all := make([]int, 0, total)
for _, chunk := range chunks {
    all = append(all, chunk...)
}
```
</examples>

<patterns>
- `append(dst, src...)` inside range loop over `src`
- Building a result slice by appending from iterated slices
- Appending from the same slice being ranged over (infinite loop risk)
- Growing slices in loops without pre-allocation
</patterns>

<related>
appendAssign, sloppyLen, rangeAppendAll
</related>
