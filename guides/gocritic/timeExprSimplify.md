# gocritic: timeExprSimplify

<instructions>
Detects verbose time duration expressions that can be simplified. For example, `time.Second * 2` should be `2 * time.Second` (or simply `2 * time.Second`), and `time.Minute * 60` should be `time.Hour`.

Use the most concise and readable time duration expression.
</instructions>

<examples>
## Good
```go
timeout := 30 * time.Second
dur := time.Hour
```
</examples>

<patterns>
- Replace `time.Second * N` with `N * time.Second`
- Replace `time.Minute * 60` with `time.Hour`
- Replace `time.Millisecond * 1000` with `time.Second`
- Replace `time.Minute * N` with `time.Hour` when N is divisible by 60
</patterns>

<related>
gocritic/assignOp, gocritic/hexLiteral
</related>
