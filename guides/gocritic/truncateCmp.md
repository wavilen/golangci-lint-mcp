# gocritic: truncateCmp

<instructions>
Detects truncation in integer comparison expressions, such as comparing an `int32` with an `int64` where the wider value is implicitly truncated. This can cause incorrect comparisons when the truncated value wraps around.

Use explicit conversion to the wider type before comparison, or ensure both operands are the same type.
</instructions>

<examples>
## Bad
```go
var x int64 = 1<<32
var y int32 = 0
if int32(x) == y { // truncation loses high bits
    slog.Info("equal")
}
```

## Good
```go
var x int64 = 1 << 32
var y int32 = 0
if x == int64(y) {
    slog.Info("equal")
}
```
</examples>

<patterns>
- Comparing `int32(x)` with an `int32` when `x` is `int64`
- Implicit truncation in comparisons between different integer widths
- `int8`/`int16` comparisons that silently wrap
- Casting to narrower type before comparison
</patterns>

<related>
offBy1, badCond, sloppyTypeAssert
</related>
