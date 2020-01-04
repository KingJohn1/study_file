// Copyright (c) liuhao. 2012-2050. All rights reserved.
package main

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"sync"
	"sync/atomic"
	"testing"
)

// todo 思考：这是once的实现，那么使用cas 会有问题吗. 应该是没问题的，但性能差很多，用例见 通道、同步与并发/锁与原子操作/锁与原子操作性能对比_test.go TestBenchMark_CasCompareLoadAndMutex
//func do() {
//	if !atomic.CompareAndSwapInt32(&o.done, 0, 1) {
//		return
//	}
//	f()
//}

func Test_Once使用(t *testing.T) {
	var once = sync.Once{}
	var a int

	for i := 0; i < 100; i++ {
		once.Do(func() {
			a++
		})
	}
	assert.True(t, a == 1)
	fmt.Println(a)

}

// 扩展：go实现单例的方法
// 常量， 变量 init()初始变量或 Getinstance 得到变量 ,类似 sync.once

type One struct {
	num        int32
	mutex      sync.Mutex
	Field, Opt int
}

// 这个函数只能一次生效 。如果得到对象 可以将这个函数作为生产函数，并添加返回值。如果希望方法只执行一次，可以在次执行方法
func (o One) ExecSpeacialOpt() {
	val := atomic.LoadInt32(&o.num)
	if val == 1 {
		return
	}

	o.mutex.Lock()
	defer
	// 注意 语法上 允许 defer 与其执行的语句有间隔，o.mutex.Unlock() 其实是defer的内容
		o.mutex.Unlock()

	// todo 执行只执行一次函数，或返回单例对象
	//func()

	// 在调用函数后需要确保函数已执行完毕，或初始化完全执行。 最好先执行函数再控制val ++ ，这样虽然会有更多的goroutine阻塞，
	// 但是确保了调用此函数后 函数已完整执行（如果无错误），而不是调用此函数后 此函数初始化的东西仍在执行执行过程中
	if val == 0 {
		val++
	}
}

func Test_单例模式(t *testing.T) {
	// 类型once的实现 就是一种单例模式
	var o One
	o.ExecSpeacialOpt()
}
