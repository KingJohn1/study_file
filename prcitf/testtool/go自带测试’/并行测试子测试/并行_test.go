// Copyright (c) Huawei Technologies Co., Ltd. 2012-2019. All rights reserved.
package 并行测试子测试

import (
	"fmt"
	"testing"
	"time"
)

type aa func(t*testing.T)

// todo 尚未搞明白，怎么能并行
func Test_并行与子测试(t *testing.T) {
	fmt.Println("main",t.Name(),time.Now().UTC())
	t.Run("first",func(t*testing.T){
		time.Sleep(1*time.Second)
	fmt.Println("name",t.Name(),time.Now().UTC())
	t.Parallel()

	})
	t.Run("second",func(t*testing.T){
		fmt.Println("name",t.Name(),time.Now().UTC())
		t.Parallel()
	})

	for i:=1;i<3;i++{
		t.Run("second1",func(t*testing.T){
			fmt.Println("name",t.Name(),time.Now().UTC())
			t.Parallel()
		})
	}
	var a aa=func(t*testing.T){
		fmt.Println("name2",t.Name(),time.Now().UTC())
		t.Parallel()
	}
	t.Run("新定义类型",a)
	fmt.Println("main",t.Name(),time.Now().UTC())
}
