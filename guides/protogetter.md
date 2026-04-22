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
- Use `msg.GetField()` instead of direct `msg.Field` access on proto messages
- Guard nested proto field access with nil checks or use generated getters
- Replace `if msg.Status == 1` with `if msg.GetStatus() == 1` for safe access
- Set proto fields via generated setter methods instead of direct assignment
</patterns>

<related>
musttag, nilnil, errcheck
</related>
