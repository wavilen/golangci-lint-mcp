# forcetypeassert

<instructions>
Forcetypeassert finds type assertions without the comma-ok safety check (`v := i.(T)` instead of `v, ok := i.(T)`). A forced assertion panics at runtime if the underlying type doesn't match.

Use the two-value comma-ok form and handle the failure case explicitly.
</instructions>

<examples>
## Bad
```go
var i any = getValue()
s := i.(string)
```

## Good
```go
var i any = getValue()
s, ok := i.(string)
if !ok {
    return fmt.Errorf("expected string, got %T", i)
}
```
</examples>

<patterns>
- Direct type assertions without comma-ok: `v.(Type)`
- Unmarshaled `any` values asserted without safety checks
- Interface values from external packages asserted to concrete types
- Switch case bodies that re-assert instead of using the matched type
</patterns>

<related>
errcheck, govet, staticcheck
