# govet: shift

<instructions>
Reports shift amounts that equal or exceed the word size of the integer type. Shifting a 64-bit integer by 64 or more bits always produces 0 (or all bits for negative values), which is almost certainly a logic error.

Fix the shift amount to a correct value that is less than the width of the type.
</instructions>

<examples>
## Good
```go
var x uint64
result := x >> 63 // shifting by less than word size
```
</examples>

<patterns>
- Use shift amounts under 32 for 32-bit types — mask or bounds-check the value
- Use shift amounts under 64 for 64-bit types — mask or bounds-check the value
- Validate or mask variable shift amounts to stay within the integer type width
</patterns>

<related>
govet/stringintconv, govet/assign
</related>
