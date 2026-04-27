# promlinter

<instructions>
Promlinter checks Prometheus metric naming conventions. Metrics must follow the Prometheus naming best practices: lowercase with underscores, include a unit suffix, and avoid colons in names (reserved for aggregation rules).

Rename metrics to follow the `[namespace]_[subsystem]_[name]_[unit]` convention.
</instructions>

<examples>
## Good
```go
var httpRequestDuration = prometheus.NewHistogram(prometheus.HistogramOpts{
    Name: "http_request_duration_seconds",
})
```
</examples>

<patterns>
- Add unit suffixes to metric names: `duration` → `duration_seconds`
- Rename CamelCase metric and label names to `snake_case`
- Avoid colons in metric names — reserve them for recording rules
- Use lowercase only for metric and label names
</patterns>

<related>
godot, govet
</related>
