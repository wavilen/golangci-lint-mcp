# mirror

<instructions>
Mirror reports wrong `reflect.Value` method usage, where a value-receiver method is called on a pointer `reflect.Value` or vice versa. These mistakes cause panics at runtime.

Use the correct `reflect.Value` method: call `Elem()` to dereference pointer values before calling value-receiver methods, or use pointer methods directly on pointer values.
</instructions>

<examples>
## Bad
```go
v := reflect.ValueOf(&x)
name := v.Type().Name() // returns empty string for pointer
```

## Good
```go
v := reflect.ValueOf(&x)
name := v.Elem().Type().Name() // dereferences pointer first
```
</examples>

<patterns>
- Call `Elem()` on pointer `reflect.Value` before using value-receiver methods
- Replace `ValueOf(ptr).FieldByName()` with `ValueOf(ptr).Elem().FieldByName()`
- Dereference pointer values with `Elem()` before calling `NumField()` or `Field()`
- Use `Elem()` to get an addressable value before checking `CanInterface()`
</patterns>

<related>
forcetypeassert, musttag, exptostd
