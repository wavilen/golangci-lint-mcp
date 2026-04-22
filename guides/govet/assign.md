# govet: assign

<instructions>
Detects useless assignments where a variable is assigned to itself (`x = x`) or the result of an expression that always equals the variable. These are typically typos or leftover debugging code that serve no purpose.

Remove the redundant assignment entirely.
</instructions>

<examples>
## Bad
```go
x := 10
x = x // useless assignment
```

## Good
```go
x := 10
// no self-assignment needed
```
</examples>

<patterns>
- Remove self-assignments (`x = x`) — either delete the line or fix the intended target
- Remove struct field self-assignments (`s.Field = s.Field`) — fix the intended field or delete
- Remove pure function reassignment to the same variable (`x = len(x)`)
</patterns>

<related>
unusedresult, appends
</related>
