# gocritic: commentedOutImport

<instructions>
Detects import declarations that are commented out rather than removed. Commented-out imports clutter the import block and can mislead readers into thinking a dependency is still in use.

Remove commented-out imports entirely. Use version control history if you need to recover them later.
</instructions>

<examples>
## Bad
```go
import (
	"fmt"
	// "os"
	"strings"
)
```

## Good
```go
import (
	"fmt"
	"strings"
)
```
</examples>

<patterns>
- Single commented-out import line inside an import block
- Entire import group commented out with `//` prefix
- Import with alias commented out: `// uuid "github.com/..."`
</patterns>

<related>
commentedOutCode, dupImport
</related>
