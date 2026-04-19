# gocritic: sliceClear

<instructions>
Detects patterns that clear a slice by re-slicing to zero length but keep the old capacity, or by assigning nil. Since Go 1.21, the `clear` builtin efficiently zeroes all elements and releases references. For older versions, `s = s[:0]` is idiomatic for reuse but does not zero elements.

Use `clear(s)` (Go 1.21+) to zero all elements and release references, or `s = s[:0]` when reusing the backing array is intentional.
</instructions>

<examples>
## Bad
```go
s = s[:0]          // elements still referenced, not zeroed
s = nil             // loses capacity
s = make([]int, 0)  // discards backing array entirely
```

## Good
```go
clear(s)  // Go 1.21+: zeroes elements, keeps length/capacity
```
</examples>

<patterns>
- Clearing slices with `s = s[:0]` when element zeroing is needed for GC
- Reassigning slices to `nil` or `make([]T, 0)` to reset them
- Resetting buffers or pools where old data should not be retained
- Loop iterations that "clear" a slice inefficiently between passes
</patterns>

<related>
zeroByteRepeat, appendCombine, makeLen
