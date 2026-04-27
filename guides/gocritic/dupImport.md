# gocritic: dupImport

<instructions>
Detects duplicate import declarations where the same package is imported multiple times, potentially with different aliases. This typically results from automated import management or manual merging of code.

Consolidate duplicate imports into a single import declaration. If aliases differ, choose one or rename to avoid confusion.
</instructions>

<examples>
## Good
```go
import (
	"fmt"
)
```
</examples>

<patterns>
- Remove duplicate imports of the same package — keep one declaration
- Remove aliased and non-aliased duplicates of the same package
- Combine multiple import blocks containing the same package into one
</patterns>

<related>
gocritic/commentedOutImport, gocritic/importShadow
</related>
