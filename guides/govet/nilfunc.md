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
- Comparing a non-nil function variable with nil
- Comparing a method value with nil
- Function nil check that is always true or always false
</patterns>

<related>
ifaceassert, bools
</related>
