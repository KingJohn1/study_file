// Copyright (c) Huawei Technologies Co., Ltd. 2012-2019. All rights reserved.
package 结构体copy问题

import (
	"fmt"
	"testing"
)

//func (man *implUpTopoManager) GetTerminateN3UpNode() upcele.UpNodeInfoItf {
//	for i, upNodeInfo := range man.stableData.UpNodeList {
//		if upNodeInfo.GetUpNodeRole().IsTerminateN3Up() {
//			return &man.stableData.UpNodeList[i]
//		}
//	}
//	return nil
//}

func Test_指针引用(t *testing.T) {
	type A struct {
		a int
	}
	var a = []A{}
	a = append(a, A{2})
	b := &a[0]
	b.a = 1
	fmt.Println(b, a)
}
