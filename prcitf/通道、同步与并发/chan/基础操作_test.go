// Copyright (c) Huawei Technologies Co., Ltd. 2012-2019. All rights reserved.
package main

import (
	"fmt"
	"github.com/smartystreets/goconvey/convey"
	"prcitf/util"
	"testing"
	"time"
)

func Test_空通道(t *testing.T) {
	convey.Convey(
		"向nil通道读或者写通道不会panic，但是如果这个通道未被创建会阻塞", t, func() {
			var chan3 chan int
			go func() {
				// ！！！这个用例很有意思，如果注释掉休眠时间，将被阻塞。因为，如果没有休眠 chan3还是空通道；如果休眠chan3 会被赋值到有空间的通道，
				// 注意是值拷贝，所以在值拷贝之前 chan3无效，在值拷贝后有效，这就是休眠后chan3生效的原因
				util.SleepLit()
				select {

				case <-chan3:
				}

				fmt.Println(chan3)
			}()

			chan3 = make(chan int)
			chan3 <- 1

			time.Sleep(1 * time.Second)
		},
	)
	convey.Convey(
		"关闭nil通道panic", t, func() {
			defer util.HasPanic()
			var chan1 chan int
			close(chan1)
		},
	)

}

func Test_关闭通道(t *testing.T) {
	convey.Convey(
		"关闭nil通道panic", t, func() {
			defer util.HasPanic()
			var chan1 chan int
			close(chan1)
		},
	)
	convey.Convey(
		"关闭通道后继续输入panic", t, func() {
			defer util.HasPanic()
			var chan1 chan int = make(chan int)
			close(chan1)
			chan1 <- 1
		},
	)
	convey.Convey(
		"关闭通道后继续输出无问题", t, func() {
			var chan1 chan int = make(chan int)
			close(chan1)
			a := <-chan1
			fmt.Println("a:", a)
			convey.So(a, convey.ShouldAlmostEqual, 0)

		},
	)
	convey.Convey(
		"存在已经阻塞的send goroutine，则此阻塞的goroutine在关闭通道时会panic", t, func() {

			defer util.HasPanic()
			var chan1 chan int = make(chan int)
			go func() {
				util.SleepLit()

				close(chan1)
			}()
			chan1 <- 1
		},
	)
}

func Test_chan输出多参数(t *testing.T) {
	convey.Convey(
		"chan 输出可以含有第二参数ok true标识有通道中有无剩余数据（而非通道是否关闭），可以区分0值是有效值还是close后默认值.close 后不会阻塞", t, func() {
			var ch =make(chan struct{},1)
			ch<- struct{}{}
			a,ok:=<-ch
			fmt.Println(a,ok)
			close(ch)
			a,ok=<-ch
			fmt.Println(a,ok)
			a,ok=<-ch
			fmt.Println(a,ok)
		},
	)

	convey.Convey(
		"chan 输出可以含有第二参数ok true标识有通道中有无剩余数据，而非通道是否会阻塞。即便加ok依然阻塞", t, func() {
			var ch =make(chan struct{},1)
			//ch<- struct{}{}
			_,ok:=<-ch
			fmt.Println(ok)
		},
	)
}

