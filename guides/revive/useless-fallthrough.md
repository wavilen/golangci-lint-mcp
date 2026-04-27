# revive: useless-fallthrough

<instructions>
Detects unnecessary `fallthrough` statements in switch cases. A `fallthrough` at the end of the final `case` or `default` clause does nothing because there's no next case to fall into. Even in non-final cases, `fallthrough` is rarely needed and often indicates logic that could be restructured.

Remove the unnecessary `fallthrough`. If both cases should execute the same logic, combine them with comma-separated case labels.
</instructions>

<examples>
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
- Remove `fallthrough` from the final `case` or `default` clause — there is no next case to fall into
- Combine cases with comma syntax instead of using `fallthrough` to a case with identical logic
- Replace fallthrough chains where each case does the same thing with multi-value case labels
- Use comma-separated case labels instead of `fallthrough` for proper case grouping
- Remove `fallthrough` from type switches — it is not allowed in Go
</patterns>

<related>
revive/useless-break, revive/unnecessary-stmt
</related>
