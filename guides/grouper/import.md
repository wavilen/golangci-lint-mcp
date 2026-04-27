# grouper: import

<instructions>
Detects multiple `import` declarations that can be grouped into a single `import` block. Multiple separate import statements add noise. Use `import ( ... )` to group all imports together, with standard library imports first, then third-party, then local imports.
</instructions>

<examples>
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
- Group multiple `import "x"` statements into a single `import ( ... )` block
- Merge two or more import blocks at the top of a file into one
- Group imports when there are multiple — single imports can remain ungrouped
</patterns>

<related>
grouper/const, grouper/type, grouper/var
</related>
