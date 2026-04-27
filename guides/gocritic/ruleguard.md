# gocritic: ruleguard

<instructions>
Detects code that matches custom rules defined via `ruleguard` DSL files. The `ruleguard` checker applies user-defined patterns expressed as Go AST matching rules, enabling project-specific linting beyond built-in gocritic checks.

Fix the code according to the specific ruleguard rule that was triggered. Refer to the rule's documentation in your project's ruleguard files for details.
</instructions>

<examples>
## Good
```go
// Use a mock or fake clock instead
func TestSomething(t *testing.T) {
	fakeClock.Advance(5 * time.Second)
}
```
</examples>

<patterns>
- Examine project `rules.go` or `.go` ruleguard files for defined patterns
- Identify project-specific conventions enforced via AST pattern matching
- Identify custom anti-patterns not covered by built-in linters
</patterns>

<related>
revive
</related>
