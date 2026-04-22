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
- Remove import aliases that exactly match the package name — use the plain import path
- Set auto-import tools to skip adding redundant aliases
- Remove the alias from refactored imports where it used to differ but now matches
- Simplify generated code by removing explicit aliases that duplicate the package name
- Remove redundant aliases from import groups while keeping necessary ones
</patterns>

<related>
redundant-build-tag, redundant-test-main-exit, dot-imports
