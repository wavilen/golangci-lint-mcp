# govet: composites

<instructions>
Reports unkeyed composite literals — struct literals where fields are specified by position rather than by name (`T{1, "foo"}` instead of `T{A: 1, B: "foo"}`). Unkeyed literals are fragile: adding or reordering fields silently breaks the code.

Always use field names in struct composite literals.
</instructions>

<examples>
## Bad
```go
type Point struct{ X, Y int }
p := Point{1, 2} // unkeyed: fragile if fields change
```

## Good
```go
type Point struct{ X, Y int }
p := Point{X: 1, Y: 2} // keyed: safe and readable
```
</examples>

<patterns>
- Struct literal without field names
- Return statement returning unkeyed struct literal
- Nested unkeyed struct literals in composite literals
</patterns>

<related>
structtag, stdmethods
</related>
