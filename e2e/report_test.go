package e2e_test

import (
	"fmt"
	"html/template"
	"os"
	"time"
)

// ReportData is the template data for the HTML report.
type ReportData struct {
	Timestamp    time.Time
	Results      []*EvalResult
	Models       []string
	Fixtures     []string
	PassCount    int
	TotalCount   int
	PassRate     float64
	AvgReduction float64
	Errors       []string
}

// GenerateReport generates an HTML report from test results.
func GenerateReport(results []*EvalResult, outputPath string) error {
	modelsSeen := make(map[string]bool)
	fixturesSeen := make(map[string]bool)
	passCount := 0
	totalReduction := 0.0
	var errors []string

	for _, r := range results {
		modelsSeen[r.Model] = true
		fixturesSeen[r.TestCase] = true
		if r.Pass {
			passCount++
		}
		totalReduction += r.IssueReduction
		if r.Error != "" {
			errors = append(errors, fmt.Sprintf("%s/%s: %s", r.TestCase, r.Model, r.Error))
		}
	}

	models := make([]string, 0, len(modelsSeen))
	for m := range modelsSeen {
		models = append(models, m)
	}
	fixtures := make([]string, 0, len(fixturesSeen))
	for f := range fixturesSeen {
		fixtures = append(fixtures, f)
	}

	avgReduction := 0.0
	if len(results) > 0 {
		avgReduction = totalReduction / float64(len(results))
	}

	data := ReportData{
		Timestamp:    time.Now(),
		Results:      results,
		Models:       models,
		Fixtures:     fixtures,
		PassCount:    passCount,
		TotalCount:   len(results),
		PassRate:     0,
		AvgReduction: avgReduction,
		Errors:       errors,
	}
	if len(results) > 0 {
		data.PassRate = float64(passCount) / float64(len(results)) * 100
	}

	tmpl, err := template.ParseFiles("report.html.tpl")
	if err != nil {
		return fmt.Errorf("parsing template: %w", err)
	}

	f, err := os.Create(outputPath)
	if err != nil {
		return fmt.Errorf("creating report: %w", err)
	}
	defer f.Close()

	return tmpl.Execute(f, data)
}
