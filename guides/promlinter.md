# promlinter

<instructions>
Promlinter checks Prometheus metric naming conventions. Metrics must follow the Prometheus naming best practices: lowercase with underscores, include a unit suffix, and avoid colons in names (reserved for aggregation rules).

Rename metrics to follow the `[namespace]_[subsystem]_[name]_[unit]` convention.
</instructions>

<examples>
## Bad
```go
var httpRequestDuration = prometheus.NewHistogram(prometheus.HistogramOpts{
    Name: "http_request_duration",
})
```

## Good
```go
var httpRequestDuration = prometheus.NewHistogram(prometheus.HistogramOpts{
    Name: "http_request_duration_seconds",
})
```
</examples>

<patterns>
- Missing unit suffix: `duration` instead of `duration_seconds`
- CamelCase metric or label names instead of snake_case
- Colons in metric names (reserved for recording rules)
- Non-lowercase characters in metric or label names
</patterns>

<related>
godot, govet
</related>
