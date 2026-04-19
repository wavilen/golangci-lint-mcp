# copyloopvar

<instructions>
Copyloopvar detects loop variables that are copied inside the loop body instead of using the loop variable directly. This pattern was necessary before Go 1.22 to avoid closure capture issues, but is now unnecessary since Go 1.22 fixed loop variable semantics.

Remove the unnecessary copy and use the loop variable directly.
</instructions>

<examples>
## Bad
```go
for _, item := range items {
    item := item // unnecessary copy
    process(item)
}
```

## Good
```go
for _, item := range items {
    process(item)
}
```
</examples>

<patterns>
- `val := val` or `item := item` shadow copies inside range loops
- Closure capture workarounds that are obsolete since Go 1.22
- Goroutine launches with copied loop variables no longer needed
</patterns>

<related>
gosimple, staticcheck, govet
