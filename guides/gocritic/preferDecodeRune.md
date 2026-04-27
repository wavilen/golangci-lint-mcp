# gocritic: preferDecodeRune

<instructions>
Detects manual UTF-8 rune decoding using `utf8.DecodeRuneInString` or `utf8.DecodeRune` where a simple `range` loop over the string would be clearer and idiomatic. The `range` keyword on strings automatically decodes runes.

Replace explicit `utf8.DecodeRune` calls in iteration with a `for _, r := range s` loop.
</instructions>

<examples>
## Good
```go
for _, r := range s {
    process(r)
}
```
</examples>

<patterns>
- Replace manual `utf8.DecodeRune` loops with `for _, r := range s`
- Replace byte-indexed rune iteration with `range` over the string
- Replace `[]rune(s)` iteration with `for _, r := range s` — avoid the allocation
- Separate byte offsets from rune values — use `for i, r := range s`
</patterns>

<related>
gocritic/indexAlloc, gocritic/stringXbytes, gocritic/preferWriteByte
</related>
