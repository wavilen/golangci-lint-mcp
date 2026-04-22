// Package testdata contains synthetic Go files with modernization issues from compound linters.
// This file covers modernize, testifylint, ginkgolinter, errorlint, and grouper patterns.
// This file does NOT need to compile — it is synthetic test data.
//
// Triggered linters:
//   modernize     - errorf (errors.New + Sprintf), simplifyrange (index-based range)
//   testifylint   - bool-compare (True with ==), formatter (no format verbs)
//   ginkgolinter  - async-assertion (Expect in goroutine), focus-container (FDescribe)
//   errorlint     - asserts (error equality), comparison (== on errors)
//   grouper       - const (ungrouped consts), import (ungrouped imports)
package testdata

import (
	"errors"
	"fmt"
	"io"
)

// --- modernize issues ---

// errorfExample uses errors.New + fmt.Sprintf instead of fmt.Errorf.
// modernize errorf: use fmt.Errorf with %w instead.
func errorfExample(path string, err error) error {
	// errorf: should be fmt.Errorf("failed to open %s: %w", path, err)
	return errors.New(fmt.Sprintf("failed to open %s: %s", path, err))
}

// simplifyrangeExample uses index-based range instead of value-based range.
// modernize simplifyrange: use for _, v := range instead of for i := range + [i].
func simplifyrangeExample(items []string) []string {
	result := make([]string, 0, len(items))
	// simplifyrange: should use for _, item := range items
	for i := range items {
		result = append(result, items[i])
	}
	return result
}

// --- testifylint issues ---

// boolCompareExample uses assert.True with equality comparison.
// testifylint bool-compare: use assert.Equal instead of assert.True(x == y).
func boolCompareExample() {
	// bool-compare: should be assert.Equal(t, expected, actual)
	// (synthetic — would need import "github.com/stretchr/testify/assert")
	_ = "assert.True(t, x == y) should be assert.Equal(t, x, y)"
}

// formatterExample uses assert.Equalf without format verbs.
// testifylint formatter: assert.Equalf with no format verbs, use assert.Equal.
func formatterExample() {
	// formatter: assert.Equalf(t, x, y, "message") with no format verbs
	// should be assert.Equal(t, x, y, "message") or add format verb
	_ = "assert.Equalf(t, got, want, 'no verbs here')"
}

// --- ginkgolinter issues ---

// asyncAssertionExample uses Expect in a goroutine.
// ginkgolinter async-assertion: Expect should use Eventually or Consistently in goroutines.
func asyncAssertionExample() {
	// async-assertion: using Expect in a goroutine is flaky
	// should use Eventually(func() ...) instead
	_ = "go func() { Expect(result).To(Equal(expected)) }()"
}

// focusContainerExample uses a focus container (FDescribe).
// ginkgolinter focus-container: FDescribe should not be committed.
func focusContainerExample() {
	// focus-container: FDescribe/FContext/FIt are debug aids, not for production
	_ = "FDescribe('debug focus', func() { ... })"
}

// --- errorlint issues ---

// errorlintAssertsExample compares errors using direct equality.
// errorlint asserts: use errors.Is() or errors.As() for error comparison.
func errorlintAssertsExample(err error, target error) bool {
	// asserts: should use errors.Is(err, target) instead of err == target
	return err == target
}

// errorlintComparisonExample compares errors with == operator.
// errorlint comparison: use errors.Is() for error comparison.
func errorlintComparisonExample(err error) bool {
	// comparison: should use errors.Is(err, io.EOF) instead of ==
	return err == io.EOF
}

// --- grouper issues ---

// grouperConstExample has ungrouped const declarations.
// grouper const: related constants should be grouped.
func grouperConstExample() int {
	// grouper const: these related constants should be grouped in a single const block
	const maxRetries = 3
	const defaultTimeout = 30
	const bufferSize = 4096
	return maxRetries + defaultTimeout + bufferSize
}

// grouperImportExample has ungrouped import blocks.
// grouper import: import statements should be grouped.
// (This is a synthetic representation — actual imports are at file top)
func grouperImportExample() {
	// grouper import: multiple import blocks should be merged into one
	// In real code: separate import (  ) blocks for "fmt" and "errors"
	// should be combined into a single import block
	_ = fmt.Sprintf("example")
	_ = errors.New("example")
}
