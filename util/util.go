package util

import (
	"go/ast"
	"go/token"
	"go/types"
	"strings"

	"golang.org/x/tools/go/analysis"
)

// check the interface n is simple or not
func Check(pass *analysis.Pass, n ast.Node) {
	switch iface := n.(type) {
	case *ast.InterfaceType:
		itp := pass.TypesInfo.TypeOf(iface).(*types.Interface)

		// factor out clearly ok interfaces
		if itp.Empty() || itp.IsMethodSet() {
			return
		}

		// black list
		checkMethodWithRegularType(pass, n, itp)
		checkSameEmbeddedType(pass, n, itp)
		checkEmpty(pass, n, itp)
	}
}

// インターフェースがメソッドと普通の型の両方を含まないことを確認
func checkMethodWithRegularType(pass *analysis.Pass, n ast.Node, itp *types.Interface) {
	if itp.NumMethods() == 0 {
		return
	}

	// factor out underlying types
	for i := 0; i < itp.NumEmbeddeds(); i++ {
		etp := itp.EmbeddedType(i)
		if !underlying(etp) {
			pass.Reportf(n.Pos(), "method and embedded type\n")
			return
		}
	}

	// TODO check type sets of interface elements contain regular type or not
}

// XXX is it robust?
func underlying(etp types.Type) bool {
	return strings.Contains(etp.String(), "~")
}

// blame interface { x; x }, interface { x | y; x }, and interface { x | y; x | z }
func checkSameEmbeddedType(pass *analysis.Pass, n ast.Node, itp *types.Interface) {
	elems := extractIfaceElem(n)
	set := make(map[string]int) // int for detecting second occurrence
	for _, ielem := range elems {
		for _, elem := range ielem {
			if set[elem] == 1 {
				pass.Reportf(n.Pos(), "overwrap %s", elem)
			}
			set[elem] += 1
		}
	}
}

// check the type set of the interface is accidentally empty set
func checkEmpty(pass *analysis.Pass, n ast.Node, itp *types.Interface) {
	// TODO
	// とても長いunionを書くときに、行末の | をうっかり消したときに役立つ
	// interface { int | int64 | int32 | float32 | float64 | string | rune
	//             byte } // --> type setは空集合!
	// interface { int | int64 | int32 | float32 | float64 | string | rune |
	//             byte } // --> ok
}

func extractIfaceElem(n ast.Node) [][]string {
	switch n := n.(type) {
	case *ast.InterfaceType:
		res := make([][]string, len(n.Methods.List))
		if n.Methods == nil {
			return res
		}
		for _, field := range n.Methods.List {
			res = append(res, extractOrElems(field.Type))
		}
		return res
	default:
		return nil
	}
}

func extractOrElems(n ast.Expr) (res []string) {
	switch n := n.(type) {
	case *ast.BinaryExpr:
		if n.Op != token.OR {
			return
		}
		x := n.X
		y := n.Y
		s := y.(*ast.Ident).Name
		res = append(res, s)
		res = append(res, extractOrElems(x)...)
	case *ast.Ident:
		return []string{n.Name}
	}
	return res
}
