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
- Convert to the wider type before comparing — avoid truncating `int64` to `int32`
- Use same-width integer types for comparisons to avoid implicit truncation
- Avoid `int8`/`int16` comparisons that silently wrap — widen the type first
- Avoid narrowing casts before comparison — compare in the original type
</patterns>

<related>
offBy1, badCond, sloppyTypeAssert
</related>
