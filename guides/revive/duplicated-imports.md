# revive: duplicated-imports

<instructions>
Detects the same package imported multiple times under different aliases. This is usually a mistake or the result of merging code branches. Multiple imports of the same package waste namespace and confuse readers about which alias to use.

Consolidate to a single import. If you need different names, pick one clear alias and remove the other.
</instructions>

<examples>
## Bad
```go
import (
    "encoding/json"
    j "encoding/json" // same package imported twice
)
```

## Good
```go
import (
    "encoding/json"
)
```
</examples>

<patterns>
- Merging branches that both added the same import with different aliases
- Renaming an import alias but forgetting to remove the old one
- Helper functions copied from another file bringing duplicate imports
- IDE auto-import adding a second copy of an already-imported package
- Refactoring that moved code between files without cleaning imports
</patterns>

<related>
dot-imports, import-alias-naming
