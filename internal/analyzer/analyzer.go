package analyzer

import "golang.org/x/tools/go/analysis"

var (
	name = "loglint"
	doc  = "linter for checking logs by aidashov amir"
)

var Analyzer = &analysis.Analyzer{
	Name: name,
	Doc:  doc,
	Run:  run,
}
