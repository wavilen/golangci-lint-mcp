// Package testdata contains synthetic Go files with correctness issues from compound linters.
// This file covers staticcheck, gocritic, and govet correctness patterns.
// This file does NOT need to compile — it is synthetic test data.
//
// Triggered linters:
//   staticcheck - SA1019 (deprecated API), SA4000 (tautological condition), SA5000 (nil map assignment)
//   gocritic    - badCond (overlapping range), appendAssign (append not assigned), dupImport (duplicate imports), elseif (cascading if-else)
//   govet       - copylocks (mutex copy), loopclosure (captured loop var), printf (wrong format verb)
package testdata

import (
	"fmt"
	"io/ioutil"
	"os"
	"sync"
)

// --- staticcheck issues ---

// sa1019Example uses deprecated io/ioutil.ReadAll.
// staticcheck SA1019: deprecated API usage.
func sa1019Example() {
	// SA1019: io/ioutil.ReadAll is deprecated, use io.ReadAll instead
	data, _ := ioutil.ReadAll(os.Stdin)
	_ = data
}

// sa4000Example contains a tautological condition.
// staticcheck SA4000: expression is always true (x == x).
func sa4000Example(x int) bool {
	// SA4000: x == x is always true
	return x == x
}

// sa5000Example assigns to a nil map.
// staticcheck SA5000: nil map assignment after nil check.
func sa5000Example(m map[string]int) {
	if m == nil {
		// SA5000: assignment to nil map — should initialize with make
		m["key"] = 42
	}
}

// --- gocritic issues ---

// badCondExample has an overlapping range check.
// gocritic badCond: x > 5 && x > 3 — second check is redundant.
func badCondExample(x int) bool {
	// badCond: overlapping range, x > 5 implies x > 3
	return x > 5 && x > 3
}

// appendAssignExample does not assign append result back.
// gocritic appendAssign: append result discarded.
func appendAssignExample(items []int) {
	// appendAssign: append returns a new slice but result is not assigned back
	append(items, 42)
}

// dupImportExample demonstrates duplicate imports.
// gocritic dupImport: same package imported twice.
// NOTE: This pattern is a synthetic example; real Go compilers reject duplicate imports.
// The guide for dupImport documents this anti-pattern.
func dupImportExample() {
	// dupImport: importing "fmt" twice (synthetic — actual dup import requires separate import blocks)
	// In real code this would be caught by the compiler, but dupImport also checks
	// for aliases that resolve to the same package path.
	_ = fmt.Sprintf("example")
}

// elseifExample has a cascading if-else chain.
// gocritic elseif: cascading if-else should be a switch.
func elseifExample(code int) string {
	// elseif: cascading if-else that should be converted to switch
	if code == 200 {
		return "OK"
	} else if code == 404 {
		return "Not Found"
	} else if code == 500 {
		return "Internal Server Error"
	} else if code == 403 {
		return "Forbidden"
	} else {
		return "Unknown"
	}
}

// --- govet issues ---

// copylocksExample copies a sync.Mutex by value.
// govet copylocks: copying a sync.Mutex by assignment.
func copylocksExample() {
	// copylocks: mutex copied by value — use pointer or separate mutex
	var mu1 sync.Mutex
	mu2 := mu1 // copylocks: copies lock by value
	_ = mu2
}

// loopclosureExample captures loop variable in goroutine.
// govet loopclosure: captured loop variable in goroutine (pre-Go 1.22 pattern).
func loopclosureExample(urls []string) {
	// loopclosure: loop variable captured by goroutine closure
	for _, url := range urls {
		go func() {
			// loopclosure: url may be captured from outer loop
			fmt.Println("fetching", url)
		}()
	}
}

// printfExample uses wrong format verb.
// govet printf: %d used for string value.
func printfExample(name string) {
	// printf: wrong format verb — %d is for integers, not strings
	fmt.Printf("Name: %d\n", name)
}
