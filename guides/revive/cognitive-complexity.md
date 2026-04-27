# revive: cognitive-complexity

<instructions>
Measures cognitive complexity — how difficult code is for a human to understand. Unlike cyclomatic complexity, cognitive complexity penalizes nesting depth, sequential boolean operators, and control flow interruptions. High scores mean code is hard to read and maintain.

The root cause is often a god function accumulating complexity from mixing unrelated responsibilities. SRP decomposition addresses this: split the function into focused helpers, each responsible for one rule or concern. The main function becomes a simple orchestrator that reads naturally top-to-bottom.
</instructions>

<examples>
## Good
```go
func calculateDiscount(c Customer, o Order) float64 {
    return loyaltyDiscount(c) +
        seasonalPromo(o) +
        bulkDiscount(o) +
        membershipDiscount(c)
}

func loyaltyDiscount(c Customer) float64 {
    if c.Tier == "gold" && c.Years > 5 {
        return 0.15
    }
    if c.Tier == "gold" {
        return 0.10
    }
    if c.Tier == "silver" && c.Years > 3 {
        return 0.08
    }
    return 0
}

func seasonalPromo(o Order) float64 {
    if o.Season != "holiday" {
        return 0
    }
    if o.Total > 200 {
        return 0.12
    }
    return 0.05
}

func bulkDiscount(o Order) float64 {
    if len(o.Items) >= 10 {
        return 0.07
    }
    if len(o.Items) >= 5 {
        return 0.03
    }
    return 0
}

func membershipDiscount(c Customer) float64 {
    if c.Member && c.Points > 1000 {
        return 0.05
    }
    return 0
}
```
</examples>

<patterns>
- Separate business rule functions with accumulated special cases into focused rule functions
- Separate pricing or discount logic into individual functions per rule category (loyalty, season, bulk)
- Decompose approval workflows into distinct validation, authorization, and notification steps
- Flatten up god functions by extracting interleaved responsibilities into helper functions
- Define granular interfaces instead of monolithic structs to keep each function focused (ISP)
- Define granular interfaces instead of monolithic structs to keep each function focused (ISP)
</patterns>

<related>
revive/cyclomatic, revive/function-length, revive/max-control-nesting, gocognit
</related>
