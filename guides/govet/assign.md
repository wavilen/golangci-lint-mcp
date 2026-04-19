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
- Variable assigned to itself (`x = x`)
- Struct field assigned to itself (`s.Field = s.Field`)
- Result of pure function assigned back to same variable
</patterns>

<related>
unusedresult, appends
</related>
