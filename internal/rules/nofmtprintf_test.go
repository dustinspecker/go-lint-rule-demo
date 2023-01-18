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
