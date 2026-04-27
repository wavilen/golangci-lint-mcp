# revive: identical-switch-conditions

<instructions>
Detects switch statements where two case expressions evaluate to the same value. The second case is unreachable — execution will always match the first one.

Remove the duplicate case expression. If the intent was to check something different, fix the condition.
</instructions>

<examples>
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
- Remove duplicate case expressions caused by copy-paste errors
- Remove constant alias conflicts where different names resolve to the same value in separate cases
- Remove duplicate enum values mapped to the same integer across separate switch cases
- Simplify variable-based switch expressions that evaluate identically
- Replace generated switch code that produces duplicate constant values
</patterns>

<related>
revive/identical-switch-branches, revive/identical-ifelseif-conditions
</related>
