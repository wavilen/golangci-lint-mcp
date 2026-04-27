# modernize: sliceclear

<instructions>
Detects manual slice clearing patterns like `for i := range s { s[i] = zero }` or `s = s[:0]` where the goal is to zero out or truncate a slice. Use Go 1.21's `clear(s)` builtin or `slices.Delete` to express intent more clearly and avoid verbose loops.
</instructions>

<examples>
## Good
```go
clear(buf)
```
</examples>

<patterns>
- Use `clear(s)` instead of a loop that zeroes every element
- Use `clear(s)` for zeroing instead of `s = s[:0]` truncation without releasing memory
- Use `clear(s)` instead of manually nil-ing slice elements before release
</patterns>

<related>
modernize/stringappend, modernize/slicesort
</related>
