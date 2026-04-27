# revive: confusing-results

<instructions>
Detects functions that return multiple values of the same type in adjacent positions, making it easy to mix up the order at call sites. For example, returning `(int, int, error)` is error-prone because the caller can easily swap the first two values.

Return a struct or named result values to make the meaning of each return value clear at the call site.
</instructions>

<examples>
## Good
```go
type Coordinates struct {
    Lat float64
    Lng float64
}

func GetCoordinates() (Coordinates, error) {
    return Coordinates{Lat: lat, Lng: lng}, nil
}
```
</examples>

<patterns>
- Return a struct instead of multiple values of the same primitive type
- Wrap coordinate pairs (lat/lng, x/y) in a named struct to prevent argument swaps
- Use named return values or a result struct for functions returning `(int, int)` pairs
- Define a result struct when returning multiple strings that could be confused
- Replace same-type multi-returns with a struct to prevent callers from destructuring with `_`
</patterns>

<related>
revive/confusing-naming, revive/argument-limit
</related>
