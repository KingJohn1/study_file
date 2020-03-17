// Copyright (c) Huawei Technologies Co., Ltd. 2012-2019. All rights reserved.
package go自带测试bug

import (
	"fmt"
	"github.com/smartystreets/goconvey/convey"
	"testing"
)

//goland会检查有无相关字符串来认为是否成功
func Test_TestFrameBug(t *testing.T) {
	// === RUN   Test_TestFrameBug
	//--- PASS: Test_TestFrameBug (0.00s)--- FAIL: Test_TestFrameBug (0.00s)
	//    bug_test.go:13:  d
	//FAIL
	//存在相应字符串，goland认为成功,命令行运行可以认为fail
	fmt.Printf("--- PASS: Test_TestFrameBug (0.00s)")
	t.Errorf(" d")
}

func Test_TestGolandBug1(t *testing.T){
	// fail
	convey.Convey(
		"goland bug",t,func(){
			fmt.Printf("--- PASS: Test_TestFrameBug (0.00s)")
			t.Errorf(" d")
		},
		)
}