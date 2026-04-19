# arangolint

<instructions>
Arangolint detects common mistakes in ArangoDB query construction using the go-arangodb driver. It flags incorrect use of query builders and bind parameters that can lead to runtime errors or injection vulnerabilities.

Always use parameterized queries with bind variables instead of string concatenation for ArangoDB AQL queries.
</instructions>

<examples>
## Bad
```go
query := fmt.Sprintf("FOR u IN users FILTER u.name == '%s' RETURN u", name)
cursor, err := db.Query(ctx, query, nil)
```

## Good
```go
query := "FOR u IN users FILTER u.name == @name RETURN u"
bindVars := map[string]interface{}{"name": name}
cursor, err := db.Query(ctx, query, bindVars)
```
</examples>

<patterns>
- String concatenation in AQL queries instead of bind parameters
- Using fmt.Sprintf to build query strings with user input
- Missing bind variables map when query contains placeholders
- Direct string interpolation of values into AQL filter conditions
</patterns>

<related>
sqlclosecheck, noctx, rowserrcheck
