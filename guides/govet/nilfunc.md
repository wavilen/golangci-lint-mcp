# govet: nilfunc

<instructions>
Reports comparisons of functions with `nil` that are always false or always true. Functions are never nil in certain contexts (e.g., methods with value receivers), making the comparison misleading.

Remove the unnecessary nil comparison or restructure to use a function pointer that can actually be nil.
</instructions>

<examples>
## Bad
```go
var fn func() = someFunc
if fn == nil { // fn is never nil — always false
    return
}
```

## Good
```go
var fn func()
if fn == nil { // fn is actually nil — valid check
    fn = defaultHandler
}
```
</examples>

<patterns>
- Remove nil comparisons on non-nil function variables — the check is always false
- Remove nil comparisons on method values — methods with value receivers are never nil
- Remove tautological function nil checks — use a function pointer that can actually be nil
</patterns>

<related>
ifaceassert, bools
</related>
