# govet: appends

<instructions>
Detects `append` calls where the result is not used or assigned to a different variable than the original slice. Since `append` may reallocate the backing array, discarding the result can lose data or cause the original slice to reference stale memory.

Always assign the result of `append` back to the same slice variable: `s = append(s, elem)`.
</instructions>

<examples>
## Good
```go
items := make([]string, 0, 3)
items = append(items, "a", "b", "c") // result assigned back to same variable
```
</examples>

<patterns>
- Assign the `append` result back to the original slice variable (`s = append(s, elem)`)
- Use the return value of every `append` call — never discard it
- Assign the `append` result back to the same struct field, not a local variable
</patterns>

<related>
govet/assign, govet/copylocks
</related>
