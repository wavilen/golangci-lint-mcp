# modernize: stringappend

<instructions>
Detects string concatenation using `+=` inside loops, which creates O(n²) allocations because strings are immutable. Each concatenation copies the entire existing string. Replace with `strings.Builder` which amortizes allocations and runs in O(n) time. The builder also minimizes copies with its internal buffer strategy.
</instructions>

<examples>
## Bad
```go
var result string
for _, part := range parts {
    result += part
}
```

## Good
```go
var b strings.Builder
for _, part := range parts {
    b.WriteString(part)
}
result := b.String()
```
</examples>

<patterns>
- Use `strings.Builder` instead of `s += x` inside a loop
- Use `b.WriteString(x)` instead of `s = s + x` in loop bodies
- Use `strings.Builder` for building query strings, CSV rows, or error messages in loops
</patterns>

<related>
errorf, sliceclear
