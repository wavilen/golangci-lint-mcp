# gocritic: sliceClear

<instructions>
Detects patterns that clear a slice by re-slicing to zero length but keep the old capacity, or by assigning nil. Since Go 1.21, the `clear` builtin efficiently zeroes all elements and releases references. For older versions, `s = s[:0]` is idiomatic for reuse but does not zero elements.

Use `clear(s)` (Go 1.21+) to zero all elements and release references, or `s = s[:0]` when reusing the backing array is intentional.
</instructions>

<examples>
## Good
```go
clear(s)  // Go 1.21+: zeroes elements, keeps length/capacity
```
</examples>

<patterns>
- Use `clear(s)` or `for i := range s { s[i] = zero }` when GC needs zeroed elements — not `s = s[:0]`
- Use `s = s[:0]` or `clear(s)` instead of reassigning to `nil` or `make([]T, 0)`
- Use `clear(s)` when resetting buffers or pools where old data should not be retained
- Replace inefficient per-element clearing in loops with `clear(s)`
</patterns>

<related>
gocritic/zeroByteRepeat, gocritic/appendCombine
</related>
