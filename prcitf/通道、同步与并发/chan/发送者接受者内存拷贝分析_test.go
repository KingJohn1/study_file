// Copyright (c) Huawei Technologies Co., Ltd. 2012-2019. All rights reserved.
package main

import (
	"fmt"
	"github.com/smartystreets/goconvey/convey"
	"prcitf/util"
	"testing"
	"time"
)

type user struct {
	name string
	age int8
}

var u = user{name: "Ankur", age: 25}
var g = &u

func modifyUser(pu *user) {
	fmt.Println("modifyUser Received Vaule", pu)
	pu.name = "Anand"
}

func printUser(u <-chan *user) {
	time.Sleep(2 * time.Second)
	fmt.Println("printUser goRoutine called", <-u)
}

func Test_发送者将内存发送到缓冲区_接受者从缓冲区读数据(t*testing.T) {
	// todo 仔细品味下的用例，这方便理解chan的内存拷贝
	// 通用原则：没有从chan接收数据前，传入chan 是一个指针，对传入指针的所指向的内存的修改都是有效的
	convey.Convey(
		"chan 没有buf时 ，receive goroutine 直接从send goroutine 的内存拷贝" +
			"因此 在chan接收数据前，对此数据进行修改有效",t, func() {
			c := make(chan *user, 0)
			fmt.Println("main begin fork go ",time.Now())
			go func(){
				fmt.Println("go begin work",time.Now())
				c <- g
			}()
			fmt.Println("main end fork go",time.Now())
			// 为了确保 先执行c <- g后执行 g = &user{name: "Ankur Anand", age: 100}；
			// 如果不休眠大概率先执行g = &user{name: "Ankur Anand", age: 100} 因为新建并启动goroutine需要耗费时间
			util.SleepLit()
			fmt.Println(g)
			// modify g
			g = &user{name: "Ankur Anand", age: 100}
			go printUser(c)
			go modifyUser(g)
			time.Sleep(5 * time.Second)
			fmt.Println(g)
		},
		)
	convey.Convey(
		"chan 有buf时 buf为生产者消费者模型，此模型有有空间读和写时 ，receive goroutine从buf读， send goroutine 写入buf" +
			"在 chan<- 后 ，没有从chan接收数据前；对数据的修改修改都是无效的，因为对于传入的数据直接写入buf，读者直接从buf读，而不在从goroutine读 ",t, func() {
			c := make(chan *user, 0)
			go func(){
				c <- g
			}()
			util.SleepLit()
			fmt.Println(g)
			// modify g
			g = &user{name: "Ankur Anand", age: 100}
			go printUser(c)
			go modifyUser(g)
			time.Sleep(5 * time.Second)
			fmt.Println(g)
		},
	)

}
