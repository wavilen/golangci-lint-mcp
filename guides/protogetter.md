# protogetter

<instructions>
Protogetter detects incorrect direct access to protobuf message fields. Generated protobuf getters handle nil-safety and default values, while direct field access can panic on nil messages or return zero values silently.

Use the generated getter method (`msg.GetName()`) instead of direct field access (`msg.Name`).
</instructions>

<examples>
## Bad
```go
name := req.Name
```

## Good
```go
name := req.GetName()
```
</examples>

<patterns>
- Direct field access on proto messages: `msg.Field` instead of `msg.GetField()`
- Accessing nested proto fields without nil checks
- Direct access in comparisons: `if msg.Status == 1` instead of `if msg.GetStatus() == 1`
- Assigning to proto fields directly instead of using setters
</patterns>

<related>
musttag, nilnil, errcheck
</related>
