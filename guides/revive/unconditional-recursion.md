# revive: unconditional-recursion

<instructions>
Detects functions that unconditionally call themselves without a base case, causing infinite recursion and eventual stack overflow. This typically happens when a method receiver name is a typo of the type name or when a wrapper function forgets to delegate to the inner implementation.

Add a termination condition, or fix the call to target the correct function/method instead of recursing into itself.
</instructions>

<examples>
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
- Call the wrapped field instead of recursing into the method itself
- Add a base case to recursive functions that currently recurse unconditionally
- Replace typos where the receiver name is confused with a same-named struct field
- Call the inner method in interface satisfaction wrappers instead of the wrapper itself
- Avoid calling `String()` or `Error()` on the same type within their own implementations
</patterns>

<related>
revive/unreachable-code, revive/datarace
</related>
