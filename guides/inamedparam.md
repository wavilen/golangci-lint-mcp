# inamedparam

<instructions>
Inamedparam enforces named parameters on interface types. Unnamed interface parameters like `fn(int, string)` make it unclear what each argument represents.

Name all parameters in interface method signatures. Use descriptive names that convey the purpose of each parameter.
</instructions>

<examples>
## Bad
```go
type Processor interface {
    Process(context.Context, []byte) error
}
```

## Good
```go
type Processor interface {
    Process(ctx context.Context, payload []byte) error
}
```
</examples>

<patterns>
- Name all parameters in interface method signatures
- Add descriptive names to callback parameters in interface definitions
- Replace bare type signatures with named parameters in interfaces
- Give descriptive names even to single-parameter interface methods
</patterns>

<related>
funcorder, nonamedreturns, tagliatelle
</related>
