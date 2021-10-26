package decl

import (
	"go/parser"
	"testing"
)

func TestImportDecl(t *testing.T) {
	const src = `package foo
import "pkg-a"
import pkg_b_v2 "pkg-b"
import _ "pkg-c"
import . "pkg-d"
`
	importDecl(src, parser.ImportsOnly)
}

func TestTypeBasicDecl(t *testing.T) {
	/*
	//base decl
	type MyInt1 int
	type MyInt2 = int

	//复合decl
	type Int2 pkg.int
	type IntPtr *int
	type IntPtrPtr **int
	 */
	const src = `package foo
//type MyInt1 int
//type MyInt2 = int
//type Int2 pkg.int
type IntPtr *int
type IntPtrPtr **int
`
	typeBaseDecl(src, parser.AllErrors)
}

func TestConstDecl(t *testing.T) {
	const src = `package foo
const Pi = 3.14
const E float64 = 2.71828
`
	constDecl(src, parser.AllErrors)
}

func TestVarDecl(t *testing.T) {
	const src = `package foo
var Pi = 3.14
`
	varDecl(src, parser.AllErrors)
}

func TestFuncDecl(t *testing.T) {
	const src = `package foo
func hello(string,  string)
func hello2(a string, b string)
func hello3(s0, s1 string,  string, f func(a, b int))
`
	funcDecl(src, parser.AllErrors)
}

func TestStructDecl(t *testing.T) {
	const src = `package foo
type Node struct {
	Next *Node
}
`
	structDecl(src, parser.AllErrors)
}