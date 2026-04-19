# nonamedreturns

<instructions>
Nonamedreturns bans named return values. While named returns can serve as documentation, they are frequently misused for naked returns or accidentally shadowed by local variables.

Remove named return values and use explicit returns instead. Only use named returns when documenting the meaning of return values in short, clear functions.
</instructions>

<examples>
## Bad
```go
func parse(input string) (result *Config, err error) {
    result = &Config{}
    err = json.Unmarshal([]byte(input), result)
    return
}
```

## Good
```go
func parse(input string) (*Config, error) {
    result := &Config{}
    err := json.Unmarshal([]byte(input), result)
    return result, err
}
```
</examples>

<patterns>
- Named returns used solely to enable naked returns
- Named returns that shadow local variable names
- Named returns in defer closures that modify return values implicitly
- Unnecessary named returns on single-return-value functions
</patterns>

<related>
nakedret, nlreturn, errname
</related>
