# grouper: import

<instructions>
Detects multiple `import` declarations that can be grouped into a single `import` block. Multiple separate import statements add noise. Use `import ( ... )` to group all imports together, with standard library imports first, then third-party, then local imports.
</instructions>

<examples>
## Bad
```go
import "fmt"
import "strings"
import "github.com/stretchr/testify/assert"
```

## Good
```go
import (
    "fmt"
    "strings"

    "github.com/stretchr/testify/assert"
)
```
</examples>

<patterns>
- Multiple `import "x"` statements — group into `import ( ... )`
- Two or more import blocks at the top of a file — merge into one
- Single import can remain ungrouped, but multiple should be grouped
</patterns>

<related>
const, type, var
