# modernize: reloop

<instructions>
Detects patterns where a loop can be simplified by using a different range form or removing unnecessary nesting. This includes ranging over a pointer to a collection when a nil check followed by ranging over the value directly is clearer, or simplifying nested loops that could be flattened.
</instructions>

<examples>
## Bad
```go
if items != nil {
    for i := range *items {
        process((*items)[i])
    }
}
```

## Good
```go
for _, item := range items {
    process(item)
}
```
</examples>

<patterns>
- Ranging over `*slice` with nil check — range over value directly (nil slices range zero times)
- Unnecessary pointer dereference inside loop — simplify binding
- Nested range where outer has single iteration — flatten the loop
</patterns>

<related>
simplifyrange, loopvar
