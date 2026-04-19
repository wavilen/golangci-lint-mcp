// Package testdata contains a well-formatted Go file demonstrating proper Go style.
// This is the "clean" counterpart for Phase 14 verification.
//
// Demonstrates: named error returns, no magic numbers, shallow nesting,
// descriptive variable names, short functions.
package testdata

import (
	"errors"
	"fmt"
	"io"
	"os"
	"strconv"
	"time"
)

const (
	maxRetries     = 3
	requestTimeout = 5 * time.Second
	bufferSize     = 4096
)

// UserData represents a user record with validated fields.
type UserData struct {
	ID    int
	Name  string
	Email string
}

// FetchUser retrieves a user by ID with proper error handling.
func FetchUser(id int) (*UserData, error) {
	if id <= 0 {
		return nil, fmt.Errorf("invalid user id: %d", id)
	}
	return &UserData{ID: id, Name: "example", Email: "user@example.com"}, nil
}

// ParseAndValidate parses a string to int and validates the range.
func ParseAndValidate(input string) (int, error) {
	value, err := strconv.Atoi(input)
	if err != nil {
		return 0, fmt.Errorf("parsing input %q: %w", input, err)
	}
	if value < 0 || value > maxRetries {
		return 0, fmt.Errorf("value %d out of range [0, %d]", value, maxRetries)
	}
	return value, nil
}

// ReadConfig reads a configuration file with proper resource cleanup.
func ReadConfig(path string) ([]byte, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("opening config: %w", err)
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		return nil, fmt.Errorf("reading config: %w", err)
	}
	return data, nil
}

// ProcessItems processes a list of items with descriptive error messages.
func ProcessItems(items []string) ([]string, error) {
	if len(items) == 0 {
		return nil, errors.New("no items to process")
	}

	results := make([]string, 0, len(items))
	for _, item := range items {
		result, err := processOne(item)
		if err != nil {
			return nil, fmt.Errorf("processing item %q: %w", item, err)
		}
		results = append(results, result)
	}
	return results, nil
}

func processOne(item string) (string, error) {
	if item == "" {
		return "", errors.New("empty item")
	}
	return fmt.Sprintf("processed: %s", item), nil
}
