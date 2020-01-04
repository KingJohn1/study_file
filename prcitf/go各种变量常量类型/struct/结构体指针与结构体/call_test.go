// Copyright (c) Huawei Technologies Co., Ltd. 2012-2019. All rights reserved.
package 结构体指针与结构体

import (
	"github.com/smartystreets/goconvey/convey"
	"testing"
)

type A struct {
	a int
}

func (a *A) Set(arg int) {
	*a = A{a: arg}
}

func Test_结构体调用结构体指针的方法(t *testing.T) {

	convey.Convey(
		"结构体调用指针的方发，函数中可以改变指针指向对象，即更新方法调用者 (这里不是调用者的参数，而是调用者)", t, func() {
			var sourceValue = 1
			var targetValue = 2
			var a1 = A{sourceValue}
			a1.Set(targetValue)
			convey.So(a1.a, convey.ShouldNotEqual, sourceValue)
			convey.So(a1.a, convey.ShouldEqual, targetValue)
		},
	)
}
