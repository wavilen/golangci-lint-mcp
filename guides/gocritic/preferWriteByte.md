# gocritic: preferWriteByte

<instructions>
Detects `w.Write([]byte{b})` or `w.Write([]byte{c})` patterns where the writer implements `io.ByteWriter`. Using `WriteByte(b)` avoids creating a single-element byte slice on every call. `bytes.Buffer` and `bufio.Writer` both implement `io.ByteWriter`.

Replace `w.Write([]byte{b})` with `w.WriteByte(b)` when writing a single byte.
</instructions>

<examples>
## Bad
```go
var buf bytes.Buffer
buf.Write([]byte{'\n'})
```

## Good
```go
var buf bytes.Buffer
buf.WriteByte('\n')
```
</examples>

<patterns>
- Writing single delimiter or separator bytes via `Write([]byte{','})`
- Flushing newline or space characters through a `bytes.Buffer`
- Byte-by-byte writing in serialization loops using `Write([]byte{b})`
- Any writer known to implement `io.ByteWriter` receiving single-byte slices
</patterns>

<related>
preferStringWriter, preferFprint, zeroByteRepeat
