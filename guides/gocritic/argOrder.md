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
- Swap `copy()` arguments to `copy(dst, src)` — destination comes first
- Check `strings.Replace` argument order — use `strings.Replace(s, old, new, n)`
- Reorder same-typed parameters when variable names suggest wrong order
- Swap `append` arguments — use `append(dst, src...)` not `append(src, dst...)`
- Swap `io.Copy` arguments — use `io.Copy(dst, src)` not `io.Copy(src, dst)`
</patterns>

<related>
badCall, evalOrder, dupArg
</related>
