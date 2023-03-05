package linters

import (
	"go/ast"
	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
	"log"
)

var WgAddAnalyzer = &analysis.Analyzer{
	Name: "wgAddAnalyzer",
	Doc:  "Check if calling WaitGroup.add() in anonymous goroutine",
	Requires: []*analysis.Analyzer{
		inspect.Analyzer,
	},
	Run: run,
}

type findWaitAddVisitor struct {
	ident []*ast.Ident
	pass  *analysis.Pass
}

func (c *findWaitAddVisitor) Visit(node ast.Node) ast.Visitor {
	switch n := node.(type) {
	case *ast.SelectorExpr:
		if m, ok := isWaitAddNode(n); ok && c.pass.TypesInfo.TypeOf(m) != nil && c.pass.TypesInfo.TypeOf(m).String() == "sync.WaitGroup" {
			c.ident = append(c.ident, m)
		}
		return c
	case *ast.GoStmt:
		c.ident = nil
	}
	return c
}

/*
*判断当前的node是否是选择表达式，以及其sel节点的Name是否是“Add”，X.Sel:wg.Add()
 */
func isWaitAddNode(node ast.Node) (*ast.Ident, bool) {
	if m, ok := node.(*ast.SelectorExpr); ok && m.Sel.Name == "Add" {
		ident, ok := m.X.(*ast.Ident)
		return ident, ok
	}
	return nil, false
}

func run(pass *analysis.Pass) (interface{}, error) {
	inspect := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)

	fn := func(node ast.Node) {
		defer func() {
			if err := recover(); err != nil {
				log.Println("recover error")
			}
		}()

		gostmt, ok := node.(*ast.GoStmt).Call.Fun.(*ast.FuncLit)
		if !ok {
			return
		}
		fcv := &findWaitAddVisitor{pass: pass}
		ast.Walk(fcv, gostmt)
		for _, ident := range fcv.ident {
			err := analysis.Diagnostic{
				Pos:     ident.Pos(),
				End:     ident.End(),
				Message: "Variable " + ident.Name + " calls Add() in anonymous goroutine",
			}
			pass.Report(err)
		}
	}

	nodeFilter := []ast.Node{
		(*ast.GoStmt)(nil),
	}
	inspect.Preorder(nodeFilter, fn)

	return nil, nil
}
