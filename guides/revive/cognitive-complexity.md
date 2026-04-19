# revive: cognitive-complexity

<instructions>
Measures cognitive complexity — how difficult code is for a human to understand. Unlike cyclomatic complexity, cognitive complexity penalizes nesting depth, sequential boolean operators, and control flow interruptions. High scores mean code is hard to read and maintain.

The root cause is often a god function accumulating complexity from mixing unrelated responsibilities. SRP decomposition addresses this: split the function into focused helpers, each responsible for one rule or concern. The main function becomes a simple orchestrator that reads naturally top-to-bottom.
</instructions>

<examples>
## Bad
```go
func calculateDiscount(c Customer, o Order) float64 {
    var disc float64
    if c.Tier == "gold" {
        if c.Years > 5 {
            disc += 0.15
        } else {
            disc += 0.10
        }
    } else if c.Tier == "silver" {
        if c.Years > 3 {
            disc += 0.08
        }
    }
    if o.Season == "holiday" {
        if o.Total > 200 {
            disc += 0.12
        } else {
            disc += 0.05
        }
    }
    if len(o.Items) >= 10 {
        disc += 0.07
    } else if len(o.Items) >= 5 {
        disc += 0.03
    }
    if c.Member && c.Points > 1000 {
        disc += 0.05
    }
    return disc
}
```

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
- Business rule functions with accumulated special cases — split into focused rule functions
- Pricing or discount logic mixing multiple rule categories (loyalty, season, bulk)
- Approval workflows combining validation, authorization, and notification in one body
- God functions where nesting arises from interleaved unrelated responsibilities
- Add new discount rules without modifying the calculator (OCP)
- Define small DiscountRule interface instead of monolithic pricing struct (ISP)
</patterns>

<related>
cyclomatic, function-length, max-control-nesting, gocognit
</related>
