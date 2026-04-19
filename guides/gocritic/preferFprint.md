# gocritic: preferFprint

<instructions>
Detects `fmt.Sprintf` followed by `io.Writer.Write` or similar patterns where `fmt.Fprint`, `fmt.Fprintf`, or `fmt.Fprintln` would write directly to the writer without allocating an intermediate string. Using `Fprint` family functions avoids the temporary string allocation.

Replace `w.Write([]byte(fmt.Sprintf(...)))` with `fmt.Fprintf(w, ...)`.
</instructions>

<examples>
## Bad
```go
w.Write([]byte(fmt.Sprintf("name: %s, age: %d\n", name, age)))
```

## Good
```go
fmt.Fprintf(w, "name: %s, age: %d\n", name, age)
```
</examples>

<patterns>
- `w.Write([]byte(fmt.Sprintf(...)))` — formats to string then converts to bytes for writing
- `buf.WriteString(fmt.Sprintf(...))` — formats to string then writes to buffer
- Using `Sprintf` to build output that is immediately written somewhere
- Response writers or buffers receiving `Sprintf` output indirectly
</patterns>

<related>
preferStringWriter, preferWriteByte, appendCombine
