# govet: directive

<instructions>
Reports issues with `//go:` compiler directives such as `//go:noinline`, `//go:nosplit`, `//go:linkname`, etc. Common problems include placing directives on non-function declarations, using unknown directives, or having incorrect syntax.

Place directives on the correct declaration they are meant to affect, and use only valid compiler directives.
</instructions>

<examples>
## Good
```go
//go:noinline
func expensive() int {
    return 42
}
```
</examples>

<patterns>
- Place `//go:noinline` and similar directives only on function declarations
- Use correct directive names — fix misspelled or unknown directives
- Place directives directly above their target without intervening blank lines
- Fix malformed `//go:generate` commands — ensure valid shell syntax
</patterns>

<related>
govet/buildtag, govet/stdversion
</related>
