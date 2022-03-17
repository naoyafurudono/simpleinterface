package util

import (
	"go/ast"
	"go/token"
)

func ExtractIfaceElem(n ast.Node) [][]string {
	switch n := n.(type) {
	case *ast.InterfaceType:
		res := make([][]string, len(n.Methods.List))
		for _, field := range n.Methods.List { // nil check
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
