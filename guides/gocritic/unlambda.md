# gocritic: unlambda

<instructions>
Detects unnecessary anonymous function wrappers (lambdas) where the function body simply calls another function with the same arguments. For example, `func(x int) { f(x) }` can be replaced with `f` directly.

Replace the lambda with a direct reference to the target function.
</instructions>

<examples>
## Bad
```go
slices.SortFunc(items, func(a, b Item) int {
	return compare(a, b)
})
```

## Good
```go
slices.SortFunc(items, compare)
```
</examples>

<patterns>
- `func(x T) { f(x) }` → `f`
- Lambda wrapping a method call: `func() { obj.Method() }` when no deferred eval needed
- Anonymous function in `sort.Slice` that just calls a comparator
</patterns>

<related>
deferUnlambda, wrapperFunc, methodExprCall
</related>
