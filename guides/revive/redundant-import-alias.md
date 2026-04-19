# revive: redundant-import-alias

<instructions>
Flags import aliases that are identical to the imported package name. Writing `import fmt "fmt"` is redundant — the default identifier is already `fmt`. This creates unnecessary noise in import declarations.

Remove the alias and use the plain import path, letting Go use the default package name.
</instructions>

<examples>
## Bad
```go
import (
    fmt "fmt"
    http "net/http"
)
```

## Good
```go
import (
    "fmt"
    "net/http"
)
```
</examples>

<patterns>
- Import aliases that exactly match the package name
- Auto-import tools adding redundant aliases
- Refactored imports where the alias used to differ but was later changed
- Generated code with explicit aliases for clarity that duplicate the name
- Import grouping where some aliases are needed but others are redundant
</patterns>

<related>
redundant-build-tag, redundant-test-main-exit, dot-imports
