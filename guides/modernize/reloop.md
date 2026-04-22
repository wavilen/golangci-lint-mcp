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
- Simplify by ranging over slice values directly instead of `*slice` with nil check (nil slices range zero times)
- Simplify binding by removing unnecessary pointer dereferences inside loops
- Flatten nested range loops where the outer loop has a single iteration
</patterns>

<related>
simplifyrange, loopvar
