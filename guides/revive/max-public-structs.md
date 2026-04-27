# revive: max-public-structs

<instructions>
Enforces a maximum number of exported (public) structs per file. Too many public structs in one file suggests the file is a dumping ground for unrelated types, making it hard to find related code and increasing merge conflicts.

Split the file into multiple files, grouping related types by domain, feature, or responsibility.
</instructions>

<examples>
## Good
```go
// user.go
type User struct { ... }
type UserProfile struct { ... }

// order.go
type Order struct { ... }
type OrderItem struct { ... }

// product.go
type Product struct { ... }
type ProductCategory struct { ... }
```
</examples>

<patterns>
- Separate monolithic "models" or "types" files into domain-specific files
- Organize auto-generated structs into separate files by domain or feature
- Group API request/response types by endpoint or feature in separate files
- Separate legacy type accumulation into files grouped by responsibility
- Organize types from shared registry files into their respective domain files
</patterns>

<related>
revive/file-length-limit, revive/function-length
</related>
