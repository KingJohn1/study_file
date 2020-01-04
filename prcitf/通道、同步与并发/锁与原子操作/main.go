package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

// 需要用main测试。这些场景也需有自己的测试用例
func main() {
	aaaaa()
	var wg = sync.WaitGroup{}
	var a, b int32

	for i := 0; i < 100; i++ {
		go func() {
			for {
				wg.Add(1)
				a++
				atomic.AddInt32(&b, 1)
				wg.Done()
			}
		}()
	}
	for i := 0; i < 100; i++ {
		go func() {
			time.Sleep(1 * time.Second)
			for {
				wg.Wait()
			}
		}()
	}
	wg.Wait()
	fmt.Println(a, b)
	//util.SleepLit()
}

// 不能导入main包
func aaaaa(){
	var c pool.AAAA
	c.A=1
	fmt.Println(c)
}