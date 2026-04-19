# revive: max-public-structs

<instructions>
Enforces a maximum number of exported (public) structs per file. Too many public structs in one file suggests the file is a dumping ground for unrelated types, making it hard to find related code and increasing merge conflicts.

Split the file into multiple files, grouping related types by domain, feature, or responsibility.
</instructions>

<examples>
## Bad
```go
// models.go — 25 exported structs covering users, orders,
// products, payments, and notifications
type User struct { ... }
type Order struct { ... }
type Product struct { ... }
// ... 22 more structs
```

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
- "Models" or "types" files containing all domain types
- Auto-generated code producing many structs in one file
- API schema files with many request/response types ungrouped
- Legacy files that accumulated types over time without splitting
- Shared constant or type definition files acting as registries
</patterns>

<related>
file-length-limit, function-length
