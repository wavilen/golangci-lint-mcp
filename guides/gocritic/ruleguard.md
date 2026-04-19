# gocritic: ruleguard

<instructions>
Detects code that matches custom rules defined via `ruleguard` DSL files. The `ruleguard` checker applies user-defined patterns expressed as Go AST matching rules, enabling project-specific linting beyond built-in gocritic checks.

Fix the code according to the specific ruleguard rule that was triggered. Refer to the rule's documentation in your project's ruleguard files for details.
</instructions>

<examples>
## Bad
```go
// Depends on your project's ruleguard rules.
// Example rule: disallow time.Sleep in tests
func TestSomething(t *testing.T) {
	time.Sleep(5 * time.Second)
}
```

## Good
```go
// Use a mock or fake clock instead
func TestSomething(t *testing.T) {
	fakeClock.Advance(5 * time.Second)
}
```
</examples>

<patterns>
- Any pattern defined in your project's `rules.go` or `.go` ruleguard files
- Project-specific conventions enforced via AST pattern matching
- Custom anti-patterns not covered by built-in linters
</patterns>

<related>
revive
</related>
