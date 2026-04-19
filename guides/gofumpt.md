# gofumpt

<instructions>
Gofumpt enforces stricter formatting rules than `gofmt`. It applies additional style rules: simplified control flow, consistent spacing, grouped imports, and normalized declarations. Code is correct but doesn't follow community formatting conventions.

Run `gofumpt -w .` to auto-fix all formatting issues.
</instructions>

<examples>
## Bad
```go
import(
"os"
"fmt"
)
func main(){
if  x  >  0  {
fmt.Println("positive")}
}
```

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
- Inconsistent spacing around operators and after keywords
- Imports not grouped (stdlib, third-party, local)
- `else` and `if` on same line with different indentation
- Multiline function signatures without aligned parameters
</patterns>

<related>
whitespace, nlreturn, gofmt, govet
</related>
