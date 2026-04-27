# gocritic: preferWriteByte

<instructions>
Detects `w.Write([]byte{b})` or `w.Write([]byte{c})` patterns where the writer implements `io.ByteWriter`. Using `WriteByte(b)` avoids creating a single-element byte slice on every call. `bytes.Buffer` and `bufio.Writer` both implement `io.ByteWriter`.

Replace `w.Write([]byte{b})` with `w.WriteByte(b)` when writing a single byte.
</instructions>

<examples>
## Good
```go
var buf bytes.Buffer
buf.WriteByte('\n')
```
</examples>

<patterns>
- Replace `Write([]byte{','})` with `WriteByte(',')` for single-byte writes
- Replace `Write([]byte{'\n'})` with `WriteByte('\n')` for flushing newline/separator bytes
- Replace byte-by-byte `Write([]byte{b})` in serialization loops with `WriteByte(b)`
- Use `WriteByte` for any single-byte write to a writer implementing `io.ByteWriter`
</patterns>

<related>
gocritic/preferStringWriter, gocritic/preferFprint, gocritic/zeroByteRepeat
</related>
