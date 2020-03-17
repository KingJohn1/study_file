// Copyright (c) Huawei Technologies Co., Ltd. 2012-2019. All rights reserved.
package goroutine

import (
	"fmt"
	"prcitf/util"
	"testing"
)

//
func Test_主gorountine退出_其余gorountinue都退出(t *testing.T) {

	var chan3 chan int = make(chan int)
	go func() {
		util.SleepLit()
		fmt.Println(chan3)
	}()

	fmt.Println("liuhao")

}

