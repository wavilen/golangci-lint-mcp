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
- Flatten long function call chains across multiple lines at natural points
- Separate complex boolean expressions across lines before logical operators
- Flatten URL or path string literals that exceed the limit into concatenated segments
- Separate long type definitions or struct literals across multiple lines
- Add line breaks between chained method calls to stay within the limit
</patterns>

<related>
file-length-limit, function-length, lll
