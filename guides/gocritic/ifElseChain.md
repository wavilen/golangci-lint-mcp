# gocritic: ifElseChain

<instructions>
Detects long `if`/`else if`/`else` chains that would be clearer as a `switch` statement. When three or more conditions check the same variable or expression, a `switch` is more readable and idiomatic in Go.

Convert the chain to a `switch` statement on the repeated expression.
</instructions>

<examples>
## Bad
```go
if color == "red" {
	red()
} else if color == "green" {
	green()
} else if color == "blue" {
	blue()
} else {
	other()
}
```

## Good
```go
switch color {
case "red":
	red()
case "green":
	green()
case "blue":
	blue()
default:
	other()
}
```
</examples>

<patterns>
- 3+ `else if` branches checking the same variable
- Multiple `if`/`else if` comparing against constants
- Sequential equality checks that map to `switch` cases
</patterns>

<related>
elseif, singleCaseSwitch, switchTrue
</related>
