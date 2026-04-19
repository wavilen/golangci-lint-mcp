# revive: package-naming

<instructions>
Enforces Go package naming conventions: package names must be lowercase, single-word identifiers without underscores or mixed-case. Well-named packages are concise, descriptive, and avoid stutter when used with exported types (e.g., `user.User` not `user.UserModel`).

Rename the package to a single lowercase word. Avoid `common`, `util`, `helpers`, or generic names.
</instructions>

<examples>
## Bad
```go
package my_package
package UserAPI
package commonUtils
```

## Good
```go
package user
package api
package auth
```
</examples>

<patterns>
- Package names containing underscores or hyphens
- Mixed-case or CamelCase package names
- Generic names like `common`, `util`, `helpers`, `misc`
- Package names that stutter with exported symbols (e.g., `package http` with `type HTTPServer`)
- Names longer than one word when a single word suffices
</patterns>

<related>
package-comments, package-directory-mismatch, var-naming
