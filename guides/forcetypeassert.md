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
- Replace direct type assertions with comma-ok form: `v, ok := i.(T)`
- Guard unmarshaled `any` value assertions with comma-ok checks
- Use comma-ok when asserting external package interfaces to concrete types
- Use the matched type directly in switch case bodies instead of re-asserting
</patterns>

<related>
errcheck, govet, staticcheck
