# gocritic: importShadow

<instructions>
Detects variables, parameters, or type parameters whose names shadow an imported package name. Shadowing makes the package inaccessible within that scope, causing confusing "undefined" compile errors or silent use of the wrong value. Rename the local identifier to avoid colliding with the imported package name.
</instructions>

<examples>
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
gocritic/builtinShadow, gocritic/builtinShadowDecl, gocritic/dupImport
</related>
