# gocritic: rangeExprCopy

<instructions>
Detects `range` expressions where a large value type is copied for every iteration. When ranging over a function call that returns a large struct, Go copies the result once for the range expression. If the result is an array or large struct, this copy can be expensive. Take the address before ranging or use a pointer.

Store the expression in a variable and range over its address, or change the range expression to return a pointer.
</instructions>

<examples>
## Good
```go
arr := getArray()
for i, v := range &arr {
    _ = i
    _ = v
}
```
</examples>

<patterns>
- Use a pointer to large arrays returned by functions before ranging — avoid copying the value
- Replace `for i, v := range someFunc()` with a pointer when `someFunc` returns ≥80 bytes
- Avoid ranging over large stack-allocated value types — use a pointer instead
- Replace value-type collection iteration with pointer-based iteration for large types
</patterns>

<related>
gocritic/rangeValCopy, gocritic/hugeParam, gocritic/appendCombine
</related>
