# gocritic: paramTypeCombine

<instructions>
Detects consecutive function parameters that share the same type but are declared individually. Go allows combining parameters of the same type into a comma-separated list with a single type annotation.

Combine consecutive same-typed parameters into a single type declaration.
</instructions>

<examples>
## Bad
```go
func draw(x int, y int, color string) {}
func copy(dst []byte, src []byte) int { return 0 }
```

## Good
```go
func draw(x, y int, color string) {}
func copy(dst, src []byte) int { return 0 }
```
</examples>

<patterns>
- `x int, y int` → `x, y int`
- `a string, b string, c string` → `a, b, c string`
- `src []byte, dst []byte` → `src, dst []byte`
</patterns>

<related>
unnamedResult, typeDefFirst
</related>
