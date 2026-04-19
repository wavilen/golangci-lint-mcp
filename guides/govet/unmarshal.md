# govet: unmarshal

<instructions>
Reports passing non-pointer values to `json.Unmarshal`, `xml.Unmarshal`, or similar decode functions. Unmarshal needs a pointer to write the decoded data into the target variable. Passing a value or nil is always an error.

Pass a pointer to the target: `json.Unmarshal(data, &target)`.
</instructions>

<examples>
## Bad
```go
var cfg Config
json.Unmarshal(data, cfg) // value, not pointer — compile error at runtime
```

## Good
```go
var cfg Config
json.Unmarshal(data, &cfg) // pointer to target
```
</examples>

<patterns>
- Passing value (not pointer) to `Unmarshal`
- Passing nil as the target
- Passing pointer to interface instead of pointer to concrete type
</patterns>

<related>
structtag, composites
</related>
