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
- Use bind parameters (`@name`) instead of string concatenation in AQL queries
- Replace `fmt.Sprintf` query building with parameterized queries and a bind variables map
- Pass a `bindVars` map when query contains `@placeholder` tokens
- Avoid direct string interpolation of values into AQL filter conditions
</patterns>

<related>
sqlclosecheck, noctx, rowserrcheck
