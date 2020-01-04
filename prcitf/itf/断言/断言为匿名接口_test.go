// Copyright (c) Huawei Technologies Co., Ltd. 2012-2019. All rights reserved.
package 断言

import (
	"fmt"
)

type S struct {
}

func (s S) get() {
	fmt.Println("get")
}
func Example_断言匿名函数() {

	var s S
	var b interface{} = s
	i, ok := b.(interface{ get() })
	i.get()
	fmt.Println(ok, i )
	// output:
	// get
	// true {}

}
