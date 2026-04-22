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
- Remove duplicate imports that appear when merging branches with different aliases
- Remove the old alias when renaming an import — keep only the new one
- Remove duplicate imports brought in by copied helper functions
- Set IDE auto-import to detect and skip already-imported packages
- Remove duplicate imports when moving code between files
</patterns>

<related>
dot-imports, import-alias-naming
