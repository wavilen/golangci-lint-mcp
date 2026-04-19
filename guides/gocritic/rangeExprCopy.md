# gocritic: rangeExprCopy

<instructions>
Detects `range` expressions where a large value type is copied for every iteration. When ranging over a function call that returns a large struct, Go copies the result once for the range expression. If the result is an array or large struct, this copy can be expensive. Take the address before ranging or use a pointer.

Store the expression in a variable and range over its address, or change the range expression to return a pointer.
</instructions>

<examples>
## Bad
```go
type BigArray [4096]int

func getArray() BigArray { return BigArray{} }

for i, v := range getArray() {
    // getArray() result is copied once for range expression (32KB)
    _ = i
    _ = v
}
```

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
- Ranging over function calls that return large arrays or structs by value
- `for i, v := range someFunc()` where `someFunc` returns a type ≥80 bytes
- Ranging over large stack-allocated arrays that get copied into the range expression
- Functions returning value-type collections iterated via `range`
</patterns>

<related>
rangeValCopy, hugeParam, appendCombine
