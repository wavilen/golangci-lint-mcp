// Package testdata contains synthetic Go files with security and style issues from compound linters.
// This file covers gosec and revive security/style patterns.
// This file does NOT need to compile — it is synthetic test data.
//
// Triggered linters:
//   gosec  - G101 (hardcoded credentials), G201 (SQL injection), G301 (permissive mkdir), G401 (weak crypto), G501 (blocklisted import)
//   revive - add-constant (magic numbers), bare-return (bare returns), cognitive-complexity (complex function), unreachable-func (dead code), unused-parameter (unused params)
package testdata

import (
	"crypto/md5"
	"crypto/sha1"
	"database/sql"
	"errors"
	"fmt"
	"os"
)

// --- gosec issues ---

// g101Example has a hardcoded credential variable.
// gosec G101: hardcoded credential detected in variable name.
func g101Example() {
	// G101: password is hardcoded — move to environment variable or secrets manager
	password := "super-secret-password-123"
	_ = password
}

// g201Example builds a SQL query with string formatting.
// gosec G201: SQL query built with fmt.Sprintf — SQL injection risk.
func g201Example(db *sql.DB, userID string) {
	// G201: SQL injection via fmt.Sprintf
	query := fmt.Sprintf("SELECT * FROM users WHERE id = '%s'", userID)
	_, _ = db.Query(query)
}

// g301Example creates a directory with overly permissive permissions.
// gosec G301: mkdir with 0777 permissions.
func g301Example(path string) {
	// G301: overly permissive mkdir — use 0750 or more restrictive
	_ = os.MkdirAll(path, 0777)
}

// g401Example uses weak cryptographic primitive.
// gosec G401: use of weak crypto (MD5).
func g401Example(data []byte) {
	// G401: MD5 is cryptographically broken, use SHA-256
	_ = md5.Sum(data)
}

// g501Example imports a blocklisted package.
// gosec G501: import of crypto/sha1 (blocklisted).
func g501Example(data []byte) {
	// G501: SHA-1 is cryptographically weak, use SHA-256
	_ = sha1.Sum(data)
}

// --- revive issues ---

// addConstantExample uses magic numbers without named constants.
// revive add-constant: magic numbers should be named constants.
func addConstantExample(timeout int) bool {
	// add-constant: magic numbers 30, 404, 500 should be named constants
	if timeout > 30 {
		return true
	}
	_ = 404
	_ = 500
	return false
}

// bareReturnExample uses bare return with named return values.
// revive bare-return: named return with bare return statement.
func bareReturnExample(input string) (result string, err error) {
	// bare-return: explicit return values are preferred for clarity
	if input == "" {
		err = errors.New("empty input")
		return // bare-return: implicit return of named values
	}
	result = "processed: " + input
	return // bare-return: implicit return of named values
}

// cognitiveComplexityExample has high cognitive complexity.
// revive cognitive-complexity: nested conditions exceed threshold.
func cognitiveComplexityExample(a, b, c int) string {
	// cognitive-complexity: deeply nested conditions — refactor to reduce complexity
	if a > 0 {
		if b > 0 {
			if c > 0 {
				if a > b {
					if b > c {
						return "a > b > c"
					} else {
						return "nested else 1"
					}
				} else {
					if a > c {
						return "b >= a > c"
					} else {
						return "nested else 2"
					}
				}
			}
		}
	}
	return "default"
}

// unreachableFuncExample has unreachable code after return.
// revive unreachable-func: unreachable function call after return.
func unreachableFuncExample(x int) int {
	if x < 0 {
		return 0
	}
	return x
	// unreachable-func: this call is unreachable due to return above
	_ = fmt.Sprintf("unreachable: %d", x)
}

// unusedParameterExample has an unused function parameter.
// revive unused-parameter: parameter 'verbose' is never used.
func unusedParameterExample(data string, verbose bool) error {
	// unused-parameter: verbose is never referenced in function body
	_ = data
	return nil
}
