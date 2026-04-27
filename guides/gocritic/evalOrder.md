# gocritic: evalOrder

<instructions>
Detects expressions where evaluation order affects correctness, particularly in multi-assignment statements where the same variable is read and modified. Go evaluates right-hand side expressions before assignment, but complex expressions with side effects can lead to subtle bugs when the same variable appears on both sides.

Simplify the expression or use temporary variables to make the data flow explicit and avoid order-dependent surprises.
</instructions>

<examples>
## Good
```go
tmp := a[i+1]
a[i] = tmp
i = i + 1
```
</examples>

<patterns>
- Separate multi-assignments where the same variable is written twice into distinct statements
- Avoid slice indexing with side-effecting index updates in the same expression
- Avoid concurrent map read/write on the same line — separate into sequential operations
- Replace `x, x = 1, 2` (bug) with distinct variables; keep `x, y = y, x` (correct swap)
</patterns>

<related>
gocritic/argOrder, gocritic/dupSubExpr, gocritic/appendAssign
</related>
