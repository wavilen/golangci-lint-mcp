# revive: import-shadowing

<instructions>
Detects import names that shadow package-level identifiers (variables, constants, types, or functions). Shadowing makes code confusing — readers may not realize a name refers to an import rather than a local declaration.

Rename either the import alias or the shadowed identifier to be distinct. Avoid naming variables the same as imported packages.
</instructions>

<examples>
## Bad
```go
import "fmt"

var fmt = "a string" // shadows the fmt import

func log() {
    fmt.Println(fmt) // confusing — which fmt?
}
```

## Good
```go
import "fmt"

var format = "a string"

func log() {
    fmt.Println(format)
}
```
</examples>

<patterns>
- Rename variables that shadow imported package names (e.g., rename `json` to `jsonData`)
- Avoid function parameter names that match import aliases
- Use distinct names for local variables in inner scopes that would shadow imports
- Use receiver names that differ from imported package names
- Rename loop variables or short variable declarations that shadow import names
</patterns>

<related>
import-alias-naming, confusing-naming
