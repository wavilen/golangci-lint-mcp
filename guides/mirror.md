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
- Calling value-receiver methods on a pointer reflect.Value without Elem()
- Using `reflect.ValueOf(ptr).FieldByName()` instead of `Elem().FieldByName()`
- Accessing `NumField()` or `Field()` on pointer values
- Comparing `CanInterface()` results on non-addressable pointer values
</patterns>

<related>
forcetypeassert, musttag, exptostd
