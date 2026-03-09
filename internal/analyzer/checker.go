package analyzer

import (
	"go/ast"

	"golang.org/x/tools/go/analysis"
)

func run(pass *analysis.Pass) (any, error) {
	for _, file := range pass.Files {
		ast.Inspect(file, func(n ast.Node) bool {
			call, ok := n.(*ast.CallExpr)
			if !ok {
				return true
			}

			sel, ok := call.Fun.(*ast.SelectorExpr)
			if !ok {
				return true
			}

			switch sel.Sel.Name {
			case "Info", "Debug", "Warn", "Error":
				if s, ok := pass.TypesInfo.Selections[sel]; ok {
					pkgPath := s.Obj().Pkg().Path()
					if pkgPath == "log/slog" || pkgPath == "go.uber.org/zap" {
						check(pass, call)
					}
				}
			default:
				return true
			}

			return true
		})
	}

	return nil, nil
}

func check(pass *analysis.Pass, call *ast.CallExpr) {

}
