# gocritic: preferStringWriter

<instructions>
Detects `io.Writer.Write([]byte(s))` calls where the writer also implements `io.StringWriter`. Calling `WriteString(s)` avoids the `[]byte(s)` conversion and potential allocation. Most buffered writers and builders in the standard library implement `io.StringWriter`.

Use `WriteString(s)` (or `fmt.Fprint(w, s)`) instead of `w.Write([]byte(s))` when writing string data.
</instructions>

<examples>
## Bad
```go
var buf bytes.Buffer
buf.Write([]byte("hello world"))
```

## Good
```go
var buf bytes.Buffer
buf.WriteString("hello world")
```
</examples>

<patterns>
- Writing string constants or variables via `w.Write([]byte(s))` to a `bytes.Buffer`
- HTTP response writes using `rw.Write([]byte(str))` where `WriteString` avoids allocation
- `bufio.Writer.Write([]byte(s))` when the writer supports `io.StringWriter`
- Any `Write([]byte(stringValue))` pattern where the argument originates as a string
</patterns>

<related>
preferFprint, preferWriteByte, stringXbytes
