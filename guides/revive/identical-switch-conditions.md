# revive: identical-switch-conditions

<instructions>
Detects switch statements where two case expressions evaluate to the same value. The second case is unreachable — execution will always match the first one.

Remove the duplicate case expression. If the intent was to check something different, fix the condition.
</instructions>

<examples>
## Bad
```go
switch statusCode {
case 404:
    handleNotFound()
case 404: // unreachable — duplicate of above
    handleGone()
}
```

## Good
```go
switch statusCode {
case 404:
    handleNotFound()
case 410:
    handleGone()
}
```
</examples>

<patterns>
- Copy-paste errors duplicating a case value
- Constants that alias the same value used in different cases
- Enum values mapped to the same integer used in separate cases
- Variable-based switch where two expressions evaluate identically
- Generated switch code with duplicate constant values
</patterns>

<related>
identical-switch-branches, identical-ifelseif-conditions
