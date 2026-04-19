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
- `"" + s` → `s`
- `s + ""` → `s`
- `"a" + "b"` → `"ab"` (constant folding)
- Multiple concatenations that could use `strings.Builder`
</patterns>

<related>
redundantSprint, emptyStringTest, assignOp
</related>
