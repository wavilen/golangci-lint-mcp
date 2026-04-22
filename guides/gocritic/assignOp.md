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
- Replace `x = x + y` with `x += y`
- Replace `x = x - y` with `x -= y`
- Replace `x = x * y` with `x *= y`
- Replace `x = x & mask` with `x &= mask`
- Replace `s = s + t` loop concatenation with `s += t`
</patterns>

<related>
stringConcatSimplify, yodaStyleExpr
</related>
