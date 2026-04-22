# durationcheck

<instructions>
Durationcheck detects cases where two `time.Duration` values are multiplied together, producing a result in nanoseconds² rather than the intended duration. This happens because `time.Duration` is `int64` nanoseconds, so multiplying two durations squares the unit.

Convert one operand to a plain number before multiplying, or restructure the expression.
</instructions>

<examples>
## Bad
```go
delay := time.Second * time.Duration(count)
```

## Good
```go
delay := time.Second * time.Duration(count)
// If both are durations, convert one:
// delay := time.Duration(a.Seconds() * b.Seconds()) * time.Second
```
</examples>

<patterns>
- Avoid multiplying two `time.Duration` values — the result is in nanoseconds²
- Convert one operand to a plain number before multiplying durations in backoff calculations
- Use `time.Duration(float64(a) * float64(b))` or extract seconds before multiplying
</patterns>

<related>
govet, staticcheck, gosec
