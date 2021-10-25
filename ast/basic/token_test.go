package basic

import (
	"go/token"
	"testing"
)

func TestParseString(t *testing.T) {
	parseString("hello, 涂鸭鸭", "hi.go")
	println(`~~~~~~~~~~~~~~~~~~~~~`)
	parseString("你好,鸭鸭", "hi.go")
}

func TestPosition(t *testing.T) {
	a := token.Position{Filename: "hi.go", Line: 1, Column: 10}
	position(a)
}

func TestExpr(t *testing.T) {
	exprLite()
	Eval(`2+3*4+ 33`)
	EvalBinary(`2+3*4+x`)
}
