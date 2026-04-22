# iotamixing

<instructions>
Iotamixing detects `const` blocks where multiple `iota` values are mixed in the same declaration group, leading to confusing or incorrect constant values. Mixing different iota expressions in one block makes the sequence hard to reason about.

Split mixed iota declarations into separate const blocks or use explicit values.
</instructions>

<examples>
## Bad
```go
const (
    ReadPerm  = 1 << iota // 1
    WritePerm             // 2
    AdminRole = iota      // 2 — same value as WritePerm!
)
```

## Good
```go
const (
    ReadPerm  = 1 << iota // 1
    WritePerm             // 2
    ExecPerm              // 4
)

const (
    RoleAdmin = iota // 0
    RoleUser         // 1
)
```
</examples>

<patterns>
- Split const blocks that mix different `iota` patterns into separate groups
- Use explicit values or separate const blocks after an iota break
- Separate bit-shift iota and plain iota into distinct const blocks
</patterns>

<related>
govet, goconst, staticcheck
