package main

import (
	"golang.org/x/tools/go/analysis/singlechecker"

	"go-lint-rule-demo/internal/rules"
)

func main() {
	singlechecker.Main(rules.NoFmtPrintf)
}
