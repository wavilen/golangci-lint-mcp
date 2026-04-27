# depguard

<instructions>
Depguard blocks imports of disallowed packages. Teams use it to enforce dependency policy — preventing use of deprecated packages, discouraging direct database access outside a DAL layer, or blocking known problematic libraries.

Configure `depguard.rules` in `.golangci.yml` with lists of denied and allowed packages. Replace blocked imports with approved alternatives.
</instructions>

<examples>
## Good
```go
import (
    "io"
    "os"
)
```
</examples>

<patterns>
- Replace deprecated standard library packages (e.g., `io/ioutil`) with their current equivalents
- Remove disallowed third-party packages and use approved alternatives per project policy
- Replace direct database driver imports with calls through the mandated data access layer
- Replace third-party imports that duplicate standard library functionality
</patterns>

<related>
gomodguard, gomoddirectives, importas
</related>
