package copy

import (
	"fmt"
	"sync"
)

type Counter struct {
	sync.Mutex
	Count int
	}

 func copy() {
	var c Counter
	c.Lock()
	defer c.Unlock()
	c.Count++
	foo(c) // 复制锁
	}

 // 这里Counter的参数是通过复制的方式传入的
 func foo(c Counter) {
	 c.Lock()
	 defer c.Unlock()
	 fmt.Println("in foo")
}