# gomoddirectives

<instructions>
Gomoddirectives checks `go.mod` for problematic directives: `replace`, `exclude`, `retract`, and `go` version declarations. These can hide dependency issues or pin to outdated versions.

Remove unnecessary `replace` directives once upstream fixes are released. Avoid local `replace` directives in published modules. Keep the `go` directive set to the minimum required version.
</instructions>

<examples>
## Bad
```go
module example.com/app

go 1.23

require example.com/lib v0.0.0

replace example.com/lib => ../lib
```

## Good
```go
module example.com/app

go 1.23

require example.com/lib v1.2.3
```
</examples>

<patterns>
- Local `replace` directives pointing to relative paths (breaks consumers)
- `replace` directives for upstream bugs that have since been fixed
- `exclude` directives that mask version conflicts instead of resolving them
- `retract` directives missing rationale comments
</patterns>

<related>
gomodguard, gocheckcompilerdirectives, goheader
</related>
