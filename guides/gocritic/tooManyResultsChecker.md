# gocritic: tooManyResultsChecker

<instructions>
Detects functions that return too many values. Functions with many return values are hard to use correctly — callers must remember the order and meaning of each value. The default threshold is 5 return values.

Reduce the number of return values by grouping related values into a struct, or split the function into smaller focused functions.
</instructions>

<examples>
## Good
```go
type Config struct {
	Host    string
	Port    int
	User    string
	Pass    string
	DB      string
	SSL     bool
	Timeout int
}

func parseConfig() (*Config, error) {
	return &Config{
		Host: "localhost", Port: 5432, User: "admin",
		Pass: "secret", DB: "mydb", SSL: true, Timeout: 30,
	}, nil
}
```
</examples>

<patterns>
- Reduce function return values to fewer than 6 — use a result struct for grouping
- Wrap multiple return values of the same type in a named struct to prevent mix-ups
- Move the error return to the last position — avoid burying it among many values
</patterns>

<related>
gocritic/unnamedResult, gocritic/paramTypeCombine
</related>
