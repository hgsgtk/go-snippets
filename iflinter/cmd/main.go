package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"log"
)

func main() {
	fst := token.NewFileSet()
	f, err := parser.ParseFile(fst, "simple.go", nil, 0)
	if err != nil {
		log.Fatal("Error:", err)
	}

	ast.Inspect(f, func(n ast.Node) bool {

		expr, ok := n.(*ast.BinaryExpr)
		if !ok {
			return true
		}

		if expr.Op == token.EQL {
			xIdent, ok := expr.X.(*ast.Ident)
			if !ok {
				return true
			}
			yIdent, ok := expr.Y.(*ast.Ident)
			if !ok {
				return true
			}

			xString := xIdent.String()
			yString := yIdent.String()

			switch yString {
			case "true":
				fmt.Printf("- got: %s == %s\n+want: %s\n", xString, yString, xString)
			case "false":
				fmt.Printf("- got: %s == %s\n+ want: !%s\n", xString, yString, xString)
			}

			return true
		}
		return true
	})
}
