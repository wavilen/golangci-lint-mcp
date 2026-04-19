# govet: appends

<instructions>
Detects `append` calls where the result is not used or assigned to a different variable than the original slice. Since `append` may reallocate the backing array, discarding the result can lose data or cause the original slice to reference stale memory.

Always assign the result of `append` back to the same slice variable: `s = append(s, elem)`.
</instructions>

<examples>
## Bad
```go
items := []string{"a", "b"}
more := append(items, "c") // more may share backing array with items
```

## Good
```go
items := []string{"a", "b"}
items = append(items, "c") // result assigned back to same variable
```
</examples>

<patterns>
- Assigning append result to a different variable instead of the original slice
- Calling append without using the return value
- Appending to a field but assigning to a local variable
</patterns>

<related>
assign, copylocks
</related>
