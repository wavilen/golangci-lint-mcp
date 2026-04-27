# gocritic: paramTypeCombine

<instructions>
Detects consecutive function parameters that share the same type but are declared individually. Go allows combining parameters of the same type into a comma-separated list with a single type annotation.

Combine consecutive same-typed parameters into a single type declaration.
</instructions>

<examples>
## Good
```go
func draw(_, _ int, _ string) {}
func cloneData(_, _ []byte) int { return 0 }
```
</examples>

<patterns>
- Combine `x int, y int` into `x, y int`
- Combine `a string, b string, c string` into `a, b, c string`
- Combine `src []byte, dst []byte` into `src, dst []byte`
</patterns>

<related>
gocritic/unnamedResult, gocritic/typeDefFirst
</related>
