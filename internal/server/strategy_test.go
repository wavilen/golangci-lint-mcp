package server

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExtractPackagesFromIssues(t *testing.T) {
	issues := []lintIssue{
		{Pos: struct {
			Filename string `json:"Filename"`
			Line     int    `json:"Line"`
			Column   int    `json:"Column"`
		}{Filename: "pkg/auth/handler.go", Line: 1, Column: 1}},
		{Pos: struct {
			Filename string `json:"Filename"`
			Line     int    `json:"Line"`
			Column   int    `json:"Column"`
		}{Filename: "pkg/auth/middleware.go", Line: 2, Column: 1}},
		{Pos: struct {
			Filename string `json:"Filename"`
			Line     int    `json:"Line"`
			Column   int    `json:"Column"`
		}{Filename: "pkg/db/connection.go", Line: 3, Column: 1}},
		{Pos: struct {
			Filename string `json:"Filename"`
			Line     int    `json:"Line"`
			Column   int    `json:"Column"`
		}{Filename: "main.go", Line: 4, Column: 1}},
	}

	packages := extractPackagesFromIssues(issues)

	assert.Len(t, packages, 3)

	// pkg/auth has 2 issues, should be first
	assert.Equal(t, "pkg/auth", packages[0].path)
	assert.Equal(t, 2, packages[0].count)

	// root ("." for main.go) has 1 issue — "." sorts before "pkg/" alphabetically
	assert.Equal(t, ".", packages[1].path)
	assert.Equal(t, 1, packages[1].count)

	// pkg/db has 1 issue
	assert.Equal(t, "pkg/db", packages[2].path)
	assert.Equal(t, 1, packages[2].count)
}

func TestExtractPackagesFromIssues_Empty(t *testing.T) {
	packages := extractPackagesFromIssues(nil)
	assert.Empty(t, packages)
}

func TestBuildPackageBreakdown(t *testing.T) {
	packages := []packageEntry{
		{path: "pkg/auth", count: 15},
		{path: "pkg/db", count: 8},
		{path: "pkg/util", count: 3},
	}

	result := buildPackageBreakdown(packages)

	assert.Contains(t, result, "pkg/auth: 15 issues")
	assert.Contains(t, result, "pkg/db: 8 issues")
	assert.Contains(t, result, "pkg/util: 3 issues")
	assert.Contains(t, result, "TOTAL: 26 issues across 3 packages")
}

func TestBuildPackageBreakdown_Empty(t *testing.T) {
	result := buildPackageBreakdown(nil)
	assert.Empty(t, result)
}

func TestRecommendStrategy_SingleAgent(t *testing.T) {
	name, reason := recommendStrategy(20, 2)
	assert.Equal(t, "single-agent", name)
	assert.Contains(t, reason, "≤20")
	assert.Contains(t, reason, "2 packages")
	assert.Contains(t, reason, "single-agent flow")
}

func TestRecommendStrategy_SingleAgent_EdgeCase(t *testing.T) {
	name, _ := recommendStrategy(30, 3)
	assert.Equal(t, "single-agent", name)
}

func TestRecommendStrategy_Subagent(t *testing.T) {
	name, reason := recommendStrategy(45, 2)
	assert.Equal(t, "subagent-per-package", name)
	assert.Contains(t, reason, ">45")
	assert.Contains(t, reason, "2 packages")
	assert.Contains(t, reason, "subagent-per-package strategy")
}

func TestRecommendStrategy_SubagentByPackages(t *testing.T) {
	name, reason := recommendStrategy(5, 5)
	assert.Equal(t, "subagent-per-package", name)
	assert.Contains(t, reason, ">5")
	assert.Contains(t, reason, "5 packages")
}

func TestIsLargeOutput_TrueByIssues(t *testing.T) {
	assert.True(t, isLargeOutput(50, 1))
}

func TestIsLargeOutput_TrueByPackages(t *testing.T) {
	assert.True(t, isLargeOutput(5, 5))
}

func TestIsLargeOutput_False(t *testing.T) {
	assert.False(t, isLargeOutput(20, 2))
}

func TestIsLargeOutput_EdgeCase(t *testing.T) {
	// At exactly the threshold should be false (> not >=)
	assert.False(t, isLargeOutput(30, 3))
}

func TestBuildPackageBreakdown_Ordering(t *testing.T) {
	// Input is already sorted by count desc, path asc (as extractPackagesFromIssues returns)
	packages := []packageEntry{
		{path: "pkg/c", count: 10},
		{path: "pkg/a", count: 5},
		{path: "pkg/b", count: 5},
	}

	result := buildPackageBreakdown(packages)
	lines := strings.Split(result, "\n")

	// First line should be pkg/c (highest count)
	assert.Contains(t, lines[0], "pkg/c: 10 issues")
	// Next two should be sorted by path (a before b) — already sorted by caller
	assert.Contains(t, lines[1], "pkg/a: 5 issues")
	assert.Contains(t, lines[2], "pkg/b: 5 issues")
}
