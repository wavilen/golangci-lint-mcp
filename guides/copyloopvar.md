# copyloopvar

<instructions>
Copyloopvar detects loop variables that are copied inside the loop body instead of using the loop variable directly. This pattern was necessary before Go 1.22 to avoid closure capture issues, but is now unnecessary since Go 1.22 fixed loop variable semantics.

Remove the unnecessary copy and use the loop variable directly.
</instructions>

<examples>
## Good
```go
for _, item := range items {
    process(item)
}
```
</examples>

<patterns>
- Remove `val := val` shadow copies inside range loops (unnecessary since Go 1.22)
- Eliminate closure capture workarounds like local copies in goroutine launches
- Use the loop variable directly in closures launched inside `for range`
</patterns>

<related>
staticcheck, govet
</related>
