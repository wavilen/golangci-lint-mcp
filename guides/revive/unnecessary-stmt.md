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
- Remove standalone expressions with no side effects (arithmetic or function calls without assignment)
- Remove empty statements (lone semicolons or blank lines inside blocks)
- Remove assignments to blank identifier where the right side has no side effects
- Eliminate increment/decrement operations on values that are never used afterward
- Remove leftover debug statements like `count + 1` from removed code
</patterns>

<related>
unnecessary-if, unreachable-code, unnecessary-format
