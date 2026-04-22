# revive: range-val-address

<instructions>
Detects taking the address of a range loop variable (`&v` inside `for _, v := range`). Before Go 1.22, range loop variables were reused across iterations, so taking their address gave the same pointer every time — a subtle bug. Even with Go 1.22+ semantics, this pattern signals unclear intent.

Assign the value to a new variable inside the loop before taking its address, or index directly into the slice to get a stable pointer.
</instructions>

<examples>
## Bad
```go
for _, v := range items {
    ptrs = append(ptrs, &v)
}
```

## Good
```go
for i := range items {
    ptrs = append(ptrs, &items[i])
}
```
</examples>

<patterns>
- Use `&items[i]` instead of `&v` when appending pointers from a range loop
- Pass values directly to goroutines as parameters instead of taking `&v` inside range loops
- Use indexing into the original slice instead of storing pointers to range variables in a map
- Use the range value to a local variable before taking its address
- Use a copy of the value in a closure instead of `&v` when the closure runs after the loop
</patterns>

<related>
range-val-in-closure, datarace
