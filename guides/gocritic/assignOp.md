# gocritic: assignOp

<instructions>
Detects assignments that can be simplified using compound assignment operators. For example, `x = x + 1` should be written as `x += 1`, and `s = s + t` as `s += t`.

Use the compound operator form for clarity and brevity.
</instructions>

<examples>
## Bad
```go
count = count + 1
name = name + suffix
bits = bits & mask
```

## Good
```go
count += 1
name += suffix
bits &= mask
```
</examples>

<patterns>
- `x = x + y` → `x += y`
- `x = x - y` → `x -= y`
- `x = x * y` → `x *= y`
- `x = x & mask` → `x &= mask`
- String concatenation in loops: `s = s + t` → `s += t`
</patterns>

<related>
stringConcatSimplify, yodaStyleExpr
</related>
