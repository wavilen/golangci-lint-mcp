# gocritic: preferDecodeRune

<instructions>
Detects manual UTF-8 rune decoding using `utf8.DecodeRuneInString` or `utf8.DecodeRune` where a simple `range` loop over the string would be clearer and idiomatic. The `range` keyword on strings automatically decodes runes.

Replace explicit `utf8.DecodeRune` calls in iteration with a `for _, r := range s` loop.
</instructions>

<examples>
## Bad
```go
for i := 0; i < len(s); {
    r, size := utf8.DecodeRuneInString(s[i:])
    process(r)
    i += size
}
```

## Good
```go
for _, r := range s {
    process(r)
}
```
</examples>

<patterns>
- Manual rune decoding loops using `utf8.DecodeRune` or `utf8.DecodeRuneInString`
- Byte-indexed string iteration that manually advances by rune size
- Converting a string to `[]rune` just to iterate over characters
- Code that mixes byte offsets with rune values unnecessarily
</patterns>

<related>
indexAlloc, stringXbytes, preferWriteByte
