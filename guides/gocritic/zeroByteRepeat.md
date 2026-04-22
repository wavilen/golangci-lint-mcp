# gocritic: zeroByteRepeat

<instructions>
Detects `strings.Repeat("\x00", n)` or `strings.Repeat(string(byte(0)), n)` used to create a zero-filled string. Use `strings.Builder` with `Grow` or simply allocate a zero-valued byte slice (zero-initialized by Go) and convert once. The `strings.Repeat` approach is unnecessarily complex for zero bytes.

Use `string(make([]byte, n))` for a zero-filled string, or `make([]byte, n)` directly if a byte slice suffices.
</instructions>

<examples>
## Bad
```go
padding := strings.Repeat("\x00", 16)
```

## Good
```go
padding := string(make([]byte, 16)) // all zeros, no Repeat needed
```
</examples>

<patterns>
- Replace `strings.Repeat("\x00", n)` with `make([]byte, n)` for zero-padded buffers
- Replace `bytes.Repeat([]byte{0}, n)` with `make([]byte, n)` — simpler and equivalent
- Replace repeated zero-byte initialization with `make([]byte, n)` — zero-valued by default
- Use `make([]byte, n)` for null-terminated or zero-filled protocol messages
</patterns>

<related>
sliceClear, stringXbytes, appendCombine
