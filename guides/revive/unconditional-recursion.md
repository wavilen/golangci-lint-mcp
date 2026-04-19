# revive: unconditional-recursion

<instructions>
Detects functions that unconditionally call themselves without a base case, causing infinite recursion and eventual stack overflow. This typically happens when a method receiver name is a typo of the type name or when a wrapper function forgets to delegate to the inner implementation.

Add a termination condition, or fix the call to target the correct function/method instead of recursing into itself.
</instructions>

<examples>
## Bad
```go
func (s *Server) Process(req *Request) error {
    return s.Process(req) // infinite recursion
}

func factorial(n int) int {
    return n * factorial(n-1) // no base case
}
```

## Good
```go
func (s *Server) Process(req *Request) error {
    return s.handler.Process(req) // delegate correctly
}

func factorial(n int) int {
    if n <= 1 {
        return 1
    }
    return n * factorial(n-1)
}
```
</examples>

<patterns>
- Method calling itself instead of the wrapped/delegated field
- Recursive functions missing a base case
- Typos where the receiver name is confused with the struct field name
- Interface satisfaction wrappers that call themselves instead of the inner method
- Accidental recursion in `String()` or `Error()` methods
</patterns>

<related>
unreachable-code, datarace
