# gocritic: unslice

<instructions>
Detects expressions that unnecessarily convert a slice to a slice, such as `x[:]` when `x` is already a slice, or `string(b[:]])` when `string(b)` suffices. These redundant slice operations add noise.

Remove the unnecessary slice operation and use the value directly.
</instructions>

<examples>
## Good
```go
header := strings.Split(line, ":")
key := strings.TrimSpace(header[0])
data := string(bytes)
```
</examples>

<patterns>
- Replace `slice[:]` with `slice` — full-length reslice is redundant
- Replace `string(b[:])` with `string(b)` when `b` is `[]byte`
- Replace `append(s[:], ...)` with `append(s, ...)` — remove redundant reslice
- Remove redundant reslices of full-length slices
</patterns>

<related>
gocritic/typeUnparen, gocritic/underef
</related>
