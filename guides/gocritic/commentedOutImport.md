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
- Remove commented-out single import lines inside import blocks
- Remove entire commented-out import groups — delete them or move to tool directives
</patterns>

<related>
commentedOutCode, dupImport
</related>
