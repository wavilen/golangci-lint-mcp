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
- Rename parameters that shadow imports — avoid `func (fmt string)`
- Rename loop variables that shadow imports — avoid `for path := range paths` when `path` is a package
- Rename local variables that shadow imported package names
- Rename type parameters that collide with package names
</patterns>

<related>
builtinShadow, builtinShadowDecl, dupImport
</related>
