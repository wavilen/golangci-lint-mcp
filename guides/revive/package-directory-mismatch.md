# revive: package-directory-mismatch

<instructions>
Detects when the package name declared in `.go` files doesn't match the directory name containing those files. Go convention is that the directory name and package name should be identical so that import paths are intuitive and the package is discoverable.

Rename the package declaration to match the directory name, or rename the directory to match the package.
</instructions>

<examples>
## Bad
```go
// file: pkg/helpers/util.go
package utils // mismatch: directory is "helpers"
```

## Good
```go
// file: pkg/helpers/helpers.go
package helpers
```
</examples>

<patterns>
- Use the package declaration name with the directory name — e.g., rename `utils` to `helpers` if the directory is `helpers`
- Ensure test packages with `_test` suffix match the directory convention
- Rename the package declaration when renaming a directory during refactoring
- Use consistent naming between multi-word directories and their package declarations
- Set code generation to output the correct package name for the target directory
</patterns>

<related>
package-comments, package-naming
