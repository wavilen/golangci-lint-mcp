# govet: unmarshal

<instructions>
Reports passing non-pointer values to `json.Unmarshal`, `xml.Unmarshal`, or similar decode functions. Unmarshal needs a pointer to write the decoded data into the target variable. Passing a value or nil is always an error.

Pass a pointer to the target: `json.Unmarshal(data, &target)`.
</instructions>

<examples>
## Good
```go
var cfg Config
json.Unmarshal(data, &cfg) // pointer to target
```
</examples>

<patterns>
- Pass a pointer (not value) to `json.Unmarshal` as the second argument
- Pass a valid non-nil pointer as the `Unmarshal` target
- Pass a pointer to a concrete type (not `*interface{}`) to `Unmarshal`
</patterns>

<related>
govet/structtag, govet/composites
</related>
