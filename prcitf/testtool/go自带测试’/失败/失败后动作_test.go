// Copyright (c) Huawei Technologies Co., Ltd. 2012-2019. All rights reserved.
package 失败

import (
	"fmt"
	"testing"
)

func Test_Error后会继续执行(t *testing.T) {
	//output :
	// === RUN   Test_Error后是否会继续执行
	//end
	//--- FAIL: Test_Error后是否会继续执行 (0.00s)
	//    失败后动作_test.go:10: liuhao
	//FAIL
	t.Error("liuhao")
	fmt.Println("end")
}

func Test_Fail后会继续执行(t *testing.T) {
	//output:
	//=== RUN   Test_Fail后不会继续执行
	//end
	//--- FAIL: Test_Fail后不会继续执行 (0.00s)
	//FAIL
	t.Fail()
	fmt.Println("end")
}

func Test_FailNow不后会继续执行(t *testing.T) {
	//output:
	//=== RUN   Test_FailNow不后会继续执行
	//--- FAIL: Test_FailNow不后会继续执行 (0.00s)
	//FAIL
	t.FailNow()
	fmt.Println("end")
}

func Test_Failed是报告是否有失败_但不会引起失败(t *testing.T) {
	t.Failed()
	fmt.Println("end")
}

func Test_Failed(t *testing.T) {
	//=== RUN   Test_Failed
	//true
	//end
	//--- FAIL: Test_Failed (0.00s)
	//FAIL
	t.Fail()
	failed:= t.Failed()
	fmt.Println(failed)
	fmt.Println("end")
}