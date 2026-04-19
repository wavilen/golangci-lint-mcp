# gocritic: newDeref

<instructions>
Detects `*new(T)` expressions that allocate a zero value and immediately dereference it. This is equivalent to a simple `var` declaration or a composite literal `T{}` but is harder to read.

Replace `*new(T)` with `var x T` or use a composite literal `T{}`.
</instructions>

<examples>
## Bad
```go
val := *new(int)
config := *new(http.Client)
```

## Good
```go
var val int
config := http.Client{}
```

</examples>

<patterns>
- `*new(int)` → `var x int` or `0`
- `*new(MyStruct)` → `MyStruct{}`
- `*new([]byte)` → `var b []byte`
</patterns>

<related>
underef, typeUnparen
</related>
