// Copyright (c) Huawei Technologies Co., Ltd. 2012-2019. All rights reserved.
package _144递减元素使数组呈锯齿状

import (
	"github.com/smartystreets/goconvey/convey"
	"testing"
)

func Test_Code(t *testing.T) {
	convey.Convey("1", t, func() {
		var nums = []int{1, 2, 3}
		convey.So(movesToMakeZigzag(nums), convey.ShouldEqual, 2)
	})

	convey.Convey("11", t, func() {
		var nums = []int{9, 6, 1, 6, 2}
		convey.So(movesToMakeZigzag(nums), convey.ShouldEqual, 4)
	})
}
