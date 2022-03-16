package simpleinterface

import (
	"fmt"
	"go/ast"

	"github.com/naoyafurudono/simpleinterface/util"
	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
)

const doc = "simpleinterface is ..."

// Analyzer is ...
var Analyzer = &analysis.Analyzer{
	Name: "simpleinterface",
	Doc:  doc,
	Run:  run,
	Requires: []*analysis.Analyzer{
		inspect.Analyzer,
	},
}

func run(pass *analysis.Pass) (interface{}, error) {
	inspect := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)

	inspect.Preorder(nil, func(n ast.Node) {
		switch n := n.(type) {
		case *ast.TypeSpec:
			pass.Reportf(n.Pos(), "type spec")
		case *ast.InterfaceType:
			fmt.Println(util.ExtractIfaceElem(n))
			lst := util.ExtractIfaceElem(n)
			s := make(map[string]bool)
			for _, ors := range lst {
				for _, elem := range ors {
					if s[elem] {
						pass.Reportf(n.Pos(), "overwrap %s", elem)
					}
					s[elem] = true
				}
			}
			// for i :=0; i < len(lst); i++ {
			// 	for j:=0; j< len(lst); j++ {

			// 	}
			// }
			// pass.Reportf(n.Pos(), "iface")
		}
	})

	return nil, nil
}
