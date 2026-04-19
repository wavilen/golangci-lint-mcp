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
- Package name `utils` in a directory named `helpers`
- Test packages with `_test` suffix not matching directory
- Refactored packages where directory was renamed but package declaration was not
- Multi-word directories with underscores but package names without
- Generated code with a different package name than the target directory
</patterns>

<related>
package-comments, package-naming
