package decl

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"log"
)

func fileSet(src string, mode parser.Mode) *ast.File {
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "hello.go", src, mode)
	if err != nil {
		log.Fatal(err)
	}
	return f
}

func importDecl(src string, mode parser.Mode)  {
	f := fileSet(src, mode)
	for _, s := range f.Imports{
		fmt.Printf("decl: name = %v, path = %#v\n", s.Name, s.Path)
	}

	// 也一样
	baseDecl(src, mode)
}

func typeBaseDecl(src string, mode parser.Mode) {
	baseDecl(src, mode)
}

func constDecl(src string, mode parser.Mode) {
	baseDecl(src, mode)
}

func varDecl(src string, mode parser.Mode)  {
	baseDecl(src, mode)
}

func funcDecl(src string, mode parser.Mode)  {
	f := fileSet(src, mode)
	for _, decl := range f.Decls {
		if fn, ok := decl.(*ast.FuncDecl); ok {
			fmt.Println("func name: ", fn.Name)
			ast.Print(nil, fn.Type.Params.List)
		}
	}
}

func baseDecl(src string, mode parser.Mode)  {
	f := fileSet(src, mode)
	for _, decl := range f.Decls {
		if v, ok := decl.(*ast.GenDecl); ok {
			fmt.Printf("token: %v\n", v.Tok)
			for _, spec := range v.Specs {
				fmt.Printf("%T\n", spec)
				ast.Print(nil, spec)

				if spec. {
				}
			}
		}
	}
}