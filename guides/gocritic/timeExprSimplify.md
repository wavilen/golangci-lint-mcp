# gocritic: timeExprSimplify

<instructions>
Detects verbose time duration expressions that can be simplified. For example, `time.Second * 2` should be `2 * time.Second` (or simply `2 * time.Second`), and `time.Minute * 60` should be `time.Hour`.

Use the most concise and readable time duration expression.
</instructions>

<examples>
## Bad
```go
timeout := time.Second * 30
dur := time.Minute * 60
```

## Good
```go
timeout := 30 * time.Second
dur := time.Hour
```
</examples>

<patterns>
- `time.Second * N` → `N * time.Second`
- `time.Minute * 60` → `time.Hour`
- `time.Millisecond * 1000` → `time.Second`
- `time.Minute * N` where N is divisible by 60 → `time.Hour`
</patterns>

<related>
assignOp, hexLiteral
</related>
