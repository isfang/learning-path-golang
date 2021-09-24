package mutex

import (
	"fmt"
	"sync"
)

/*
 资源竞争有问题方法
*/
func badCounter()  {
	var count = 0
	// 使用WaitGroup等待10个goroutine完成
	var wg sync.WaitGroup
	wg.Add(10)
	for i:=0;i<10;i++{
		go func() {
			defer wg.Done()
			// 对变量count执行10次加1
			for j := 0; j < 100000; j++ {
				count++ }
		}()
	}
	// 等待10个goroutine完成
	wg.Wait()
	fmt.Println(count)
}