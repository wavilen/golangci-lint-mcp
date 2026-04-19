# revive: line-length-limit

<instructions>
Enforces a maximum line length per source file. Long lines are hard to read in standard terminal widths and in code review tools. This is a style-only rule; Go code is not affected functionally by line length.

Break long lines at natural points: after operators, before function arguments, or by extracting sub-expressions into named variables.
</instructions>

<examples>
## Bad
```go
result, err := someService.DoSomethingVeryComplex(ctx, param1, param2, param3, param4, param5, param6, param7)
```

## Good
```go
result, err := someService.DoSomethingVeryComplex(
    ctx,
    param1,
    param2,
    param3,
    param4,
    param5,
)
```
</examples>

<patterns>
- Long function call chains on a single line
- Complex boolean expressions spanning past the margin
- URL or path string literals exceeding the limit
- Long type definitions or struct literals on one line
- Chained method calls without line breaks
</patterns>

<related>
file-length-limit, function-length, lll
