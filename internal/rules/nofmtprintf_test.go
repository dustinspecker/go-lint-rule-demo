package rules_test

import (
	"testing"

	"golang.org/x/tools/go/analysis/analysistest"

	"go-lint-rule-demo/internal/rules"
)

func TestNoFmtPrintf(t *testing.T) {
	testDataDir := analysistest.TestData()

	analysistest.Run(t, testDataDir, rules.NoFmtPrintf, "./src/nofmtprintf/")
}

func TestNoFmtPrintfAutoFix(t *testing.T) {
	testDataDir := analysistest.TestData()

	results := analysistest.RunWithSuggestedFixes(t, testDataDir, rules.NoFmtPrintf, "./src/nofmtprintf/")

	suggestedFixProvided := false
	for _, result := range results {
		for _, diagnostic := range result.Diagnostics {
			for _, suggestedFix := range diagnostic.SuggestedFixes {
				if len(suggestedFix.TextEdits) != 0 {
					suggestedFixProvided = true
				}
			}
		}
	}

	if !suggestedFixProvided {
		t.Errorf("expected a suggested fix to be provided, but didn't have any in %+v", results)
	}
}
