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
- Run `gofmt -w .` to normalize indentation to Go standard (tabs)
- Apply `gofmt` to fix irregular spacing around operators and keywords
- Run `gofmt` to add required blank lines between functions
- Use `gofmt` to auto-align struct fields and multi-line declarations
</patterns>

<related>
gofumpt, govet, whitespace
</related>
