# gocritic: stringConcatSimplify

<instructions>
Detects string concatenation expressions that can be simplified. This includes chains of `+` with empty strings, redundant `"" + x`, or repeated concatenation that could use `fmt.Sprintf` or `strings.Join`.

Remove empty-string concatenations and simplify the expression to its minimal form.
</instructions>

<examples>
## Bad
```go
msg := "" + name + ""
msg := "Hello" + " " + name
```

## Good
```go
msg := name
msg := "Hello " + name
```
</examples>

<patterns>
- Remove `"" + s` — use `s` directly
- Remove `s + ""` — use `s` directly
- Combine `"a" + "b"` into `"ab"` at compile time
- Replace multiple concatenations with `strings.Builder` for performance
</patterns>

<related>
redundantSprint, emptyStringTest, assignOp
</related>
