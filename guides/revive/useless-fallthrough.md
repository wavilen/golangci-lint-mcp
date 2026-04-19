# revive: useless-fallthrough

<instructions>
Detects unnecessary `fallthrough` statements in switch cases. A `fallthrough` at the end of the final `case` or `default` clause does nothing because there's no next case to fall into. Even in non-final cases, `fallthrough` is rarely needed and often indicates logic that could be restructured.

Remove the unnecessary `fallthrough`. If both cases should execute the same logic, combine them with comma-separated case labels.
</instructions>

<examples>
## Bad
```go
switch status {
case "active":
    log.Println("active")
    fallthrough
case "pending":
    log.Println("pending")
    fallthrough // useless — last case
default:
    log.Println("unknown")
}
```

## Good
```go
switch status {
case "active", "pending":
    log.Println("processing")
default:
    log.Println("unknown")
}
```
</examples>

<patterns>
- `fallthrough` in the final `case` or `default` clause
- `fallthrough` to a case that would be better combined with comma syntax
- Fallthrough chains where each case does the same thing
- Using `fallthrough` as a substitute for proper case grouping
- Fallthrough in type switches (which is not allowed in Go)
</patterns>

<related>
useless-break, unnecessary-stmt
