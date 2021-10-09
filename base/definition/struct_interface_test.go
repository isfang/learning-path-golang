package definition

import "testing"

func TestInterface(t *testing.T) {
	name := "Lee"
	a := A{}
	a.Hello(name) //hello Lee, i am a

	b := B{&A{}}
	b.Hello(name) //hello Lee, i am b

	b.IHello.Hello(name) //hello Lee, i am a

	c := C{&A{}}
	c.Hello(name) //hello Lee, i am a

	c.IHello = &D{}
	c.Hello(name) //hello Lee, i am d
}
