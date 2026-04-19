# revive: time-equal

<instructions>
Detects time comparisons using `==` or `!=` operators. Time values in Go contain internal monotonic clock readings and location data that make direct equality comparison unreliable. Use the `Equal` method which correctly compares the instant in time.

Replace `t1 == t2` with `t1.Equal(t2)` and `t1 != t2` with `!t1.Equal(t2)`.
</instructions>

<examples>
## Bad
```go
if startTime == endTime {
    return errors.New("zero duration")
}
if t != time.Time{} {
    process(t)
}
```

## Good
```go
if startTime.Equal(endTime) {
    return errors.New("zero duration")
}
if !t.IsZero() {
    process(t)
}
```
</examples>

<patterns>
- Comparing two `time.Time` values with `==` or `!=`
- Checking if a time is the zero value with `t == time.Time{}`
- Using `==` to compare times from different time zones
- Comparing time values returned from different sources (clock vs parsed)
- Testing time equality in table-driven tests
</patterns>

<related>
time-date, time-naming, staticcheck/SA9004
