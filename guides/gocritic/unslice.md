# gocritic: unslice

<instructions>
Detects expressions that unnecessarily convert a slice to a slice, such as `x[:]` when `x` is already a slice, or `string(b[:]])` when `string(b)` suffices. These redundant slice operations add noise.

Remove the unnecessary slice operation and use the value directly.
</instructions>

<examples>
## Bad
```go
header := strings.Split(line, ":")
key := strings.TrimSpace(header[0][:])
data := string(bytes[:])
```

## Good
```go
header := strings.Split(line, ":")
key := strings.TrimSpace(header[0])
data := string(bytes)
```
</examples>

<patterns>
- `slice[:]` → `slice`
- `string(b[:])` → `string(b)` when `b` is `[]byte`
- `append(s[:], ...)` → `append(s, ...)`
- Redundant reslice of full-length slice
</patterns>

<related>
typeUnparen, underef
</related>
