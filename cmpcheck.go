// Package cmpcheck define checker for cmp package
package cmpcheck

import (
	"go/ast"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
	"golang.org/x/tools/go/types/typeutil"
)

var Doc = `cmpcheck is checker of cmp package.
This checker checks arguments of cmp package.
This will warn if a1 and a2 are different when using cmp.Diff(a1, a2) or
there is no option in the error comparison.`

var Analyzer = &analysis.Analyzer{
	Name:     "cmpcheck",
	Doc:      "check for unused results of calls to some functions",
	Requires: []*analysis.Analyzer{inspect.Analyzer},
	Run:      run,
}

const (
	cmpDiff   = "cmp.Diff"
	deepEqual = "reflect.DeepEqual"
)

func run(pass *analysis.Pass) (interface{}, error) {
	inspect := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)
	nodeFilter := []ast.Node{
		(*ast.CallExpr)(nil),
	}
	inspect.Preorder(nodeFilter, func(n ast.Node) {
		call := n.(*ast.CallExpr)
		fn := typeutil.StaticCallee(pass.TypesInfo, call)
		if fn == nil {
			return
		}
		if len(call.Args) < 2 {
			return
		}
		fname := fn.FullName()
		if fname != cmpDiff && fname != deepEqual {
			return
		}

		left, right := pass.TypesInfo.Types[call.Args[0]], pass.TypesInfo.Types[call.Args[1]]
		if left.Type.String() != right.Type.String() {
			pass.Reportf(n.Pos(), "%s and %s are not same type", left.Type.String(), right.Type.String())
			return
		}

		if left.Type.String() == "error" && len(call.Args) == 2 {
			switch fname {
			case cmpDiff:
				pass.Reportf(n.Pos(), "required cmpopts.EquateErrors when compare error")
			case deepEqual:
				pass.Reportf(n.Pos(), "error comparison should use errors.Is")
			}
		}
	})

	return nil, nil
}
