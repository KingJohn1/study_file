// Copyright (c) Huawei Technologies Co., Ltd. 2012-2019. All rights reserved.
package 函数练习_test

import (
	"errors"
	"fmt"
	"log"
	"testing"
)

func Test_Print(t *testing.T) {
	var printFormat="liuhao %d dd"
	fmt.Printf(printFormat,666)
	t.SkipNow()
	t.Error(errors.New("dd"))
	t.SkipNow()
	log.Fatal()
}

func BenchmarkSprintf(b* testing.B){
	for i:=0;i<b.N;i++ {
		fmt.Sprintf("liuhao")
	}
}