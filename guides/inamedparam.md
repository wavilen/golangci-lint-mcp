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
- Interface methods with only type parameters and no names
- Callback function signatures in interfaces with unnamed params
- `func(int, error)` style signatures in interface definitions
- Single-parameter interfaces where the name adds context
</patterns>

<related>
funcorder, nonamedreturns, tagliatelle
</related>
