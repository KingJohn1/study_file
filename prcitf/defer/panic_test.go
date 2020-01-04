// Copyright (c) Huawei Technologies Co., Ltd. 2012-2019. All rights reserved.
package _defer

import (
	"fmt"
	"testing"
)

func Test_Panic(t *testing.T) {
	// 抛出的panic只会进入defer的recover函数,其他defer不会被执行，直到被回复
	func(){
		var a []int

		defer func(){
			if err := recover();err!=nil{
				fmt.Println(err, 2 )
				//debug.PrintStack()
			}

		}()
		defer fmt.Println(a[1])
		defer fmt.Println("sdfasd 2")
		defer func(){
			if err := recover();err!=nil{
				fmt.Println(err,1  )
			//debug.PrintStack()
			}
		}()
		defer fmt.Println("sdfasd 1")
		fmt.Println(a[2])
	}()
}
