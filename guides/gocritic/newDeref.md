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
- Replace `*new(int)` with `var x int` or the zero value `0`
- Replace `*new(MyStruct)` with `MyStruct{}`
- Replace `*new([]byte)` with `var b []byte`
</patterns>

<related>
underef, typeUnparen
</related>
