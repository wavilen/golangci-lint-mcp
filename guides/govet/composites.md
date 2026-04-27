# govet: composites

<instructions>
Reports unkeyed composite literals — struct literals where fields are specified by position rather than by name (`T{1, "foo"}` instead of `T{A: 1, B: "foo"}`). Unkeyed literals are fragile: adding or reordering fields silently breaks the code.

Always use field names in struct composite literals.
</instructions>

<examples>
## Good
```go
type Point struct{ X, Y int }
p := Point{X: 1, Y: 2} // keyed: safe and readable
```
</examples>

<patterns>
- Use keyed struct literals with field names (`T{A: 1}`) instead of positional (`T{1}`)
- Use keyed fields in return struct literals — never return unkeyed composites
- Use keyed fields for nested struct literals in composite expressions
</patterns>

<related>
govet/structtag, govet/stdmethods
</related>
