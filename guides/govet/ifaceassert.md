# govet: ifaceassert

<instructions>
Reports impossible interface-to-interface type assertions that can be statically proven to always fail or always succeed. For example, asserting `io.Reader` to `io.Writer` always fails since no interface relationship exists. These are logic errors.

Remove the impossible assertion or correct the interface types to reflect a valid relationship.
</instructions>

<examples>
## Bad
```go
var r io.Reader
w := r.(io.Writer) // Reader never implements Writer — always panics
```

## Good
```go
var r io.ReadWriter // a type that implements both
w := r.(io.Writer)  // valid assertion
```
</examples>

<patterns>
- Asserting to an interface the source cannot implement
- Asserting to the same interface (always succeeds — useless)
- Interface assertion where neither is a subset of the other
</patterns>

<related>
nilfunc, errorsas
</related>
