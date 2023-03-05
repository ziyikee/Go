package utils

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"testing"
)

var srcCode = `
	package hello
	
	import "fmt"
	
	func greet() {
		var msg = "Hello World!"
		fmt.Println(msg)
	}
`

func Test(t *testing.T) {
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "", srcCode, 0)
	if err != nil {
		fmt.Printf("err = %s", err)
	}
	ast.Print(fset, f)
}
