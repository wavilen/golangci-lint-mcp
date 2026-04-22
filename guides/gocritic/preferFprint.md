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
- Replace `w.Write([]byte(fmt.Sprintf(...)))` with `fmt.Fprintf(w, ...)`
- Replace `buf.WriteString(fmt.Sprintf(...))` with `fmt.Fprintf(buf, ...)`
- Replace `Sprintf` + immediate write with `Fprintf` directed to the writer
- Use `fmt.Fprintf` when writing formatted output directly to a response writer or buffer
</patterns>

<related>
preferStringWriter, preferWriteByte, appendCombine
