// Copyright (c) Huawei Technologies Co., Ltd. 2012-2019. All rights reserved.
package main

import (
	"fmt"
	"testing"
)

func TestAdd(t *testing.T) {

	r := Add(1, 2)
	if r !=3 {

		t.Errorf("Add(1, 2) failed. Got %d, expected 3.", r)
	}
	t.SkipNow()

}

func TestAdd1(t *testing.T) {

	r := Add(1, 2)
	if r !=3 {

		t.Errorf("Add(1, 2) failed. Got %d, expected 3.", r)
	}

}

// testmain 可以为整个包的测试做预处理，和结束处理
func TestMain(m *testing.M) {
	fmt.Println("begin")
	m.Run()
	fmt.Println("end")
}
