# gocritic: dupImport

<instructions>
Detects duplicate import declarations where the same package is imported multiple times, potentially with different aliases. This typically results from automated import management or manual merging of code.

Consolidate duplicate imports into a single import declaration. If aliases differ, choose one or rename to avoid confusion.
</instructions>

<examples>
## Bad
```go
import (
	"fmt"
	"fmt" // duplicate
)
```

## Good
```go
import (
	"fmt"
)
```
</examples>

<patterns>
- Same package imported twice without alias
- Same package imported with and without alias: `"os"` and `os "os"`
- Multiple import blocks containing the same package
</patterns>

<related>
commentedOutImport, importShadow
</related>
