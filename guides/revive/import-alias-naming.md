# revive: import-alias-naming

<instructions>
Enforces a naming convention for import aliases. Import aliases should follow the configured style — typically lowercase, no underscores, and matching the package's canonical name or a well-known abbreviation.

Rename import aliases to match the project convention. Only use aliases when needed to resolve name collisions or when the package name is ambiguous.
</instructions>

<examples>
## Bad
```go
import (
    f "fmt"                    // unnecessary alias
    http_utils "net/http"      // underscores not allowed
    MyLib "github.com/x/y"     // PascalCase not allowed
)
```

## Good
```go
import (
    "fmt"
    "net/http"
    y "github.com/x/y" // lowercase, meaningful alias
)
```
</examples>

<patterns>
- Unnecessary aliases that duplicate the package name
- Aliases with underscores or mixed case
- Aliases that shadow standard library names
- Long aliases that are less readable than the original package path
- Inconsistent alias styles across the codebase
</patterns>

<related>
import-shadowing, duplicated-imports, dot-imports
