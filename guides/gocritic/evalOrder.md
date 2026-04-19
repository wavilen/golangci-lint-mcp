# gocritic: evalOrder

<instructions>
Detects expressions where evaluation order affects correctness, particularly in multi-assignment statements where the same variable is read and modified. Go evaluates right-hand side expressions before assignment, but complex expressions with side effects can lead to subtle bugs when the same variable appears on both sides.

Simplify the expression or use temporary variables to make the data flow explicit and avoid order-dependent surprises.
</instructions>

<examples>
## Bad
```go
a[i] = a[i+1] // reads a[i+1] and writes a[i]
i = i + 1     // but both happen on same line in original
x, x = 1, 2   // x assigned twice
```

## Good
```go
tmp := a[i+1]
a[i] = tmp
i = i + 1
```
</examples>

<patterns>
- Multi-assignment where same variable is written twice
- Slice indexing with side-effecting index updates
- Map access with concurrent read/write on same line
- Complex assignments like `x, y = y, x` (fine) vs `x, x = 1, 2` (bug)
</patterns>

<related>
argOrder, dupSubExpr, appendAssign
</related>
