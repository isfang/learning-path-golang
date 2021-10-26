package chain

import "fmt"

type Handler interface {
	Handle(Content string)
	next(handler Handler, content string)
}

type IOSHandler struct {
	handler Handler
}

func (h *IOSHandler)Handle(content string) {
	fmt.Println("iOS 开始过滤")
	h.next(h.handler, content)
}

func (h *IOSHandler)next(handler Handler, content string)  {
	if nil != h.handler{
		h.handler.Handle(content)
	}
}



type AndroidHandler struct {
	handler Handler
}

func (h *AndroidHandler)Handle(content string) {
	fmt.Println("Android 开始过滤")
	h.next(h.handler, content)
}

func (h *AndroidHandler)next(handler Handler, content string)  {
	if nil != h.handler{
		h.handler.Handle(content)
	}
}