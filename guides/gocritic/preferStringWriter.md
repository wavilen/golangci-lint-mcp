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
- Replace `w.Write([]byte(s))` with `w.WriteString(s)` when writing strings to `bytes.Buffer`
- Replace `rw.Write([]byte(str))` with `io.WriteString(rw, str)` for HTTP responses
- Replace `bufio.Writer.Write([]byte(s))` with `w.WriteString(s)` when the writer supports `io.StringWriter`
- Use `WriteString` for any `Write([]byte(stringValue))` pattern to avoid allocation
</patterns>

<related>
preferFprint, preferWriteByte, stringXbytes
