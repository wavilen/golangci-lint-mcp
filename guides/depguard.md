# depguard

<instructions>
Depguard blocks imports of disallowed packages. Teams use it to enforce dependency policy — preventing use of deprecated packages, discouraging direct database access outside a DAL layer, or blocking known problematic libraries.

Configure `depguard.rules` in `.golangci.yml` with lists of denied and allowed packages. Replace blocked imports with approved alternatives.
</instructions>

<examples>
## Bad
```go
import (
    "io/ioutil" // deprecated since Go 1.16
)
```

## Good
```go
import (
    "io"
    "os"
)
```
</examples>

<patterns>
- Importing deprecated standard library packages like `io/ioutil`
- Using disallowed third-party packages that violate project dependency policy
- Direct database driver imports when a data access layer is mandated
- Importing packages that duplicate functionality available in the standard library
</patterns>

<related>
gomodguard, gomoddirectives, importas
