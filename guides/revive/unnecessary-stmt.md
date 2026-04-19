# revive: unnecessary-stmt

<instructions>
Detects statements that have no effect, such as expressions evaluated but whose result is discarded, or empty statements that serve no purpose. These are usually leftover artifacts from refactoring or debugging.

Remove the useless statement entirely, or if it was meant to have side effects, assign the result to the blank identifier with a comment explaining why.
</instructions>

<examples>
## Bad
```go
x + 1           // result discarded
len(items)      // no side effect
_ = 42          // assigning constant to blank identifier
```

## Good
```go
result := x + 1
if len(items) > 0 {
    process(items)
}
```
</examples>

<patterns>
- Standalone expressions with no side effects (arithmetic, function calls without assignment)
- Empty statements (lone semicolons or blank lines inside blocks)
- Assignments to blank identifier where the right side has no side effects
- Increment/decrement operations on values never used
- Leftover debug statements like `count + 1` from removed code
</patterns>

<related>
unnecessary-if, unreachable-code, unnecessary-format
