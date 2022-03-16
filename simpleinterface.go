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
		case *ast.InterfaceType:
			fmt.Println(util.ExtractIfaceElem(n))
			lst := util.ExtractIfaceElem(n)

			s := make(map[string]int)
			for _, ors := range lst {
				for _, elem := range ors {
					if s[elem] == 1 {
						pass.Reportf(n.Pos(), "overwrap %s", elem)
					}
					s[elem] += 1
				}
			}

			ast.Print(nil, n)
		}
	})

	return nil, nil
}
