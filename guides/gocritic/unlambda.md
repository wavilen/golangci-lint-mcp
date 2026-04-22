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
- Replace `func(x T) { f(x) }` with `f` when the function signature matches
- Replace lambda-wrapped method calls `func() { obj.Method() }` with `obj.Method` when no deferred evaluation needed
- Replace anonymous functions in `sort.Slice` that just call a comparator with a direct reference
</patterns>

<related>
deferUnlambda, wrapperFunc, methodExprCall
</related>
