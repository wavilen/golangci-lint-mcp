# govet: directive

<instructions>
Reports issues with `//go:` compiler directives such as `//go:noinline`, `//go:nosplit`, `//go:linkname`, etc. Common problems include placing directives on non-function declarations, using unknown directives, or having incorrect syntax.

Place directives on the correct declaration they are meant to affect, and use only valid compiler directives.
</instructions>

<examples>
## Bad
```go
//go:noinline
const x = 1 // noinline only applies to functions
```

## Good
```go
//go:noinline
func expensive() int {
    return 42
}
```
</examples>

<patterns>
- `//go:noinline` on non-function declarations
- Unknown or misspelled directive names
- Directive separated from its target by blank lines
- `//go:generate` with malformed command
</patterns>

<related>
buildtag, stdversion
</related>
