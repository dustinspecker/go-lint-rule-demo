package rules

import (
	"go/ast"
	"go/types"

	"golang.org/x/tools/go/analysis"
)

var NoFmtPrintf = &analysis.Analyzer{
	Name: "nofmtprintf",
	Doc:  "Avoid fmt.Printf",
	Run:  noFmtPrintfRun,
}

func noFmtPrintfRun(pass *analysis.Pass) (interface{}, error) {
	var fmtPrintfType types.Type

	for _, pkg := range pass.Pkg.Imports() {
		if pkg.Name() == "fmt" {
			fmtPrintfType = pkg.Scope().Lookup("Printf").Type()
		}
	}

	for _, file := range pass.Files {
		ast.Inspect(file, func(node ast.Node) bool {
			// example all function calls
			callExpr, isCallExpr := node.(*ast.CallExpr)
			if !isCallExpr {
				return true
			}

			callExprFunType := pass.TypesInfo.TypeOf(callExpr.Fun)
			if callExprFunType == fmtPrintfType {
				pass.Report(analysis.Diagnostic{
					Pos:     node.Pos(),
					End:     node.End(),
					Message: "Don't use fmt.Printf",
				})
			}

			return true
		})
	}

	return nil, nil
}
