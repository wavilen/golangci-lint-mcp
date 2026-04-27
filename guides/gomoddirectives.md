# gomoddirectives

<instructions>
Gomoddirectives checks `go.mod` for problematic directives: `replace`, `exclude`, `retract`, and `go` version declarations. These can hide dependency issues or pin to outdated versions.

Remove unnecessary `replace` directives once upstream fixes are released. Avoid local `replace` directives in published modules. Keep the `go` directive set to the minimum required version.
</instructions>

<examples>
## Good
```go
module example.com/app

go 1.23

require example.com/lib v1.2.3
```
</examples>

<patterns>
- Remove local `replace` directives that point to relative paths
- Remove `replace` directives for upstream bugs that are now fixed
- Resolve version conflicts directly instead of using `exclude` directives
- Add rationale comments to `retract` directives
</patterns>

<related>
gomodguard, gocheckcompilerdirectives, goheader
</related>
