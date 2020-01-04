// Copyright (c) Huawei Technologies Co., Ltd. 2012-2019. All rights reserved.
package main_test

import (
	"fmt"
	"testing"
)

type B struct {
	a int
	b string
c []int
}

func (b B)String()string{
return fmt.Sprintf("c:%v",b.c[1])
}

func TestPanic(t *testing.T) {


		var s =B{
			a :1,
			b:"liu",
		}
		//t.Logf("%v",s)
		fmt.Printf("%v",s)


	//time.Sleep(1000*time.Millisecond)
	fmt.Println()
	fmt.Println("end")
}
