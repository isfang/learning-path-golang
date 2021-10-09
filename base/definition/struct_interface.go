package definition

import (
	"fmt"
)

type IHello interface {
	Hello(name string)
}

type A struct {
}

func (*A) Hello(name string) {
	fmt.Println("hello " + name + ", i am a")
}

type D struct {
}

func (*D) Hello(name string) {
	fmt.Println("hello " + name + ", i am d")
}

type B struct {
	IHello
}

func (*B) Hello(name string) {
	fmt.Println("hello " + name + ", i am b")
}

type C struct {
	IHello
}
