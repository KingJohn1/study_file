// Copyright (c) Huawei Technologies Co., Ltd. 2012-2019. All rights reserved.
package runtime

import (
	"fmt"
	"os"
	"runtime"
	"testing"
)

func Test_获取系统相关数据(t *testing.T) {
	fmt.Println(runtime.GOOS)
	fmt.Println(runtime.GOARCH)

	name, err := os.Hostname()
	fmt.Println(err, name)
	name, err = os.UserHomeDir()
	fmt.Println(err, name)

	// 设置并行cpu数 返回之前设置，如果是<1 表示不修改
	fmt.Println(runtime.GOMAXPROCS(8))
	fmt.Println(runtime.GOMAXPROCS(1))
	fmt.Println(runtime.NumCPU())
}
