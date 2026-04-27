# gofumpt

<instructions>
Gofumpt enforces stricter formatting rules than `gofmt`. It applies additional style rules: simplified control flow, consistent spacing, grouped imports, and normalized declarations. Code is correct but doesn't follow community formatting conventions.

Run `gofumpt -w .` to auto-fix all formatting issues.
</instructions>

<examples>
## Good
```go
import (
    "fmt"
    "os"
)

func main() {
    if x > 0 {
        fmt.Println("positive")
    }
}
```
</examples>

<patterns>
- Run `gofumpt -w .` to enforce consistent spacing
- Group imports into stdlib, third-party, and local blocks
- Move `else` blocks to a new line with consistent indentation
- Align multiline function signature parameters
</patterns>

<related>
whitespace, nlreturn, gofmt, govet
</related>
