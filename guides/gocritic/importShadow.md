# gocritic: importShadow

<instructions>
Detects identifiers in local scope (variables, type parameters, function parameters) that shadow an imported package name. This makes the package inaccessible within that scope and can lead to confusing "undefined" errors.

Rename the local identifier to avoid colliding with the imported package name.
</instructions>

<examples>
## Bad
```go
import "fmt"

func process(fmt string) {
	fmt.Println(fmt) // fmt is now a string, not the package
}
```

## Good
```go
import "fmt"

func process(format string) {
	fmt.Println(format)
}
```
</examples>

<patterns>
- Function parameter shadows an import: `func (fmt string)`
- Loop variable shadows an import: `for path := range paths`
- Local variable declaration shadows an import
- Type parameter name collides with a package name
</patterns>

<related>
builtinShadow, builtinShadowDecl, dupImport
</related>
