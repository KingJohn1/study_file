// Copyright (c) Huawei Technologies Co., Ltd. 2012-2019. All rights reserved.
package main

import (
	"fmt"
	"sync/atomic"
)

func main(){

	var a *int32=new(int32)
	*a=2
	i := atomic.AddInt32(a, 1)
	fmt.Println(i)
	val := atomic.LoadInt32(a)
	fmt.Println(val)
}
