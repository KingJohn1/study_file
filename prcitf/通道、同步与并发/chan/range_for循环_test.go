// Copyright (c) Huawei Technologies Co., Ltd. 2012-2019. All rights reserved.
package main

import (
	"fmt"
	"testing"
)

func Test_ForRange(t *testing.T) {

	chan1:=make([]chan int,10)

	for i:= range chan1{
		chan1[i]=make(chan int,1)
		chan1[i]<-i+1
		close(chan1[i])
	}

	// range : i<-chan[i]
	for i:= range chan1[1] {

		fmt.Printf("%T %v",i,i)
	}

}
