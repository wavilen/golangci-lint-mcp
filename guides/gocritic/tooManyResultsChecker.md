# gocritic: tooManyResultsChecker

<instructions>
Detects functions that return too many values. Functions with many return values are hard to use correctly — callers must remember the order and meaning of each value. The default threshold is 5 return values.

Reduce the number of return values by grouping related values into a struct, or split the function into smaller focused functions.
</instructions>

<examples>
## Bad
```go
func parseConfig() (host string, port int, user string, pass string, db string, ssl bool, timeout int, err error) {
	return "localhost", 5432, "admin", "secret", "mydb", true, 30, nil
}
```

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
- Functions returning 6+ values
- Multiple return values of the same type (easy to mix up order)
- Error buried among many return values
</patterns>

<related>
unnamedResult, paramTypeCombine
</related>
