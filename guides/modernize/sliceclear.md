# modernize: sliceclear

<instructions>
Detects manual slice clearing patterns like `for i := range s { s[i] = zero }` or `s = s[:0]` where the goal is to zero out or truncate a slice. Use Go 1.21's `clear(s)` builtin or `slices.Delete` to express intent more clearly and avoid verbose loops.
</instructions>

<examples>
## Bad
```go
for i := range buf {
    buf[i] = 0
}
```

## Good
```go
clear(buf)
```
</examples>

<patterns>
- Loop that zeroes every element — use `clear(s)`
- `s = s[:0]` to truncate without releasing memory — use `clear(s)` if zeroing is also needed
- Manual nil-ing of slice elements before release — use `clear(s)`
</patterns>

<related>
stringappend, slicesort
