package decl

import (
	"fmt"
	"go/ast"
	"go/parser"
)

func structDecl(src string, mode parser.Mode)  {
	f := fileSet(src, mode)
	ast.Print(nil, f.Decls[0].(*ast.GenDecl).Specs[0])
	fmt.Println("~~~~~~~~~~~~~")
	ast.Print(nil, f.Decls[0].(*ast.GenDecl))
	fmt.Println("~~~~~~~~~~~~~")
	ast.Print(nil, f.Decls)
}