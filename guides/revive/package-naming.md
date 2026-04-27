# revive: package-naming

<instructions>
Enforces Go package naming conventions: package names must be lowercase, single-word identifiers without underscores or mixed-case. Well-named packages are concise, descriptive, and avoid stutter when used with exported types (e.g., `user.User` not `user.UserModel`).

Rename the package to a single lowercase word. Avoid `common`, `util`, `helpers`, or generic names.
</instructions>

<examples>
## Good
```go
package user
package api
package auth
```
</examples>

<patterns>
- Use lowercase single-word package names without underscores or hyphens
- Rename mixed-case or CamelCase packages to lowercase
- Replace generic names like `common`, `util`, `helpers`, `misc` with specific domain names
- Avoid package names that stutter with exported symbols (e.g., rename package `http` when it exports `HTTPServer`)
- Simplify multi-word package names to a single descriptive word
</patterns>

<related>
revive/package-comments, revive/package-directory-mismatch, revive/var-naming
</related>
