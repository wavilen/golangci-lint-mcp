# revive: confusing-results

<instructions>
Detects functions that return multiple values of the same type in adjacent positions, making it easy to mix up the order at call sites. For example, returning `(int, int, error)` is error-prone because the caller can easily swap the first two values.

Return a struct or named result values to make the meaning of each return value clear at the call site.
</instructions>

<examples>
## Bad
```go
func GetCoordinates() (float64, float64, error) {
    return lat, lng, nil
    // Caller might write: lng, lat, _ := GetCoordinates()
}
```

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
- Functions returning two or more values of the same primitive type
- Lat/lng or x/y coordinate pairs returned as bare floats
- Functions returning `(int, int)` for index and count
- Multiple string returns that can be confused (name vs path, key vs value)
- Functions where callers frequently destructure with `_` for confusion
</patterns>

<related>
confusing-naming, argument-limit
