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
- Replace `fmt.Println(fmt.Sprintf(...))` with `fmt.Printf`
- Use `fmt.Printf` with format verbs instead of `fmt.Print` with string concatenation
- Use `log.Printf` for debug logging instead of `fmt.Println`
- Simplify `fmt.Sprint` followed by `fmt.Println` to a direct `fmt.Printf` or `fmt.Println`
- Combine multiple `Print` family calls into a single `Printf` when possible
</patterns>

<related>
use-errors-new, unnecessary-format, string-format
