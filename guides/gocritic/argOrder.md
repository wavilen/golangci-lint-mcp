# gocritic: argOrder

<instructions>
Detects function calls where arguments of the same type are likely swapped. This commonly occurs with functions that accept multiple parameters of identical types (e.g., `strings.Replace(s, old, new, n)`, `strings.Index(s, substr)`, or `copy(dst, src)`). The checker flags suspicious argument ordering based on heuristics.

Review the flagged call and verify the argument order matches the function signature. Swap arguments if they are reversed.
</instructions>

<examples>
## Bad
```go
// copy expects (dst, src), but arguments look swapped
n := copy(src, dst)
```

## Good
```go
n := copy(dst, src)
```
</examples>

<patterns>
- `copy(src, dst)` where the first argument should be the destination
- `strings.Replace(old, new, ...)` with swapped old/new arguments
- Multi-argument calls with same-typed params where names suggest wrong order
- `append(dst, src...)` vs `append(src, dst...)` confusion
- `io.Copy(src, dst)` instead of `io.Copy(dst, src)`
</patterns>

<related>
badCall, evalOrder, dupArg
</related>
