# gofmt

<instructions>
Gofmt enforces standard Go formatting per the spec. Run `gofmt -w .` to auto-fix all issues — indentation, spacing, alignment, and line breaks are normalized automatically.
</instructions>

<examples>
## Bad
```go
type Config struct{
	Name string
	Age  int
	Email string
}
func process( c Config )(int,error){
	if  c.Age>18  {
	return c.Age,nil}
	return 0,fmt.Errorf("too young")
}
```

## Good
```go
type Config struct {
	Name  string
	Age   int
	Email string
}

func process(c Config) (int, error) {
	if c.Age > 18 {
		return c.Age, nil
	}
	return 0, fmt.Errorf("too young")
}
```
</examples>

<patterns>
- Inconsistent indentation (tabs vs spaces, wrong tab width)
- Irregular spacing around operators and after keywords
- Missing blank lines between functions or after imports
- Misaligned struct fields or multi-line declarations
</patterns>

<related>
gofumpt, govet, whitespace
</related>
