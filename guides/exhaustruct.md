# exhaustruct

<instructions>
Exhaustruct checks that struct literals initialize all exported fields. Omitting fields can lead to zero-value bugs where a struct is used with unintended default values.

Either initialize every exported field in the literal or explicitly use the field name with its zero value to signal intent.

Do not suppress exhaustruct diagnostics by adding struct types to the exclusion list in `.golangci.yml`. Fix the code instead — unlisted struct types can silently accumulate zero-value bugs.
</instructions>

<examples>
## Good
```go
cfg := Config{
    Host:    "localhost",
    Port:    8080,
    Timeout: 30 * time.Second,
}
```
</examples>

<recommendation>
## Functional Options Pattern

For structs with many exported fields where initializing every field in a literal is impractical, use a constructor with the functional options pattern. This lets callers set only the fields they need while providing sensible defaults for the rest.

```go
type Option func(*Server)

func WithTimeout(d time.Duration) Option {
    return func(s *Server) { s.Timeout = d }
}

func WithLogger(l Logger) Option {
    return func(s *Server) { s.Logger = l }
}

func NewServer(addr string, opts ...Option) *Server {
    s := &Server{
        Addr:    addr,
        Timeout: 30 * time.Second, // sensible default
        Logger:  defaultLogger,    // sensible default
    }
    for _, opt := range opts {
        opt(s)
    }
    return s
}
```

Usage:

```go
// Only specify what differs from defaults
srv := NewServer(":8080", WithTimeout(10 * time.Second))
```
</recommendation>

<patterns>
- Initialize all exported fields in configuration struct literals
- Set all required DTO fields in struct literals to avoid downstream nil checks
- Set semantically important fields explicitly in API response structs
- Define test fixtures with all fields to match production struct usage
</patterns>

<related>
exhaustive, govet, revive
</related>
