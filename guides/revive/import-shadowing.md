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
- Variables named after imported packages (e.g., `json := ...`)
- Function parameters with names matching import aliases
- Local variables in inner scopes shadowing imports
- Receiver names matching import package names
- Loop variables or short variable declarations shadowing imports
</patterns>

<related>
import-alias-naming, confusing-naming
