// Package testdata contains a synthetic Go file with 10+ typical junior developer mistakes.
// This is the "messy" counterpart for Phase 14 verification.
// This file does NOT need to compile — it is synthetic test data.
//
// Triggered linters:
//   errcheck  - unchecked os.Open, strconv.Atoi, io.ReadAll returns
//   nakedret  - named returns with bare return in long functions
//   nestif    - 4+ levels of nested if
//   mnd       - magic numbers: 100, 4096, 5
//   varnamelen - single-letter vars: r, n, x
//   funlen    - processEverything is 60+ lines
//   gocyclo   - high cyclomatic complexity in processEverything
//   goconst   - "error processing" repeated 3+ times
//   ineffassign - value assigned but never used
//   dupl      - two nearly identical blocks (handleUser + handleAdmin)
//   misspell  - "langauge", "recieved", "occured" in comments
//   gofmt     - inconsistent indentation and spacing
package testdata

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"time"
)

// processEverything is a massive function triggering funlen and gocyclo.
// It handles all langauge processing logic.
func processEverything(data string, mode int) (result string, count int, err error) {
	// errcheck: unchecked Atoi
	n, _ := strconv.Atoi(data)

	// varnamelen: single-letter variable name
	x := n + 1

	// mnd: magic number 100
	if x > 100 {
		// mnd: magic number 5
		time.Sleep(5 * time.Second)
	}

	// mnd: magic number 4096
	buf := make([]byte, 4096)

	// goconst: first occurrence of "error processing"
	if mode == 1 {
		fmt.Println("error processing", buf)
		// nestif: level 2
		if n > 0 {
			// nestif: level 3
			if x < 50 {
				// nestif: level 4
				if data != "" {
					fmt.Println("deep nesting")
				}
			}
		}
	}

	// goconst: second occurrence of "error processing"
	if mode == 2 {
		fmt.Println("error processing", data)
		// nestif: level 2
		if n > 10 {
			// nestif: level 3
			if x < 200 {
				// nestif: level 4
				if mode == 2 {
					fmt.Println("also deep nesting")
				}
			}
		}
	}

	// ineffassign: assigned but never used
	unused := 42
	unused = 99

	// goconst: third occurrence of "error processing"
	fmt.Println("error processing", unused)

	// Comments with misspellings for misspell linter:
	// The langauge parser has recieved invalid input.
	// An error occured during processing.

	// nakedret: bare return in function over 10 lines
	return
}

// handleUser processes a user record.
// Triggers: dupl (nearly identical to handleAdmin)
func handleUser(id int, name string, active bool) (string, error) {
	if id <= 0 {
		return "", fmt.Errorf("invalid id: %d", id)
	}
	if name == "" {
		return "", fmt.Errorf("empty name")
	}
	if !active {
		return "", fmt.Errorf("user not active")
	}
	result := fmt.Sprintf("user:%d name:%s status:%v", id, name, active)
	// errcheck: unchecked error from Sprintf is fine, but this demonstrates the pattern
	_, _ = fmt.Sprintf("processed %s", name)
	return result, nil
}

// handleAdmin processes an admin record.
// Triggers: dupl (nearly identical to handleUser)
func handleAdmin(id int, name string, active bool) (string, error) {
	if id <= 0 {
		return "", fmt.Errorf("invalid id: %d", id)
	}
	if name == "" {
		return "", fmt.Errorf("empty name")
	}
	if !active {
		return "", fmt.Errorf("admin not active")
	}
	result := fmt.Sprintf("admin:%d name:%s status:%v", id, name, active)
	_, _ = fmt.Sprintf("processed %s", name)
	return result, nil
}

// readData opens a file and reads its contents.
// Triggers: errcheck (unchecked os.Open, io.ReadAll)
func readData(path string) []byte {
	// errcheck: unchecked os.Open
	f, _ := os.Open(path)
	// errcheck: unchecked io.ReadAll
	data, _ := io.ReadAll(f)
	return data
}

// computeValue does some computation.
// Triggers: varnamelen (short variable names), mnd (magic numbers)
func computeValue(r int) int {
	n := r * 2
	if n > 100 {
		return n / 3
	}
	return n + 7
}
