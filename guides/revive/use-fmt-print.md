# revive: use-fmt-print

<instructions>
Detects incorrect usage of `fmt.Print`/`fmt.Println` where `fmt.Printf` would be more appropriate, or vice versa. Using `fmt.Println(fmt.Sprintf(...))` is redundant because `fmt.Printf` combines formatting and printing. Similarly, `fmt.Print` with string concatenation could use `fmt.Printf`.

Use `fmt.Printf` when format verbs are needed. Use `fmt.Println` or `fmt.Print` for simple output without formatting. Avoid mixing `Sprintf` with `Print` — use `Printf` directly.
</instructions>

<examples>
## Bad
```go
fmt.Println(fmt.Sprintf("Hello %s", name))
fmt.Print("Count: " + strconv.Itoa(n) + "\n")
```

## Good
```go
fmt.Printf("Hello %s\n", name)
fmt.Printf("Count: %d\n", n)
```
</examples>

<patterns>
- `fmt.Println(fmt.Sprintf(...))` instead of `fmt.Printf`
- `fmt.Print` with string concatenation instead of `fmt.Printf` with verbs
- `fmt.Println` for debug logging where `log.Printf` would be better
- `fmt.Sprint` followed by `fmt.Println` for simple values
- Mixing `Print` family calls when a single `Printf` suffices
</patterns>

<related>
use-errors-new, unnecessary-format, string-format
