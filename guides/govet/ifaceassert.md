# govet: ifaceassert

<instructions>
Reports impossible interface-to-interface type assertions that can be statically proven to always fail or always succeed. For example, asserting `io.Reader` to `io.Writer` always fails since no interface relationship exists. These are logic errors.

Remove the impossible assertion or correct the interface types to reflect a valid relationship.
</instructions>

<examples>
## Good
```go
var r io.ReadWriter // a type that implements both
w := r.(io.Writer)  // valid assertion
```
</examples>

<patterns>
- Remove impossible interface assertions where the source cannot implement the target
- Remove redundant same-interface assertions (`r.(io.Reader)` on an `io.Reader`) — always succeeds
- Remove interface assertions with no overlapping method set — use a valid type or interface
</patterns>

<related>
govet/nilfunc, govet/errorsas
</related>
