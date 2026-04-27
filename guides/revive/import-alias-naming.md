# revive: import-alias-naming

<instructions>
Enforces a naming convention for import aliases. Import aliases should follow the configured style — typically lowercase, no underscores, and matching the package's canonical name or a well-known abbreviation.

Rename import aliases to match the project convention. Only use aliases when needed to resolve name collisions or when the package name is ambiguous.
</instructions>

<examples>
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
- Remove unnecessary aliases that duplicate the package name
- Use lowercase aliases without underscores for import aliases
- Avoid aliases that shadow standard library package names
- Simplify verbose aliases that are less readable than the original package path
- Use a consistent alias style across the codebase
</patterns>

<related>
revive/import-shadowing, revive/duplicated-imports, revive/dot-imports
</related>
