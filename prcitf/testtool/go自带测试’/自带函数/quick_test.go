// Copyright (c) Huawei Technologies Co., Ltd. 2012-2019. All rights reserved.
package 自带函数

import (
	"testing"
	"testing/quick"
)

// 对test中quick包使用
func reportError(property string, err error, t *testing.T) {
	if err != nil {
		t.Errorf("%s: %s", property, err)
	}
}

type MyStruct struct {
	X int
}

func myStructProperty(in MyStruct) bool { return in.X == 42 }

func TestCheckProperty(t *testing.T) {
	reportError("myStructProperty", quick.Check(myStructProperty, nil), t)
}
