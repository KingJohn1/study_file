// Copyright (c) Huawei Technologies Co., Ltd. 2012-2019. All rights reserved.
package main

import (
	"fmt"
	"github.com/smartystreets/goconvey/convey"
	"prcitf/util"
	"testing"
)

func Test_Select(t *testing.T) {
	convey.Convey(
		"select中只有一个close的chan可选 会panic", t, func() {
			defer util.HasPanic()
			var chan1 chan int = make(chan int)
			// 关闭chann后，继续给chann输入 会panic

			close(chan1)
			select {
			case chan1 <- 1:
			default:
				fmt.Println("default")
			}
		},
	)
	convey.Convey(
		"通道被close后仍可以不阻塞输出，select中有向close的通道写数据，当select选到此分支则panic", t, func() {
			var chan1 chan int = make(chan int)
			// 关闭chann后，继续给chann输入 会panic

			var chan2 chan int = make(chan int)
			close(chan1)
			select {
			case <-chan1:
				fmt.Println("chan1")

				//	可以从关闭的chan获取输出，因此会被选中
			case chan1 <- 1:
			case chan2 <- 1:
				fmt.Println("chan2")
			default:
				fmt.Println("default")
			}
		},
	)

}

func getChan()chan int{
	fmt.Println("get chan called")
	return make(chan int,0)
}

func Test_SelectCase如果是函数一定会被执行(t *testing.T) {
	ch:=make(chan int,1)
	ch<-1
	select {
	case <-getChan():
		fmt.Println("case 1")
	case <-ch:
		fmt.Println("case 2")
	default:
		fmt.Println("default")
	}
}