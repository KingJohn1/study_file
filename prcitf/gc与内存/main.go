// Copyright (c) Huawei Technologies Co., Ltd. 2012-2019. All rights reserved.
package main

import (
	"fmt"
	"runtime"
)

func main(){
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Printf("%d Kb\n", m.Alloc / 1024)
	var d runtime.MemStats
	runtime.ReadMemStats(&d)
	fmt.Printf("%#v",d)
}

