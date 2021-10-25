package basic

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/scanner"
	"go/token"
	"strconv"
)

func parseString(content string, fileName string) {
	var fset = token.NewFileSet()
	var file = fset.AddFile(fileName, fset.Base(), len(content))

	var s scanner.Scanner

	s.Init(file, []byte(content), nil, scanner.ScanComments)

	for {
		pos, tok, lit := s.Scan()
		if tok == token.EOF {
			break
		}
		fmt.Printf("%s\t%s\t%q\n", fset.Position(pos), tok, lit)
	}
}

func position(p token.Position) {
	fmt.Println(p.String())
}

/**

 */
func exprLite() {

	/*
	   0  *ast.BasicLit {
	   1  .  ValuePos: 1
	   2  .  Kind: INT
	   3  .  Value: "9999"
	   4  }

	*/
	//ast.BasicLit{}

	e, _ := parser.ParseExpr(`9999`)
	ast.Print(nil, e)
	e, _ = parser.ParseExpr(`x`)
	ast.Print(nil, e)

	ast.Print(nil, ast.NewIdent(`x`))

}

func Eval(content string) {
	expr, _ := parser.ParseExpr(content)
	fmt.Println(EvalConstant(expr))
}

func EvalConstant(exp ast.Expr) float64 {
	switch exp := exp.(type) {
	case *ast.BinaryExpr:
		return EvalBinaryExpr(exp)
	case *ast.BasicLit:
		f, _ := strconv.ParseFloat(exp.Value, 64)
		return f
	}
	return 0
}

func EvalBinaryExpr(exp *ast.BinaryExpr) float64 {
	switch exp.Op {
	case token.ADD:
		return EvalConstant(exp.X) + EvalConstant(exp.Y)
	case token.MUL:
		return EvalConstant(exp.X) * EvalConstant(exp.Y)
	}
	return 0
}

func EvalBinary(content string) {
	expr, _ := parser.ParseExpr(content)
	fmt.Println(EvalConstantBinary(expr, map[string]float64{
		"x": 100,
	}))
}

func EvalConstantBinary(exp ast.Expr, vars map[string]float64) float64 {
	switch exp := exp.(type) {
	case *ast.BinaryExpr:
		return EvalBinaryExpr2(exp, vars)
	case *ast.BasicLit:
		f, _ := strconv.ParseFloat(exp.Value, 64)
		return f
	case *ast.Ident:
		return vars[exp.Name]
	}
	return 0
}

func EvalBinaryExpr2(exp *ast.BinaryExpr, vars map[string]float64) float64 {
	switch exp.Op {
	case token.ADD:
		return EvalConstantBinary(exp.X, vars) + EvalConstantBinary(exp.Y, vars)
	case token.MUL:
		return EvalConstantBinary(exp.X, vars) * EvalConstantBinary(exp.Y, vars)
	}
	return 0
}
