// Copyright (c) Huawei Technologies Co., Ltd. 2012-2019. All rights reserved.
package main

import (
	"fmt"
	"runtime"
	"testing"
)

// 查看分配内存

func TestMemStats(t *testing.T){
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Printf("%d Kb\n", m.Alloc / 1024)

}

// 如果需要在一个对象 obj 被从内存移除前执行一些特殊操作，比如写到日志文件中，可以通过如下方式调用函数来实现：
func TestSetFinalizer(t *testing.T){
	//runtime.SetFinalizer(obj, func(obj *typeObj))
}
